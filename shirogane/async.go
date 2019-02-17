package shirogane

import (
	"context"
	"errors"
	"sync"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
)

type Async struct {
	Conn *marionette.Conn

	mapLock sync.Mutex
	pending map[uint32]chan *marionette.Message

	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Async) Send(cmd ito.Ito) (resp chan *marionette.Message, err error) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()

	if s.ctx == nil {
		err = errors.New("async client has not started")
		return
	}
	if !cmd.Validate() {
		return nil, errors.New("invalid command")
	}

	id, err := s.Conn.Send(cmd.Command(), cmd.Param)
	if err != nil {
		return
	}

	resp = make(chan *marionette.Message, 1)
	s.pending[id] = resp

	return
}

func (s *Async) Start() {
	s.pending = map[uint32]chan *marionette.Message{}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.mainLoop()
}

func (s *Async) Stop() {
	s.cancel()
}

func (s *Async) mainLoop() {
	for {
		select {
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
			s.pending = map[uint32]chan *marionette.Message{}
			return
		case x := <-s.Conn.ResultChan():
			s.dispatch(x)
		}
	}
}

func (s *Async) dispatch(msg *marionette.Message) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()

	ch, ok := s.pending[msg.Serial]
	if !ok {
		// should never happen
		// but if it does happen
		// all we can do is ㄇㄉㄈㄎ
		return
	}

	delete(s.pending, msg.Serial)
	ch <- msg
	close(ch)
}
