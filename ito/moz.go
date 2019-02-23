package ito

import marionette "github.com/raohwork/marionette-go"

// MozGetContext defines "Marionette:GetContext" command
//
// See GeckoDriver.prototype.getContext
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L844
type MozGetContext struct {
	noParam
	returnStr
}

func (c *MozGetContext) Command() (ret string) {
	return "Marionette:GetContext"
}

// MozSetContext defines "Marionette:SetContext" command
//
// See GeckoDriver.prototype.setContext
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L822
type MozSetContext struct {
	Context string `json:"value"`
	returnStr
}

func (c *MozSetContext) Command() (ret string) {
	return "Marionette:SetContext"
}

func (c *MozSetContext) Param() (ret interface{}) {
	return c
}

func (c *MozSetContext) Validate() (ok bool) {
	return c.Context == marionette.ChromeContext ||
		c.Context == marionette.ContentContext
}

// MozGetScreenOrientation defines "Marionette:GetScreenOrientation" command
//
// See GeckoDriver.prototype.getScreenOrientation
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2970
type MozGetScreenOrientation struct {
	returnStr
}

func (c *MozGetScreenOrientation) Command() (ret string) {
	return "Marionette:GetScreenOrientation"
}

func (c *MozGetScreenOrientation) Param() (ret interface{}) {
	return
}

func (c *MozGetScreenOrientation) Validate() (ok bool) {
	return true
}

// MozSetScreenOrientation defines "Marionette:SetScreenOrientation" command
//
// See GeckoDriver.prototype.setScreenOrientation
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2984
type MozSetScreenOrientation struct {
	Value string `json:"orientation"`
}

func (c *MozSetScreenOrientation) Command() (ret string) {
	return "Marionette:SetScreenOrientation"
}

func (c *MozSetScreenOrientation) Param() (ret interface{}) {
	return c
}

func (c *MozSetScreenOrientation) Validate() (ok bool) {
	return c.Value != ""
}

// MozGetWindowType defines "Marionette:GetWindowType" command
//
// See GeckoDriver.prototype.getWindowType
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1138
type MozGetWindowType struct {
	returnStr
}

func (c *MozGetWindowType) Command() (ret string) {
	return "Marionette:GetWindowType"
}

func (c *MozGetWindowType) Param() (ret interface{}) {
	return
}

func (c *MozGetWindowType) Validate() (ok bool) {
	return true
}
