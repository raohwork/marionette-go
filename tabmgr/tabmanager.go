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

// Colubine is a client focus on multi-tab/window environment
//
// You give her a list of tab names, she will take care of your commands, making
// sure they are executed in correct tab.
//
// As commands are mean to be executed in specified tab only, some commands are not
// supported like NewWindow and SwitchToWindow (see Tab for full list).
//
// Only content window and conext are supported.
//
// The name comes from Japanese comic "Karakuri circus", which is an tabmgr
// created by Faceless commander.
type TabManager struct {
	currentTab string
	allTabs    map[string]string
	client     *mnclient.Commander
	tabClients map[string]*Tab

	lock sync.Mutex
}

// NewColumbine creates a Columbine instance
//
// You have to start/issue new session/stop the *real client* (passed in m) by
// your self.
//
// Passing empty tab names leads to panic!
func New(m mnsender.Sender, tabs []string) (ret *TabManager, err error) {
	if len(tabs) == 0 {
		panic(errors.New("tabs cannot be empty"))
	}

	ret = &TabManager{
		currentTab: tabs[0],
		allTabs:    map[string]string{},
		tabClients: map[string]*Tab{},
		client:     &mnclient.Commander{Sender: m},
	}

	// get current tabs
	curTabs, err := ret.client.GetWindowHandles()
	if err != nil {
		return
	}

	wanted := len(tabs)
	// open new tabs if needed
	// Firefox 65 does not supports NewWindow command yet
	fx65workround := false
	for x := len(curTabs); x < wanted; x++ {
		if !fx65workround {
			_, _, err = ret.client.NewWindow("tab", false)
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
			err = ret.client.ExecuteScript(
				`window.open('about:blank')`, nil,
			)
			if err != nil {
				return
			}
		}
	}
	// close old tabs if needed
	for x := len(curTabs); x > wanted; x-- {
		list, err := ret.client.CloseWindow()
		if err != nil {
			return nil, err
		}
		ret.client.SwitchToWindow(list[0])
	}
	// no matter what, just fetch tabs list again
	if curTabs, err = ret.client.GetWindowHandles(); err != nil {
		return
	}

	for idx, tab := range tabs {
		ret.allTabs[tab] = curTabs[idx]
	}

	// switch to current tab
	if err = ret.client.SwitchToWindow(ret.allTabs[ret.currentTab]); err != nil {
		return
	}

	// create tab clients
	for _, tab := range tabs {
		c := &mySender{
			name:   tab,
			mgr:    ret,
			Sender: m,
		}
		ret.tabClients[tab] = &Tab{
			mySender:  c,
			Commander: &mnclient.Commander{Sender: c},
		}
	}

	return
}

func (c *TabManager) allocateTab(tab string) (err error) {
	c.lock.Lock()
	if c.currentTab == tab {
		return
	}

	if err = c.client.SwitchToWindow(c.allTabs[tab]); err != nil {
		return
	}

	c.currentTab = tab
	return
}

func (c *TabManager) releaseTab() {
	c.lock.Unlock()
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
