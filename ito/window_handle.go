package ito

// GetWindowHandles represents "WebDriver:GetWindowHandles" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1345
type GetWindowHandles struct {
}

func (c *GetWindowHandles) Command() (ret string) {
	return "WebDriver:GetWindowHandles"
}

func (c *GetWindowHandles) Param() (ret map[string]interface{}) {
	return
}

func (c *GetWindowHandles) Validate() (ok bool) {
	return true
}

// GetWindowHandle represents "WebDriver:GetWindowHandle" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1325
type GetWindowHandle struct {
}

func (c *GetWindowHandle) Command() (ret string) {
	return "WebDriver:GetWindowHandle"
}

func (c *GetWindowHandle) Param() (ret map[string]interface{}) {
	return
}

func (c *GetWindowHandle) Validate() (ok bool) {
	return true
}

// GetChromeWindowHandles represents "WebDriver:GetChromeWindowHandles" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1388
type GetChromeWindowHandles struct {
}

func (c *GetChromeWindowHandles) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandles"
}

func (c *GetChromeWindowHandles) Param() (ret map[string]interface{}) {
	return
}

func (c *GetChromeWindowHandles) Validate() (ok bool) {
	return true
}

// GetChromeWindowHandle represents "WebDriver:GetChromeWindowHandle" command
//
// See https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1360
type GetChromeWindowHandle struct {
}

func (c *GetChromeWindowHandle) Command() (ret string) {
	return "WebDriver:GetChromeWindowHandle"
}

func (c *GetChromeWindowHandle) Param() (ret map[string]interface{}) {
	return
}

func (c *GetChromeWindowHandle) Validate() (ok bool) {
	return true
}
