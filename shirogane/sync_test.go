package shirogane

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mncmd"
)

var addr string

func init() {
	addr = os.Getenv("MARIONETTE_ADDR")
	if addr == "" {
		log.Fatal("You must set envvar MARIONETTE_ADDR to run tests")
	}
}

func connect(t *testing.T) (ret *marionette.Conn) {
	ret, err := marionette.ConnectTo(addr)
	if err != nil {
		t.Fatalf("Cannot connect to remote end: %s", err)
	}

	return
}

func TestSyncClient(t *testing.T) {
	conn := connect(t)
	defer conn.Close()

	s := &Sync{Conn: conn}

	try := func(cmd mncmd.Command) *marionette.Message {
		resp, err := s.Send(cmd)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		t.Logf("Result: %+v", resp)

		return resp
	}

	cSess := &mncmd.NewSession{}
	msg := try(cSess)
	id, caps, err := cSess.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding NewSession response: %s", err)
	} else {
		t.Logf("session id: %s", id)
		t.Logf("capabilities: %+v", caps)
	}

	cHandles := &mncmd.GetWindowHandles{}
	msg = try(cHandles)
	handles, err := cHandles.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowHandles response: %s", err)
	} else {
		t.Logf("window handles: %+v", handles)
	}

	cHandle := &mncmd.GetWindowHandle{}
	msg = try(cHandle)
	curid, err := cHandle.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowHandle response: %s", err)
	} else {
		t.Logf("current handle: %s", curid)
	}

	try(&mncmd.GetChromeWindowHandles{})
	try(&mncmd.GetChromeWindowHandle{})

	cRect := &mncmd.GetWindowRect{}
	msg = try(cRect)
	rect, err := cRect.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowRect response: %s", err)
	} else {
		t.Logf("window rect: %+v", rect)
	}
	try(&mncmd.FullscreenWindow{})
	// try(&mncmd.MinimizeWindow{})
	try(&mncmd.MaximizeWindow{})
	try(&mncmd.SetWindowRect{Rect: rect})

	cNewWin := &mncmd.NewWindow{Type: "tab", Focus: true}
	msg = try(cNewWin)
	newID, _, err := cNewWin.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding NewWindow response: %s", err)
	} else {
		t.Logf("new window handle: %s", newID)
	}

	try(&mncmd.SwitchToWindow{Name: curid})
	try(&mncmd.SwitchToWindow{Name: newID})
	try(&mncmd.CloseWindow{})
	try(&mncmd.SwitchToWindow{Name: curid})
	try(&mncmd.SetTimeouts{Timeouts: &marionette.Timeouts{
		Implicit: 30000,
		PageLoad: 30000,
		Script:   30000,
	}})
	(&mncmd.GetTimeouts{}).Decode(try(&mncmd.GetTimeouts{}))
	(&mncmd.GetCapabilities{}).Decode(try(&mncmd.GetCapabilities{}))

	try(&mncmd.Refresh{})
	(&mncmd.GetPageSource{}).Decode(try(&mncmd.GetPageSource{}))
	try(&mncmd.Navigate{URL: "https://google.com"})
	time.Sleep(time.Second)
	try(&mncmd.Back{})
	try(&mncmd.Forward{})
	(&mncmd.GetTitle{}).Decode(try(&mncmd.GetTitle{}))
	(&mncmd.GetCurrentURL{}).Decode(try(&mncmd.GetCurrentURL{}))

	cFindEl := &mncmd.FindElement{
		Using: marionette.Selector,
		Value: `input[name="q"]`,
	}
	msg = try(cFindEl)
	el, err := cFindEl.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding FindElement response: %s", err)
	} else {
		t.Logf("element: %+v", el)
	}
	try(&mncmd.ElementClick{Element: el})
	try(&mncmd.ElementSendKeys{Element: el, Text: "test"})
	try(&mncmd.ElementClear{Element: el})

	cFindEl.Value = "form"
	el, _ = cFindEl.Decode(try(cFindEl))
	msg = try(&mncmd.FindElements{
		Using:       marionette.Selector,
		Value:       "a",
		RootElement: el,
	})
	if _, err := (&mncmd.FindElements{}).Decode(msg); err != nil {
		t.Errorf("Error decoding FindElements response: %s", err)
	}

	try(&mncmd.GetActiveElement{})

	cGetProp := &mncmd.GetElementAttribute{
		Element: el,
		Name:    "id",
	}
	if id, err := cGetProp.Decode(try(cGetProp)); err != nil {
		t.Errorf("Error decoding GetElementAttribute response: %s", err)
	} else {
		t.Logf(`element id="%s"`, id)
	}

	cGetCSS := &mncmd.GetElementCSSValue{
		Element: el,
		Prop:    "max-width",
	}
	if ret, err := cGetCSS.Decode(try(cGetCSS)); err != nil {
		t.Errorf("Error decoding GetElementCSSValue response: %s", err)
	} else {
		t.Logf(`element css="%s"`, ret)
	}

	(&mncmd.GetElementProperty{}).Decode(try(&mncmd.GetElementProperty{Element: el, Name: "isConnected"}))
	(&mncmd.GetElementTagName{}).Decode(try(&mncmd.GetElementTagName{Element: el}))
	(&mncmd.GetElementText{}).Decode(try(&mncmd.GetElementText{Element: el}))
	(&mncmd.IsElementDisplayed{}).Decode(try(&mncmd.IsElementDisplayed{Element: el}))
	(&mncmd.IsElementEnabled{}).Decode(try(&mncmd.IsElementEnabled{Element: el}))
	(&mncmd.IsElementSelected{}).Decode(try(&mncmd.IsElementSelected{Element: el}))

	cFindEl.Value = "a"
	el, _ = cFindEl.Decode(try(cFindEl))

	cElRect := &mncmd.GetElementRect{Element: el}
	rect, err = cElRect.Decode(try(cElRect))
	if err != nil {
		t.Errorf("Error decoding GetElementRect response: %s", err)
	}

	t.Run("screenshot", func(t *testing.T) {
		cScrShot := &mncmd.TakeScreenshot{
			Highlights: []*marionette.WebElement{el},
		}
		if png, err := cScrShot.Decode(try(cScrShot)); err != nil {
			t.Errorf("Error decoding TakeScreenshot response: %s", err)
		} else {
			buf, err := base64.StdEncoding.DecodeString(png)
			if err == nil {
				ioutil.WriteFile("shot.png", buf, 0600)
			}
		}
	})

	// js
	cJS := &mncmd.ExecuteScript{Script: `return {x:"test"}`}
	msg = try(cJS)
	var jsResp map[string]string
	if err := cJS.Decode(msg, &jsResp); err != nil {
		t.Errorf("Error decoding ExecuteJavascript response: %s", err)
	} else {
		t.Logf("js reply: %+v", jsResp)
	}
	cAsyncJS := &mncmd.ExecuteAsyncScript{Script: `arguments[0]("test")`}
	msg = try(cAsyncJS)
	var asyncJSResp string
	if err := cAsyncJS.Decode(msg, &asyncJSResp); err != nil {
		t.Errorf("Error decoding ExecuteAsyncJavascript response: %s", err)
	} else {
		t.Logf("js reply: %s", asyncJSResp)
	}

	// open a alert dialog
	cJS.Script = `setTimeout(() => alert('test'), 0)`
	try(cJS)
	time.Sleep(100 * time.Millisecond)
	cDlgText := &mncmd.GetAlertText{}
	if data, err := cDlgText.Decode(try(cDlgText)); err != nil {
		t.Errorf("Error decoding GetAlertText response: %s", err)
	} else {
		if data != "test" {
			t.Errorf("expect alert('test'), got %s", data)
		}
	}
	try(&mncmd.AcceptAlert{})

	try(cJS)
	time.Sleep(100 * time.Millisecond)
	try(&mncmd.DismissAlert{})

	cJS.Script = `setTimeout(() => prompt('test'), 0)`
	try(cJS)
	time.Sleep(100 * time.Millisecond)
	try(&mncmd.SendAlertText{Text: "test2"})
	try(&mncmd.DismissAlert{})

	cGetCookies := &mncmd.GetCookies{}
	msg = try(cGetCookies)
	cookies, err := cGetCookies.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetCookies response: %s", err)
	}
	try(&mncmd.AddCookie{Cookie: &marionette.Cookie{
		Name:  "marionette",
		Value: "puppet",
	}})
	msg = try(cGetCookies)
	ncookies, err := cGetCookies.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetCookies response: %s", err)
	}
	if x, y := len(cookies), len(ncookies); x != y-1 {
		t.Errorf("Expected %d cookies, got %d", x+1, y)
	}
	try(&mncmd.DeleteAllCookies{})

	// marionette
	try(&mncmd.MozSetContext{Context: marionette.ChromeContext})
	(&mncmd.MozGetContext{}).Decode(try(&mncmd.MozGetContext{}))
	try(&mncmd.MozSetContext{Context: marionette.ContentContext})
	try(&mncmd.MozGetScreenOrientation{})
	try(&mncmd.MozGetWindowType{})

	// actions
	chain := marionette.ActionChain{}
	chain.
		MouseMoveTo(int(rect.X+rect.W/2), int(rect.Y+rect.H/2), 100).
		MouseDown(marionette.MouseLeft).
		MouseUp(marionette.MouseLeft)
	cAct := &mncmd.PerformActions{
		Actions: chain,
	}
	try(cAct)

	try(&mncmd.CloseWindow{})
}
