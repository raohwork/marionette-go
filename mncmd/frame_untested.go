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
