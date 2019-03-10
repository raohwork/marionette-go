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

func (tc *cmdrTestCase) testMozGetContext(t *testing.T) {
	ctx, err := tc.MozGetContext()
	if err != nil {
		t.Fatalf("cannot get context: %s", err)
	}

	if ctx != marionette.ContentContext {
		t.Fatalf("unexpected context: %s", ctx)
	}
}

func (tc *cmdrTestCase) testMozSetContext(t *testing.T) {
	err := tc.MozSetContext(marionette.ChromeContext)
	if err != nil {
		t.Fatalf("cannot set context: %s", err)
	}
	defer tc.MozSetContext(marionette.ContentContext)

	ctx, _ := tc.MozGetContext()
	if ctx != marionette.ChromeContext {
		t.Fatalf("unexpected context: %s", ctx)
	}
}
