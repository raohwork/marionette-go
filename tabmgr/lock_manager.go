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
	"sync"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnclient"
	"github.com/raohwork/marionette-go/mncmd"
	"github.com/raohwork/marionette-go/mnsender"
)

// LockManager ensures later commands are issued to correct tab
type LockManager interface {
	allocateTab(tab string) error
	releaseTab()
	currentTab() string
}

type lockMgr struct {
	curTab string
	tabs   map[string]string
	lock   sync.Locker
	cl     *mnclient.Commander
}

func (c *lockMgr) allocateTab(tab string) (err error) {
	c.lock.Lock()
	if c.curTab == tab {
		return
	}

	if err = c.cl.SwitchToWindow(c.tabs[tab]); err != nil {
		return
	}

	c.curTab = tab
	return
}

func (c *lockMgr) releaseTab() {
	c.lock.Unlock()
}

func (c *lockMgr) currentTab() (ret string) {
	return c.curTab
}

// NewLockManager creates a LockManager instance
//
//   - The tabs maps tab name to marionette handle.
//   - cl is used to issue SwitchToWindow command.
//   - lock ensures only one tab is in active.
func NewLockManager(
	tabs map[string]string, cl *mnclient.Commander, lock sync.Locker,
) (ret LockManager) {
	return &lockMgr{
		tabs: tabs,
		lock: lock,
		cl:   cl,
	}
}

type lockedSender struct {
	name string
	mgr  LockManager
	mnsender.Sender
}

func (s *lockedSender) GetName() (ret string) {
	return s.name
}

func (s *lockedSender) Sync(cmd mncmd.Command) (msg *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	msg, err = s.Sender.Sync(cmd)
	s.mgr.releaseTab()

	return
}

func (s *lockedSender) Async(cmd mncmd.Command) (ch chan *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	ch, err = s.Sender.Async(cmd)
	s.mgr.releaseTab()

	return
}
