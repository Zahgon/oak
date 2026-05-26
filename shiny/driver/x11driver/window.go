// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x11driver

// TODO: implement a back buffer.

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/render"
	"github.com/BurntSushi/xgb/xproto"

	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/driver/internal/lifecycler"
	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
)

type Window struct {
	s *screenImpl

	xw xproto.Window
	xg xproto.Gcontext
	xp render.Picture

	event.Deque
	xevents chan xgb.Event

	// This next group of variables are mutable, but are only modified in the
	// screenImpl.run goroutine.
	width, height uint32

	lifecycler lifecycler.State

	mu sync.Mutex

	lastMouseX, lastMouseY int16

	x, y     uint32
	released bool
}

func (w *Window) Release() { _ = "STUB: not implemented"; return }

// TODO: call w.lifecycler.SetDead and w.lifecycler.SendEvent, a la
// handling atomWMDeleteWindow?

func (w *Window) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Copy(dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Publish() {
	_ = "STUB: not implemented"
	// TODO: implement a back buffer, and copy or flip that here to the front
	// buffer.
	return
}

// This sync isn't needed to flush the outgoing X11 requests. Instead, it
// acts as a form of flow control. Outgoing requests can be quite small on
// the wire, e.g. draw this texture ID (an integer) to this rectangle (four
// more integers), but much more expensive on the server (blending a
// million source and destination pixels). Without this sync, the Go X11
// client could easily end up sending work at a faster rate than the X11
// server can serve.

func (w *Window) SetFullScreen(fullscreen bool) error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetBorderless(borderless bool) error { _ = "STUB: not implemented"; return nil }

func (w *Window) handleConfigureNotify(ev xproto.ConfigureNotifyEvent) {
	_ = "STUB: not implemented"
	// TODO: does the order of these lifecycle and size events matter? Should
	// they really be a single, atomic event?
	return
}

func (w *Window) handleExpose() { _ = "STUB: not implemented"; return }

func (w *Window) handleKey(detail xproto.Keycode, state uint16, dir key.Direction) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) handleMouse(x, y int16, b xproto.Button, state uint16, dir mouse.Direction) {
	_ = "STUB: not implemented"
	return
}

// TODO: should a mouse.Event have a separate MouseModifiers field, for
// which buttons are pressed during a mouse move?

func (w *Window) MoveWindow(x, y, width, height int) error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetTitle(title string) error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetTopMost(topMost bool) error { _ = "STUB: not implemented"; return nil }

func (w *Window) HideCursor() error {
	_ = "STUB: not implemented"
	// ask X for a pixmap id
	return nil
}

// Create a 1x1 pixmap with that pixmap id
// depth has to be 1, otherwise you get BadMatch
// the drawable has to be this root window. I don't know why.
// You can't make a pixmap with less than 1x1 dimensions.
// I don't even know if this pixmap is black or transparent

// ask X for a cursor id

// create a cursor from the pixmap with that cursor id.
// the zeros are colors (r,g,b,r,g,b) and the hotspot of the cursor (x,y)
// the second px is a mask which we ignore.

// change the cursor of the window to be the created cursor.

// free the things we created
// it does not make sense to me that we can free these, and still persist our created
// cursor, but it works

func (w *Window) SetIcon(icon image.Image) error { _ = "STUB: not implemented"; return nil }

// 4 bytes, b/g/r/a, per pixel

// prepend width and height

// 32 here is the bit size of a cardinal, which is a bgra pixel
// we divide our length by 4 because we're sending a byte slice,
// not a cardinal slice

func (w *Window) GetCursorPosition() (x, y float64) {
	_ = "STUB: not implemented"
	// it's really not easy to do this with X
	// we're just caching the last values we got
	return 0, 0
}
