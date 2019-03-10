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

// MozGetContext defines "Marionette:GetContext" command
//
// See GeckoDriver.prototype.getContext
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L844
type MozGetContext struct {
	noParam
	returnStr
}

func (c *MozGetContext) Command() (ret string) {
	return "Marionette:GetContext"
}

// MozSetContext defines "Marionette:SetContext" command
//
// See GeckoDriver.prototype.setContext
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L822
type MozSetContext struct {
	Context string `json:"value"`
}

func (c *MozSetContext) Command() (ret string) {
	return "Marionette:SetContext"
}

func (c *MozSetContext) Param() (ret interface{}) {
	return c
}

func (c *MozSetContext) Validate() (ok bool) {
	return c.Context == marionette.ChromeContext ||
		c.Context == marionette.ContentContext
}

// MozGetScreenOrientation defines "Marionette:GetScreenOrientation" command
//
// See GeckoDriver.prototype.getScreenOrientation
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2970
type MozGetScreenOrientation struct {
	returnStr
}

func (c *MozGetScreenOrientation) Command() (ret string) {
	return "Marionette:GetScreenOrientation"
}

func (c *MozGetScreenOrientation) Param() (ret interface{}) {
	return
}

func (c *MozGetScreenOrientation) Validate() (ok bool) {
	return true
}

// MozSetScreenOrientation defines "Marionette:SetScreenOrientation" command
//
// See GeckoDriver.prototype.setScreenOrientation
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2984
type MozSetScreenOrientation struct {
	Value string `json:"orientation"`
}

func (c *MozSetScreenOrientation) Command() (ret string) {
	return "Marionette:SetScreenOrientation"
}

func (c *MozSetScreenOrientation) Param() (ret interface{}) {
	return c
}

func (c *MozSetScreenOrientation) Validate() (ok bool) {
	return c.Value != ""
}

// MozGetWindowType defines "Marionette:GetWindowType" command
//
// See GeckoDriver.prototype.getWindowType
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1138
type MozGetWindowType struct {
	returnStr
}

func (c *MozGetWindowType) Command() (ret string) {
	return "Marionette:GetWindowType"
}

func (c *MozGetWindowType) Param() (ret interface{}) {
	return
}

func (c *MozGetWindowType) Validate() (ok bool) {
	return true
}

// MozAcceptConnections defines "Marionette:AcceptConnections" command
//
// See GeckoDriver.prototype.acceptConnections
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3291
type MozAcceptConnections struct {
	Accept bool `json:"value"`
}

func (c *MozAcceptConnections) Command() (ret string) {
	return "Marionette:AcceptConnections"
}

func (c *MozAcceptConnections) Param() (ret interface{}) {
	return c
}

func (c *MozAcceptConnections) Validate() (ok bool) {
	return true
}

// MozQuit defines "Marionette:Quit" command
//
// See GeckoDriver.prototype.quit
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3312
type MozQuit struct {
	Flags []string `json:"flags,omitempty"`
	returnStr
}

func (c *MozQuit) Command() (ret string) {
	return "Marionette:Quit"
}

func (c *MozQuit) Param() (ret interface{}) {
	return c
}

func (c *MozQuit) Validate() (ok bool) {
	return true
}

// MozInstallAddon defines "Addon:Install" command
//
// See GeckoDriver.prototype.installAddon
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3381
type MozInstallAddon struct {
	Path      string `json:"path"`
	Temporary bool   `json:"temp,omitempty"`
	returnStr
}

func (c *MozInstallAddon) Command() (ret string) {
	return "Addon:Install"
}

func (c *MozInstallAddon) Param() (ret interface{}) {
	return c
}

func (c *MozInstallAddon) Validate() (ok bool) {
	return c.Path != ""
}

// MozUninstallAddon defines "Addon:Uninstall" command
//
// See GeckoDriver.prototype.uninstallAddon
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3394
type MozUninstallAddon struct {
	ID string `json:"id"`
}

func (c *MozUninstallAddon) Command() (ret string) {
	return "Addon:Uninstall"
}

func (c *MozUninstallAddon) Param() (ret interface{}) {
	return c
}

func (c *MozUninstallAddon) Validate() (ok bool) {
	return c.ID != ""
}
