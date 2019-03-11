package mnclient

import marionette "github.com/raohwork/marionette-go"

type WindowResizer struct {
	*Commander
}

func (r *WindowResizer) Outer(rect marionette.Rect) (ret marionette.Rect, err error) {
	return r.SetWindowRect(rect)
}

func (r *WindowResizer) Inner(rect marionette.Rect) (ret marionette.Rect, err error) {
	ret, err = r.Outer(rect)
	if err != nil {
		return
	}

	var w, h float64
	err = r.ExecuteScript(
		`return window.innerWidth`,
		&w,
	)
	if err != nil {
		return
	}
	err = r.ExecuteScript(
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
