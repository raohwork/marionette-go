// Package mncmd contains supported commands to control marionette
//
// All commands has related entry in
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js
//
// You may refer to the link above for further info of how the command works.
package mncmd

import (
	"encoding/json"

	marionette "github.com/raohwork/marionette-go"
)

// Command defines a command
//
// Even not defined here, many commands implement a "Decode" method, which decodes
// marionette.Message to corresponding data types.
type Command interface {
	Command() string
	Param() interface{}
	Validate() bool
}

// simple function to save some time
func recode(msg *marionette.Message, resp interface{}) (err error) {
	buf, _ := json.Marshal(msg.Data)
	return json.Unmarshal(buf, resp)
}

// non object return value
//
// Marionette server returns object values (object, array) as-is, but wraps non-obj
// values in fake object like {"value": "return value"}.
type nonObjResp struct {
	Value interface{} `json:"value"`
}

// simple helper to deal with complex parameter
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

func (p parameter) SetNotB(key string, data bool) {
	if !data {
		p[key] = !data
	}
}

func (p parameter) SetP(key string, data interface{}) {
	if data != nil {
		p[key] = data
	}
}

// mixin for commands which needs no parameter
type noParam struct{}

func (c noParam) Param() (ret interface{}) {
	return
}

func (c noParam) Validate() (ok bool) {
	return true
}

// mixin for commands which returns []string
type returnStrArr struct {
}

func (c returnStrArr) Decode(msg *marionette.Message) (ret []string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &ret)
	return
}

// mixin for commands which returns bool
type returnBool struct {
}

func (c returnBool) Decode(msg *marionette.Message) (ret bool, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	resp := nonObjResp{Value: &ret}
	if err = recode(msg, &resp); err != nil {
		return
	}

	return
}

// mixin for commands which returns string
type returnStr struct {
}

func (c returnStr) Decode(msg *marionette.Message) (ret string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	resp := nonObjResp{Value: &ret}
	if err = recode(msg, &resp); err != nil {
		return
	}

	return
}

// mixin for commands which returns WebElement
type returnElem struct {
}

func (c returnElem) Decode(msg *marionette.Message) (el *marionette.WebElement, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	arr := map[string]string{}
	resp := nonObjResp{Value: &arr}
	if err = recode(msg, &resp); err != nil {
		return
	}

	for k, v := range arr {
		el = &marionette.WebElement{
			Type: k,
			UUID: v,
		}
		break
	}

	return
}

// mixin for commands which returns []WebElement
type returnElems struct {
}

func (c returnElems) Decode(msg *marionette.Message) (el []*marionette.WebElement, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	arr := []map[string]string{}
	if err = recode(msg, &arr); err != nil {
		return
	}

	dec := func(m map[string]string) {
		for k, v := range m {
			el = append(el, &marionette.WebElement{
				Type: k,
				UUID: v,
			})
			break
		}
	}

	el = make([]*marionette.WebElement, 0, len(arr))
	for _, m := range arr {
		dec(m)
	}

	return
}

// mixin for commands which returns string/bool/number
type returnMixed struct {
}

func (c returnMixed) Decode(msg *marionette.Message) (ret interface{}, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	resp := nonObjResp{Value: &ret}
	if err = recode(msg, &resp); err != nil {
		return
	}

	return
}
