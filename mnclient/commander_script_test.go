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

import "testing"

func (tc *cmdrTestCase) testExecuteScript(t *testing.T) {
	t.Run("set-n-get", func(t *testing.T) {
		js := `window.testProp = arguments[0]; return window.testProp`
		var ret int
		err := tc.ExecuteScript(js, &ret, 2)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		if ret != 2 {
			t.Fatalf("unexpected value: %d", ret)
		}
	})

	t.Run("same-sandbox", func(t *testing.T) {
		js := `return window.testProp`
		var ret interface{}
		err := tc.ExecuteScript(js, &ret)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		f, ok := ret.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret)
		}
		if f != 2.0 {
			t.Fatalf("unexpected value: %f", f)
		}
	})
}

func (tc *cmdrTestCase) testExecuteScriptIn(t *testing.T) {
	t.Run("set-n-get", func(t *testing.T) {
		js := `window.testProp = arguments[0]; return window.testProp`
		var ret int
		err := tc.ExecuteScriptIn("box1", js, &ret, 2)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		if ret != 2 {
			t.Fatalf("unexpected value: %d", ret)
		}
	})

	t.Run("same-sandbox", func(t *testing.T) {
		js := `return window.testProp`
		var ret interface{}
		err := tc.ExecuteScriptIn("box1", js, &ret)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		f, ok := ret.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret)
		}
		if f != 2.0 {
			t.Fatalf("unexpected value: %f", f)
		}
	})

	t.Run("new-sandbox", func(t *testing.T) {
		js := `return window.testProp`
		var ret interface{}
		err := tc.ExecuteScriptIn("box2", js, &ret)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		if ret != nil {
			t.Fatalf("unexpected value: %#v", ret)
		}
	})
}

func (tc *cmdrTestCase) testExecuteAsyncScript(t *testing.T) {
	t.Run("set-n-get", func(t *testing.T) {
		js := `window.testProp = arguments[0]; setTimeout(() => arguments[1](window.testProp), 10)`
		ch, err := tc.ExecuteAsyncScript(js, 2)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		ret := <-ch
		if ret.Err != nil {
			t.Fatalf("js error: %s", ret.Err)
		}

		val, ok := ret.Result.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret.Result)
		}
		if val != 2 {
			t.Fatalf("unexpected value: %f", val)
		}
	})

	t.Run("same-sandbox", func(t *testing.T) {
		js := `setTimeout(() => arguments[0](window.testProp), 10)`
		ch, err := tc.ExecuteAsyncScript(js)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		ret := <-ch
		if ret.Err != nil {
			t.Fatalf("js error: %s", ret.Err)
		}

		val, ok := ret.Result.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret.Result)
		}
		if val != 2 {
			t.Fatalf("unexpected value: %f", val)
		}
	})
}

func (tc *cmdrTestCase) testExecuteAsyncScriptIn(t *testing.T) {
	t.Run("set-n-get", func(t *testing.T) {
		js := `window.testProp = arguments[0]; setTimeout(() => arguments[1](window.testProp), 10)`
		ch, err := tc.ExecuteAsyncScriptIn("async1", js, 2)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		ret := <-ch
		if ret.Err != nil {
			t.Fatalf("js error: %s", ret.Err)
		}

		val, ok := ret.Result.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret.Result)
		}
		if val != 2 {
			t.Fatalf("unexpected value: %f", val)
		}
	})

	t.Run("same-sandbox", func(t *testing.T) {
		js := `setTimeout(() => arguments[0](window.testProp), 10)`
		ch, err := tc.ExecuteAsyncScriptIn("async1", js)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		ret := <-ch
		if ret.Err != nil {
			t.Fatalf("js error: %s", ret.Err)
		}

		val, ok := ret.Result.(float64)
		if !ok {
			t.Fatalf("unexpected value: %#v", ret.Result)
		}
		if val != 2 {
			t.Fatalf("unexpected value: %f", val)
		}
	})

	t.Run("new-sandbox", func(t *testing.T) {
		js := `setTimeout(() => arguments[0](window.testProp), 10)`
		ch, err := tc.ExecuteAsyncScriptIn("async2", js)
		if err != nil {
			t.Fatalf("cannot exec js: %s", err)
		}
		ret := <-ch
		if ret.Err != nil {
			t.Fatalf("js error: %s", ret.Err)
		}

		if ret.Result != nil {
			t.Fatalf("unexpected value: %#v", ret.Result)
		}
	})
}
