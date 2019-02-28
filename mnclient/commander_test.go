package mnclient

import (
	"net/http"
	"os"
	"strconv"
	"testing"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnsender"
)

var addr = "127.0.0.1:2828"

func init() {
	if x := os.Getenv("MARIONETTE_ADDR"); x != "" {
		addr = x
	}
}

func TestCommander(t *testing.T) {
	// start a test webserver
	mux := http.NewServeMux()
	websrv := &http.Server{
		Addr:    "127.0.0.1:9487",
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir("./testdata")))
	go websrv.ListenAndServe()
	defer websrv.Close()

	sender, err := mnsender.NewTCPSender(addr, 0)
	if err != nil {
		t.Fatalf("unexpected error in NewTCPSender(): %s", err)
	}
	sender.Start()
	defer sender.Close()
	cl := &Commander{Sender: sender}

	tc := &cmdrTestCase{Commander: cl}

	t.Run("NewSessionWith", tc.testNewSessionWith)

	// basic commands, need these to use tc.with()
	t.Run("NewWindow", tc.testNewWindow)
	t.Run("SwitchToWindow", tc.testSwitchToWindow)
	t.Run("GetWindowHandles", tc.testGetWindowHandles)
	t.Run("GetWindowHandle", tc.testGetWindowHandle)
	t.Run("CloseWindow", tc.testCloseWindow)
	t.Run("Navigate", tc.testNavigate)
	t.Run("GetCurrentURL", tc.testGetCurrentURL)

	// browser toolbar actions
	t.Run("Back", tc.with(tc.testBack, tc.testNavigate))
	t.Run("Forward", tc.with(tc.testForward, tc.testBack))
	t.Run("Refresh", tc.with(tc.testRefresh, tc.testBack))
	t.Run("GetSetWindowRect", tc.testGetSetWindowRect)
	t.Run("FullscreenWindow", tc.testFullscreenWindow)
	t.Run("MaximizeWindow", tc.testMaximizeWindow)
	// minimize is not tested as it might cause weird result in some os

	var prereq []func(*testing.T)
	// elements
	t.Run("GetActiveElement", tc.with(tc.testGetActiveElement))
	prereq = []func(*testing.T){tc.loadTestHTML("element.html")}
	t.Run("FindElement", tc.with(tc.testFindElement, prereq...))
	t.Run("FindElements", tc.with(tc.testFindElements, prereq...))
	prereq = append(prereq, tc.testFindElement)
	t.Run("GetElementAttribute", tc.with(tc.testGetElementAttribute, prereq...))
	t.Run("GetElementCSSValue", tc.with(tc.testGetElementCSSValue, prereq...))
	t.Run("GetElementProperty", tc.with(tc.testGetElementProperty, prereq...))
	t.Run("GetElementRect", tc.with(tc.testGetElementRect, prereq...))
	t.Run("GetElementText", tc.with(tc.testGetElementText, prereq...))
	t.Run("GetElementTagName", tc.with(tc.testGetElementTagName, prereq...))
	t.Run("IsElementDisplayed", tc.with(tc.testIsElementDisplayed, prereq...))
	t.Run("IsElementSelected", tc.with(tc.testIsElementSelected, prereq...))
	t.Run("IsElementEnabled", tc.with(tc.testIsElementEnabled, prereq...))
	t.Run("ElementClick", tc.with(
		tc.testElementClick,
		append(prereq, tc.testGetElementText)...,
	))
	t.Run("ElementClear", tc.with(
		tc.testElementClear,
		append(prereq, tc.testGetElementProperty)...,
	))
	t.Run("ElementSendKeys", tc.with(
		tc.testElementSendKeys,
		append(prereq, tc.testGetElementProperty)...,
	))

	// informational commands
	t.Run("GetCapabilities", tc.testGetCapabilities)
	t.Run("GetTimeouts", tc.testGetTimeouts)
	t.Run("SetTimeouts", tc.with(tc.testSetTimeouts, tc.testGetTimeouts))
	prereq = []func(*testing.T){tc.loadTestHTML("element.html")}
	t.Run("GetTitle", tc.with(tc.testGetTitle, prereq...))
	t.Run("GetPageSource", tc.with(tc.testGetPageSource, prereq...))

	// js commands
	t.Run("ExecuteScript", tc.with(tc.testExecuteScript))
	t.Run("ExecuteScriptIn", tc.with(tc.testExecuteScriptIn))
	t.Run("ExecuteAsyncScript", tc.with(tc.testExecuteAsyncScript))
	t.Run("ExecuteAsyncScriptIn", tc.with(tc.testExecuteAsyncScriptIn))

	// cookie
	t.Run("Cookies", tc.with(tc.testCookies, prereq...))

	// dialog
	prereq = []func(*testing.T){
		tc.loadTestHTML("element.html"),
		tc.testExecuteScript,
	}
	t.Run("AcceptAlert", tc.with(tc.testAcceptAlert, prereq...))
	t.Run("DismissAlert", tc.with(tc.testDismissAlert, prereq...))
	t.Run("GetAlertText", tc.with(tc.testGetAlertText, prereq...))
	t.Run("SendAlertText", tc.with(tc.testSendAlertText, append(
		prereq,
		tc.testAcceptAlert,
		tc.testDismissAlert,
	)...))

	// screenshot
	prereq = []func(*testing.T){
		func(t *testing.T) {
			tc.must(t, "can-script", tc.testExecuteScript)
			tc.must(t, "can-resize", tc.testGetSetWindowRect)
			tc.SetWindowRect(marionette.Rect{
				X: 100, Y: 100,
				W: 800, H: 600,
			})
			var w, h float64
			_ = tc.ExecuteScript(
				`return window.innerWidth`,
				&w,
			)
			_ = tc.ExecuteScript(
				`return window.innerHeight`,
				&h,
			)

			ok := w == 800 && h == 600
			if !ok {
				tc.SetWindowRect(marionette.Rect{
					X: 100, Y: 100,
					W: 800 + (800 - w),
					H: 600 + (600 - h),
				})
			}
		},
		tc.loadTestHTML("element.html"),
	}
	t.Run("ScreenshotDocument", tc.with(tc.testScreenshotDocument, prereq...))
	t.Run("ScreenshotViewport", tc.with(tc.testScreenshotViewport, prereq...))
	prereq = append(prereq, tc.testFindElement)
	t.Run("ScreenshotElement", tc.with(tc.testScreenshotElement, prereq...))
}

type cmdrTestCase struct {
	*Commander
}

func (tc *cmdrTestCase) loadTestHTML(fn string) (ret func(*testing.T)) {
	return func(t *testing.T) {
		uri := "http://localhost:9487/" + fn
		if err := tc.Navigate(uri); err != nil {
			t.Fatalf("cannot navigate to specified file: %s", err)
		}
	}
}

// ensuring only one tab opened
func (tc *cmdrTestCase) ensureSingleTab(t *testing.T) {
	tabs, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("Setup:cannot get tab list: %s", err)
	}

	if err = tc.SwitchToWindow(tabs[0]); err != nil {
		t.Fatalf("Setup: cannot switch to first tab: %s", err)
	}
	if l := len(tabs); l > 1 {
		for l > 1 {
			tabs, err = tc.CloseWindow()
			l = len(tabs)
			if err != nil {
				t.Fatalf("Setup: cannot close window: %s", err)
			}
			if err = tc.SwitchToWindow(tabs[0]); err != nil {
				t.Fatalf("Setup: cannot switch to next tab: %s", err)
			}
		}
	}
}

// ensuring only one tab opened and navigates to about:blank
func (tc *cmdrTestCase) setup(t *testing.T) {
	t.Run("tab-count", tc.ensureSingleTab)
	if err := tc.Navigate("about:blank"); err != nil {
		t.Fatalf("cannot navigate to blank page: %s", err)
	}
}

// ensuring only one tab opened and navigates to about:logo
func (tc *cmdrTestCase) teardown(t *testing.T) {
	t.Run("tab-count", tc.ensureSingleTab)
	if err := tc.Navigate("about:logo"); err != nil {
		t.Fatalf("cannot navigate to logo page: %s", err)
	}
}

// navigate to a page
func (tc *cmdrTestCase) nav(uri string) (ret func(*testing.T)) {
	return func(t *testing.T) {
		if err := tc.Navigate(uri); err != nil {
			t.Fatalf("failed to goto %s: %s", uri, err)
		}
	}
}

func (tc *cmdrTestCase) must(t *testing.T, name string, f func(*testing.T)) {
	if !t.Run(name, f) {
		t.SkipNow()
	}
}

// run selected test case with setup and teardown
func (tc *cmdrTestCase) with(
	f func(*testing.T), reqs ...func(*testing.T),
) (ret func(*testing.T)) {
	return func(t *testing.T) {
		ok := t.Run("Setup", tc.setup)
		if !ok {
			t.Skip("setup failed, skipping")
		}

		if len(reqs) > 0 {
			ok = t.Run("PreReq", func(t *testing.T) {
				for idx, f := range reqs {
					t.Run(
						"#"+strconv.Itoa(idx),
						f,
					)
				}
			})
			if !ok {
				t.Skip("pre-requirement failed, skipping")
			}
		}

		t.Run("Case", f)
		t.Run("Teardown", tc.teardown)
	}
}
