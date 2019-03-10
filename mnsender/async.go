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
	"context"
	"errors"
	"sync"

	marionette "github.com/raohwork/marionette-go/v3"
	"github.com/raohwork/marionette-go/v3/mncmd"
)

// Async is very basic asynchronized command send/receiver
type Async struct {
	Conn *Conn

	mapLock sync.Mutex
	pending map[uint32]chan *marionette.Message

	ctx     context.Context
	cancel  context.CancelFunc
	running chan struct{}
}

// Send sends a command to server, returns a channel immediately
//
// The client will close the channel once message is transmitted.
//
// Calling Send() on a stopped client leads to nil pointer panic!
func (s *Async) Send(cmd mncmd.Command) (resp chan *marionette.Message, err error) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()

	if s.ctx == nil {
		err = errors.New("async client has not started")
		return
	}
	if !cmd.Validate() {
		return nil, errors.New("invalid command")
	}

	id, err := s.Conn.Send(cmd.Command(), cmd.Param())
	if err != nil {
		return
	}

	resp = make(chan *marionette.Message, 1)
	s.pending[id] = resp

	return
}

// Start runs the main loop at background to receive/dispatch messages
func (s *Async) Start() {
	s.mapLock.Lock()
	s.pending = map[uint32]chan *marionette.Message{}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.running = make(chan struct{})
	s.mapLock.Unlock()
	go s.mainLoop()
}

// Wait blocks until main loop stops
func (s *Async) Wait() {
	<-s.running
}

// Stop stops the main loop and clear pending requests
func (s *Async) Stop() {
	s.cancel()
	s.Wait()
}

func (s *Async) mainLoop() {
	for {
		select {
		case x := <-s.Conn.ResultChan():
			if x != nil {
				s.dispatch(x)
			}
		case <-s.ctx.Done():
			s.mapLock.Lock()
			defer s.mapLock.Unlock()
			s.ctx = nil
			errMsg := errors.New("client exit")
			for id, x := range s.pending {
				if len(x) == 0 {
					x <- &marionette.Message{
						Error:  errMsg,
						Serial: id,
					}
				}
				close(x)
			}
			s.pending = nil
			close(s.running)
			return
		}
	}
}

func (s *Async) dispatch(msg *marionette.Message) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()

	if s.pending == nil {
		// stopped
		return
	}
	ch := s.pending[msg.Serial]
	delete(s.pending, msg.Serial)
	ch <- msg
	close(ch)
}
