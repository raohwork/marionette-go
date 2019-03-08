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

// TabCreator is easy to use tab creation toolkit
type TabCreator interface {
	// Returns mapping of chrome window handle to its tab handles.
	CurrentTabMapping() (ret map[string][]string, err error)

	// Open or close window to ensure number of windows
	//
	// Given a list of window handles (like, from GetChromeWindowHandles()) in
	// have, open/close window till the number you want.
	//
	// Passing empty array to have leads to panic. It returns empty array if
	// want = 0, but there might be one window left due to Marionette design,
	// which is mean to prevent disconnecting. Passing negative numbers to want
	// leads to panic.
	EnsureWindowNumber(have []string, want int) (windows []string, err error)

	// Open or close tabs to ensure number of tabs
	//
	// Given a list of tab handles in have, open/close tabs till the number you
	// want. It opens tabs in the window of first tab when more tabs are needed,
	// and close from the list if there's too much tabs.
	//
	// Passing empty array to have leads to panic. It returns empty array if
	// want = 0, but there might be one tab left due to Marionette design, which
	// is mean to prevent disconnecting. Passing negative numbers to want leads
	// to panic.
	EnsureTabNumber(have []string, want int) (tabs []string, err error)
}

type tabCreator struct {
	cl *Commander
}

// NewTabCreator creates a TabCreator instance
func NewTabCreator(cl *Commander) (ret TabCreator) {
	return &tabCreator{cl: cl}
}

func (c *tabCreator) CurrentTabMapping() (ret map[string][]string, err error) {
	curTabs, err := c.cl.GetWindowHandles()
	if err != nil {
		return
	}
	ret = map[string][]string{}
	for _, tab := range curTabs {
		if err = c.cl.SwitchToWindow(tab); err != nil {
			return
		}

		var win string
		win, err = c.cl.GetChromeWindowHandle()
		if err != nil {
			return
		}
		ret[win] = append(ret[win], tab)
	}

	return
}

func (c *tabCreator) EnsureTabNumber(
	have []string, want int,
) (tabs []string, err error) {
	err = c.cl.SwitchToWindow(have[0])
	if err != nil {
		return
	}
	for len(have) < want {
		var id string
		id, _, err = c.cl.NewWindow("tab", false)
		if err != nil {
			return
		}
		have = append(have, id)
	}
	for len(have) > want {
		if err = c.cl.SwitchToWindow(have[0]); err != nil {
			return
		}
		if _, err = c.cl.CloseWindow(); err != nil {
			return
		}
		have = have[1:]
	}
	tabs = have

	return
}

func (c *tabCreator) EnsureWindowNumber(
	have []string, want int,
) (windows []string, err error) {
	err = c.cl.SwitchToWindow(have[0])
	if err != nil {
		return
	}
	for len(have) < want {
		var id string

		id, _, err = c.cl.NewWindow("window", false)
		if err != nil {
			return
		}

		// the id it returned is tab id, need more work
		if err = c.cl.SwitchToWindow(id); err != nil {
			return
		}
		if id, err = c.cl.GetChromeWindowHandle(); err != nil {
			return
		}
		have = append(have, id)
	}
	for len(have) > want {
		if err = c.cl.SwitchToWindow(have[0]); err != nil {
			return
		}
		if _, err = c.cl.CloseChromeWindow(); err != nil {
			return
		}
		have = have[1:]
	}
	windows = have

	return
}
