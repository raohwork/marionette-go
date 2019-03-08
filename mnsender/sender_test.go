// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

package mnsender

import (
	"testing"
)

// it's quite hard to write test cases for concurrent program, so I use benchmark
// instead
func BenchmarkSenderAsync(b *testing.B) {
	srv, rw := newFakeServer()
	srv.Start()
	defer srv.Stop()

	sender := NewSender(rw, 0)

	if err := sender.Start(); err != nil {
		b.Fatalf("unexpected error in Start(): %s", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch, err := sender.Async(fakeCmd{})
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

func BenchmarkSenderSync(b *testing.B) {
	srv, rw := newFakeServer()
	srv.Start()
	defer srv.Stop()

	sender := NewSender(rw, 0)

	if err := sender.Start(); err != nil {
		b.Fatalf("unexpected error in Start(): %s", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msg, err := sender.Sync(fakeCmd{})
			if err != nil {
				b.Fatalf("unexpected error in Send(): %s", err)
			}

			if msg == nil || msg.Error != nil {
				b.Fatalf("invalid response: %+v", msg)
			}
		}
	})
}
