package tabmgr

import (
	"fmt"
	"os"
	"testing"
	"time"

	marionette "github.com/raohwork/marionette-go"
	"github.com/raohwork/marionette-go/mnclient"
	"github.com/raohwork/marionette-go/mnsender"
)

var addr = "127.0.0.1:2828"

func init() {
	if x := os.Getenv("MARIONETTE_ADDR"); x != "" {
		addr = x
	}
}

func connect(t *testing.T) (m mnsender.Sender, c *mnclient.Commander) {
	m, err := mnsender.NewTCPSender(addr, 0)
	if err != nil {
		t.Fatalf("Unexpected error when establishing tcp conn: %s", err)
	}

	if err := m.Start(); err != nil {
		t.Fatalf("Unexpected error when connecting to server: %s", err)
	}
	c = &mnclient.Commander{Sender: m}
	if _, _, err := c.NewSession(); err != nil {
		t.Fatalf("Unexpected error when create new session: %s", err)
	}

	return
}

func TestInit(t *testing.T) {
	sender, client := connect(t)
	defer sender.Close()

	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
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

func TestOrdered(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating TabManager: %s", err)
	}
	c.Reset("")
	defer c.Reset("about:logo")

	js := `window.title=arguments[0];` +
		`document.querySelector('head').innerHTML=` +
		`'<title>' + arguments[0] + '</title>'`
	for _, l := range lbl {
		if err = c.GetTab(l).ExecuteScript(js, nil, l); err != nil {
			t.Fatalf("failed to set title for %s: %s", l, err)
		}
		// wait some time
		time.Sleep(100 * time.Millisecond)

		var title string
		if title, err = c.GetTab(l).GetTitle(); err != nil {
			t.Errorf("failed to get title for %s: %s", l, err)
		}
		if title != l {
			t.Errorf("expected title to be %s, got %s", l, title)
		}
	}
}

func TestIntersect(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating TabManager: %s", err)
	}
	c.Reset("")
	defer c.Reset("about:logo")

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

func TestConcurrent(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating TabManager: %s", err)
	}
	c.Reset("")
	defer c.Reset("about:logo")

	htmlTmpl := `<html><head></head><body>%s</body></html>`
	js := `document.querySelector('body').innerHTML=arguments[0];`
	t.Run("group", func(t *testing.T) {
		for _, l := range lbl {
			l := l
			t.Run("thread-"+l, func(t *testing.T) {
				t.Parallel()
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
				if s := fmt.Sprintf(htmlTmpl, l); src != s {
					t.Fatalf(
						"unexpected source for tab %s: %s",
						l, src,
					)
				}

			})
		}
	})
}

func TestWaitForOK(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating TabManager: %s", err)
	}
	c.Reset("about:about")
	defer c.Reset("about:logo")

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

func TestWaitForFail(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()
	lbl := []string{"a", "b", "c", "d", "e"}
	c, err := New(sender, lbl)
	if err != nil {
		t.Fatalf("Unexpected error when creating TabManager: %s", err)
	}
	c.Reset("")
	defer c.Reset("about:logo")

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
