package shirogane

import (
	"errors"
	"sync"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mncmd"
)

// Sync is very basic synchronized client
type Sync struct {
	Conn *marionette.Conn

	lock sync.Mutex
}

// Send sends a command to marionette server, blocks until it gets the response
func (s *Sync) Send(cmd mncmd.Command) (resp *marionette.Message, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !cmd.Validate() {
		return nil, errors.New("invalid command")
	}

	_, err = s.Conn.Send(cmd.Command(), cmd.Param())
	if err == nil {
		resp = <-s.Conn.ResultChan()
	}

	return
}
