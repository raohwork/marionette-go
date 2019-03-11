package mnclient

import (
	"testing"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnsender"
)

func TestWindowResizer(t *testing.T) {
	sender, err := mnsender.NewTCPSender(addr, 0)
	if err != nil {
		t.Fatalf("unexpected error in NewTCPSender(): %s", err)
	}
	sender.Start()
	defer sender.Close()
	cl := &Commander{Sender: sender}
	cl.NewSession()

	resize := &WindowResizer{Commander: cl}

	t.Run("Outer", func(t *testing.T) {
		ret, err := resize.Outer(marionette.Rect{
			W: 800,
			H: 600,
		})
		if err != nil {
			t.Fatalf("cannot set outer size: %s", err)
		}
		if ret.W != 800 {
			t.Errorf("unexpected width: %f", ret.W)
		}
		if ret.H != 600 {
			t.Errorf("unexpected height: %f", ret.H)
		}
	})
	t.Run("Inner", func(t *testing.T) {
		ret, err := resize.Inner(marionette.Rect{
			W: 800,
			H: 600,
		})
		if err != nil {
			t.Fatalf("cannot set inner size: %s", err)
		}
		if ret.W < 800 {
			t.Errorf("unexpected width: %f", ret.W)
		}
		if ret.H < 600 {
			t.Errorf("unexpected height: %f", ret.H)
		}

		var w, h float64
		err = cl.ExecuteScript(
			`return window.innerWidth`,
			&w,
		)
		if err != nil {
			return
		}
		if w != 800 {
			t.Errorf("unexpected width: %f", w)
		}

		err = cl.ExecuteScript(
			`return window.innerHeight`,
			&h,
		)
		if err != nil {
			return
		}
		if h != 600 {
			t.Errorf("unexpected height: %f", h)
		}
	})
}
