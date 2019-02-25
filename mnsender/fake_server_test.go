package mnsender

import (
	"context"
	"encoding/json"
	"io"
)

// fakeCmd is a command mock
type fakeCmd struct{}

func (c fakeCmd) Command() (ret string) {
	return "fake command"
}

func (c fakeCmd) Param() (ret interface{}) {
	return
}

func (c fakeCmd) Validate() (ok bool) {
	return true
}

// pipedRWC is a connection between client and fake server
type pipedRWC struct {
	io.Reader
	io.Writer
}

func (c *pipedRWC) Close() (err error) {
	return
}

// fakeServer mocks marionette server, the only supported command is fakeCmd
type fakeServer struct {
	ctx       context.Context
	cancel    context.CancelFunc
	transport *transport
}

func (s *fakeServer) Start() {
	go s.mainloop()
}

func (s *fakeServer) Stop() {
	s.cancel()
}

func (s *fakeServer) mainloop() {
	// init packet
	s.transport.Send(map[string]string{"test": "test"})

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.mainfunc()
		}
	}
}

func (s *fakeServer) mainfunc() {
	data, err := s.transport.Receive()
	if err != nil {
		return
	}

	var msg []interface{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return
	}

	id, ok := msg[1].(float64)
	if !ok {
		return
	}

	// due to pipe design, we have to send message in another goroutine
	go func() {
		s.transport.Send([]interface{}{
			1, int(id), nil, map[string]string{},
		})
	}()
}

func newFakeServer() (srv *fakeServer, rwc *pipedRWC) {
	r, send := io.Pipe()
	recv, w := io.Pipe()

	rwc = &pipedRWC{
		Reader: r,
		Writer: w,
	}

	ctx, cancel := context.WithCancel(context.Background())
	conn := &pipedRWC{
		Reader: recv,
		Writer: send,
	}
	trans := &transport{conn: conn}
	srv = &fakeServer{
		transport: trans,
		ctx:       ctx,
		cancel:    cancel,
	}

	return
}
