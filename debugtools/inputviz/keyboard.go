package inputviz

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type KeyboardLayout interface {
	KeyRect(k key.Code) floatgeom.Rect2
}

type LayoutKey interface {
	Pos() LayoutPosition
}

type LayoutPosition struct {
	Key    key.Code
	Gap    bool
	Width  float64
	Height float64
	Row    float64
	Col    float64
}

type gap float64

func (g gap) Pos() LayoutPosition { _ = "STUB: not implemented"; return *new(LayoutPosition) }

type standardKey key.Code

func (s standardKey) Pos() LayoutPosition { _ = "STUB: not implemented"; return *new(LayoutPosition) }

type wideKey struct {
	k key.Code
	w float64
}

func (w wideKey) Pos() LayoutPosition { _ = "STUB: not implemented"; return *new(LayoutPosition) }

type tallKey struct {
	k key.Code
	h float64
}

func (h tallKey) Pos() LayoutPosition { _ = "STUB: not implemented"; return *new(LayoutPosition) }

type LayoutQWERTY struct {
	Bounds    floatgeom.Rect2
	layoutMap map[key.Code]LayoutPosition
}

func (l *LayoutQWERTY) init() {
	if l.layoutMap != nil {
		return
	}
	type sk = standardKey

	l.layoutMap = make(map[key.Code]LayoutPosition)
	qwertyRows := [][]LayoutKey{
		{sk(key.Escape), gap(1), sk(key.F1), sk(key.F2), sk(key.F3), sk(key.F4), gap(.5), sk(key.F5), sk(key.F6), sk(key.F7), sk(key.F8), gap(.5), sk(key.F9), sk(key.F10), sk(key.F11), sk(key.F12), gap(2.1), sk(key.Pause)},
		{sk(key.GraveAccent), sk(key.Num1), sk(key.Num2), sk(key.Num3), sk(key.Num4), sk(key.Num5), sk(key.Num6), sk(key.Num7), sk(key.Num8), sk(key.Num9), sk(key.Num0), sk(key.HyphenMinus), sk(key.EqualSign), wideKey{key.DeleteBackspace, 2.0}, gap(.1), sk(key.Insert), sk(key.Home), sk(key.PageUp), gap(.1), sk(key.KeypadNumLock), sk(key.KeypadSlash), sk(key.KeypadAsterisk), sk(key.KeypadHyphenMinus)},
		{wideKey{key.Tab, 1.5}, sk(key.Q), sk(key.W), sk(key.E), sk(key.R), sk(key.T), sk(key.Y), sk(key.U), sk(key.I), sk(key.O), sk(key.P), sk(key.LeftSquareBracket), sk(key.RightSquareBracket), wideKey{key.Backslash, 1.5}, gap(.1), sk(key.DeleteForward), sk(key.End), sk(key.PageDown), gap(.1), sk(key.Keypad7), sk(key.Keypad8), sk(key.Keypad9), tallKey{key.KeypadPlusSign, 2}},
		{wideKey{key.CapsLock, 1.5}, sk(key.A), sk(key.S), sk(key.D), sk(key.F), sk(key.G), sk(key.H), sk(key.J), sk(key.K), sk(key.L), sk(key.Semicolon), sk(key.Apostrophe), wideKey{key.ReturnEnter, 2.5}, gap(3.2), sk(key.Keypad4), sk(key.Keypad5), sk(key.Keypad6)},
		{wideKey{key.LeftShift, 2.0}, sk(key.Z), sk(key.X), sk(key.C), sk(key.V), sk(key.B), sk(key.N), sk(key.M), sk(key.Comma), sk(key.FullStop), sk(key.Slash), wideKey{key.RightShift, 3.0}, gap(1.1), sk(key.UpArrow), gap(1.1), sk(key.Keypad1), sk(key.Keypad2), sk(key.Keypad3), tallKey{key.KeypadEnter, 2.0}},
		{wideKey{key.LeftControl, 1.5}, sk(key.LeftGUI), wideKey{key.LeftAlt, 1.5}, wideKey{key.Spacebar, 7.0}, wideKey{key.RightAlt, 1.5}, sk(key.RightGUI), wideKey{key.RightControl, 1.5}, gap(.1), sk(key.LeftArrow), sk(key.DownArrow), sk(key.RightArrow), gap(.1), wideKey{key.Keypad0, 2.0}, sk(key.KeypadFullStop)},
	}
	rowFloats := []float64{0.0, 1.1, 2.1, 3.1, 4.1, 5.1}
	for row, cols := range qwertyRows {
		rf := rowFloats[row]
		cf := 0.0
		for _, v := range cols {
			ps := v.Pos()
			if ps.Key != 0 {
				l.layoutMap[ps.Key] = LayoutPosition{
					Row:    rf,
					Col:    cf,
					Width:  ps.Width - .1,
					Height: ps.Height - .1,
				}
			}
			cf += ps.Width
		}
	}
}

func (l *LayoutQWERTY) KeyRect(k key.Code) floatgeom.Rect2 {
	_ = "STUB: not implemented"
	return *new(floatgeom.Rect2)
}

// max row = 5.1, drawn down to 6
// max col = 21.2, drawn right to 22.1

var defaultColors = map[key.Code]color.Color{}

type Keyboard struct {
	Rect      floatgeom.Rect2
	BaseLayer int
	Colors    map[key.Code]color.Color
	KeyboardLayout

	RenderCharacters bool
	Font             *render.Font

	event.CallerID
	ctx *scene.Context

	rs map[key.Code]*render.Switch

	bindings []event.Binding
}

func (k *Keyboard) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

func (k *Keyboard) RenderAndListen(ctx *scene.Context, layer int) error {
	_ = "STUB: not implemented"
	return nil
}

// Only render strings that will stay inside their boundaries

func (k *Keyboard) Destroy() { _ = "STUB: not implemented"; return }
