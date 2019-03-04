// This file is part of marionette-go
//
// marionette-go is free software: you can redistribute it and/or modify it
// under the terms of the GNU Lesser General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with marionette-go. If not, see <https://www.gnu.org/licenses/>.

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
