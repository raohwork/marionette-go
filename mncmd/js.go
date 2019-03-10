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

// ExecuteScript defines "WebDriver:ExecuteScript" command
//
// Unlike other commands, the Decode() method of ExecuteScript accepts dest, which
// works as go idiom, json.Unmarshal().
//
// See GeckoDriver.prototype.executeScript
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L859
type ExecuteScript struct {
	Script       string
	Args         []interface{}
	Sandbox      string
	ReuseSandbox bool
	Filename     string
	Line         int
}

func (c *ExecuteScript) Decode(msg *marionette.Message, dest interface{}) (err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var resp nonObjResp
	if dest != nil {
		resp.Value = dest
	}
	return recode(msg, &resp)
}

func (c *ExecuteScript) Command() (ret string) {
	return "WebDriver:ExecuteScript"
}

func (c *ExecuteScript) Param() (ret interface{}) {
	x := parameter{}
	x["script"] = c.Script
	args := c.Args
	if args == nil {
		args = []interface{}{}
	}
	x["args"] = args
	x.SetS("sandbox", c.Sandbox)
	x.SetNotB("newSandbox", c.ReuseSandbox)
	x.SetS("filename", c.Filename)
	x.SetI("line", c.Line)

	return x
}

func (c *ExecuteScript) Validate() (ok bool) {
	return c.Script != ""
}

// ExecuteAsyncScript defines "WebDriver:ExecuteAsyncScript" command
//
// Unlike other commands, the Decode() method of ExecuteAsyncScript accepts dest,
// which works as go idiom, json.Unmarshal().
//
// See GeckoDriver.prototype.executeAsyncScript
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L914
type ExecuteAsyncScript struct {
	Script       string
	Args         []interface{}
	Sandbox      string
	ReuseSandbox bool
	Filename     string
	Line         int
}

func (c *ExecuteAsyncScript) Decode(msg *marionette.Message, dest interface{}) (err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var resp nonObjResp
	resp.Value = dest
	return recode(msg, &resp)
}

func (c *ExecuteAsyncScript) Command() (ret string) {
	return "WebDriver:ExecuteAsyncScript"
}

func (c *ExecuteAsyncScript) Param() (ret interface{}) {
	x := parameter{}
	x["script"] = c.Script
	args := c.Args
	if args == nil {
		args = []interface{}{}
	}
	x["args"] = args
	x.SetS("sandbox", c.Sandbox)
	x.SetNotB("newSandbox", c.ReuseSandbox)
	x.SetS("filename", c.Filename)
	x.SetI("line", c.Line)

	return x
}

func (c *ExecuteAsyncScript) Validate() (ok bool) {
	return c.Script != ""
}
