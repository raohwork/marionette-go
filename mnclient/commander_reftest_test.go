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

	rule := &marionette.ReftestRef{
		URL: "http://localhost:9487/ref_expect.html",
		Rel: marionette.RelEQ,
	}
	rules := marionette.ReftestRefList([]*marionette.ReftestRef{rule})

	result, err := tc.ReftestRun(
		victim,
		marionette.ReftestPass,
		rules,
		30000,
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
