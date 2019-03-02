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

// GetActiveFrame defines "WebDriver:GetActiveFrame" command
//
// See GeckoDriver.prototype.getActiveFrame
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1620
type GetActiveFrame struct {
	noParam
	returnElem
}

func (c *GetActiveFrame) Command() (ret string) {
	return "WebDriver:GetActiveFrame"
}

// SwitchToFrame defines "WebDriver:SwitchToFrame" command
//
// See GeckoDriver.prototype.switchToFrame
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1656
type SwitchToFrame struct {
	Element *marionette.WebElement `json:"element,omitempty"`
	ID      interface{}            `json:"id,omitempty"` // must be int/uint/string
	Focus   bool                   `json:"focus,omitempty"`
}

func (c *SwitchToFrame) Command() (ret string) {
	return "WebDriver:SwitchToFrame"
}

func (c *SwitchToFrame) Param() (ret interface{}) {
	return c
}

func (c *SwitchToFrame) Validate() (ok bool) {
	if c.Element == nil && c.ID == nil {
		return
	}

	if c.ID != nil {
		switch c.ID.(type) {
		case int:
		case uint:
		case string:
			// safe above
		default:
			return
		}
	}

	return true
}

// SwitchToParentFrame defines "WebDriver:SwitchToParentFrame" command
//
// See GeckoDriver.prototype.switchToParentFrame
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1640
type SwitchToParentFrame struct {
	noParam
}

func (c *SwitchToParentFrame) Command() (ret string) {
	return "WebDriver:SwitchToParentFrame"
}
