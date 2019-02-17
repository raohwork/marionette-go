package ito

import marionette "github.com/raohwork/marionette-go"

type returnRect struct{}

func (c returnRect) Decode(msg *marionette.Message) (rect marionette.Rect, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &rect)
	return
}

// GetWindowRect defines "WebDriver:GetWindowRect" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1399
type GetWindowRect struct {
	returnRect
	noParam
}

func (c *GetWindowRect) Command() (ret string) {
	return "WebDriver:GetWindowRect"
}

// SetWindowRect defines "WebDriver:SetWindowRect" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1422
type SetWindowRect struct {
	returnRect
	Rect marionette.Rect
}

func (c *SetWindowRect) Command() (ret string) {
	return "WebDriver:SetWindowRect"
}

func (c *SetWindowRect) Param() (ret map[string]interface{}) {
	ret = map[string]interface{}{
		"x":      c.Rect.X,
		"y":      c.Rect.Y,
		"width":  c.Rect.W,
		"height": c.Rect.H,
	}
	return
}

func (c *SetWindowRect) Validate() (ok bool) {
	return true
}

// MinimizeWindow defines "WebDriver:MinimizeWindow" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3017
type MinimizeWindow GetWindowRect

func (c *MinimizeWindow) Command() (ret string) {
	return "WebDriver:MinimizeWindow"
}

// MaximizeWindow defines "WebDriver:MaximizeWindow" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3066
type MaximizeWindow GetWindowRect

func (c *MaximizeWindow) Decode(msg *marionette.Message) (rect marionette.Rect, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &rect)
	return
}

func (c *MaximizeWindow) Command() (ret string) {
	return "WebDriver:MaximizeWindow"
}

// FullscreenWindow defines "WebDriver:FullscreenWindow" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3114
type FullscreenWindow GetWindowRect

func (c *FullscreenWindow) Decode(msg *marionette.Message) (rect marionette.Rect, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &rect)
	return
}

func (c *FullscreenWindow) Command() (ret string) {
	return "WebDriver:FullscreenWindow"
}
