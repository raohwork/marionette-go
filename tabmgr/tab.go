// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

package tabmgr

import (
	"errors"
	"time"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnclient"
	"github.com/raohwork/marionette-go/mnsender"
)

// Tab represents a tab under manager's control
//
// It will check if active tab is desired one right before executing commands, and
// switch to it if needed.
//
// Here's a list of unsupported commands:
//
//   - CloseChromeWindow
//   - CloseWindow
//   - NewSession
//   - NewWindow
//   - SwitchToWindow
//   - SwitchToWindowBG
//
// And a list of "USE AT YOUR OWN RISK" commands:
//
//   - Moz commands like MozGetContext
//   - Reftest commands
//
// Also, modal dialogs may cause undesired effect to many commands.
//
// Executing unsupported commands leads to panic!
type Tab struct {
	mySender *lockedSender
	*mnclient.Commander
}

// NewTab creates a new Tab instance
func NewTab(name string, mgr LockManager, sender mnsender.Sender) (ret *Tab) {
	s := &lockedSender{
		name:   name,
		mgr:    mgr,
		Sender: sender,
	}
	return &Tab{
		mySender:  s,
		Commander: &mnclient.Commander{Sender: s},
	}
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
	return t.mySender.GetName()
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
