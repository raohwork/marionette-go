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

func (s *Ashihana) AcceptAlert() (err error) {
	cmd := &ito.AcceptAlert{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) AddCookie(cookie *marionette.Cookie) (err error) {
	cmd := &ito.AddCookie{Cookie: cookie}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) Back() (err error) {
	cmd := &ito.Back{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) CloseChromeWindow() (handles []string, err error) {
	cmd := &ito.CloseChromeWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Ashihana) CloseWindow() (handles []string, err error) {
	cmd := &ito.CloseWindow{}
	msg, err := s.Sync(cmd)
	return cmd.Decode(msg)
}

func (s *Ashihana) DeleteAllCookies() (err error) {
	cmd := &ito.DeleteAllCookies{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) DeleteCookie(name string) (err error) {
	cmd := &ito.DeleteCookie{Name: name}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) DismissAlert() (err error) {
	cmd := &ito.DismissAlert{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) ElementClear(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClear{Element: el}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) ElementClick(el *marionette.WebElement) (err error) {
	cmd := &ito.ElementClick{Element: el}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) ElementSendKeys(el *marionette.WebElement, text string) (err error) {
	cmd := &ito.ElementSendKeys{Element: el, Text: text}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) ExecuteAsyncScript(script string, args ...interface{}) (
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

func (s *Ashihana) ExecuteAsyncScriptIn(
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

func (s *Ashihana) FindElement(
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

func (s *Ashihana) FindElements(
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

func (s *Ashihana) Forward() (err error) {
	cmd := &ito.Forward{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) FullscreenWindow() (err error) {
	cmd := &ito.FullscreenWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) GetActiveElement() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveElement{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetActiveFrame() (ret *marionette.WebElement, err error) {
	cmd := &ito.GetActiveFrame{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetAlertText() (ret string, err error) {
	cmd := &ito.GetAlertText{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetCapabilities() (ret *marionette.Capabilities, err error) {
	cmd := &ito.GetCapabilities{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetChromeWindowHandle() (ret string, err error) {
	cmd := &ito.GetChromeWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetChromeWindowHandles() (ret []string, err error) {
	cmd := &ito.GetChromeWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetCookies() (ret []*marionette.Cookie, err error) {
	cmd := &ito.GetCookies{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetCurrentURL() (ret string, err error) {
	cmd := &ito.GetCurrentURL{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

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

func (s *Ashihana) GetElementRect(
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

func (s *Ashihana) GetElementTagName(
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

func (s *Ashihana) GetPageSource() (ret string, err error) {
	cmd := &ito.GetPageSource{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetTimeouts() (ret *marionette.Timeouts, err error) {
	cmd := &ito.GetTimeouts{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetTitle() (ret string, err error) {
	cmd := &ito.GetTitle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetWindowHandle() (ret string, err error) {
	cmd := &ito.GetWindowHandle{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetWindowHandles() (ret []string, err error) {
	cmd := &ito.GetWindowHandles{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) GetWindowRect() (ret marionette.Rect, err error) {
	cmd := &ito.GetWindowRect{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}
	return cmd.Decode(msg)
}

func (s *Ashihana) IsElementDisplayed(
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

func (s *Ashihana) IsElementEnabled(
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

func (s *Ashihana) IsElementSelected(
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

func (s *Ashihana) MaximizeWindow() (err error) {
	cmd := &ito.MaximizeWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) MinimizeWindow() (err error) {
	cmd := &ito.MinimizeWindow{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) Navigate(url string) (err error) {
	cmd := &ito.Navigate{URL: url}
	_, err = s.Sync(cmd)
	return
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

func (s *Ashihana) NewSession() (id string, cap *marionette.Capabilities, err error) {
	cmd := &ito.NewSession{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Ashihana) NewWindow(typ string, focus bool) (id, winType string, err error) {
	cmd := &ito.NewWindow{Type: typ, Focus: focus}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Ashihana) Refresh() (err error) {
	cmd := &ito.Refresh{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) SendAlertText(text string) (err error) {
	cmd := &ito.SendAlertText{Text: text}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) SetTimeouts(t *marionette.Timeouts) (err error) {
	cmd := &ito.SetTimeouts{Timeouts: t}
	_, err = s.Sync(cmd)
	return
}

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
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) SwitchToParentFrame() (err error) {
	cmd := &ito.SwitchToParentFrame{}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) SwitchToWindow(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) SwitchToWindowBG(handle string) (err error) {
	cmd := &ito.SwitchToWindow{Name: handle, NoFocus: true}
	_, err = s.Sync(cmd)
	return
}

func (s *Ashihana) TakeScreenshot() (img string, err error) {
	cmd := &ito.TakeScreenshot{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Ashihana) TakeScreenshotBytes() (img []byte, err error) {
	str, err := s.TakeScreenshot()
	if err != nil {
		return
	}

	return base64.StdEncoding.DecodeString(str)
}

func (s *Ashihana) MozGetContext() (ret string, err error) {
	cmd := &ito.MozGetContext{}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}

func (s *Ashihana) MozSetContext(context string) (ret string, err error) {
	cmd := &ito.MozSetContext{Context: context}
	msg, err := s.Sync(cmd)
	if err != nil {
		return
	}

	return cmd.Decode(msg)
}
