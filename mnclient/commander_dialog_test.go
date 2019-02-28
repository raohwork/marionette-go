package mnclient

import (
	"testing"
)

func (tc *cmdrTestCase) testAcceptAlert(t *testing.T) {
	cases := map[string]string{
		"Alert":  "alert",
		"Propmt": "prompt",
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			err := tc.ExecuteScript(c+`("test")`, nil)
			if err != nil {
				t.Fatalf("cannot open %s: %s", c, err)
			}

			if err := tc.AcceptAlert(); err != nil {
				t.Fatalf("cannot accept alert: %s", err)
			}
		})
	}
}

func (tc *cmdrTestCase) testDismissAlert(t *testing.T) {
	cases := map[string]string{
		"Alert":  "alert",
		"Propmt": "prompt",
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			err := tc.ExecuteScript(c+`("test")`, nil)
			if err != nil {
				t.Fatalf("cannot open %s: %s", c, err)
			}

			if err := tc.DismissAlert(); err != nil {
				t.Fatalf("cannot dismiss alert: %s", err)
			}
		})
	}
}

func (tc *cmdrTestCase) testGetAlertText(t *testing.T) {
	cases := map[string]string{
		"Alert":  "alert",
		"Propmt": "prompt",
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			err := tc.ExecuteScript(c+`("test")`, nil)
			if err != nil {
				t.Fatalf("cannot open %s: %s", c, err)
			}
			defer tc.DismissAlert()

			str, err := tc.GetAlertText()
			if err != nil {
				t.Fatalf("cannot get alert text: %s", err)
			}
			if str != "test" {
				t.Fatalf("unexpected alert text: %s", str)
			}
		})
	}
}

func (tc *cmdrTestCase) testSendAlertText(t *testing.T) {
	err := tc.ExecuteScript(`window.myRes = prompt("test")`, nil)
	if err != nil {
		t.Fatalf("cannot open prompt: %s", err)
	}
	defer tc.DismissAlert()

	if err := tc.SendAlertText("myVal"); err != nil {
		t.Fatalf("cannot send text to alert: %s", err)
	}

	tc.AcceptAlert()

	var str string
	if err := tc.ExecuteScript(`return window.myRes`, &str); err != nil {
		t.Fatalf("cannot fetch result: %s", err)
	}
	if str != "myVal" {
		t.Fatalf("unexpected value: %s", str)
	}
}
