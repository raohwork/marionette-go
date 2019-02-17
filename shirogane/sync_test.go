package shirogane

import (
	"log"
	"os"
	"testing"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
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

	try := func(cmd ito.Ito) *marionette.Message {
		resp, err := s.Send(cmd)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		t.Logf("Result: %+v", resp)

		return resp
	}

	cSess := &ito.NewSession{}
	msg := try(cSess)
	id, caps, err := cSess.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding NewSession response: %s", err)
	} else {
		t.Logf("session id: %s", id)
		t.Logf("capabilities: %+v", caps)
	}

	cHandles := &ito.GetWindowHandles{}
	msg = try(cHandles)
	handles, err := cHandles.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowHandles response: %s", err)
	} else {
		t.Logf("window handles: %+v", handles)
	}

	cHandle := &ito.GetWindowHandle{}
	msg = try(cHandle)
	curid, err := cHandle.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowHandle response: %s", err)
	} else {
		t.Logf("current handle: %s", curid)
	}

	try(&ito.GetChromeWindowHandles{})
	try(&ito.GetChromeWindowHandle{})

	cRect := &ito.GetWindowRect{}
	msg = try(cRect)
	rect, err := cRect.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding GetWindowRect response: %s", err)
	} else {
		t.Logf("window rect: %+v", rect)
	}
	try(&ito.FullscreenWindow{})
	// try(&ito.MinimizeWindow{})
	try(&ito.MaximizeWindow{})
	try(&ito.SetWindowRect{Rect: rect})

	cNewWin := &ito.NewWindow{Type: "tab", Focus: true}
	msg = try(cNewWin)
	newID, _, err := cNewWin.Decode(msg)
	if err != nil {
		t.Errorf("Error decoding NewWindow response: %s", err)
	} else {
		t.Logf("new window handle: %s", newID)
	}

	try(&ito.SwitchToWindow{Name: curid})
	try(&ito.SwitchToWindow{Name: newID})
	try(&ito.CloseWindow{})
	try(&ito.SwitchToWindow{Name: curid})
	try(&ito.SetTimeouts{Timeouts: &marionette.Timeouts{
		Implicit: 30000,
		PageLoad: 30000,
		Script:   30000,
	}})
	(&ito.GetTimeouts{}).Decode(try(&ito.GetTimeouts{}))
	(&ito.GetCapabilities{}).Decode(try(&ito.GetCapabilities{}))

	cFindEl := &ito.FindElement{
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
	try(&ito.ElementClick{Element: el})
	try(&ito.ElementSendKeys{Element: el, Text: "test"})
	try(&ito.ElementClear{Element: el})

	cFindEl.Value = "form"
	el, _ = cFindEl.Decode(try(cFindEl))
	msg = try(&ito.FindElements{
		Using:       marionette.Selector,
		Value:       "a",
		RootElement: el,
	})
	if _, err := (&ito.FindElements{}).Decode(msg); err != nil {
		t.Errorf("Error decoding FindElements response: %s", err)
	}

	try(&ito.GetActiveElement{})

	cGetProp := &ito.GetElementAttribute{
		Element: el,
		Name:    "id",
	}
	if id, err := cGetProp.Decode(try(cGetProp)); err != nil {
		t.Errorf("Error decoding GetElementAttribute response: %s", err)
	} else {
		t.Logf(`element id="%s"`, id)
	}

	cGetCSS := &ito.GetElementCSSValue{
		Element: el,
		Prop:    "max-width",
	}
	if ret, err := cGetCSS.Decode(try(cGetCSS)); err != nil {
		t.Errorf("Error decoding GetElementCSSValue response: %s", err)
	} else {
		t.Logf(`element css="%s"`, ret)
	}

	(&ito.GetElementProperty{}).Decode(try(&ito.GetElementProperty{Element: el, Name: "isConnected"}))
	(&ito.GetElementRect{}).Decode(try(&ito.GetElementRect{Element: el}))
	(&ito.GetElementTagName{}).Decode(try(&ito.GetElementTagName{Element: el}))
	(&ito.GetElementText{}).Decode(try(&ito.GetElementText{Element: el}))
	(&ito.IsElementDisplayed{}).Decode(try(&ito.IsElementDisplayed{Element: el}))
	(&ito.IsElementEnabled{}).Decode(try(&ito.IsElementEnabled{Element: el}))
	(&ito.IsElementSelected{}).Decode(try(&ito.IsElementSelected{Element: el}))
}
