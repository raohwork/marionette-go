package ito

import marionette "github.com/raohwork/marionette-go"

// NewSession represents "WebDriver:NewSession" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L587
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

func (c *NewSession) Command() (ret string) {
	return "WebDriver:NewSession"
}

func (c *NewSession) Param() (ret map[string]interface{}) {
	ret = map[string]interface{}{}
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

	return
}

func (c *NewSession) Validate() (ok bool) {
	return true
}
