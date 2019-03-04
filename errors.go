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

package marionette

import "fmt"

// ErrType denotes type of error returned from marionette server
type ErrType string

const (
	ErrElementClickIntercepted = ErrType("element click intercepted")
	ErrElementNotInteractable  = ErrType("element not interactable")
	ErrInsecureCertificate     = ErrType("insecure certificate")
	ErrInvalidArgument         = ErrType("invalid argument")
	ErrInvalidCookieDomain     = ErrType("invalid cookie domain")
	ErrInvalidElementState     = ErrType("invalid element state")
	ErrInvalidSelector         = ErrType("invalid selector")
	ErrInvalidSessionId        = ErrType("invalid session id")
	ErrJavascriptError         = ErrType("javascript error")
	ErrMoveTargetOutOfBounds   = ErrType("move target out of bounds")
	ErrNoSuchAlert             = ErrType("no such alert")
	ErrNoSuchCookie            = ErrType("no such cookie")
	ErrNoSuchElement           = ErrType("no such element")
	ErrNoSuchFrame             = ErrType("no such frame")
	ErrNoSuchWindow            = ErrType("no such window")
	ErrScriptTimeout           = ErrType("script timeout")
	ErrSessionNotCreated       = ErrType("session not created")
	ErrStaleElementReference   = ErrType("stale element reference")
	ErrTimeout                 = ErrType("timeout")
	ErrUnableToSetCookie       = ErrType("unable to set cookie")
	ErrUnableToCapturescreen   = ErrType("unable to capture screen")
	ErrUnexpecteDalertopen     = ErrType("unexpected alert open")
	ErrUnknownCommand          = ErrType("unknown command")
	ErrUnknownError            = ErrType("unknown error")
	ErrUnknownMethod           = ErrType("unknown method")
	ErrUnsupportedOperation    = ErrType("unsupported operation")
)

// ErrDriver is error returned from marionette server
type ErrDriver struct {
	Type       ErrType `json:"error"`
	Message    string  `json:"message"`
	StackTrace string  `json:"stacktrace"`
}

func (e *ErrDriver) Error() (ret string) {
	return "mario error: " + string(e.Type) + ": " + e.Message
}

func (e *ErrDriver) String() (ret string) {
	return fmt.Sprintf("%+v", map[string]interface{}{
		"type":       e.Type,
		"message":    e.Message,
		"stacktrace": e.StackTrace,
	})
}

// ErrResponseDecode denotes we are failed to decode incoming data as Message
type ErrResponseDecode struct {
	Err error
}

func (e *ErrResponseDecode) Error() (ret string) {
	return e.Err.Error()
}

func (e *ErrResponseDecode) String() (ret string) {
	return e.Err.Error()
}

// ErrConnection denotes some network related problem occurred
type ErrConnection struct {
	When   string
	Origin error
}

func (e *ErrConnection) Error() (ret string) {
	return e.Origin.Error()
}

func (e *ErrConnection) String() (ret string) {
	return "connection error when " + e.When + ": " + e.Origin.Error()
}
