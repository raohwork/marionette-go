package ito

import marionette "github.com/raohwork/marionette-go"

// Perform defines "WebDriver:PerformActions" command
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
