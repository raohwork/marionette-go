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

// AddCookie defines "WebDriver:AddCookie" command
//
// See GeckoDriver.prototype.addCookie
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2611
type AddCookie struct {
	Cookie *marionette.Cookie `json:"cookie"`
}

func (c *AddCookie) Command() (ret string) {
	return "WebDriver:AddCookie"
}

func (c *AddCookie) Param() (ret interface{}) {
	return c
}

func (c *AddCookie) Validate() (ok bool) {
	return c.Cookie != nil
}

// DeleteCookie defines "WebDriver:DeleteCookie" command
//
// See GeckoDriver.prototype.deleteCookie
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2688
type DeleteCookie struct {
	Name string `json:"name"`
}

func (c *DeleteCookie) Command() (ret string) {
	return "WebDriver:DeleteCookie"
}

func (c *DeleteCookie) Param() (ret interface{}) {
	return c
}

func (c *DeleteCookie) Validate() (ok bool) {
	return c.Name != ""
}

// DeleteAllCookies defines "WebDriver:DeleteAllCookies" command
//
// See GeckoDriver.prototype.deleteAllCookies
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2667
type DeleteAllCookies struct {
	noParam
}

func (c *DeleteAllCookies) Command() (ret string) {
	return "WebDriver:DeleteAllCookies"
}

// GetCookies defines "WebDriver:GetCookies" command
//
// See GeckoDriver.prototype.getCookies
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2645
type GetCookies struct {
	noParam
}

func (c *GetCookies) Decode(msg *marionette.Message) (ret []*marionette.Cookie, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	err = recode(msg, &ret)
	return
}

func (c *GetCookies) Command() (ret string) {
	return "WebDriver:GetCookies"
}
