// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x11driver

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

	"github.com/BurntSushi/xgbutil"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/render"
	"github.com/BurntSushi/xgb/shm"
	"github.com/BurntSushi/xgb/xproto"

	"github.com/oakmound/oak/v4/shiny/driver/internal/x11key"
	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/math/f64"
)

type screenImpl struct {
	*xgbutil.XUtil
	xc      *xgb.Conn
	xsi     *xproto.ScreenInfo
	keysyms x11key.KeysymTable

	atoms      map[string]xproto.Atom
	numLockMod uint16

	pixelsPerPt  float32
	pictformat24 render.Pictformat
	pictformat32 render.Pictformat

	// window32 and its related X11 resources is an unmapped window so that we
	// have a depth-32 window to create depth-32 pixmaps from, i.e. pixmaps
	// with an alpha channel. The root window isn't guaranteed to be depth-32.
	gcontext32 xproto.Gcontext
	window32   xproto.Window

	// opaqueP is a fully opaque, solid fill picture.
	opaqueP render.Picture

	uniformMu sync.Mutex
	uniformC  render.Color
	uniformP  render.Picture

	mu              sync.Mutex
	buffers         map[shm.Seg]*bufferImpl
	uploads         map[uint16]chan struct{}
	windows         map[xproto.Window]*Window
	nPendingUploads int
	completionKeys  []uint16
}

var (
	initialAtoms = []string{
		"_NET_WM_NAME",
		"UTF8_STRING",
		"WM_DELETE_WINDOW",
		"WM_PROTOCOLS",
		"WM_TAKE_FOCUS",
		"_NET_WM_ICON",
	}
)

const (
	millimetersPerInch = 25.4
	pointsPerInch      = 72
)

func newScreenImpl(xutil *xgbutil.XUtil) (s *screenImpl, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *screenImpl) run() { _ = "STUB: not implemented"; return }

// ~~~
// Auto repeat disabling nonsense:
// Auto repeats via X come in this form:
// [real-press] [.] [.] [.] [.] [.] [.] [auto-release] [auto-press] [.] [auto-release] [auto-press] [.] [real-release]
// If thisEv here represents an auto-release, then that means
// there will be swiftly incoming another event which has the same sequence
// ID and is an auto press.
// When we PollForEvent, we don't block and will get a nil event back
// if nothing is waiting for us. So theoretically, because the auto-release
// and auto-press come in simultaneously, we just do a single Poll, check
// if its an auto press and discard them both if that's the case.
//
// Problem 1: Theoretically, another event could come in in the middle
// of the release and press-> [auto-release] [mouse-press] [auto-press]
// we need to both handle these messages and poll again after handling them
// to see if our auto-press came in yet.
//
// Problem 2: In practice, auto release and auto-press do -not- come in
// at the same time, so this goroutine could get the first event and hit
// Poll before the next one is added to the event queue. This leads to the
// addition of a sleep, to allow event originator goroutines to empty
// their queues before we check for the press. This is still imprecise,
// and requires more testing to see if it is sufficient. An overloaded
// system of goroutines could lead to this sleep not being enough. We would
// need to fork XGB to change how events are read to properly fix this in
// this manner.
//
// This approach obviously means input releases are not being as accurately
// processed as would be ideal.

// Auto repeat press/release. Skip.

// ~~~

func (s *screenImpl) handleSecondLayerEvent(ev xgb.Event) { _ = "STUB: not implemented"; return }

// A non-zero Count means that there are more expose events
// coming. For example, a non-rectangular exposure (e.g. from a
// partially overlapped window) will result in multiple expose
// events whose dirty rectangles combine to define the dirty
// region. Go's paint events do not provide dirty regions, so
// we only pass on the final X11 expose event.

// TODO: is findBuffer and the s.buffers field unused? Delete?

func (s *screenImpl) findBuffer(key shm.Seg) *bufferImpl { _ = "STUB: not implemented"; return nil }

func (s *screenImpl) findWindow(key xproto.Window) *Window { _ = "STUB: not implemented"; return nil }

// handleCompletions must only be called while holding s.mu.
func (s *screenImpl) handleCompletions() { _ = "STUB: not implemented"; return }

const (
	maxShmSide = 0x00007fff // 32,767 pixels.
	maxShmSize = 0x10000000 // 268,435,456 bytes.
)

func (s *screenImpl) NewImage(size image.Point) (retBuf screen.Image, retErr error) {
	_ = "STUB: not implemented"
	// TODO: detect if the X11 server or connection cannot support SHM pixmaps,
	// and fall back to regular pixmaps.
	return *new(screen.Image), nil
}

// No-op, but we can't take the else path because the minimum shmget
// size is 1.

// readOnly is whether the shared memory is read-only from the X11 server's
// point of view. We need false to use SHM pixmaps.

func (s *screenImpl) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

//render.SetPictureFilter(s.xc, xp, uint16(len("bilinear")), "bilinear", nil)
// The X11 server doesn't zero-initialize the pixmap. We do it ourselves.

func (s *screenImpl) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	return *new(screen.Window), nil
}

func (s *screenImpl) initKeyboardMapping() error { _ = "STUB: not implemented"; return nil }

// Figure out which modifier is the numlock modifier (see chapter 12.7 of the XLib Manual).

// XK_Num_Lock from /usr/include/X11/keysymdef.h.

func (s *screenImpl) initPictformats() error { _ = "STUB: not implemented"; return nil }

func findPictformat(fs []render.Pictforminfo, depth byte) (render.Pictformat, error) {
	_ = "STUB: not implemented"
	// This presumes little-endian BGRA.
	return *new(render.Pictformat), nil
}

func (s *screenImpl) initWindow32() error { _ = "STUB: not implemented"; return nil }

// The CwBorderPixel attribute seems necessary for depth == 32. See
// http://stackoverflow.com/questions/3645632/how-to-create-a-window-with-a-bit-depth-of-32

func findVisual(xsi *xproto.ScreenInfo, depth byte) (xproto.Visualid, error) {
	_ = "STUB: not implemented"
	return *new(xproto.Visualid), nil
}

func (s *screenImpl) setProperty(xw xproto.Window, prop xproto.Atom, values ...xproto.Atom) {
	_ = "STUB: not implemented"
	return
}

func (s *screenImpl) drawUniform(xp render.Picture, src2dst *f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// We implement draw.Src as render.PictOpOutReverse followed by
// render.PictOpOver, for the same reason as in textureImpl.draw.
