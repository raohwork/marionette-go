package mnclient

import "testing"

func (tc *cmdrTestCase) testGetTitle(t *testing.T) {
	str, err := tc.GetTitle()
	if err != nil {
		t.Fatalf("cannot get title: %s", err)
	}
	if str != "Element Test" {
		t.Fatalf("unexpected title: %s", str)
	}
}

func (tc *cmdrTestCase) testGetCapabilities(t *testing.T) {
	cap, err := tc.GetCapabilities()
	if err != nil {
		t.Fatalf("cannot get capabilities: %s", err)
	}
	if cap == nil {
		t.Fatal("empty caps")
	}

	t.Log(cap)
}

func (tc *cmdrTestCase) testGetTimeouts(t *testing.T) {
	tim, err := tc.GetTimeouts()
	if err != nil {
		t.Fatalf("cannot get timeouts: %s", err)
	}
	if tim == nil {
		t.Fatal("empty timeouts")
	}

	t.Log(tim)
}

func (tc *cmdrTestCase) testSetTimeouts(t *testing.T) {
	tim, _ := tc.GetTimeouts()
	err := tc.SetTimeouts(tim)
	if err != nil {
		t.Fatalf("cannot set timeouts: %s", err)
	}
}

func (tc *cmdrTestCase) testGetPageSource(t *testing.T) {
	src, err := tc.GetPageSource()
	if err != nil {
		t.Fatalf("cannot get proxy: %s", err)
	}
	expected := `<html><head>
    <title>Element Test</title>
  </head>
  <body>
    <div id="hidden" style="display: none">
    </div>
    <div id="ctrl">
      try
      <input id="text" name="text" type="text" value="demo">
      <input id="checked" type="checkbox" value="check" checked="">
      <input id="unchecked" type="checkbox" value="unchecked">
      <button id="run" onclick="document.querySelector('#result').innerHTML = document.querySelector('#text').value">go</button>
    </div>
    <div id="result"></div>
    <button id="enabled">enabled</button>
    <button id="disabled" disabled="">disabled</button>
  

</body></html>`
	if src != expected {
		t.Fatalf("unexpected source: [%s]", src)
	}
}
