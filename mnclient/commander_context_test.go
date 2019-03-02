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
