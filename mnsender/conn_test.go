package mnsender

import (
	"testing"
)

func TestConn(t *testing.T) {
	rw, r, w := newTestRW(nil)

	// simulate system info packet
	r.WriteString("1:1")
	// some data packets
	r.WriteString(`15:[1,1,null,null]`)
	r.WriteString(`15:[1,2,null,null]`)

	cl, err := NewConn(rw, 0)
	if err != nil {
		t.Fatalf("unexpected error in NewConn(): %s", err)
	}

	t.Run("send", func(t *testing.T) {
		id, err := cl.Send("test", nil)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if id != 1 {
			t.Errorf("invalid id: %d", id)
		}

		if str := w.String(); str != `17:[0,1,"test",null]` {
			t.Errorf("unexpected packet: %s", str)
		}
	})

	t.Run("receive", func(t *testing.T) {
		msg := <-cl.ResultChan()
		if msg.Type != 1 {
			t.Errorf("unexpected type: %d", msg.Type)
		}
		if msg.Serial != 1 {
			t.Errorf("unexpected serial: %d", msg.Serial)
		}
		if msg.Data != nil {
			t.Errorf("unexpected data: %+v", msg.Data)
		}
		if msg.Error != nil {
			t.Errorf("unexpected error message: %+v", msg.Error)
		}
	})

	cl.Close()
	cl.Cleanup()

	if err = cl.Wait(); err != nil {
		t.Fatalf("unexpected error in Wait(): %s", err)
	}
}
