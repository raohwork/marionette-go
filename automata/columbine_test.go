package automata

import (
	"fmt"
	"log"
	"os"
	"testing"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/shirogane"
)

var addr string

func init() {
	addr = os.Getenv("MARIONETTE_ADDR")
	if addr == "" {
		log.Fatal("You must set envvar MARIONETTE_ADDR to run tests")
	}
}

func connect(t *testing.T) (m *shirogane.Mixed, c *shirogane.Ashihana) {
	m = &shirogane.Mixed{Addr: addr}

	if err := m.Start(); err != nil {
		t.Fatalf("Unexpected error when connecting to server: %s", err)
	}
	c = &shirogane.Ashihana{Kuroga: m}
	if _, _, err := c.NewSession(); err != nil {
		t.Fatalf("Unexpected error when create new session: %s", err)
	}

	return
}

func TestColumbineInit(t *testing.T) {
	mixed, client := connect(t)
	defer mixed.Close()

	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}
	_ = c

	tabs, err := client.GetWindowHandles()
	if err != nil {
		t.Fatalf("unexpected error when checking tabs: %s", err)
	}
	if l := len(tabs); l != 5 {
		t.Fatalf("expected 5 tabs, got %d", l)
	}
}

func TestColumbineOrdered(t *testing.T) {
	mixed, _ := connect(t)
	defer mixed.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}

	js := `window.title=arguments[0];` +
		`document.querySelector('head').innerHTML=` +
		`'<title>' + arguments[0] + '</title>'`
	for _, l := range lbl {
		if err = c.GetTab(l).ExecuteScript(js, nil, l); err != nil {
			t.Fatalf("failed to set title for %s: %s", l, err)
		}
		var title string
		if title, err = c.GetTab(l).GetTitle(); err != nil {
			t.Errorf("failed to get title for %s: %s", l, err)
		}
		if title != l {
			t.Errorf("expected title to be %s, got %s", l, title)
		}
	}
}

func TestColumbineIntersect(t *testing.T) {
	mixed, _ := connect(t)
	defer mixed.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}

	js := `window.title=arguments[0];` +
		`document.querySelector('head').innerHTML=` +
		`'<title>' + arguments[0] + '</title>'`
	for _, l := range lbl {
		if err = c.GetTab(l).ExecuteScript(js, nil, l+l); err != nil {
			t.Fatalf("failed to set title for %s: %s", l, err)
		}
	}
	for _, l := range lbl {
		var title string
		if title, err = c.GetTab(l).GetTitle(); err != nil {
			t.Errorf("failed to get title for %s: %s", l, err)
		}
		if title != l+l {
			t.Errorf("expected title to be %s, got %s", l+l, title)
		}
	}
}

func TestColumbineConcurrent(t *testing.T) {
	mixed, _ := connect(t)
	defer mixed.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}

	htmlTmpl := `<html><head><title>%s</title></head><body><div>%s</div></body></html>`
	js := `const el = document.createElement('div');` +
		`el.innerText=arguments[0];` +
		`document.querySelector('body').appendChild(el);`
	for _, l := range lbl {
		t.Run("thread-"+l, func(t *testing.T) {
			tab := c.GetTab(l)

			err := tab.ExecuteScript(js, nil, l)
			if err != nil {
				t.Fatalf(
					"unexpected error when running js in tab %s: %s",
					l,
					err,
				)
			}

			src, err := tab.GetPageSource()
			if err != nil {
				t.Fatalf(
					"unexpected error when retrieving page source: %s",
					err,
				)
			}
			if s := fmt.Sprintf(htmlTmpl, l+l, l); src != s {
				t.Fatalf(
					"unexpected source for tab %s: %s", l, src,
				)
			}

		})
	}
}

func TestColumbineWaitForOK(t *testing.T) {
	mixed, _ := connect(t)
	defer mixed.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}

	tab := c.GetTab("a")
	js := `setTimeout(() => document.querySelector('div').setAttribute('class', 'test'), 3000);`

	ch := make(chan error, 1)
	go func() {
		_, err := tab.WaitFor("div.test", 5)
		ch <- err
		close(ch)
	}()

	tab.ExecuteScript(js, nil)
	if err = <-ch; err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestColumbineWaitForFail(t *testing.T) {
	mixed, _ := connect(t)
	defer mixed.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := NewColumbine(mixed, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating Columbine: %s", err)
	}

	tab := c.GetTab("b")
	_, err = tab.WaitFor("div.test", 5)
	if err == nil {
		t.Fatal("expected driver error, got nothing")
	}

	e, ok := err.(*marionette.ErrDriver)
	if !ok {
		t.Fatalf("expected driver error, got %s", err)
	}

	if e.Type != marionette.ErrNoSuchElement {
		t.Fatalf("unexpected err: %s", err)
	}
}
