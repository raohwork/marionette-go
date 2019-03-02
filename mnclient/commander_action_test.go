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
