package mncmd

// AcceptAlert defines "WebDriver:AcceptAlert" command
//
// See GeckoDriver.prototype.acceptDialog
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3177
type AcceptAlert struct {
	noParam
}

func (c *AcceptAlert) Command() (ret string) {
	return "WebDriver:AcceptAlert"
}

// DismissAlert defines "WebDriver:DismissAlert" command
//
// See GeckoDriver.prototype.dismissDialog
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3159
type DismissAlert struct {
	noParam
}

func (c *DismissAlert) Command() (ret string) {
	return "WebDriver:DismissAlert"
}

// GetAlertText defines "WebDriver:GetAlertText" command
//
// See GeckoDriver.prototype.getTextFromDialog
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3195
type GetAlertText struct {
	noParam
	returnStr
}

func (c *GetAlertText) Command() (ret string) {
	return "WebDriver:GetAlertText"
}

// SendAlertText defines "WebDriver:SendAlertText" command
//
// See GeckoDriver.prototype.sendKeysToDialog
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3206
type SendAlertText struct {
	Text string `json:"text"`
}

func (c *SendAlertText) Command() (ret string) {
	return "WebDriver:SendAlertText"
}

func (c *SendAlertText) Param() (ret interface{}) {
	return c
}

func (c *SendAlertText) Validate() (ok bool) {
	return c.Text != ""
}
