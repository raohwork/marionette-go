package mnclient

import "testing"

func (tc *cmdrTestCase) testNewSessionWith(t *testing.T) {
	id, cap, err := tc.NewSessionWith("normal", true)
	if err != nil {
		t.Fatalf("unexpected error in NewSessionWith(): %s", err)
	}
	if id == "" {
		t.Fatal("empty id")
	}
	if cap == nil {
		t.Fatal("empty cap")
	}

	if cap.PageLoadStrategy != "normal" {
		t.Errorf("unexpected page load strategy: %s", cap.PageLoadStrategy)
	}

	if !cap.AcceptInsecureCerts {
		t.Errorf("unexpected insecure cert: %t", cap.AcceptInsecureCerts)
	}

	t.Logf("id: %s", id)
	t.Logf("cap: %+v", cap)
}

func (tc *cmdrTestCase) testNavigate(t *testing.T) {
	err := tc.Navigate("about:about")
	if err != nil {
		t.Fatalf("unexpected error in Navigate(): %s", err)
	}
}

func (tc *cmdrTestCase) testGetWindowHandles(t *testing.T) {
	h, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("unexpected error in GetWindowHandles(): %s", err)
	}

	if len(h) < 1 {
		t.Fatal("empty handles")
	}

	t.Logf("handles: %+v", h)
}

func (tc *cmdrTestCase) testNewWindow(t *testing.T) {
	me, _ := tc.GetWindowHandle()
	id, typ, err := tc.NewWindow("tab", true)
	if err != nil {
		t.Fatalf("unexpected error in NewWindow(): %s", err)
	}
	if id == "" {
		t.Error("empty id")
	}
	if typ != "tab" {
		t.Errorf("unexpected type: %s", typ)
	}

	newMe, _ := tc.GetWindowHandle()
	if me != newMe {
		t.Error("marionette switched to new tab")
	}
}

func (tc *cmdrTestCase) testSwitchToWindow(t *testing.T) {
	tabs, err := tc.GetWindowHandles()
	if err != nil {
		t.Fatalf("cannot get all tabs: %s", err)
	}

	for _, h := range tabs {
		if err = tc.SwitchToWindow(h); err != nil {
			t.Errorf("cannot switch to tab#%s: %s", h, err)
		}
	}
}

func (tc *cmdrTestCase) testGetWindowHandle(t *testing.T) {
	h, err := tc.GetWindowHandle()
	if err != nil {
		t.Fatalf("unexpected error in GetWindowHandle(): %s", err)
	}
	if h == "" {
		t.Error("empty handler")
	}
	t.Logf("handle: %s", h)
}

func (tc *cmdrTestCase) testCloseWindow(t *testing.T) {
	handles, err := tc.CloseWindow()
	if err != nil {
		t.Fatalf("unexpected error in CloseWindow(): %s", err)
	}
	if len(handles) < 1 {
		t.Errorf("invalid handles returned: %+v", handles)
	}

	// clean up: switch to another window
	t.Run("CleanUp", tc.testSwitchToWindow)
}

func (tc *cmdrTestCase) testGetCurrentURL(t *testing.T) {
	tc.must(t, "ensure-tabs", tc.ensureSingleTab)
	if err := tc.Navigate("about:about"); err != nil {
		t.Fatalf("cannot go to about:about: %s", err)
	}

	uri, err := tc.GetCurrentURL()
	if err != nil {
		t.Fatalf("cannot get current url: %s", err)
	}
	if uri != "about:about" {
		t.Errorf("unexpected url: %s", uri)
	}
}
