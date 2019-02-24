package tabmgr

import (
	"errors"
	"time"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnclient"
	"github.com/raohwork/marionette-go/mncmd"
	"github.com/raohwork/marionette-go/mnsender"
)

type tabManager interface {
	allocateTab(tab string) error
	releaseTab()
}

type mySender struct {
	name string
	mgr  tabManager
	mnsender.Sender
}

func (s *mySender) Sync(cmd mncmd.Command) (msg *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	msg, err = s.Sender.Sync(cmd)
	s.mgr.releaseTab()

	return
}

func (s *mySender) Async(cmd mncmd.Command) (ch chan *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	ch, err = s.Sender.Async(cmd)
	s.mgr.releaseTab()

	return
}

// Tab represents a tab under Columbine's controll
//
// It will check if active tab is desired one right before executing commands, and
// switch to it if needed.
//
// Commands not in following list are supported:
//
//   - CloseChromeWindow
//   - CloseWindow
//   - NewSession
//   - NewWindow
//   - SwitchToWindow
//   - SwitchToWindowBG
//
// Executing unsupported commands leads to panic!
type Tab struct {
	mySender *mySender
	*mnclient.Commander
}

func (t *Tab) CloseChromeWindow() (handles []string, err error) {
	panic(errors.New("CloseChromeWindow is not supported in Columbine"))
}
func (t *Tab) CloseWindow() (handles []string, err error) {
	panic(errors.New("CloseWindow is not supported in Columbine"))
}
func (t *Tab) NewSession() (a string, b *marionette.Capabilities, err error) {
	panic(errors.New("NewSession is not supported in Columbine"))
}
func (t *Tab) NewWindow(typ string, focus bool) (a, b string, err error) {
	panic(errors.New("NewWindow is not supported in Columbine"))
}
func (t *Tab) SwitchToWindow(s string) (err error) {
	panic(errors.New("SwitchToWindow is not supported in Columbine"))
}
func (t *Tab) SwitchToWindowBG(s string) (err error) {
	panic(errors.New("SwitchToWindowBG is not supported in Columbine"))
}

// GetName returns current tab name
func (t *Tab) GetName() (ret string) {
	return t.mySender.name
}

// WaitFor periodically check if specified element presents
//
// It makes few attempts (specified in "tries") to run FindElement command, and
// waits a second between each attempt.
//
// Should be useful if you're manipulating dynamic generated pages like SPA.
func (t *Tab) WaitFor(qstr string, tries int) (ret *marionette.WebElement, err error) {
	if tries < 1 {
		tries = 1
	}
	for x := 0; x < tries; x++ {
		time.Sleep(time.Second)

		ret, err = t.FindElement(
			marionette.Selector,
			qstr,
			nil,
		)
		if ret != nil && err == nil {
			return
		}

		x, ok := err.(*marionette.ErrDriver)
		if !ok {
			return
		}

		if x.Type != marionette.ErrNoSuchElement {
			return
		}
	}

	return
}
