package mnclient

import (
	"reflect"
	"testing"
)

func (tc *cmdrTestCase) testBack(t *testing.T) {
	t.Run("goto-google", tc.nav("https://google.com"))
	t.Run("goto-logo", tc.nav("about:logo"))
	if err := tc.Back(); err != nil {
		t.Fatalf("cannot go back: %s", err)
	}
	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "https://www.google.com/" {
		t.Fatalf("unexpected url: %s", uri)
	}
}

func (tc *cmdrTestCase) testForward(t *testing.T) {
	t.Run("back", tc.testBack)
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
	t.Run("back", tc.testBack)
	if err := tc.Refresh(); err != nil {
		t.Fatalf("cannot refresh: %s", err)
	}
	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "https://www.google.com/" {
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
