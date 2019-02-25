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
