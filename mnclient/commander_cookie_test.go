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

func (tc *cmdrTestCase) testCookies(t *testing.T) {
	ok := t.Run("DeleteAll-error-free", func(t *testing.T) {
		if err := tc.DeleteAllCookies(); err != nil {
			t.Fatalf("cannot delete all cookies: %s", err)
		}
	})
	if !ok {
		t.Skip("unexpected error in DeleteAllCookies(), skip now")
	}

	myCookie := &marionette.Cookie{
		Name:  "myCookie",
		Value: "myValue",
	}
	ok = t.Run("AddCookie-error-free", func(t *testing.T) {
		if err := tc.AddCookie(myCookie); err != nil {
			t.Fatalf("cannot add cookie: %s", err)
		}
	})
	if !ok {
		t.Skip("unexpected error in AddCookie(), skip now")
	}

	ok = t.Run("GetCookies", func(t *testing.T) {
		c, err := tc.GetCookies()
		if err != nil {
			t.Fatalf("cannot get all cookies: %s", err)
		}

		if l := len(c); l != 1 {
			t.Fatalf("expected 1 cookie, got %+v", c)
		}

		if c[0].Name != myCookie.Name || c[0].Value != myCookie.Value {
			t.Fatalf("unexpected cookie: %+v", c[0])
		}
	})
	if !ok {
		t.Skip("unexpected error in GetCookies(), skip now")
	}

	myCookie2 := &marionette.Cookie{
		Name:  "myCookie2",
		Value: "myValue2",
	}
	t.Run("AddAnotherCookie", func(t *testing.T) {
		if err := tc.AddCookie(myCookie2); err != nil {
			t.Fatalf("cannot add cookie: %s", err)
		}
	})

	t.Run("DeleteCookie", func(t *testing.T) {
		if err := tc.DeleteCookie("myCookie"); err != nil {
			t.Fatalf("cannot delete cookie: %s", err)
		}

		c, err := tc.GetCookies()
		if err != nil {
			t.Fatalf("cannot get all cookies: %s", err)
		}

		if len(c) != 1 {
			t.Fatalf("expected 1 cookie, got %+v", c)
		}

		if c[0].Name != "myCookie2" || c[0].Value != "myValue2" {
			t.Fatalf("unexpected cookie: %+v", c[0])
		}
	})

	t.Run("DeleteAllCookie", func(t *testing.T) {
		if err := tc.DeleteAllCookies(); err != nil {
			t.Fatalf("cannot delete all cookies: %s", err)
		}

		c, err := tc.GetCookies()
		if err != nil {
			t.Fatalf("cannot get all cookies: %s", err)
		}

		if len(c) != 0 {
			t.Fatalf("expected no cookie, got %+v", c)
		}
	})
}
