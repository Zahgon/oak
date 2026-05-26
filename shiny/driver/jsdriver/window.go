//go:build js
// +build js

package jsdriver

import (
	"image"
	"image/draw"
	"syscall/js"

	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
)

type Window struct {
	screen *screenImpl
	cvs    *Canvas2D
	event.Deque
}

func (w *Window) Release() { _ = "STUB: not implemented"; return }
func (w *Window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Publish() { _ = "STUB: not implemented"; return }

func (w *Window) sendMouseEvent(mouseEvent js.Value, dir mouse.Direction) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) sendKeyEvent(keyEvent js.Value, dir key.Direction) {
	_ = "STUB: not implemented"
	return
}

func parseKeyCode(cd string) key.Code { _ = "STUB: not implemented"; return *new(key.Code) }

// TODO: more keys
