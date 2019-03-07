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
	"sync"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnclient"
	"github.com/raohwork/marionette-go/mnsender"
)

// TabManager is a client focus on multi-tab environment
//
// You give it a list of tab names, it will take care of your commands, making
// sure they are executed in correct tab.
//
// As commands are mean to be executed in specified tab only, some commands are not
// supported like NewWindow and SwitchToWindow (see Tab for full list).
//
// Only content window and conext are supported.
type TabManager struct {
	tabClients map[string]*Tab
	LockManager
}

// New is identical to NewWithLock(m, tabs, &sync.Mutex{})
//
// See NewWithLock() for further information.
func New(m mnsender.Sender, tabs []string) (ret *TabManager, err error) {
	return NewWithLock(m, tabs, &sync.Mutex{})
}

// NewWithLock creates a TabManager with predefined lock instance
//
// You have to start/issue new session/stop the *real client* (passed in m) before
// calling New(). It will open tabs in current window or close other window/tab if
// needed.
//
// Passing empty tab names leads to panic!
func NewWithLock(m mnsender.Sender, tabs []string, lock sync.Locker) (ret *TabManager, err error) {
	if len(tabs) == 0 {
		panic(errors.New("tabs cannot be empty"))
	}

	cl := &mnclient.Commander{Sender: m}
	allTabs := map[string]string{}

	// get current tabs
	curTabs, err := cl.GetWindowHandles()
	if err != nil {
		return
	}

	wanted := len(tabs)
	// open new tabs if needed
	// Firefox 65 does not supports NewWindow command yet
	fx65workround := false
	for x := len(curTabs); x < wanted; x++ {
		if !fx65workround {
			_, _, err = cl.NewWindow("tab", false)
			if err != nil {
				e, ok := err.(*marionette.ErrDriver)
				if !ok {
					return
				}
				if e.Type != marionette.ErrUnknownCommand {
					return
				}
				fx65workround = true
			}
		}
		if fx65workround {
			err = cl.ExecuteScript(
				`window.open('about:blank')`, nil,
			)
			if err != nil {
				return
			}
		}
	}
	// close old tabs if needed
	for x := len(curTabs); x > wanted; x-- {
		list, err := cl.CloseWindow()
		if err != nil {
			return nil, err
		}
		cl.SwitchToWindow(list[0])
	}
	// no matter what, just fetch tabs list again
	if curTabs, err = cl.GetWindowHandles(); err != nil {
		return
	}

	for idx, tab := range tabs {
		allTabs[tab] = curTabs[idx]
	}

	ret = NewWithMap(m, allTabs, lock)

	// compatible with 3.0.0: switch to first tab
	ret.allocateTab(tabs[0])
	ret.releaseTab()
	return
}

// NewWithMap creates a TabManager with predefined tabs
//
// You have to start/issue new session/stop the *real client* (passed in m) before
// calling New(). It will open tabs in current window or close other window/tab if
// needed.
//
// Passing empty tab mapping leads to panic!
//
// It is impossible for go map to determine "first" element, NewWithMap will switch
// to randomly choosed one of managed tabs before returning created TabManager.
//
// With this function, external packages may write their own logic to create
// windows/tabs, and let TabManager take care about race conditions.
func NewWithMap(m mnsender.Sender, tabs map[string]string, lock sync.Locker) (ret *TabManager) {
	cl := &mnclient.Commander{Sender: m}
	ret = &TabManager{
		tabClients:  map[string]*Tab{},
		LockManager: NewLockManager(tabs, cl, lock),
	}

	// create tab clients
	var tab string
	for n, _ := range tabs {
		tab = n
		ret.tabClients[n] = NewTab(n, ret, m)
	}

	// switch to a random tab
	ret.allocateTab(tab)
	ret.releaseTab()

	return
}

// GetTab retrieves specified Tab instance by tab name you gave to TabManager
//
// Returns nil if not exists.
func (c *TabManager) GetTab(tab string) (ret *Tab) {
	return c.tabClients[tab]
}

// Reset set all tabs to a page (default "about:blank")
func (t *TabManager) Reset(uri string) (err error) {
	if uri == "" {
		uri = "about:blank"
	}
	for _, tab := range t.tabClients {
		err = tab.Navigate(uri)
		if err != nil {
			return
		}
	}

	return
}
