package ito

// Ito defines a command
//
// Ito is "string" in japanese. I use it here to denote the strings controlling
// marionette.
type Ito interface {
	// Below are called by client (Shirogane)
	Command() string
	Param() map[string]interface{}
	Validate() bool
}

type parameter map[string]interface{}

func (p parameter) SetS(key, data string) {
	if data != "" {
		p[key] = data
	}
}

func (p parameter) SetI(key string, data int) {
	if data != 0 {
		p[key] = data
	}
}

func (p parameter) SetB(key string, data bool) {
	if data {
		p[key] = data
	}
}

func (p parameter) SetP(key string, data interface{}) {
	if data != nil {
		p[key] = data
	}
}
