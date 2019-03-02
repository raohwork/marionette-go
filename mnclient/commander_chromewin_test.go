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

package mnclient

import (
	"testing"
)

func (tc *cmdrTestCase) testGetChromeHandles(t *testing.T) {
	handles, err := tc.GetChromeWindowHandles()
	if err != nil {
		t.Fatalf("cannot get all chrome window handles: %s", err)
	}
	if len(handles) != 1 {
		t.Fatalf("unexpected handles: %+v", handles)
	}

	h, err := tc.GetChromeWindowHandle()
	if err != nil {
		t.Fatalf("cannot get current chrome window handle: %s", err)
	}
	if h != handles[0] {
		t.Fatalf("current chrome handle != allhandle: %s", h)
	}
}

func (tc *cmdrTestCase) testCloseChromeWindow(t *testing.T) {
	me, _ := tc.GetChromeWindowHandle()
	newMe, _, _ := tc.NewWindow("window", true)
	if me == newMe {
		t.Error("new chrome window == current")
	}
	newMe, _ = tc.GetChromeWindowHandle()
	if me != newMe {
		t.Error("marionette switched to new window")
	}

	handles, _ := tc.GetChromeWindowHandles()
	if l := len(handles); l != 2 {
		t.Errorf("expected 2 windows, got %d", l)
	}

	handles, err := tc.CloseChromeWindow()
	if err != nil {
		t.Fatalf("cannot close current chrome window: %s", err)
	}

	if l := len(handles); l != 1 {
		t.Errorf("expected 1 windows, got %d", l)
	}
}
