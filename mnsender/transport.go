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
	"bytes"
	"encoding/json"
	"io"
	"strconv"
)

// handles marionette protocol format
type transport struct {
	conn io.ReadWriter
}

func (t *transport) Send(data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}

	msg := []byte(strconv.Itoa(len(buf)) + ":")
	msg = append(msg, buf...)

	_, err = t.conn.Write(msg)
	return
}

func (t *transport) Receive() (ret []byte, err error) {
	l, err := t.receiveLength()
	if err != nil {
		return
	}

	ret = make([]byte, l)
	if _, err = io.ReadFull(t.conn, ret); err == io.EOF {
		err = nil
	}

	return
}

func (t *transport) receiveLength() (ret int, err error) {
	buf := &bytes.Buffer{}
	char := make([]byte, 1)

	for {
		_, err = t.conn.Read(char)
		if err != nil {
			return
		}

		if char[0] != ':' {
			buf.Write(char)
			continue
		}

		ret, err = strconv.Atoi(buf.String())
		return
	}
}
