// This file is part of marionette-go
//
// marionette-go is free software: you can redistribute it and/or modify it
// under the terms of the GNU Lesser General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with marionette-go. If not, see <https://www.gnu.org/licenses/>.

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
