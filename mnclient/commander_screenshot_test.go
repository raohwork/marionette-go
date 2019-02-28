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
