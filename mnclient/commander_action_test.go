package mnclient

import (
	"strings"
	"testing"

	marionette "github.com/raohwork/marionette-go"
)

func (tc *cmdrTestCase) testPerformAction(t *testing.T) {
	btn, _ := tc.FindElement(marionette.ID, "run", nil)
	rect, _ := tc.GetElementRect(btn)

	actions := marionette.ActionChain{}
	actions.MouseMoveTo(int(rect.X+rect.W/2), int(rect.Y+rect.H/2), 100)
	actions.MouseClick(marionette.MouseLeft)
	if err := tc.PerformActions(actions); err != nil {
		t.Fatalf("cannot perform action: %s", err)
	}

	result, _ := tc.FindElement(marionette.ID, "result", nil)
	txt, _ := tc.GetElementText(result)
	if strings.TrimSpace(txt) != "demo" {
		// gather info
		var realRect marionette.Rect
		tc.ExecuteScript(`
const el = document.querySelector('#run');
return el.getBoundingClientRect();
`, &realRect)
		t.Fatalf("unexpected value: %s, dumping info:\nreal rect: %+v\nmy rect: %+v", txt, realRect, rect)
	}
}

func (tc *cmdrTestCase) testReleaseActions(t *testing.T) {
	if err := tc.ReleaseActions(); err != nil {
		t.Fatalf("cannot release actions: %s", err)
	}
}
