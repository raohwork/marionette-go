package tabmgr

import "testing"

func TestNewWindowMgr3Win(t *testing.T) {
	sender, _ := connect(t)
	defer sender.Close()

	windows := map[string][]string{
		"w1": {"t1"},
		"w2": {"t1", "t2"},
		"w3": {"t1", "t2", "t3"},
	}
	wm, err := WindowManager(sender, windows)
	if err != nil {
		t.Fatalf("cannot create window manager: %s", err)
	}

	getw := func(id string) (hwin string) {
		w := wm.GetTab(id)
		if w == nil {
			t.Fatalf("cannot find %s", id)
		}
		hwin, err := w.GetChromeWindowHandle()
		if err != nil {
			t.Fatalf("cannot get window handle of %s: %s", id, err)
		}
		return
	}

	w1h := getw("w1:t1")
	w2h := getw("w2:t1")
	w3h := getw("w3:t1")

	if w1h == w2h || w2h == w3h || w2h == w1h {
		t.Fatalf("same window detected: %s, %s, %s", w1h, w2h, w3h)
	}

	for _, tn := range []string{"t2", "t3"} {
		if h := getw("w3:" + tn); h != w3h {
			t.Fatalf("tab w3:%s not with w3:t1", tn)
		}
	}

	if h := getw("w2:t2"); h != w2h {
		t.Fatal("tab w2:t2 not with w2:t1")
	}
}
