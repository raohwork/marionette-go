package marionette

import "encoding/json"

const (
	// mouse buttons
	beginOfValidMouseBtn = iota - 1
	MouseLeft
	MouseCenter
	MouseRight
	MouseBack
	MouseForward
	endOfValidMouseBtn
)

// ActionType denotes a virtual input device, aka inupt source
type ActionType string

const (
	// supported ActionTypes
	NoneType    ActionType = "none"
	KeyType     ActionType = "key"
	PointerType ActionType = "pointer"
)

// ActionSubType denotes possible actions
type ActionSubType string

const (
	// possible actions
	PauseAction       ActionSubType = "pause"
	KeyDownAction     ActionSubType = "keyDown"
	KeyUpAction       ActionSubType = "keyUp"
	PointerDownAction ActionSubType = "pointerDown"
	PointerUpAction   ActionSubType = "pointerUp"
	PointerMoveAction ActionSubType = "pointerMove"
	// unsupported, leave here for future usage
	PointerCancelAction ActionSubType = "pointerCancel"
)

var validSubType = map[ActionSubType]ActionType{
	PauseAction:         NoneType,
	KeyDownAction:       KeyType,
	KeyUpAction:         KeyType,
	PointerDownAction:   PointerType,
	PointerUpAction:     PointerType,
	PointerMoveAction:   PointerType,
	PointerCancelAction: PointerType,
}

// ActionSequence represents an ui action
type ActionSequence struct {
	ID          string        `json:"id"`
	Type        ActionType    `json:"type"`
	PointerType string        `json:"pointerType,omitempty"` // currently only "mouse" is supported
	Actions     []*ActionItem `json:"actions"`
}

// ActionItem represents an input event
type ActionItem struct {
	Type     ActionSubType `json:"type"`
	Value    string        `json:"value,omitempty"`
	Button   int           `json:"button,omitempty"`
	Duration int           `json:"duration,omitempty"`
	Origin   string        `json:"origin,omitempty"` // can be "viewport" or "pointer"
	X        int           `json:"x,omitempty"`
	Y        int           `json:"y,omitempty"`
}

func (i ActionItem) MarshalJSON() (ret []byte, err error) {
	data := map[string]interface{}{
		"type": i.Type,
	}

	switch i.Type {
	case PointerUpAction, PointerDownAction:
		data["button"] = i.Button
	case PointerMoveAction:
		data["x"] = int(i.X)
		data["y"] = int(i.Y)
		data["origin"] = i.Origin
		data["duration"] = i.Duration
	case KeyUpAction, KeyDownAction:
		data["value"] = i.Value
	case PauseAction:
		data["duration"] = i.Duration
	}

	return json.Marshal(data)
}

// ActionChain helps you to build list of ui actions
type ActionChain []*ActionSequence

func (b *ActionChain) prepare(sub ActionSubType) {
	typ := validSubType[sub]
	l := len(*b)
	if l > 0 && (*b)[l-1].Type == typ {
		return
	}

	seq := &ActionSequence{
		ID:      string(typ),
		Type:    typ,
		Actions: []*ActionItem{},
	}
	if typ == PointerType {
		seq.PointerType = "mouse"
	}

	*b = append(*b, seq)
}

func (b *ActionChain) addAction(act *ActionItem) (z *ActionChain) {
	cur := (*b)[len(*b)-1]
	cur.Actions = append(cur.Actions, act)
	return b
}

// Wait creates a no-op action
func (b *ActionChain) Wait(ms int) (z *ActionChain) {
	b.prepare(PauseAction)
	act := &ActionItem{
		Type: PauseAction,
	}
	if ms > 0 {
		act.Duration = ms
	}

	return b.addAction(act)
}

// MouseDown creates a PointerDown action
func (b *ActionChain) MouseDown(btn int) (z *ActionChain) {
	// invalid button, ignore
	if btn <= beginOfValidMouseBtn || btn >= endOfValidMouseBtn {
		return b
	}

	typ := PointerDownAction
	b.prepare(typ)

	act := &ActionItem{
		Type:   typ,
		Button: btn,
	}

	return b.addAction(act)
}

// MouseUp creates a PointerUp action
func (b *ActionChain) MouseUp(btn int) (z *ActionChain) {
	// invalid button, ignore
	if btn <= beginOfValidMouseBtn || btn >= endOfValidMouseBtn {
		return b
	}

	typ := PointerUpAction
	b.prepare(typ)

	act := &ActionItem{
		Type:   typ,
		Button: btn,
	}

	return b.addAction(act)
}

// MouseClick is identical to b.MouseDown(btn).MouseUp(btn)
func (b *ActionChain) MouseClick(btn int) (z *ActionChain) {
	return b.MouseDown(btn).MouseUp(btn)
}

// MouseMoveTo creates a PointerMove action relative to vieport origin
func (b *ActionChain) MouseMoveTo(x, y, duration int) (z *ActionChain) {
	if duration < 0 {
		duration = 0
	}

	typ := PointerMoveAction
	b.prepare(typ)

	act := &ActionItem{
		Type:     typ,
		X:        x,
		Y:        y,
		Duration: duration,
		Origin:   "viewport",
	}
	return b.addAction(act)
}

// MouseMoveFor creates a PointerMove action relative to current mouse position
func (b *ActionChain) MouseMoveFor(x, y, duration int) (z *ActionChain) {
	if duration < 0 {
		duration = 0
	}

	typ := PointerMoveAction
	b.prepare(typ)

	act := &ActionItem{
		Type:     typ,
		X:        x,
		Y:        y,
		Duration: duration,
		Origin:   "pointer",
	}
	return b.addAction(act)
}

// KeyDown creates a KeyDown action
func (b *ActionChain) KeyDown(key string) (z *ActionChain) {
	typ := KeyDownAction
	b.prepare(typ)
	act := &ActionItem{
		Type:  typ,
		Value: key,
	}

	return b.addAction(act)
}

// KeyUp creates a KeyUp action
func (b *ActionChain) KeyUp(key string) (z *ActionChain) {
	typ := KeyUpAction
	b.prepare(typ)
	act := &ActionItem{
		Type:  typ,
		Value: key,
	}

	return b.addAction(act)
}

// KeyPress is identical to b.KeyDown(key).KeyUp(key)
func (b *ActionChain) KeyPress(key string) (z *ActionChain) {
	return b.KeyDown(key).KeyUp(key)
}
