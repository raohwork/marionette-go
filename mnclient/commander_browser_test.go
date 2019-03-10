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
	"reflect"
	"testing"
)

func (tc *cmdrTestCase) testBack(t *testing.T) {
	tc.must(t, "goto-about", tc.nav("about:about"))
	tc.must(t, "goto-logo", tc.nav("about:logo"))
	if err := tc.Back(); err != nil {
		t.Fatalf("cannot go back: %s", err)
	}
	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "about:about" {
		t.Fatalf("unexpected url: %s", uri)
	}
}

func (tc *cmdrTestCase) testForward(t *testing.T) {
	if err := tc.Forward(); err != nil {
		t.Fatalf("cannot go forward: %s", err)
	}
	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "about:logo" {
		t.Fatalf("unexpected url: %s", uri)
	}
}

func (tc *cmdrTestCase) testRefresh(t *testing.T) {
	if err := tc.Refresh(); err != nil {
		t.Fatalf("cannot refresh: %s", err)
	}
	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "about:about" {
		t.Fatalf("unexpected url: %s", uri)
	}
}

func (tc *cmdrTestCase) testGetSetWindowRect(t *testing.T) {
	rect, err := tc.GetWindowRect()
	if err != nil {
		t.Fatalf("cannot get window rect: %s", err)
	}

	ret, err := tc.SetWindowRect(rect)
	if err != nil {
		t.Fatalf("cannot set window rect: %s", err)
	}
	if !reflect.DeepEqual(ret, rect) {
		t.Errorf("unexpected size: %+v", ret)
	}
}

func (tc *cmdrTestCase) testFullscreenWindow(t *testing.T) {
	rect, err := tc.GetWindowRect()
	if err != nil {
		t.Fatalf("cannot get window rect: %s", err)
	}

	if err = tc.FullscreenWindow(); err != nil {
		t.Errorf("cannot goto fullscreen mode: %s", err)
	}

	_, err = tc.SetWindowRect(rect)
	if err != nil {
		t.Fatalf("cannot set window rect: %s", err)
	}
}

func (tc *cmdrTestCase) testMaximizeWindow(t *testing.T) {
	rect, err := tc.GetWindowRect()
	if err != nil {
		t.Fatalf("cannot get window rect: %s", err)
	}

	if err = tc.MaximizeWindow(); err != nil {
		t.Errorf("cannot maximize: %s", err)
	}

	_, err = tc.SetWindowRect(rect)
	if err != nil {
		t.Fatalf("cannot set window rect: %s", err)
	}
}
