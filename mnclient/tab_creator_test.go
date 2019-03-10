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

package mnclient

import (
	"testing"

	"github.com/raohwork/marionette-go/v3/mnsender"
)

func TestTabCreator(t *testing.T) {
	sender, err := mnsender.NewTCPSender(addr, 0)
	if err != nil {
		t.Fatalf("unexpected error in NewTCPSender(): %s", err)
	}
	sender.Start()
	defer sender.Close()
	cl := &Commander{Sender: sender}
	cl.NewSession()

	tc := &tcreateTestCase{
		TabCreator: NewTabCreator(cl),
		Commander:  cl,
	}

	t.Run("Mapping", tc.testGet)
	t.Run("OpenNewWin", tc.testOpenNewWin)
	t.Run("CloseOldWin", tc.testCloseOldWin)
	t.Run("OpenNewTab", tc.testOpenNewTab)
	t.Run("CloseOldTab", tc.testCloseOldTab)
}

type tcreateTestCase struct {
	TabCreator
	*Commander
}

func (tc *tcreateTestCase) testGet(t *testing.T) {
	mapping, err := tc.CurrentTabMapping()
	if err != nil {
		t.Fatalf("cannot get current tb mapping: %s", err)
	}

	// test windows
	windows, err := tc.GetChromeWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of current opened windows: %s", err)
	}

	if exp, act := len(windows), len(mapping); exp != act {
		t.Fatalf("expected %d windows, got %d", exp, act)
	}

	for _, h := range windows {
		if _, ok := mapping[h]; !ok {
			t.Fatalf("expected window %s to exist, but not found", h)
		}
	}

	// test tabs
	alltabs := map[string]bool{}
	tabcount := 0
	for _, tabs := range mapping {
		for _, h := range tabs {
			alltabs[h] = true
		}
		tabcount += len(tabs)
	}
	if l := len(alltabs); l != tabcount {
		t.Fatal("duplicated tabs detected!")
	}

	tabs, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of opened tabs: %s", err)
	}
	if exp := len(tabs); exp != tabcount {
		t.Fatalf("expected %d tabs, got %d", exp, tabcount)
	}
	for _, h := range tabs {
		if _, ok := alltabs[h]; !ok {
			t.Fatalf("expected tab %s to exist, but not found", h)
		}
	}
}

func (tc *tcreateTestCase) isWindowHandle(h string) (f func(*testing.T)) {
	return func(t *testing.T) {
		if err := tc.SwitchToWindow(h); err != nil {
			t.Fatalf("cannot switch to %s: %s", h, err)
		}

		act, err := tc.GetChromeWindowHandle()
		if err != nil {
			t.Fatalf("cannot get window handle for %s: %s", h, err)
		}
		if act != h {
			t.Fatalf("%s is not valid window handle", h)
		}
	}
}

func (tc *tcreateTestCase) isTabHandle(h string) (f func(*testing.T)) {
	return func(t *testing.T) {
		if err := tc.SwitchToWindow(h); err != nil {
			t.Fatalf("cannot switch to %s: %s", h, err)
		}

		act, err := tc.GetWindowHandle()
		if err != nil {
			t.Fatalf("cannot get tab handle for %s: %s", h, err)
		}
		if act != h {
			t.Fatalf("%s is not valid tab handle", h)
		}
	}
}

func (tc *tcreateTestCase) testOpenNewWin(t *testing.T) {
	wins, err := tc.GetChromeWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of opened windows: %s", err)
	}

	exp := len(wins) + 10
	now, err := tc.EnsureWindowNumber(wins, exp)
	if err != nil {
		t.Fatalf("cannot ensure window number: %s", err)
	}

	for _, h := range now {
		if ok := t.Run("IsWin-"+h, tc.isWindowHandle(h)); !ok {
			t.SkipNow()
		}
	}

	if act := len(now); exp != act {
		t.Fatalf("expected %d windows, got %d", exp, act)
	}

	// test if original windows are still alive
	curmap := map[string]bool{}
	for _, h := range now {
		curmap[h] = true
	}

	for _, h := range wins {
		if _, ok := curmap[h]; !ok {
			t.Fatalf("original window %s is dead", h)
		}
	}
}

func (tc *tcreateTestCase) testCloseOldWin(t *testing.T) {
	wins, err := tc.GetChromeWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of opened windows: %s", err)
	}

	exp := 1
	now, err := tc.EnsureWindowNumber(wins, exp)
	if err != nil {
		t.Fatalf("cannot ensure window number: %s", err)
	}

	for _, h := range now {
		if ok := t.Run("IsWin-"+h, tc.isWindowHandle(h)); !ok {
			t.SkipNow()
		}
	}

	if act := len(now); exp != act {
		t.Fatalf("expected %d windows, got %d", exp, act)
	}

	// test if original window is still alive
	ok := false
	for _, h := range wins {
		if h == now[0] {
			ok = true
			break
		}
	}

	if !ok {
		t.Fatal("original window is daed")
	}
}

func (tc *tcreateTestCase) testOpenNewTab(t *testing.T) {
	tabs, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of opened tabs: %s", err)
	}

	exp := len(tabs) + 10
	now, err := tc.EnsureTabNumber(tabs, exp)
	if err != nil {
		t.Fatalf("cannot ensure tab number: %s", err)
	}

	for _, h := range now {
		if ok := t.Run("IsTab-"+h, tc.isTabHandle(h)); !ok {
			t.SkipNow()
		}
	}

	if act := len(now); exp != act {
		t.Fatalf("expected %d tabs, got %d", exp, act)
	}

	// test if original tabs are still alive
	curmap := map[string]bool{}
	for _, h := range now {
		curmap[h] = true
	}

	for _, h := range tabs {
		if _, ok := curmap[h]; !ok {
			t.Fatalf("original tab %s is dead", h)
		}
	}
}

func (tc *tcreateTestCase) testCloseOldTab(t *testing.T) {
	tabs, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("cannot get list of opened tabs: %s", err)
	}

	exp := 1
	now, err := tc.EnsureTabNumber(tabs, exp)
	if err != nil {
		t.Fatalf("cannot ensure tab number: %s", err)
	}

	for _, h := range now {
		if ok := t.Run("IsTab-"+h, tc.isTabHandle(h)); !ok {
			t.SkipNow()
		}
	}

	if act := len(now); exp != act {
		t.Fatalf("expected %d tabs, got %d", exp, act)
	}

	// test if original tabs are still alive
	ok := false
	for _, h := range tabs {
		if h == now[0] {
			ok = true
			break
		}
	}

	if !ok {
		t.Fatal("original tab is daed")
	}
}
