package shirogane

import (
	"encoding/base64"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
)

// Wrapped is another Mixed client which wraps supported commands as method
type Wrapped struct {
	Mixed
}

func (s *Wrapped) AcceptAlert() (err error) {
	cmd := &ito.AcceptAlert{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) AddCookie(cookie *marionette.Cookie) (err error) {
	cmd := &ito.AddCookie{Cookie: cookie}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) Back() (err error) {
	cmd := &ito.Back{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) CloseChromeWindow() (handles []string, err error) {
	cmd := &ito.CloseChromeWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Wrapped) CloseWindow() (handles []string, err error) {
	cmd := &ito.CloseWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Wrapped) DeleteAllCookies() (err error) {
	cmd := &ito.DeleteAllCookies{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) DeleteCookie(name string) (err error) {
	cmd := &ito.DeleteCookie{Name: name}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) DismissAlert() (err error) {
	cmd := &ito.DismissAlert{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) ElementClear(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClear{Element: el}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) ElementClick(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClick{Element: el}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) ElementSendKeys(el *marionette.WebElement, text string) (err error) {
	cmd := &ito.ElementSendKeys{Element: el, Text: text}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) ExecuteAsyncScript(script string, args ...interface{}) (
	ch chan interface{}, err error,
) {
	cmd := &ito.ExecuteAsyncScript{
		Script: script,
		Args:   args,
	}

	msgch, err := s.Async(cmd)
	if err != nil {
		return
	}

	ch = make(chan interface{})
	go func() {
		defer close(ch)
		var data interface{}
		err := cmd.Decode(<-msgch, &data)
		if err != nil {
			ch <- err
			return
		}
		ch <- data
	}()

	return
}

func (s *Wrapped) ExecuteAsyncScriptIn(
	sandbox, script string, args ...interface{},
) (
	ch chan interface{}, err error,
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

	ch = make(chan interface{})
	go func() {
		defer close(ch)
		var data interface{}
		err := cmd.Decode(<-msgch, &data)
		if err != nil {
			ch <- err
			return
		}
		ch <- data
	}()

	return
}

func (s *Wrapped) ExecuteScript(
	script string, data interface{}, args ...interface{},
) (err error) {
	cmd := &ito.ExecuteScript{
		Script: script,
		Args:   args,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg, data)
}

func (s *Wrapped) ExecuteScriptIn(
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

func (s *Wrapped) FindElement(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ret *marionette.WebElement, err error) {
	cmd := &ito.FindElement{
		Using:       by,
		Value:       qstr,
		RootElement: root,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Wrapped) FindElements(
	by marionette.FindStrategy, qstr string, root *marionette.WebElement,
) (ret []*marionette.WebElement, err error) {
	cmd := &ito.FindElements{
		Using:       by,
		Value:       qstr,
		RootElement: root,
	}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Wrapped) Forward() (err error) {
	cmd := &ito.Forward{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) FullscreenWindow() (err error) {
	cmd := &ito.FullscreenWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) GetActiveElement() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveElement{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetActiveFrame() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveFrame{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetAlertText() (ret string, err error) {
	cmd := &ito.GetAlertText{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetCapabilities() (ret *marionette.Capabilities, err error) {
	cmd := &ito.GetCapabilities{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetChromeWindowHandle() (ret string, err error) {
	cmd := &ito.GetChromeWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetChromeWindowHandles() (ret []string, err error) {
	cmd := &ito.GetChromeWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetCookies() (ret []*marionette.Cookie, err error) {
	cmd := &ito.GetCookies{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetCurrentURL() (ret string, err error) {
	cmd := &ito.GetCurrentURL{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetElementAttribute(
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

func (s *Wrapped) GetElementCSSValue(
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

func (s *Wrapped) GetElementProperty(
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

func (s *Wrapped) GetElementRect(
	el *marionette.WebElement,
) (ret marionette.Rect, err error) {
	cmd := &ito.GetElementRect{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetElementTagName(
	el *marionette.WebElement,
) (ret string, err error) {
	cmd := &ito.GetElementTagName{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetElementText(el *marionette.WebElement) (ret string, err error) {
	cmd := &ito.GetElementText{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetPageSource() (ret string, err error) {
	cmd := &ito.GetPageSource{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetTimeouts() (ret *marionette.Timeouts, err error) {
	cmd := &ito.GetTimeouts{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetTitle() (ret string, err error) {
	cmd := &ito.GetTitle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetWindowHandle() (ret string, err error) {
	cmd := &ito.GetWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetWindowHandles() (ret []string, err error) {
	cmd := &ito.GetWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) GetWindowRect() (ret marionette.Rect, err error) {
	cmd := &ito.GetWindowRect{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) IsElementDisplayed(
	el *marionette.WebElement, key string,
) (ret bool, err error) {
	cmd := &ito.IsElementDisplayed{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) IsElementEnabled(
	el *marionette.WebElement, key string,
) (ret bool, err error) {
	cmd := &ito.IsElementEnabled{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) IsElementSelected(
	el *marionette.WebElement, key string,
) (ret bool, err error) {
	cmd := &ito.IsElementSelected{
		Element: el,
	}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Wrapped) MaximizeWindow() (err error) {
	cmd := &ito.MaximizeWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) MinimizeWindow() (err error) {
	cmd := &ito.MinimizeWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) Navigate(url string) (err error) {
	cmd := &ito.Navigate{URL: url}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) NewSession() (id string, cap *marionette.Capabilities, err error) {
	cmd := &ito.NewSession{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Wrapped) NewWindow(typ string, focus bool) (id, winType string, err error) {
	cmd := &ito.NewWindow{Type: typ, Focus: focus}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Wrapped) Refresh() (err error) {
	cmd := &ito.Refresh{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SendAlertText(text string) (err error) {
	cmd := &ito.SendAlertText{Text: text}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SetTimeouts(t *marionette.Timeouts) (err error) {
	cmd := &ito.SetTimeouts{Timeouts: t}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SetWindowRect(
	r marionette.Rect,
) (ret marionette.Rect, err error) {
	cmd := &ito.SetWindowRect{Rect: r}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Wrapped) SwitchToFrame(
	el *marionette.WebElement, id interface{}, focus bool,
) (err error) {
	cmd := &ito.SwitchToFrame{
		Element: el,
		Focus:   focus,
	}
	if id != nil {
		cmd.ID = id
	}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SwitchToParentFrame() (err error) {
	cmd := &ito.SwitchToParentFrame{}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SwitchToWindow(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) SwitchToWindowBG(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle, NoFocus: true}
	_, err = s.Sync(cmd)
	return
}

func (s *Wrapped) TakeScreenshot() (img string, err error) {
	cmd := &ito.TakeScreenshot{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Wrapped) TakeScreenshotBytes() (img []byte, err error) {
	str, err := s.TakeScreenshot()
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

func (s *Wrapped) MozGetContext() (ret string, err error) {
	cmd := &ito.MozGetContext{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Wrapped) MozSetContext(context string) (ret string, err error) {
	cmd := &ito.MozSetContext{Context: context}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}
