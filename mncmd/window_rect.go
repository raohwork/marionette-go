// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

package mncmd

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
// See GeckoDriver.prototype.getWindowRect
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1399
type GetWindowRect struct {
	returnRect
	noParam
}

func (c *GetWindowRect) Command() (ret string) {
	return "WebDriver:GetWindowRect"
}

// SetWindowRect defines "WebDriver:SetWindowRect" command
//
// See GeckoDriver.prototype.setWindowRect
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1422
type SetWindowRect struct {
	returnRect
	Rect marionette.Rect
}

func (c *SetWindowRect) Command() (ret string) {
	return "WebDriver:SetWindowRect"
}

func (c *SetWindowRect) Param() (ret interface{}) {
	return &c.Rect
}

func (c *SetWindowRect) Validate() (ok bool) {
	return true
}

// MinimizeWindow defines "WebDriver:MinimizeWindow" command
//
// See GeckoDriver.prototype.minimizeWindow
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3017
type MinimizeWindow GetWindowRect

func (c *MinimizeWindow) Command() (ret string) {
	return "WebDriver:MinimizeWindow"
}

// MaximizeWindow defines "WebDriver:MaximizeWindow" command
//
// See GeckoDriver.prototype.maximizeWindow
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3066
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
