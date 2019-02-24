package mncmd

import (
	marionette "github.com/raohwork/marionette-go"
)

// NewSession represents "WebDriver:NewSession" command
//
// See GeckoDriver.prototype.newSession
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L587
type NewSession struct {
	PageLoadStrategy     string // can be none/eager/normal
	AcceptInsecureCerts  bool
	Timeouts             *marionette.Timeouts
	Proxy                *marionette.Proxy
	AccessibilityChecks  bool
	SpecialPointerOrigin bool
	WebdriverClick       bool
	SessionID            string
}

func (c *NewSession) Decode(msg *marionette.Message) (
	id string, cap *marionette.Capabilities, err error,
) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var resp struct {
		ID  string                   `json:"sessionId"`
		Cap *marionette.Capabilities `json:"capabilities"`
	}
	if err = recode(msg, &resp); err == nil {
		id = resp.ID
		cap = resp.Cap
	}

	return
}

func (c *NewSession) Command() (ret string) {
	return "WebDriver:NewSession"
}

func (c *NewSession) Param() (data interface{}) {
	ret := map[string]interface{}{}
	cap := parameter(map[string]interface{}{})

	if c.SessionID != "" {
		ret["sessionId"] = c.SessionID
	}

	cap.SetS("pageLoadStrategy", c.PageLoadStrategy)
	cap.SetB("acceptInsecureCerts", c.AcceptInsecureCerts)
	cap.SetP("timeouts", c.Timeouts)
	cap.SetP("proxy", c.Proxy)
	cap.SetB("moz:accessibilityChecks", c.AccessibilityChecks)
	cap.SetB("moz:useNonSpecCompliantPointerOrigin", c.SpecialPointerOrigin)
	cap.SetB("moz:webdriverClick", c.WebdriverClick)

	if len(cap) > 0 {
		ret["capabilities"] = cap
	}

	return ret
}

func (c *NewSession) Validate() (ok bool) {
	return true
}

// SetTimeouts defines "WebDriver:SetTimeouts" command
//
// See GeckoDriver.prototype.setTimeouts
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1834
type SetTimeouts struct {
	Timeouts *marionette.Timeouts
}

func (c *SetTimeouts) Command() (ret string) {
	return "WebDriver:SetTimeouts"
}

func (c *SetTimeouts) Param() (ret interface{}) {
	return c.Timeouts
}

func (c *SetTimeouts) Validate() (ok bool) {
	return c.Timeouts != nil
}

// GetTimeouts defines "WebDriver:GetTimeouts" command
//
// See GeckoDriver.prototype.getTimeouts
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1830
type GetTimeouts struct {
	noParam
}

func (c *GetTimeouts) Decode(msg *marionette.Message) (ret *marionette.Timeouts, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var t marionette.Timeouts
	if err = recode(msg, &t); err == nil {
		ret = &t
	}

	return
}

func (c *GetTimeouts) Command() (ret string) {
	return "WebDriver:GetTimeouts"
}

// GetCapabilities defines "WebDriver:GetCapabilities" command
//
// See GeckoDriver.prototype.getSessionCapabilities
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L807
type GetCapabilities struct {
	noParam
}

func (c *GetCapabilities) Decode(msg *marionette.Message) (ret *marionette.Capabilities, err error) {
	if msg.Error != nil {
		err = msg.Error
		return
	}

	var t struct {
		Cap *marionette.Capabilities `json:"capabilities"`
	}
	if err = recode(msg, &t); err == nil {
		ret = t.Cap
	}

	return
}

func (c *GetCapabilities) Command() (ret string) {
	return "WebDriver:GetCapabilities"
}
