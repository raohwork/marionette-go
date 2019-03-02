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

// it's quite hard to write test cases for concurrent program, so I use benchmark
// instead
func BenchmarkAsync(b *testing.B) {
	srv, rw := newFakeServer()
	srv.Start()
	defer srv.Stop()

	conn, err := NewConn(rw, 0)
	if err != nil {
		b.Fatalf("unexpected error in NewConn(): %s", err)
	}

	cl := &Async{Conn: conn}
	cl.Start()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch, err := cl.Send(fakeCmd{})
			if err != nil {
				b.Fatalf("unexpected error in Send(): %s", err)
			}

			msg := <-ch
			if msg == nil || msg.Error != nil {
				b.Fatalf("invalid response: %+v", msg)
			}
		}
	})
}
