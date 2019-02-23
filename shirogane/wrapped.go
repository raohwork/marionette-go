package shirogane

import (
	"encoding/base64"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
)

// Ashihana is a client which wraps supported commands as method
//
// The name comes from Japnese comic "Karakuri circus", which is the last name of
// main characters in Kuroga.
type Ashihana struct {
	Kuroga
}

func (s *Ashihana) runSync(cmd ito.Ito) (err error) {
	msg, err := s.Sync(cmd)
	if err == nil {
		err = msg.Error
	}
	return
}

// AcceptAlert presses the "OK" button of the modal dialog
func (s *Ashihana) AcceptAlert() (err error) {
	cmd := &ito.AcceptAlert{}
	return s.runSync(cmd)
}

// AddCookie adds a cookie to the document
func (s *Ashihana) AddCookie(cookie *marionette.Cookie) (err error) {
	cmd := &ito.AddCookie{Cookie: cookie}
	return s.runSync(cmd)
}

// Back presses the "Back" button on the browser toolbar
func (s *Ashihana) Back() (err error) {
	cmd := &ito.Back{}
	return s.runSync(cmd)
}

// CloseChromeWindow closes current active chrome window
//
// It returns a list of currently opened chrome window
func (s *Ashihana) CloseChromeWindow() (handles []string, err error) {
	cmd := &ito.CloseChromeWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

// CloseChromeWindow closes current active window/tab
//
// It returns a list of currently opened window/tab
func (s *Ashihana) CloseWindow() (handles []string, err error) {
	cmd := &ito.CloseWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

// DeleteAllCookies deletes all cookie of the document
func (s *Ashihana) DeleteAllCookies() (err error) {
	cmd := &ito.DeleteAllCookies{}
	_, err = s.Sync(cmd)
	return
}

// DeleteCookie deletes specified cookie
func (s *Ashihana) DeleteCookie(name string) (err error) {
	cmd := &ito.DeleteCookie{Name: name}
	return s.runSync(cmd)
}

// DismissAlert presses "close" button of the modal dialog
func (s *Ashihana) DismissAlert() (err error) {
	cmd := &ito.DismissAlert{}
	return s.runSync(cmd)
}

// ElementClear clears the text of the element
func (s *Ashihana) ElementClear(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClear{Element: el}
	return s.runSync(cmd)
}

// ElementClick clicks the element
func (s *Ashihana) ElementClick(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClick{Element: el}
	return s.runSync(cmd)
}

// ElementSendKeys sends keystrokes to the element
func (s *Ashihana) ElementSendKeys(el *marionette.WebElement, text string) (err error) {
	cmd := &ito.ElementSendKeys{Element: el, Text: text}
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
func (s *Ashihana) ExecuteAsyncScript(script string, args ...interface{}) (
	ch chan ScriptResult, err error,
) {
	cmd := &ito.ExecuteAsyncScript{
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
func (s *Ashihana) ExecuteAsyncScriptIn(
	sandbox, script string, args ...interface{},
) (
	ch chan ScriptResult, err error,
) {
	cmd := &ito.ExecuteAsyncScript{
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
func (s *Ashihana) ExecuteScript(
	script string, data interface{}, args ...interface{},
) (err error) {
	cmd := &ito.ExecuteScript{
		Script: script,
		Args:   args,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg, data)
}

// ExecuteScriptIn executes the script in specified sandbox
//
// The sandbox is cached on window object for later use
func (s *Ashihana) ExecuteScriptIn(
	sandbox, script string, data interface{}, args ...interface{},
) (err error) {
	cmd := &ito.ExecuteScript{
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
func (s *Ashihana) FindElementAsync(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ch chan ElementResult, err error) {
	cmd := &ito.FindElement{
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
func (s *Ashihana) FindElement(
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
func (s *Ashihana) FindElementsAsync(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ch chan ElementResults, err error) {
	cmd := &ito.FindElements{
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
func (s *Ashihana) FindElements(
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
func (s *Ashihana) Forward() (err error) {
	cmd := &ito.Forward{}
	return s.runSync(cmd)
}

// FullscreenWindow switches active window to fullscreen mode
func (s *Ashihana) FullscreenWindow() (err error) {
	cmd := &ito.FullscreenWindow{}
	return s.runSync(cmd)
}

// GetActiveElement retrieves active element
func (s *Ashihana) GetActiveElement() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveElement{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetActiveFrame retrieves active frame
func (s *Ashihana) GetActiveFrame() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveFrame{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetAlertText retrieves the text label of the modal dialog
func (s *Ashihana) GetAlertText() (ret string, err error) {
	cmd := &ito.GetAlertText{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCapabilities retrieves browser capabilities
func (s *Ashihana) GetCapabilities() (ret *marionette.Capabilities, err error) {
	cmd := &ito.GetCapabilities{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetChromeWindowHandle retrieves current active chrome window handler
func (s *Ashihana) GetChromeWindowHandle() (ret string, err error) {
	cmd := &ito.GetChromeWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetChromeWindowHandles retrieves all opened chrome window handlers
func (s *Ashihana) GetChromeWindowHandles() (ret []string, err error) {
	cmd := &ito.GetChromeWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCookies retrieves all cookies of the document
func (s *Ashihana) GetCookies() (ret []*marionette.Cookie, err error) {
	cmd := &ito.GetCookies{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetCurrentURL retrieves current url
func (s *Ashihana) GetCurrentURL() (ret string, err error) {
	cmd := &ito.GetCurrentURL{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementAttribute retrieves specified attribute of the element
func (s *Ashihana) GetElementAttribute(
	el *marionette.WebElement, key string,
) (ret string, err error) {
	cmd := &ito.GetElementAttribute{
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
func (s *Ashihana) GetElementCSSValue(
	el *marionette.WebElement, key string,
) (ret string, err error) {
	cmd := &ito.GetElementCSSValue{
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
func (s *Ashihana) GetElementProperty(
	el *marionette.WebElement, key string,
) (ret interface{}, err error) {
	cmd := &ito.GetElementProperty{
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
func (s *Ashihana) GetElementRect(el *marionette.WebElement) (ret marionette.Rect, err error) {
	cmd := &ito.GetElementRect{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementTagName retrieves tag name of the element (like "div")
func (s *Ashihana) GetElementTagName(el *marionette.WebElement) (ret string, err error) {
	cmd := &ito.GetElementTagName{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetElementText retrieves text of the element
func (s *Ashihana) GetElementText(el *marionette.WebElement) (ret string, err error) {
	cmd := &ito.GetElementText{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetPageSource retrieves page source of current document
func (s *Ashihana) GetPageSource() (ret string, err error) {
	cmd := &ito.GetPageSource{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetTimeouts retrieves timeout settings of marionette server
func (s *Ashihana) GetTimeouts() (ret *marionette.Timeouts, err error) {
	cmd := &ito.GetTimeouts{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetTitle retrieves the text of title bar of current window/tab
func (s *Ashihana) GetTitle() (ret string, err error) {
	cmd := &ito.GetTitle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowHandle retrieves the handler of current window
func (s *Ashihana) GetWindowHandle() (ret string, err error) {
	cmd := &ito.GetWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowHandles retrieves the handlers of all opened window
func (s *Ashihana) GetWindowHandles() (ret []string, err error) {
	cmd := &ito.GetWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// GetWindowRect retrieves bounding box of current window
func (s *Ashihana) GetWindowRect() (ret marionette.Rect, err error) {
	cmd := &ito.GetWindowRect{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementDisplayed checks if the element is displayed
func (s *Ashihana) IsElementDisplayed(el *marionette.WebElement) (ret bool, err error) {
	cmd := &ito.IsElementDisplayed{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementEnabled checks if the element is enabled
func (s *Ashihana) IsElementEnabled(el *marionette.WebElement) (ret bool, err error) {
	cmd := &ito.IsElementEnabled{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// IsElementSelected checks if the element is selected
func (s *Ashihana) IsElementSelected(el *marionette.WebElement) (ret bool, err error) {
	cmd := &ito.IsElementSelected{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

// MaximizeWindow maximizes current window
func (s *Ashihana) MaximizeWindow() (err error) {
	cmd := &ito.MaximizeWindow{}
	return s.runSync(cmd)
}

// MinimizeWindow minimizes current window
func (s *Ashihana) MinimizeWindow() (err error) {
	cmd := &ito.MinimizeWindow{}
	return s.runSync(cmd)
}

// Navigate navigates to the url
func (s *Ashihana) Navigate(url string) (err error) {
	cmd := &ito.Navigate{URL: url}
	return s.runSync(cmd)
}

// NavigateAsync runs Navigate command asynchronously
func (s *Ashihana) NavigateAsync(url string) (ch chan error) {
	cmd := &ito.Navigate{URL: url}
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

// NewSession creates a new webdriver session
func (s *Ashihana) NewSession() (id string, cap *marionette.Capabilities, err error) {
	cmd := &ito.NewSession{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// NewWindow opens a new window
func (s *Ashihana) NewWindow(typ string, focus bool) (id, winType string, err error) {
	cmd := &ito.NewWindow{Type: typ, Focus: focus}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// Refresh presses the "refresh" button on toolbar
func (s *Ashihana) Refresh() (err error) {
	cmd := &ito.Refresh{}
	return s.runSync(cmd)
}

// SendAlertText sends keystrokes to the input area of modal dialog
func (s *Ashihana) SendAlertText(text string) (err error) {
	cmd := &ito.SendAlertText{Text: text}
	return s.runSync(cmd)
}

// SetTimeouts sets timeout settings of marionette server
func (s *Ashihana) SetTimeouts(t *marionette.Timeouts) (err error) {
	cmd := &ito.SetTimeouts{Timeouts: t}
	return s.runSync(cmd)
}

// SetWindowRect resizes and moves current window to specified size/position
func (s *Ashihana) SetWindowRect(
	r marionette.Rect,
) (ret marionette.Rect, err error) {
	cmd := &ito.SetWindowRect{Rect: r}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// SwitchToFrame swtich active frame (frameset/iframe or main frame)
func (s *Ashihana) SwitchToFrame(
	el *marionette.WebElement, id interface{}, focus bool,
) (err error) {
	cmd := &ito.SwitchToFrame{
		Element: el,
		Focus:   focus,
	}
	if id != nil {
		cmd.ID = id
	}
	return s.runSync(cmd)
}

// SwitchToParentFrame switches to parent frame
func (s *Ashihana) SwitchToParentFrame() (err error) {
	cmd := &ito.SwitchToParentFrame{}
	return s.runSync(cmd)
}

// SwitchToWindow switches to specified window/tab and bring it to foreground
func (s *Ashihana) SwitchToWindow(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle}
	return s.runSync(cmd)
}

// SwitchToWindowBG switches to specified window/tab, but does not bring it up
func (s *Ashihana) SwitchToWindowBG(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle, NoFocus: true}
	return s.runSync(cmd)
}

// ScreenshotDocument takes a screenshot of whole document in base64-encoded png
func (s *Ashihana) ScreenshotDocument(
	highlights []*marionette.WebElement,
) (img string, err error) {
	cmd := &ito.TakeScreenshot{Highlights: highlights}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// ScreenshotDocumentBytes takes a screenshot of whole document in png
func (s *Ashihana) ScreenshotDocumentBytes(
	highlights []*marionette.WebElement,
) (img []byte, err error) {
	str, err := s.ScreenshotDocument(highlights)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// ScreenshotViewport takes a screenshot of whole viewport in base64-encoded png
func (s *Ashihana) ScreenshotViewport(
	highlights []*marionette.WebElement,
) (img string, err error) {
	cmd := &ito.TakeScreenshot{
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
func (s *Ashihana) ScreenshotViewportBytes(
	highlights []*marionette.WebElement,
) (img []byte, err error) {
	str, err := s.ScreenshotViewport(highlights)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// ScreenshotElement takes a screenshot of the element in base64-encoded png
func (s *Ashihana) ScreenshotElement(el *marionette.WebElement) (img string, err error) {
	cmd := &ito.TakeScreenshot{Element: el}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// ScreenshotElementBytes takes a screenshot of the element in png
func (s *Ashihana) ScreenshotElementBytes(el *marionette.WebElement) (img []byte, err error) {
	str, err := s.ScreenshotElement(el)
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

// PerformActionsAsync sends virtual input events to current window asynchronously
func (s *Ashihana) PerformActionsAsync(act marionette.ActionChain) (errCh chan error) {
	cmd := &ito.PerformActions{Actions: act}
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
func (s *Ashihana) PerformActions(act marionette.ActionChain) (err error) {
	cmd := &ito.PerformActions{Actions: act}
	return s.runSync(cmd)
}

// ReleaseActions releases all pressed/clicked virtual input devices
func (s *Ashihana) ReleaseActions() (err error) {
	cmd := &ito.ReleaseActions{}
	return s.runSync(cmd)
}

// MozGetContext retrieves current context (content or chrome)
func (s *Ashihana) MozGetContext() (ret string, err error) {
	cmd := &ito.MozGetContext{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozSetContext sets current context (content or chrome)
func (s *Ashihana) MozSetContext(context string) (ret string, err error) {
	cmd := &ito.MozSetContext{Context: context}
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
func (s *Ashihana) MozGetWindowType() (ret string, err error) {
	cmd := &ito.MozGetWindowType{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozGetScreenOrientation retrieves screen orientation (valid only in fennec)
func (s *Ashihana) MozGetScreenOrientation() (ret string, err error) {
	cmd := &ito.MozGetScreenOrientation{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

// MozSetScreenOrientation sets screen orientation (valid only in fennec)
func (s *Ashihana) MozSetScreenOrientation(v string) (err error) {
	cmd := &ito.MozSetScreenOrientation{Value: v}
	msg, err := s.Sync(cmd)
	if err == nil {
		err = msg.Error
	}
	return
}
