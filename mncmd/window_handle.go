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

// GetWindowHandles represents "WebDriver:GetWindowHandles" command
//
// See GeckoDriver.prototype.getWindowHandles
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1345
type GetWindowHandles struct {
	noParam
	returnStrArr
}

func (c *GetWindowHandles) Command() (ret string) {
	return "WebDriver:GetWindowHandles"
}

// GetWindowHandle represents "WebDriver:GetWindowHandle" command
//
// See GeckoDriver.prototype.getWindowHandle
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1325
type GetWindowHandle struct {
	noParam
	returnStr
}

func (c *GetWindowHandle) Command() (ret string) {
	return "WebDriver:GetWindowHandle"
}

// GetChromeWindowHandles represents "WebDriver:GetChromeWindowHandles" command
//
// See GeckoDriver.prototype.getChromeWindowHandles
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1388
type GetChromeWindowHandles struct {
	noParam
	returnStrArr
}

func (c *GetChromeWindowHandles) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandles"
}

// GetChromeWindowHandle represents "WebDriver:GetChromeWindowHandle" command
//
// See GeckoDriver.prototype.getChromeWindowHandle
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1360
type GetChromeWindowHandle struct {
	noParam
	returnStr
}

func (c *GetChromeWindowHandle) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandle"
}
