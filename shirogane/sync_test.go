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
}
