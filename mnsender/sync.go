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
	"errors"
	"sync"

	marionette "github.com/raohwork/marionette-go/v3"
	"github.com/raohwork/marionette-go/v3/mncmd"
)

// Sync is very basic synchronized client
//
// This client is not tested as most of logics are in Conn, which is tested.
type Sync struct {
	Conn *Conn

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
