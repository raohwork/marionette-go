package mnclient

import "testing"

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
	_, _, _ = tc.NewWindow("window", true)

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
