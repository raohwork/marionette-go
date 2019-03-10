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
	"sync"

	"github.com/raohwork/marionette-go/v3/mnclient"
	"github.com/raohwork/marionette-go/v3/mnsender"
)

// WindowManager creates a TabManager instance to control tabs in multiple window
//
// The actual tab name will be "window_name:tab_name"
//
//   w, _ := WindowManager(sender, map[string][]string{"w1": {"t1", "t2"}})
//   tab := w.GetTab("w1:t1")
//
// It creates/closes windows and tabs automatically (using mnclient.TabCreator,
// refer there for implementation details), does not cooperate with other managers
// well. That's the reason why there's no "lock" argument here.
func WindowManager(
	s mnsender.Sender, windows map[string][]string,
) (ret *TabManager, err error) {
	if len(windows) == 0 {
		panic(errors.New("windows cannot be empty"))
	}
	for _, tabs := range windows {
		if len(tabs) == 0 {
			panic(errors.New("tabs cannot be empty"))
		}
	}

	cl := &mnclient.Commander{Sender: s}
	creator := mnclient.NewTabCreator(cl)

	// ensuring number of windows
	curWin, err := cl.GetChromeWindowHandles()
	if err != nil {
		return
	}
	curWin, err = creator.EnsureWindowNumber(curWin, len(windows))
	if err != nil {
		return
	}

	// create mapping of window name => window handle
	winMap := map[string]string{}
	idx := 0
	for name, _ := range windows {
		winMap[name] = curWin[idx]
		idx++
	}

	// get current tab mapping
	curTabs, err := creator.CurrentTabMapping()
	if err != nil {
		return
	}

	// ensuring number of tabs in each window
	for name, winH := range winMap {
		have := curTabs[winH]
		want := len(windows[name])
		got, e := creator.EnsureTabNumber(have, want)
		if e != nil {
			err = e
			return
		}
		curTabs[winH] = got
	}

	// create tab name => handle mapping
	allTabs := map[string]string{}
	for wname, tabs := range windows {
		winH := winMap[wname]
		for idx, tname := range tabs {
			allTabs[wname+":"+tname] = curTabs[winH][idx]
		}
	}

	ret = NewWithMap(s, allTabs, &sync.Mutex{})
	return
}
