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

	marionette "github.com/raohwork/marionette-go"
)

func (tc *cmdrTestCase) testReftest(t *testing.T) {
	if err := tc.MozSetContext(marionette.ChromeContext); err != nil {
		t.Fatalf("cannot set context: %s", err)
	}
	defer tc.MozSetContext(marionette.ContentContext)

	victim := "http://localhost:9487/ref_victim.html"
	err := tc.ReftestSetup(
		map[string]uint{victim: 1},
		marionette.ReftestAlwaysScreenshot,
	)
	if err != nil {
		t.Fatalf("cannot initialize reftest: %s", err)
	}

	rules := (marionette.ReftestRefList{}).
		Or("http://localhost:9487/ref_unexpect.html", marionette.RelNE)

	result, err := tc.ReftestRun(
		victim,
		marionette.ReftestPass,
		rules,
		10000,
	)
	if err != nil {
		t.Fatalf("cannot run reftest: %s", err)
	}

	if result.Status != "PASS" {
		t.Errorf("unexpected result: %+v", result)
	}

	if err := tc.ReftestTeardown(); err != nil {
		t.Fatalf("cannot deinitialize reftest: %s", err)
	}
}
