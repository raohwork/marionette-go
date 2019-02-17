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

	try := func(cmd ito.Ito) {
		resp, err := s.Send(cmd)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		t.Logf("Result: %+v", resp)
	}

	try(&ito.NewSession{})
	try(&ito.GetWindowHandles{})
	try(&ito.GetWindowHandle{})
	try(&ito.GetChromeWindowHandles{})
	try(&ito.GetChromeWindowHandle{})
}
