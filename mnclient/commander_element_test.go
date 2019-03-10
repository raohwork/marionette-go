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

	marionette "github.com/raohwork/marionette-go"
)

func (tc *cmdrTestCase) testGetActiveElement(t *testing.T) {
	el, err := tc.GetActiveElement()
	if err != nil {
		t.Fatalf("cannot get active element: %s", err)
	}
	if el == nil {
		t.Fatal("empty element")
	}
}

func (tc *cmdrTestCase) testFindElement(t *testing.T) {
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
			t.Errorf("non-empty el: %+v", el)
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
			t.Errorf("non-empty el: %+v", el)
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
	t.Run("by-selector", f(3, marionette.Selector, "div", nil))
	t.Run("none", func(t *testing.T) {
		el, err := tc.FindElements(marionette.ID, "test", nil)
		if len(el) != 0 {
			t.Errorf("non-empty el: %+v", el)
		}

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	})
}

func (tc *cmdrTestCase) testGetElementAttribute(t *testing.T) {
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
	el, _ := tc.FindElement(marionette.ID, "hidden", nil)

	val, err := tc.GetElementCSSValue(el, "display")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != "none" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testGetElementProperty(t *testing.T) {
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

func (tc *cmdrTestCase) testGetElementTagName(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "ctrl", nil)

	tag, err := tc.GetElementTagName(el)
	if err != nil {
		t.Fatalf("cannot get tag name: %s", err)
	}

	if tag != "div" {
		t.Fatalf("unexpected tag name: %s", tag)
	}
}

func (tc *cmdrTestCase) testGetElementRect(t *testing.T) {
	el, _ := tc.FindElement(marionette.Selector, "body", nil)

	rect, err := tc.GetElementRect(el)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	t.Log(rect)
}

func (tc *cmdrTestCase) testGetElementText(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "ctrl", nil)

	txt, err := tc.GetElementText(el)
	if err != nil {
		t.Fatalf("cannot get element text: %s", err)
	}

	if txt != "try go" {
		t.Fatalf("unexpected text: %s", txt)
	}
}

func (tc *cmdrTestCase) testElementClick(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "run", nil)
	t.Log(el)

	err := tc.ElementClick(el)
	if err != nil {
		t.Fatalf("cannot click on element: %s", err)
	}

	el, _ = tc.FindElement(marionette.ID, "result", nil)
	txt, _ := tc.GetElementText(el)
	if txt != "demo" {
		t.Fatalf("unexpected result: %s", txt)
	}
}

func (tc *cmdrTestCase) testElementClear(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "text", nil)

	err := tc.ElementClear(el)
	if err != nil {
		t.Fatalf("cannot clear element: %s", err)
	}

	val, _ := tc.GetElementPropertyStr(el, "value")
	if val != "" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testElementSendKeys(t *testing.T) {
	el, _ := tc.FindElement(marionette.ID, "text", nil)
	tc.ElementClear(el)

	err := tc.ElementSendKeys(el, "test")
	if err != nil {
		t.Fatalf("cannot send keys to element: %s", err)
	}

	val, _ := tc.GetElementPropertyStr(el, "value")
	if val != "test" {
		t.Fatalf("unexpected value: %s", val)
	}
}

func (tc *cmdrTestCase) testIsElementDisplayed(t *testing.T) {
	t.Run("displayed", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "text", nil)
		ok, err := tc.IsElementDisplayed(el)
		if err != nil {
			t.Fatalf("cannot check displayness: %s", err)
		}
		if !ok {
			t.Fatal("not displayed")
		}
	})
	t.Run("hidden", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "hidden", nil)
		ok, err := tc.IsElementDisplayed(el)
		if err != nil {
			t.Fatalf("cannot check displayness: %s", err)
		}
		if ok {
			t.Fatal("not hidden")
		}
	})
}

func (tc *cmdrTestCase) testIsElementSelected(t *testing.T) {
	t.Run("selected", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "checked", nil)
		ok, err := tc.IsElementSelected(el)
		if err != nil {
			t.Fatalf("cannot check if selected: %s", err)
		}
		if !ok {
			t.Fatal("not selected")
		}
	})
	t.Run("hidden", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "unchecked", nil)
		ok, err := tc.IsElementSelected(el)
		if err != nil {
			t.Fatalf("cannot check if selected: %s", err)
		}
		if ok {
			t.Fatal("not unselected")
		}
	})
}

func (tc *cmdrTestCase) testIsElementEnabled(t *testing.T) {
	t.Run("enabled", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "enabled", nil)
		ok, err := tc.IsElementEnabled(el)
		if err != nil {
			t.Fatalf("cannot check if enabled: %s", err)
		}
		if !ok {
			t.Fatal("not enabled")
		}
	})
	t.Run("hidden", func(t *testing.T) {
		el, _ := tc.FindElement(marionette.ID, "disabled", nil)
		ok, err := tc.IsElementEnabled(el)
		if err != nil {
			t.Fatalf("cannot check if disabled: %s", err)
		}
		if ok {
			t.Fatal("not disabled")
		}
	})
}
