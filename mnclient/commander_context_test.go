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
