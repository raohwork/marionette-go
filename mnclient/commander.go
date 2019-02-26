package mnclient

import (
	"encoding/base64"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mncmd"
	"github.com/raohwork/marionette-go/mnsender"
)

// Commander is a client mixin which wraps supported commands into methods
type Commander struct {
	mnsender.Sender
}

func (s *Commander) runSync(cmd mncmd.Command) (err error) {
	msg, err := s.Sync(cmd)
	if err == nil {
		err = msg.Error
	}
	return
}

// AcceptAlert presses the "OK" button of the modal dialog
func (s *Commander) AcceptAlert() (err error) {
	cmd := &mncmd.AcceptAlert{}
	return s.runSync(cmd)
}

// AddCookie adds a cookie to the document
func (s *Commander) AddCookie(cookie *marionette.Cookie) (err error) {
	cmd := &mncmd.AddCookie{Cookie: cookie}
	return s.runSync(cmd)
}

// Back presses the "Back" button on the browser toolbar
func (s *Commander) Back() (err error) {
	cmd := &mncmd.Back{}
	return s.runSync(cmd)
}

// CloseChromeWindow closes current active chrome window
//
// It returns a list of currently opened chrome window
func (s *Commander) CloseChromeWindow() (handles []string, err error) {
	cmd := &mncmd.CloseChromeWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

// CloseChromeWindow closes current active window/tab
//
// It returns a list of currently opened window/tab
func (s *Commander) CloseWindow() (handles []string, err error) {
	cmd := &mncmd.CloseWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

// DeleteAllCookies deletes all cookie of the document
func (s *Commander) DeleteAllCookies() (err error) {
	cmd := &mncmd.DeleteAllCookies{}
	return s.runSync(cmd)
}

// DeleteCookie deletes specified cookie
func (s *Commander) DeleteCookie(name string) (err error) {
	cmd := &mncmd.DeleteCookie{Name: name}
	return s.runSync(cmd)
}

// DismissAlert presses "close" button of the modal dialog
func (s *Commander) DismissAlert() (err error) {
	cmd := &mncmd.DismissAlert{}
	return s.runSync(cmd)
}

// ElementClear clears the text of the element
func (s *Commander) ElementClear(el *marionette.WebElement) (err error) {
	cmd := &mncmd.ElementClear{Element: el}
	return s.runSync(cmd)
}

// ElementClick clicks the element
func (s *Commander) ElementClick(el *marionette.WebElement) (err error) {
	cmd := &mncmd.ElementClick{Element: el}
	return s.runSync(cmd)
}

// ElementSendKeys sends keystrokes to the element
func (s *Commander) ElementSendKeys(el *marionette.WebElement, text string) (err error) {
	cmd := &mncmd.ElementSendKeys{Element: el, Text: text}
	return s.runSync(cmd)
}

// ScriptResult is the result returned from script
type ScriptResult struct {
	Result interface{}
	Err    error
}

// ExecuteAsyncScript executes the script in default, mutable sandbox
//
// It returns the value passed to callback. Callback is always the last argument.
func (s *Commander) ExecuteAsyncScript(script string, args ...interface{}) (
	ch chan ScriptResult, err error,
) {
	cmd := &mncmd.ExecuteAsyncScript{
		Script: script,
		Args:   args,
	}

	msgch, err := s.Async(cmd)
	if err != nil {
		return
	}

	ch = make(chan ScriptResult)
	go func() {
		defer close(ch)
		var data interface{}
		err := cmd.Decode(<-msgch, &data)
		ch <- ScriptResult{
			Result: data,
			Err:    err,
		}
	}()

	return
}

// ExecuteAsyncScriptIn executes the script in specified sandbox
//
// It returns the value passed to callback. Callback is always the last argument.
//
// The sandbox is cached on window object for later use
func (s *Commander) ExecuteAsyncScriptIn(
	sandbox, script string, args ...interface{},
) (
	ch chan ScriptResult, err error,
) {
	cmd := &mncmd.ExecuteAsyncScript{
		Script:       script,
		Args:         args,
		Sandbox:      sandbox,
		ReuseSandbox: true,
	}

	msgch, err := s.Async(cmd)
	if err != nil {
		return
	}

	ch = make(chan ScriptResult)
	go func() {
		defer close(ch)
		var data interface{}
		err := cmd.Decode(<-msgch, &data)
		ch <- ScriptResult{
			Result: data,
			Err:    err,
		}
	}()

	return
}

// ExecuteScript executes the script in default, mutable sandbox
func (s *Commander) ExecuteScript(
	script string, data interface{}, args ...interface{},
) (err error) {
	cmd := &mncmd.ExecuteScript{
		Script: script,
		Args:   args,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg, data)
}

// ExecuteScriptIn executes the script in specified sandbox
//
// The sandbox is cached on window object for later use
func (s *Commander) ExecuteScriptIn(
	sandbox, script string, data interface{}, args ...interface{},
) (err error) {
	cmd := &mncmd.ExecuteScript{
		Script:       script,
		Args:         args,
		Sandbox:      sandbox,
		ReuseSandbox: true,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg, data)
}

// ElementResult is the result returned from FindElement
type ElementResult struct {
	Result *marionette.WebElement
	Err    error
}

// FindElementAsync finds an element asynchronously
func (s *Commander) FindElementAsync(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ch chan ElementResult, err error) {
	cmd := &mncmd.FindElement{
		Using:       by,
		Value:       qstr,
		RootElement: root,
	}
	msgCh, err := s.Async(cmd)
	if err != nil {
		return
	}

	ch = make(chan ElementResult)
	go func() {
		defer close(ch)
		msg := <-msgCh
		el, err := cmd.Decode(msg)
		ch <- ElementResult{
			Result: el,
			Err:    err,
		}
	}()

	return
}

// FindElement finds an element
func (s *Commander) FindElement(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ret *marionette.WebElement, err error) {
	ch, err := s.FindElementAsync(by, qstr, root)
	if err != nil {
		return
	}
	res := <-ch
	return res.Result, res.Err
}

// ElementResults is the result returned from FindElements
type ElementResults struct {
	Result []*marionette.WebElement
	Err    error
}

// FindElements retrieves all matching elements asynchronously
func (s *Commander) FindElementsAsync(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ch chan ElementResults, err error) {
	cmd := &mncmd.FindElements{
		Using:       by,
		Value:       qstr,
		RootElement: root,
	}
	msgCh, err := s.Async(cmd)
	if err != nil {
		return
	}

	ch = make(chan ElementResults)
	go func() {
		defer close(ch)
		res, err := cmd.Decode(<-msgCh)
		ch <- ElementResults{
			Result: res,
			Err:    err,
		}
	}()

	return
}

// FindElemens retrieves all matching elements
func (s *Commander) FindElements(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ret []*marionette.WebElement, err error) {
	ch, err := s.FindElementsAsync(by, qstr, root)
	if err != nil {
		return
	}
	res := <-ch
	return res.Result, res.Err
}

// Forward presses the "forward" button on browser toolbar
func (s *Commander) Forward() (err error) {
	cmd := &mncmd.Forward{}
	return s.runSync(cmd)
}

// FullscreenWindow switches active window to fullscreen mode
func (s *Commander) FullscreenWindow() (err error) {
	cmd := &mncmd.FullscreenWindow{}
	return s.runSync(cmd)
}

// GetActiveElement retrieves active element
func (s *Commander) GetActiveElement() (ret *marionette.WebElement, err error) {
	cmd := &mncmd.GetActiveElement{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetActiveFrame retrieves active frame
func (s *Commander) GetActiveFrame() (ret *marionette.WebElement, err error) {
	cmd := &mncmd.GetActiveFrame{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetAlertText retrieves the text label of the modal dialog
func (s *Commander) GetAlertText() (ret string, err error) {
	cmd := &mncmd.GetAlertText{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCapabilities retrieves browser capabilities
func (s *Commander) GetCapabilities() (ret *marionette.Capabilities, err error) {
	cmd := &mncmd.GetCapabilities{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetChromeWindowHandle retrieves current active chrome window handler
func (s *Commander) GetChromeWindowHandle() (ret string, err error) {
	cmd := &mncmd.GetChromeWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetChromeWindowHandles retrieves all opened chrome window handlers
func (s *Commander) GetChromeWindowHandles() (ret []string, err error) {
	cmd := &mncmd.GetChromeWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCookies retrieves all cookies of the document
func (s *Commander) GetCookies() (ret []*marionette.Cookie, err error) {
	cmd := &mncmd.GetCookies{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCurrentURL retrieves current url
func (s *Commander) GetCurrentURL() (ret string, err error) {
	cmd := &mncmd.GetCurrentURL{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementAttribute retrieves specified attribute of the element
func (s *Commander) GetElementAttribute(
	el *marionette.WebElement, key string,
) (ret string, err error) {
	cmd := &mncmd.GetElementAttribute{
		Element: el,
		Name:    key,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementCSSValue retrieves specified css value of the element
//
// You should use css names in "key", not js variant
//
//    c.GetElementCSSValue(el, "text-align") // not "textAlign"
func (s *Commander) GetElementCSSValue(
	el *marionette.WebElement, key string,
) (ret string, err error) {
	cmd := &mncmd.GetElementCSSValue{
		Element: el,
		Prop:    key,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementProperty retrieves specified property of the element
func (s *Commander) GetElementProperty(
	el *marionette.WebElement, key string,
) (ret interface{}, err error) {
	cmd := &mncmd.GetElementProperty{
		Element: el,
		Name:    key,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementRect retrieves the bounding rect of the element
//
// The X(left) and Y(top) are computed aginst origin(top-left) of the document.
func (s *Commander) GetElementRect(el *marionette.WebElement) (ret marionette.Rect, err error) {
	cmd := &mncmd.GetElementRect{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementTagName retrieves tag name of the element (like "div")
func (s *Commander) GetElementTagName(el *marionette.WebElement) (ret string, err error) {
	cmd := &mncmd.GetElementTagName{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementText retrieves text of the element
func (s *Commander) GetElementText(el *marionette.WebElement) (ret string, err error) {
	cmd := &mncmd.GetElementText{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetPageSource retrieves page source of current document
func (s *Commander) GetPageSource() (ret string, err error) {
	cmd := &mncmd.GetPageSource{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetTimeouts retrieves timeout settings of marionette server
func (s *Commander) GetTimeouts() (ret *marionette.Timeouts, err error) {
	cmd := &mncmd.GetTimeouts{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetTitle retrieves the text of title bar of current window/tab
func (s *Commander) GetTitle() (ret string, err error) {
	cmd := &mncmd.GetTitle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowHandle retrieves the handler of current window
func (s *Commander) GetWindowHandle() (ret string, err error) {
	cmd := &mncmd.GetWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowHandles retrieves the handlers of all opened window
func (s *Commander) GetWindowHandles() (ret []string, err error) {
	cmd := &mncmd.GetWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowRect retrieves bounding box of current window
func (s *Commander) GetWindowRect() (ret marionette.Rect, err error) {
	cmd := &mncmd.GetWindowRect{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementDisplayed checks if the element is displayed
func (s *Commander) IsElementDisplayed(el *marionette.WebElement) (ret bool, err error) {
	cmd := &mncmd.IsElementDisplayed{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementEnabled checks if the element is enabled
func (s *Commander) IsElementEnabled(el *marionette.WebElement) (ret bool, err error) {
	cmd := &mncmd.IsElementEnabled{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementSelected checks if the element is selected
func (s *Commander) IsElementSelected(el *marionette.WebElement) (ret bool, err error) {
	cmd := &mncmd.IsElementSelected{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// MaximizeWindow maximizes current window
func (s *Commander) MaximizeWindow() (err error) {
	cmd := &mncmd.MaximizeWindow{}
	return s.runSync(cmd)
}

// MinimizeWindow minimizes current window
func (s *Commander) MinimizeWindow() (err error) {
	cmd := &mncmd.MinimizeWindow{}
	return s.runSync(cmd)
}

// Navigate navigates to the url
func (s *Commander) Navigate(url string) (err error) {
	cmd := &mncmd.Navigate{URL: url}
	return s.runSync(cmd)
}

// NavigateAsync runs Navigate command asynchronously
func (s *Commander) NavigateAsync(url string) (ch chan error) {
	cmd := &mncmd.Navigate{URL: url}
	ch = make(chan error, 1)
	c, err := s.Async(cmd)
	if err != nil {
		ch <- err
		close(ch)
		return
	}

	go func(c chan *marionette.Message, ch chan error) {
		msg := <-c
		if msg != nil && msg.Error != nil {
			ch <- msg.Error
		}
		close(ch)
	}(c, ch)

	return
}

// NewSession creates a new webdriver session with default options
func (s *Commander) NewSession() (id string, cap *marionette.Capabilities, err error) {
	cmd := &mncmd.NewSession{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// NewSessionWith creates a new webdriver session with specified options
//
// Valid "page" values
//
//    - none: no strategy, Navigate() will return immediately
//    - eager: return when enter "interactive" ready state (after DOMContentLoaded event)
//    - normal: return when enter "complete" ready state (after load event)
func (s *Commander) NewSessionWith(page string, insecureCert bool) (
	id string, cap *marionette.Capabilities, err error,
) {
	cmd := &mncmd.NewSession{
		PageLoadStrategy:    page,
		AcceptInsecureCerts: insecureCert,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// NewWindow opens a new window
func (s *Commander) NewWindow(typ string, focus bool) (id, winType string, err error) {
	cmd := &mncmd.NewWindow{Type: typ, Focus: focus}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// Refresh presses the "refresh" button on toolbar
func (s *Commander) Refresh() (err error) {
	cmd := &mncmd.Refresh{}
	return s.runSync(cmd)
}

// SendAlertText sends keystrokes to the input area of modal dialog
func (s *Commander) SendAlertText(text string) (err error) {
	cmd := &mncmd.SendAlertText{Text: text}
	return s.runSync(cmd)
}

// SetTimeouts sets timeout settings of marionette server
func (s *Commander) SetTimeouts(t *marionette.Timeouts) (err error) {
	cmd := &mncmd.SetTimeouts{Timeouts: t}
	return s.runSync(cmd)
}

// SetWindowRect resizes and moves current window to specified size/position
func (s *Commander) SetWindowRect(
	r marionette.Rect,
) (ret marionette.Rect, err error) {
	cmd := &mncmd.SetWindowRect{Rect: r}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// SwitchToFrame swtich active frame (frameset/iframe or main frame)
func (s *Commander) SwitchToFrame(
	el *marionette.WebElement, id interface{}, focus bool,
) (err error) {
	cmd := &mncmd.SwitchToFrame{
		Element: el,
		Focus:   focus,
	}
	if id != nil {
		cmd.ID = id
	}
	return s.runSync(cmd)
}

// SwitchToParentFrame switches to parent frame
func (s *Commander) SwitchToParentFrame() (err error) {
	cmd := &mncmd.SwitchToParentFrame{}
	return s.runSync(cmd)
}

// SwitchToWindow switches to specified window/tab and bring it to foreground
func (s *Commander) SwitchToWindow(handle string) (err error) {
	cmd := &mncmd.SwitchToWindow{Name: handle}
	return s.runSync(cmd)
}

// SwitchToWindowBG switches to specified window/tab, but does not bring it up
func (s *Commander) SwitchToWindowBG(handle string) (err error) {
	cmd := &mncmd.SwitchToWindow{Name: handle, NoFocus: true}
	return s.runSync(cmd)
}

// ScreenshotDocument takes a screenshot of whole document in base64-encoded png
func (s *Commander) ScreenshotDocument(
	highlights []*marionette.WebElement,
) (img string, err error) {
	cmd := &mncmd.TakeScreenshot{Highlights: highlights}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// ScreenshotDocumentBytes takes a screenshot of whole document in png
func (s *Commander) ScreenshotDocumentBytes(
	highlights []*marionette.WebElement,
) (img []byte, err error) {
	str, err := s.ScreenshotDocument(highlights)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// ScreenshotViewport takes a screenshot of whole viewport in base64-encoded png
func (s *Commander) ScreenshotViewport(
	highlights []*marionette.WebElement,
) (img string, err error) {
	cmd := &mncmd.TakeScreenshot{
		ViewportOnly: true,
		Highlights:   highlights,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// ScreenshotViewportBytes takes a screenshot of whole viewport in png
func (s *Commander) ScreenshotViewportBytes(
	highlights []*marionette.WebElement,
) (img []byte, err error) {
	str, err := s.ScreenshotViewport(highlights)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// ScreenshotElement takes a screenshot of the element in base64-encoded png
func (s *Commander) ScreenshotElement(el *marionette.WebElement) (img string, err error) {
	cmd := &mncmd.TakeScreenshot{Element: el}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// ScreenshotElementBytes takes a screenshot of the element in png
func (s *Commander) ScreenshotElementBytes(el *marionette.WebElement) (img []byte, err error) {
	str, err := s.ScreenshotElement(el)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// PerformActionsAsync sends virtual input events to current window asynchronously
func (s *Commander) PerformActionsAsync(act marionette.ActionChain) (errCh chan error) {
	cmd := &mncmd.PerformActions{Actions: act}
	errCh = make(chan error, 1)
	ch, err := s.Async(cmd)
	if err != nil {
		errCh <- err
		close(errCh)
		return
	}

	go func() {
		_ = <-ch
		errCh <- nil
		close(errCh)
	}()

	return
}

// PerformActions sends virtual input events to current window
func (s *Commander) PerformActions(act marionette.ActionChain) (err error) {
	cmd := &mncmd.PerformActions{Actions: act}
	return s.runSync(cmd)
}

// ReleaseActions releases all pressed/clicked virtual input devices
func (s *Commander) ReleaseActions() (err error) {
	cmd := &mncmd.ReleaseActions{}
	return s.runSync(cmd)
}

// MozGetContext retrieves current context (content or chrome)
func (s *Commander) MozGetContext() (ret string, err error) {
	cmd := &mncmd.MozGetContext{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozSetContext sets current context (content or chrome)
func (s *Commander) MozSetContext(context string) (ret string, err error) {
	cmd := &mncmd.MozSetContext{Context: context}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozGetWindowType retrieves application type
//
// Can be
//
//   - marionette.FirefoxWindow
//   - marionette.GeckoViewWindow
//   - marionette.ThunderbirdWindow
func (s *Commander) MozGetWindowType() (ret string, err error) {
	cmd := &mncmd.MozGetWindowType{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozGetScreenOrientation retrieves screen orientation (valid only in fennec)
func (s *Commander) MozGetScreenOrientation() (ret string, err error) {
	cmd := &mncmd.MozGetScreenOrientation{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozSetScreenOrientation sets screen orientation (valid only in fennec)
func (s *Commander) MozSetScreenOrientation(v string) (err error) {
	cmd := &mncmd.MozSetScreenOrientation{Value: v}
	msg, err := s.Sync(cmd)
	if err == nil {
		err = msg.Error
	}
	return
}

// MozAcceptConnections tells current server to enable/disable new connections
func (s *Commander) MozAcceptConnections(enable bool) (err error) {
	cmd := &mncmd.MozAcceptConnections{Accept: enable}
	return s.runSync(cmd)
}

// MozQuit tells current server to quit or restart
func (s *Commander) MozQuit(flags ...string) (ret string, err error) {
	cmd := &mncmd.MozQuit{Flags: flags}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozInstallAddon installs an addon to the server
func (s *Commander) MozInstallAddon(path string, temp bool) (id string, err error) {
	cmd := &mncmd.MozInstallAddon{Path: path, Temporary: temp}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozUninstallAddon uninstalls known addon on the server
func (s *Commander) MozUninstallAddon(id string) (err error) {
	cmd := &mncmd.MozUninstallAddon{ID: id}
	return s.runSync(cmd)
}
