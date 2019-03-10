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
