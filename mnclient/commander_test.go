package mnclient

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/raohwork/marionette-go/mnsender"
)

var addr = "127.0.0.1:2828"

func init() {
	if x := os.Getenv("MARIONETTE_ADDR"); x != "" {
		addr = x
	}
}

func TestCommander(t *testing.T) {
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
	t.Run("Back", tc.with(tc.testBack))
	t.Run("Forward", tc.with(tc.testForward))
	t.Run("Refresh", tc.with(tc.testRefresh))
	t.Run("GetSetWindowRect", tc.testGetSetWindowRect)
	t.Run("FullscreenWindow", tc.testFullscreenWindow)
	t.Run("MaximizeWindow", tc.testMaximizeWindow)
	// minimize is not tested as it might cause weird result in some os

	// elements
	t.Run("FindElement", tc.with(tc.testFindElement))
	t.Run("FindElements", tc.with(tc.testFindElements))
	t.Run("GetElementAttribute", tc.with(tc.testGetElementAttribute))
	t.Run("GetElementCSSValue", tc.with(tc.testGetElementCSSValue))
	t.Run("GetElementProperty", tc.with(tc.testGetElementProperty))
	t.Run("GetElementRect", tc.with(tc.testGetElementRect))
}

type cmdrTestCase struct {
	*Commander
}

func (tc *cmdrTestCase) loadTestHTML(fn string) (ret func(*testing.T)) {
	return func(t *testing.T) {
		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("cannot get working directory: %s", err)
		}
		real, err := filepath.Abs(pwd)
		if err != nil {
			t.Fatalf("cannot get absolute path: %s", err)
		}
		uri := "file://" + real + "/testdata/" + fn
		if err = tc.Navigate(uri); err != nil {
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
func (tc *cmdrTestCase) with(f func(*testing.T)) (ret func(*testing.T)) {
	return func(t *testing.T) {
		ok := t.Run("Setup", tc.setup)
		if !ok {
			t.Skip("setup failed, skipping")
		}
		t.Run("Case", f)
		t.Run("Teardown", tc.teardown)
	}
}
