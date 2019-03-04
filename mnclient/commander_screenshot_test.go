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

func (tc *cmdrTestCase) testScreenshotDocument(t *testing.T) {
	png, err := tc.ScreenshotDocumentBytes(nil)
	if err != nil {
		t.Fatalf("cannot take document screenshot: %s", err)
	}
	if len(png) == 0 {
		t.Fatal("empty png")
	}
}

func (tc *cmdrTestCase) testScreenshotViewport(t *testing.T) {
	png, err := tc.ScreenshotViewportBytes(nil)
	if err != nil {
		t.Fatalf("cannot take viewport screenshot: %s", err)
	}
	if len(png) == 0 {
		t.Fatal("empty png")
	}
}

func (tc *cmdrTestCase) testScreenshotElement(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "ctrl", nil)
	png, err := tc.ScreenshotElementBytes(el)
	if err != nil {
		t.Fatalf("cannot take element screenshot: %s", err)
	}
	if len(png) == 0 {
		t.Fatal("empty png")
	}
}
