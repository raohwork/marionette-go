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
