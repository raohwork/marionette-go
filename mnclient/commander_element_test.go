package mnclient

import (
	"testing"

	marionette "github.com/raohwork/marionette-go"
)

func (tc *cmdrTestCase) testFindElement(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))

	f := func(by marionette.FindStrategy, qstr string, root *marionette.WebElement) func(*testing.T) {
		return func(t *testing.T) {
			el, err := tc.FindElement(by, qstr, root)
			if err != nil {
				t.Fatalf("cannot find by %s: %s", by, err)
			}
			if el == nil {
				t.Error("empty element")
			}
		}
	}

	ok := true
	ok = ok && t.Run("by-id", f(marionette.ID, "text", nil))
	ok = ok && t.Run("by-selector", f(marionette.Selector, "div", nil))
	ok = ok && t.Run("none", func(t *testing.T) {
		el, err := tc.FindElement(marionette.ID, "test", nil)
		if el != nil {
			t.Errorf("non-empty el: %s", el)
		}

		if err == nil {
			t.Fatal("why no error?")
		}

		e, ok := err.(*marionette.ErrDriver)
		if !ok {
			t.Fatalf("unexpected error: %s", err)
		}

		if e.Type != marionette.ErrNoSuchElement {
			t.Fatalf("unexpected error: %s", err)
		}
	})
	t.Run("nested", func(t *testing.T) {
		if !ok {
			t.SkipNow()
		}

		root, _ := tc.FindElement(marionette.Selector, "div#ctrl", nil)

		el, err := tc.FindElement(marionette.ID, "text", root)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if el == nil {
			t.Error("empty element")
		}
	})
	t.Run("nested-none", func(t *testing.T) {
		if !ok {
			t.SkipNow()
		}
		root, _ := tc.FindElement(marionette.Selector, "#text", nil)

		el, err := tc.FindElement(marionette.ID, "ctrl", root)
		if el != nil {
			t.Errorf("non-empty el: %s", el)
		}

		if err == nil {
			t.Fatal("why no error?")
		}

		e, ok := err.(*marionette.ErrDriver)
		if !ok {
			t.Fatalf("unexpected error: %s", err)
		}

		if e.Type != marionette.ErrNoSuchElement {
			t.Fatalf("unexpected error: %s", err)
		}
	})
}

func (tc *cmdrTestCase) testFindElements(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))

	f := func(want int, by marionette.FindStrategy, qstr string, root *marionette.WebElement) func(*testing.T) {
		return func(t *testing.T) {
			el, err := tc.FindElements(by, qstr, root)
			if err != nil {
				t.Fatalf("cannot find by %s: %s", by, err)
			}
			if l := len(el); l != want {
				t.Fatalf("expected %d elements, got %d", want, l)
			}
		}
	}

	t.Run("by-id", f(1, marionette.ID, "text", nil))
	t.Run("by-selector", f(2, marionette.Selector, "div", nil))
	t.Run("none", func(t *testing.T) {
		el, err := tc.FindElements(marionette.ID, "test", nil)
		if len(el) != 0 {
			t.Errorf("non-empty el: %s", el)
		}

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	})
}

func (tc *cmdrTestCase) testGetElementAttribute(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))

	el, _ := tc.FindElement(marionette.Selector, "#text", nil)

	val, err := tc.GetElementAttribute(el, "id")
	if err != nil {
		t.Fatalf("cannot get attribute: %s", err)
	}
	if val != "text" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testGetElementCSSValue(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))
	el, _ := tc.FindElement(marionette.ID, "ctrl", nil)

	val, err := tc.GetElementCSSValue(el, "display")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != "none" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testGetElementProperty(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))
	el, _ := tc.FindElement(marionette.Selector, "body", nil)

	val, err := tc.GetElementProperty(el, "nodeName")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	str, ok := val.(string)
	if !ok {
		t.Fatalf("unexpected data type: %t", val)
	}
	if str != "BODY" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testGetElementRect(t *testing.T) {
	tc.must(t, "element-html", tc.loadTestHTML("element.html"))
	el, _ := tc.FindElement(marionette.Selector, "body", nil)

	rect, err := tc.GetElementRect(el)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	t.Log(rect)
}
