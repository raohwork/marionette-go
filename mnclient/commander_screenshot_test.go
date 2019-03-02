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
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"testing"

	marionette "github.com/raohwork/marionette-go"
)

func (tc *cmdrTestCase) testScreenshotDocument(t *testing.T) {
	png, err := tc.ScreenshotDocumentBytes(nil)
	if err != nil {
		t.Fatalf("cannot take document screenshot: %s", err)
	}
	sum := md5.Sum(png)
	sumStr := hex.EncodeToString(sum[:])
	expected := `9ac5ddf51eea2f3af318bd486a39c5b9`
	if sumStr != expected {
		ioutil.WriteFile("document.png", png, 0644)
		t.Fatal("hash mismatch, dumped png to document.png")
	}
}

func (tc *cmdrTestCase) testScreenshotViewport(t *testing.T) {
	png, err := tc.ScreenshotViewportBytes(nil)
	if err != nil {
		t.Fatalf("cannot take viewport screenshot: %s", err)
	}
	sum := md5.Sum(png)
	sumStr := hex.EncodeToString(sum[:])
	expected := `6141b756eaad8fccb46659d338271863`
	if sumStr != expected {
		ioutil.WriteFile("viewport.png", png, 0644)
		t.Fatal("hash mismatch, dumped png to viewport.png")
	}
}

func (tc *cmdrTestCase) testScreenshotElement(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "ctrl", nil)
	png, err := tc.ScreenshotElementBytes(el)
	if err != nil {
		t.Fatalf("cannot take element screenshot: %s", err)
	}
	sum := md5.Sum(png)
	sumStr := hex.EncodeToString(sum[:])
	expected := `b58e2f91b47370be2b0ead4f265f80ae`
	if sumStr != expected {
		ioutil.WriteFile("element.png", png, 0644)
		t.Fatal("hash mismatch, dumped png to element.png")
	}
}
