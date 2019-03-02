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
	"errors"
	"sync"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mncmd"
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
