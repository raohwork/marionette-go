package mnclient

import marionette "github.com/raohwork/marionette-go"

// ResizerCommands defines raw commands that WindowResizer will use
type ResizerCommands interface {
	SetWindowRect(marionette.Rect) (marionette.Rect, error)
	ExecuteScript(string, interface{}, ...interface{}) error
}

// WindowResizer wraps Commander to provide easy-to-use window resizing tool
type WindowResizer struct {
	Commander ResizerCommands
}

// Outer resize window itself, identical to Commander.SetWindowRect
func (r *WindowResizer) Outer(rect marionette.Rect) (ret marionette.Rect, err error) {
	return r.Commander.SetWindowRect(rect)
}

// Inner resize content size, ensuring window.innerWidth and window.innerHeight
func (r *WindowResizer) Inner(rect marionette.Rect) (ret marionette.Rect, err error) {
	ret, err = r.Outer(rect)
	if err != nil {
		return
	}

	var w, h float64
	err = r.Commander.ExecuteScript(
		`return window.innerWidth`,
		&w,
	)
	if err != nil {
		return
	}
	err = r.Commander.ExecuteScript(
		`return window.innerHeight`,
		&h,
	)
	if err != nil {
		return
	}

	ok := w == rect.W && h == rect.H
	if !ok {
		rect.W += rect.W - w
		rect.H += rect.H - h
		ret, err = r.Outer(rect)
	}

	return
}
