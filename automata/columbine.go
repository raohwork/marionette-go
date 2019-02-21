package automata

import (
	"errors"
	"sync"
	"time"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/ito"
	"github.com/raohwork/marionette-go/shirogane"
)

type tabManager interface {
	allocateTab(tab string) error
	releaseTab()
}

type myMixed struct {
	name string
	mgr  tabManager
	shirogane.Kuroga
}

func (s *myMixed) Sync(cmd ito.Ito) (msg *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	msg, err = s.Kuroga.Sync(cmd)
	s.mgr.releaseTab()

	return
}

func (s *myMixed) Async(cmd ito.Ito) (ch chan *marionette.Message, err error) {
	if err = s.mgr.allocateTab(s.name); err != nil {
		return
	}
	ch, err = s.Kuroga.Async(cmd)
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
//   - MozGetContext
//   - MozGetContext
//
// Executing unsupported commands leads to panic!
type Tab struct {
	myMixed *myMixed
	*shirogane.Ashihana
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
func (t *Tab) MozGetContext() (a string, err error) {
	panic(errors.New("MozGetContext is not supported in Columbine"))
}
func (t *Tab) MozSetContext(s string) (a string, err error) {
	panic(errors.New("MozSetContext is not supported in Columbine"))
}

// GetName returns current tab name
func (t *Tab) GetName() (ret string) {
	return t.myMixed.name
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
// The name comes from Japanese comic "Karakuri circus", which is an automata
// created by Faceless commander.
type Columbine struct {
	currentTab string
	allTabs    map[string]string
	client     *shirogane.Ashihana
	tabClients map[string]*Tab

	lock sync.Mutex
}

// NewColumbine creates a Columbine instance
//
// You have to start/issue new session/stop the *real client* (passed in m) by
// your self.
//
// Passing empty tab names leads to panic!
func NewColumbine(m shirogane.Kuroga, tabs []string) (ret *Columbine, err error) {
	if len(tabs) == 0 {
		panic(errors.New("tabs cannot be empty"))
	}

	ret = &Columbine{
		currentTab: tabs[0],
		allTabs:    map[string]string{},
		tabClients: map[string]*Tab{},
		client:     &shirogane.Ashihana{Kuroga: m},
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
		c := &myMixed{
			name:   tab,
			mgr:    ret,
			Kuroga: m,
		}
		ret.tabClients[tab] = &Tab{
			myMixed:  c,
			Ashihana: &shirogane.Ashihana{Kuroga: c},
		}
	}

	return
}

func (c *Columbine) allocateTab(tab string) (err error) {
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

func (c *Columbine) releaseTab() {
	c.lock.Unlock()
}

// GetTab retrieves specified Tab instance by tab name you gave to Columbine
//
// Returns nil if not exists.
func (c *Columbine) GetTab(tab string) (ret *Tab) {
	return c.tabClients[tab]
}
