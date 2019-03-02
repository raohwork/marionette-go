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

// ReftestSetup defines "reftest:setup" command
//
// See GeckoDriver.prototype.setupReftest
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3508
type ReftestSetup struct {
	URLCount   map[string]uint `json:"urlCount"`
	Screenshot string          `json:"screenshot,omitempty"`
}

func (c *ReftestSetup) Command() (ret string) {
	return "reftest:setup"
}

func (c *ReftestSetup) Param() (ret interface{}) {
	return c
}

func (c *ReftestSetup) Validate() (ok bool) {
	return len(c.URLCount) > 0
}

// ReftestRun defines "reftest:run" command
//
// See GeckoDriver.prototype.runReftest
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3533
type ReftestRun struct {
	TestURL string                    `json:"test"`
	Ref     marionette.ReftestRefList `json:"references"`
	Expect  string                    `json:"expected"`
	Timeout int                       `json:"timeout"`
	Width   int                       `json:"width,omitempty"`
	Height  int                       `json:"height,omitempty"`
}

func (c *ReftestRun) Decode(msg *marionette.Message) (
	ret *marionette.ReftestResult, err error,
) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var data struct {
		Value *marionette.ReftestResult `json:"value"`
	}
	if err = recode(msg, &data); err == nil {
		ret = data.Value
	}

	return
}

func (c *ReftestRun) Command() (ret string) {
	return "reftest:run"
}

func (c *ReftestRun) Param() (ret interface{}) {
	return c
}

func (c *ReftestRun) Validate() (ok bool) {
	return c.TestURL != "" &&
		c.Ref != nil &&
		c.Expect != "" &&
		c.Timeout > 0 &&
		c.Width >= 0 &&
		c.Height >= 0
}

// ReftestTeardown defines "reftest:teardown" command
//
// See GeckoDriver.prototype.teardownReftest
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L3550
type ReftestTeardown struct {
	noParam
}

func (c *ReftestTeardown) Command() (ret string) {
	return "reftest:teardown"
}
