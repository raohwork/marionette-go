package marionette

import (
	"context"
	"encoding/json"
	"errors"
	"net"
)

const (
	ChromeContext  = "chrome"
	ContentContext = "content"
)

// Message represents messages to/from marionette server
type Message struct {
	Type   int
	Serial uint32
	Error  error
	Data   interface{}
}

// Conn represents a cnnection to Marionette server
type Conn struct {
	conn   *net.TCPConn
	serial uint32
	ctx    context.Context
	cancel context.CancelFunc
	ch     chan *Message
	errch  chan error

	transport transport
}

// ConnectTo creates a Conn instance to connect to remote marionette server
//
// It just creates a net.TCPConn with default parameter, tries to enable tcp
// keepalive and pass it to NewConn(tcpconn, 0).
func ConnectTo(addr string) (ret *Conn, err error) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}

	tcpconn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		return
	}

	tcpconn.SetKeepAlive(true)

	return NewConn(tcpconn, 0)
}

// NewConn creates a Conn instance with user initialized tcp connection
//
// The resultBufferSize is capacity of the result channel.
func NewConn(tcpconn *net.TCPConn, resultBufferSize uint) (ret *Conn, err error) {
	ret = &Conn{
		conn:      tcpconn,
		serial:    1,
		ch:        make(chan *Message, resultBufferSize),
		errch:     make(chan error, 1),
		transport: transport{conn: tcpconn},
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

// Wait blocks until first transport error
//
// You SHOULD close the connection after Wait() return.
func (c *Conn) Wait() (err error) {
	return <-c.errch
}

// Close disconnect from marionette server and release allocated resources
//
// Since it closes the result channel, calling it multiple times leads to panic.
func (c *Conn) Close() (err error) {
	if c == nil {
		return nil
	}
	c.cancel()
	close(c.ch)
	return c.conn.Close()
}

// ResultChan retrieves the channel instance for reading results.
func (c *Conn) ResultChan() (ch chan *Message) {
	if c == nil {
		return nil
	}
	return c.ch
}

// Send sends a command message to the server
func (c *Conn) Send(cmd string, param map[string]interface{}) (id uint32, err error) {
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
		return 0, &ErrConnection{
			When:   "send",
			Origin: err,
		}
	}
	c.serial++

	return
}

func (c *Conn) receiveMessage() (id uint32, e error, resp interface{}) {
	f := func(err error) (a uint32, b error, r interface{}) {
		e := &ErrResponseDecode{
			Err: err,
		}
		return id, e, resp
	}

	var eDriver ErrDriver
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

func (c *Conn) receiver() {
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

	c.ch <- &Message{
		Type:   1,
		Serial: id,
		Data:   data,
		Error:  err,
	}

	return nil
}
