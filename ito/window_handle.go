package ito

import marionette "github.com/raohwork/marionette-go"

// GetWindowHandles represents "WebDriver:GetWindowHandles" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1345
type GetWindowHandles struct {
	noParam
}

func (c *GetWindowHandles) Decode(msg *marionette.Message) (ids []string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &ids)
	return
}

func (c *GetWindowHandles) Command() (ret string) {
	return "WebDriver:GetWindowHandles"
}

// GetWindowHandle represents "WebDriver:GetWindowHandle" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1325
type GetWindowHandle struct {
	noParam
}

func (c *GetWindowHandle) Decode(msg *marionette.Message) (id string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	resp := nonObjResp{Value: &id}
	if err = recode(msg, &resp); err != nil {
		return
	}

	return
}

func (c *GetWindowHandle) Command() (ret string) {
	return "WebDriver:GetWindowHandle"
}

// GetChromeWindowHandles represents "WebDriver:GetChromeWindowHandles" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1388
type GetChromeWindowHandles struct {
	noParam
}

func (c *GetChromeWindowHandles) Decode(msg *marionette.Message) (ids []string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &ids)
	return
}

func (c *GetChromeWindowHandles) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandles"
}

// GetChromeWindowHandle represents "WebDriver:GetChromeWindowHandle" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1360
type GetChromeWindowHandle struct {
	noParam
}

func (c *GetChromeWindowHandle) Decode(msg *marionette.Message) (id string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	resp := nonObjResp{Value: &id}
	if err = recode(msg, &resp); err != nil {
		return
	}

	return
}

func (c *GetChromeWindowHandle) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandle"
}
