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

// PerformActions defines "WebDriver:PerformActions" command
//
// I'm not link to source here as the code in driver.js heavily depends on several
// source files.
type PerformActions struct {
	Actions marionette.ActionChain `json:"actions"`
}

func (c *PerformActions) Command() (ret string) {
	return "WebDriver:PerformActions"
}

func (c *PerformActions) Param() (ret interface{}) {
	return c
}

func (c *PerformActions) Validate() (ok bool) {
	return len(c.Actions) > 0
}

// ReleaseActions defines "WebDriver:ReleaseActions" command
//
// I'm not link to source here as the code driver.js heavily depends on several
// source files.
type ReleaseActions struct {
	noParam
}

func (c *ReleaseActions) Command() (ret string) {
	return "WebDriver:ReleaseActions"
}
