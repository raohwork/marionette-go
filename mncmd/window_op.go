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

// NewWindow defines "WebDriver:NewWindow" command
//
// See GeckoDriver.prototype.newWindow
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2712
type NewWindow struct {
	Type  string `json:"type,omitempty"` // can be tab (default) or window
	Focus bool   `json:"focus,omitempty"`
}

func (c *NewWindow) Decode(msg *marionette.Message) (id, typ string, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var resp struct {
		ID   string `json:"handle"`
		Type string `json:"type"`
	}
	if err = recode(msg, &resp); err == nil {
		id = resp.ID
		typ = resp.Type
	}

	return
}

func (c *NewWindow) Command() (ret string) {
	return "WebDriver:NewWindow"
}

func (c *NewWindow) Param() (ret interface{}) {
	return c
}

func (c *NewWindow) Validate() (ok bool) {
	return true
}

// SwitchToWindow defines "WebDriver:SwitchToWindow" command
//
// See GeckoDriver.prototype.switchToWindow
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1493
type SwitchToWindow struct {
	Name    string // required
	NoFocus bool
}

func (c *SwitchToWindow) Command() (ret string) {
	return "WebDriver:SwitchToWindow"
}

func (c *SwitchToWindow) Param() (ret interface{}) {
	return map[string]interface{}{
		"name":  c.Name,
		"focus": !c.NoFocus,
	}
}

func (c *SwitchToWindow) Validate() (ok bool) {
	return c.Name != ""
}

// CloseWindow defines "WebDriver:CloseWindow" command
//
// See GeckoDriver.prototype.close
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2772
type CloseWindow struct {
	noParam
	returnStrArr
}

func (c *CloseWindow) Command() (ret string) {
	return "WebDriver:CloseWindow"
}

// CloseChromeWindow defines "WebDriver:CloseChromeWindow" command
//
// See GeckoDriver.prototype.closeChromeWindow
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2816
type CloseChromeWindow struct {
	noParam
	returnStrArr
}

func (c *CloseChromeWindow) Command() (ret string) {
	return "WebDriver:CloseChromeWindow"
}
