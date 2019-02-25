package mnsender

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

type testRWC struct {
	io.Reader
	io.Writer
	CloseErr error
}

func (t *testRWC) Close() (err error) {
	return t.CloseErr
}

func newTestRW(closeErr error) (rw *testRWC, r, w *bytes.Buffer) {
	r = &bytes.Buffer{}
	w = &bytes.Buffer{}
	rw = &testRWC{
		Reader:   r,
		Writer:   w,
		CloseErr: closeErr,
	}

	return
}

// ensuring transport works perfectly if underlying connection is stable.
// might write tests for buggy situations like connection interrupted in the future.
func TestTransport(t *testing.T) {
	rw, r, w := newTestRW(nil)
	tr := &transport{conn: rw}

	r.WriteString("3:123")
	r.WriteString(`6:"test"`)
	r.WriteString(`12:{"test":"a"}`)
	r.WriteString(`9:[4,3,2,1]`)
	t.Run("receive-int", func(t *testing.T) {
		buf, err := tr.Receive()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		var data int
		err = json.Unmarshal(buf, &data)
		if err != nil {
			t.Fatalf("cannot decode to integer: %s", err)
		}
		if data != 123 {
			t.Fatalf("expected to be 123, got %d", data)
		}
	})

	t.Run("receive-str", func(t *testing.T) {
		buf, err := tr.Receive()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		var data string
		err = json.Unmarshal(buf, &data)
		if err != nil {
			t.Fatalf("cannot decode to string: %s", err)
		}
		if data != "test" {
			t.Fatalf("expected to be test, got %s", data)
		}
	})

	t.Run("receive-object", func(t *testing.T) {
		buf, err := tr.Receive()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		var data map[string]string
		err = json.Unmarshal(buf, &data)
		if err != nil {
			t.Fatalf("cannot decode to string map: %s", err)
		}
		if len(data) != 1 || data["test"] != "a" {
			t.Fatalf("expected to be {test:a}, got %+v", data)
		}
	})

	t.Run("receive-arr", func(t *testing.T) {
		buf, err := tr.Receive()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		var data []int
		err = json.Unmarshal(buf, &data)
		if err != nil {
			t.Fatalf("cannot decode to int array: %s", err)
		}
		if len(data) != 4 ||
			data[0] != 4 ||
			data[1] != 3 ||
			data[2] != 2 ||
			data[3] != 1 {
			t.Fatalf("expected to be [4,3,2,1], got %+v", data)
		}
	})

	t.Run("send-int", func(t *testing.T) {
		defer w.Reset()
		err := tr.Send(234)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if str := w.String(); str != "3:234" {
			t.Fatalf("expected '3:234', got %s", str)
		}
	})

	t.Run("send-str", func(t *testing.T) {
		defer w.Reset()
		err := tr.Send("wheezy")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if str := w.String(); str != `8:"wheezy"` {
			t.Fatalf(`expected '8:"wheezy"', got %s`, str)
		}
	})

	t.Run("send-obj", func(t *testing.T) {
		defer w.Reset()
		err := tr.Send(map[string]bool{"test": true})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if str := w.String(); str != `13:{"test":true}` {
			t.Fatalf(`expected '13:{"test":true}', got %s`, str)
		}
	})

	t.Run("send-arr", func(t *testing.T) {
		defer w.Reset()
		err := tr.Send([]interface{}{1, false, nil, "test", map[string]string{"x": "y"}})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if str := w.String(); str != `31:[1,false,null,"test",{"x":"y"}]` {
			t.Fatalf(`expected '31:[1,false,null,"test",{"x":"y"}]', got %s`, str)
		}
	})
}
