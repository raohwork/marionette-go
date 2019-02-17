package ito

import marionette "github.com/raohwork/marionette-go"

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

func (s *NewSession) Command() (ret string) {
	return "WebDriver:NewSession"
}

func (s *NewSession) Param() (ret map[string]interface{}) {
	ret = map[string]interface{}{}
	cap := parameter(map[string]interface{}{})

	if s.SessionID != "" {
		ret["sessionId"] = s.SessionID
	}

	cap.SetS("pageLoadStrategy", s.PageLoadStrategy)
	cap.SetB("acceptInsecureCerts", s.AcceptInsecureCerts)
	cap.SetP("timeouts", s.Timeouts)
	cap.SetP("proxy", s.Proxy)
	cap.SetB("moz:accessibilityChecks", s.AccessibilityChecks)
	cap.SetB("moz:useNonSpecCompliantPointerOrigin", s.SpecialPointerOrigin)
	cap.SetB("moz:webdriverClick", s.WebdriverClick)

	if len(cap) > 0 {
		ret["capabilities"] = cap
	}

	return
}

func (s *NewSession) Validate() (ok bool) {
	return true
}
