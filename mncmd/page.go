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

import marionette "github.com/raohwork/marionette-go/v3"

// Back defines "WebDriver:Back" command
//
// See GeckoDriver.prototype.goBack
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1174
type Back struct {
	noParam
}

func (c *Back) Command() (ret string) {
	return "WebDriver:Back"
}

// Forward defines "WebDriver:Forward" command
//
// See GeckoDriver.prototype.goForward
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1216
type Forward struct {
	noParam
}

func (c *Forward) Command() (ret string) {
	return "WebDriver:Forward"
}

// Refresh defines "WebDriver:Refresh" command
//
// See GeckoDriver.prototype.refresh
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1259
type Refresh struct {
	noParam
}

func (c *Refresh) Command() (ret string) {
	return "WebDriver:Refresh"
}

// GetCurrentURL defines "WebDriver:GetCurrentURL" command
//
// See GeckoDriver.prototype.getCurrentUrl
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1098
type GetCurrentURL struct {
	noParam
	returnStr
}

func (c *GetCurrentURL) Command() (ret string) {
	return "WebDriver:GetCurrentURL"
}

// GetPageSource defines "WebDriver:GetPageSource" command
//
// See GeckoDriver.prototype.getCurrentUrl
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1145
type GetPageSource struct {
	noParam
	returnStr
}

func (c *GetPageSource) Command() (ret string) {
	return "WebDriver:GetPageSource"
}

// GetTitle defines "WebDriver:GetTitle" command
//
// See GeckoDriver.prototype.getCurrentUrl
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1120
type GetTitle struct {
	noParam
	returnStr
}

func (c *GetTitle) Command() (ret string) {
	return "WebDriver:GetTitle"
}

// Navigate defines "WebDriver:Navigate" command
//
// See GeckoDriver.prototype.get
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1040
type Navigate struct {
	URL string `json:"url"`
}

func (c *Navigate) Command() (ret string) {
	return "WebDriver:Navigate"
}

func (c *Navigate) Param() (ret interface{}) {
	return c
}

func (c *Navigate) Validate() (ok bool) {
	return c.URL != ""
}

// TakeScreenshot defines "WebDriver:TakeScreenshot" command
//
// See GeckoDriver.prototype.takeScreenshot
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2890
type TakeScreenshot struct {
	Element      *marionette.WebElement
	Highlights   []*marionette.WebElement
	ViewportOnly bool
	Hash         bool
	DontScrollTo bool
	returnStr
}

func (c *TakeScreenshot) Command() (ret string) {
	return "WebDriver:TakeScreenshot"
}

func (c *TakeScreenshot) Param() (ret interface{}) {
	x := parameter{}
	x.SetP("id", c.Element)
	if len(c.Highlights) > 0 {
		x.SetP("highlights", c.Highlights)
	}
	x.SetNotB("full", c.ViewportOnly)
	x.SetB("hash", c.Hash)
	x.SetNotB("scroll", c.DontScrollTo)

	return x
}

func (c *TakeScreenshot) Validate() (ok bool) {
	return true
}
