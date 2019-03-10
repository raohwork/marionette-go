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

	marionette "github.com/raohwork/marionette-go/v3"
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
