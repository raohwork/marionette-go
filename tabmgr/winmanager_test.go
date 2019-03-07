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
