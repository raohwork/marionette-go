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

func (tc *cmdrTestCase) testSwitchToFrame(t *testing.T) {
	if err := tc.SwitchToFrame(nil, 0, true); err != nil {
		t.Fatalf("cannot switch to first frame")
	}

	el, _ := tc.FindElement(marionette.ID, "result", nil)
	if el == nil {
		t.Fatal("no result")
	}

	txt, _ := tc.GetElementText(el)
	if str := strings.TrimSpace(txt); str != "frame1" {
		t.Fatalf("unexpected frame: %s", str)
	}
}

func (tc *cmdrTestCase) testGetActiveFrame(t *testing.T) {
	tc.SwitchToFrame(nil, 1, true)
	h, err := tc.GetActiveFrame()
	if err != nil {
		t.Fatalf("cannot get active frame: %s", err)
	}
	if h == nil {
		t.Fatal("empty handle for inner frame")
	}

	t.Logf("handle: %+v", h)
}

func (tc *cmdrTestCase) testSwitchToParentFrame(t *testing.T) {
	if err := tc.SwitchToParentFrame(); err != nil {
		t.Fatalf("cannot switch to parent frame: %s", err)
	}

	el, _ := tc.FindElement(marionette.ID, "result", nil)
	if el != nil {
		t.Fatalf("unexpected element: %+v", el)
	}
}
