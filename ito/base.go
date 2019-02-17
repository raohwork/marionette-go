// Package ito contains supported commands to control marionette
//
// The name "ito" is a Japanese word means "string". In the Japanese comic
// "Karakuri Circus", a marionette is controlled by its controller using strings.
package ito

import (
	"encoding/json"

	marionette "github.com/raohwork/marionette-go"
)

// Ito defines a command
//
// Even not defined here, many commands implement a "Decode" method, which decodes
// marionette.Message to corresponding data types.
type Ito interface {
	// Below are called by client (Shirogane)
	Command() string
	Param() interface{}
	Validate() bool
}

func recode(msg *marionette.Message, resp interface{}) (err error) {
	buf, _ := json.Marshal(msg.Data)
	return json.Unmarshal(buf, resp)
}

type nonObjResp struct {
	Value interface{} `json:"value"`
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

type noParam struct{}

func (c noParam) Param() (ret interface{}) {
	return
}

func (c noParam) Validate() (ok bool) {
	return true
}

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
