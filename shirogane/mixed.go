package shirogane

import (
	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
)

// Mixed is basic mixed mode client
type Mixed struct {
	Addr string // marionette server address, use 127.0.0.1:2828 if leave empty

	client *Async
}

// Start connects to marionette server and start the mainloop
func (s *Mixed) Start() (err error) {
	addr := s.Addr
	if addr == "" {
		addr = "127.0.0.1:2828"
	}
	conn, err := marionette.ConnectTo(addr)
	if err != nil {
		return
	}

	s.client = &Async{Conn: conn}
	s.client.Start()
	return conn.Wait()
}

// Close exits the mainloop and release all related resources
func (s *Mixed) Close() (err error) {
	s.client.Stop()
	return s.client.Conn.Close()
}

// Sync send command synchronously (block until response actually)
func (s *Mixed) Sync(cmd ito.Ito) (msg *marionette.Message, err error) {
	msgch, err := s.client.Send(cmd)
	if err != nil {
		return
	}

	msg = <-msgch
	err = msg.Error
	return
}

// Async send command asynchronously
func (s *Mixed) Async(cmd ito.Ito) (ch chan *marionette.Message, err error) {
	return s.client.Send(cmd)
}