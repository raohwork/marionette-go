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

// ErrConnection denotes some network related problem occured
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
