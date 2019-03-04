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
	"io"
	"net"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mncmd"
)

// Sender abstracts an async client which supports blocking and non-blocking call
type Sender interface {
	Start() (err error)
	Close()
	Wait()
	Sync(cmd mncmd.Command) (msg *marionette.Message, err error)
	Async(cmd mncmd.Command) (ch chan *marionette.Message, err error)
}

// NewTCPSender creates a Sender with default tcp options
func NewTCPSender(addr string, bufSize int) (ret Sender, err error) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err == nil {
		ret = NewSender(conn, bufSize)
	}

	return
}

// NewSender creates a sender with pre-made connection
func NewSender(c io.ReadWriteCloser, bufSize int) (ret Sender) {
	return &mixed{
		tcpConn: c,
		bufSize: bufSize,
	}
}

// mixed is an asynchronous client supports both blocking and non-blocking call
type mixed struct {
	tcpConn io.ReadWriteCloser
	bufSize int

	client *Async
}

// Start initializes the protocol and start the mainloop in background
func (s *mixed) Start() (err error) {
	bufSize := s.bufSize
	if bufSize < 0 {
		bufSize = 0
	}

	tcp := s.tcpConn
	if tcp == nil {
		return errors.New("mnsender.mixed: empty connection")
	}
	conn, err := NewConn(tcp, uint(bufSize))
	if err != nil {
		return
	}

	s.client = &Async{Conn: conn}
	s.client.Start()

	return
}

// Wait waits until main loop stopped and disconnected
func (s *mixed) Wait() {
	s.client.Wait()
	s.client.Conn.Wait()
}

// Close exits the mainloop and release all related resources
func (s *mixed) Close() {
	s.client.Stop()
	s.client.Conn.Close()
	s.client.Conn.Cleanup()
}

// Sync send command synchronously (block until response actually)
func (s *mixed) Sync(cmd mncmd.Command) (msg *marionette.Message, err error) {
	msgch, err := s.client.Send(cmd)
	if err != nil {
		return
	}

	msg = <-msgch
	err = msg.Error
	return
}

// Async send command asynchronously
func (s *mixed) Async(cmd mncmd.Command) (ch chan *marionette.Message, err error) {
	return s.client.Send(cmd)
}
