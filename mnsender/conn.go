package mnsender

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	marionette "github.com/raohwork/marionette-go"
)

// Conn represents a cnnection to Marionette server
type Conn struct {
	conn   io.ReadWriteCloser
	serial uint32
	ctx    context.Context
	cancel context.CancelFunc
	ch     chan *marionette.Message
	errch  chan error

	transport transport
}

// NewConn creates a Conn instance with user initialized tcp connection
//
// The resultBufferSize is capacity of the result channel.
func NewConn(c io.ReadWriteCloser, resultBufferSize uint) (ret *Conn, err error) {
	ret = &Conn{
		conn:      c,
		serial:    1,
		ch:        make(chan *marionette.Message, resultBufferSize),
		errch:     make(chan error, 1),
		transport: transport{conn: c},
	}
	ret.ctx, ret.cancel = context.WithCancel(context.Background())

	// first packet will be system info
	_, err = ret.transport.Receive()
	if err != nil {
		ret = nil
		return
	}

	go ret.receiver()
	return
}

// Wait blocks until first transport error, or nil if connnection is closed without error
func (c *Conn) Wait() (err error) {
	if c == nil {
		return
	}
	return <-c.errch
}

// Close disconnect from marionette server
func (c *Conn) Close() {
	if c == nil {
		return
	}
	c.cancel()
}

// ResultChan retrieves the channel instance for reading results.
func (c *Conn) ResultChan() (ch chan *marionette.Message) {
	if c == nil {
		return nil
	}
	return c.ch
}

// Send sends a command message to the server
func (c *Conn) Send(cmd string, param interface{}) (id uint32, err error) {
	if c == nil {
		return 0, errors.New("connection has not initialized")
	}

	id = c.serial
	data := [4]interface{}{
		int(0), // type: command
		id,     // serial number
		cmd,    // command name
		param,  // parameters
	}

	if err = c.transport.Send(data); err != nil {
		return 0, &marionette.ErrConnection{
			When:   "send",
			Origin: err,
		}
	}
	c.serial++

	return
}

func (c *Conn) receiveMessage() (id uint32, e error, resp interface{}) {
	f := func(err error) (a uint32, b error, r interface{}) {
		e := &marionette.ErrResponseDecode{
			Err: err,
		}
		return id, e, resp
	}

	var eDriver marionette.ErrDriver
	var typ uint32
	arr := [4]interface{}{
		&typ,
		&id,
		&eDriver,
		&resp,
	}
	buf, err := c.transport.Receive()
	if err != nil {
		return 0, err, nil
	}
	if err = json.Unmarshal(buf, &arr); err != nil {
		return f(err)
	}

	if typ != 1 {
		return f(errors.New("invalid response type"))
	}

	if eDriver.Type != "" {
		e = &eDriver
	}

	return
}

func (c *Conn) shutdown() {
	close(c.ch)
	close(c.errch)
	c.conn.Close()
}

func (c *Conn) receiver() {
	defer c.shutdown()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if err := c.doReceive(); err != nil {
				return
			}
		}
	}
}

func (c *Conn) doReceive() (err error) {
	id, err, data := c.receiveMessage()
	if id == 0 {
		c.errch <- err
		return
	}

	c.ch <- &marionette.Message{
		Type:   1,
		Serial: id,
		Data:   data,
		Error:  err,
	}

	return nil
}
