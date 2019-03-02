// This file is part of marionette-go
//
// marionette-go is free software: you can redistribute it and/or modify it
// under the terms of the GNU Lesser General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with marionette-go. If not, see <https://www.gnu.org/licenses/>.

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
