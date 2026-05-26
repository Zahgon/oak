// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package windriver

// TODO: implement a back buffer.

import (
	"image"
	"image/color"
	"image/draw"
	"sync"
	"syscall"

	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/driver/internal/win32"
	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	windowLock sync.RWMutex
	allWindows = make(map[win32.HWND]*Window)
)

type Window struct {
	hwnd win32.HWND

	changeLock sync.RWMutex

	event.Deque

	sz             size.Event
	lifecycleStage lifecycle.Stage
	// Todo: the windows api is confused about
	// whether styles are int32 or uint32s.
	style, exStyle int32
	fullscreen     bool
	borderless     bool
	maximized      bool
	topMost        bool
	windowRect     *win32.RECT
	clientRect     *win32.RECT

	// guid is set on intialization and converted to trayGUID
	// when the tray icon is created.
	guid     [16]byte
	trayGUID *win32.GUID
}

func (w *Window) Release() { _ = "STUB: not implemented"; return }

func (w *Window) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// TODO:

func (w *Window) SetTitle(title string) error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetBorderless(borderless bool) error {
	_ = "STUB: not implemented"
	// Don't set borderless if currently fullscreen.
	return nil
}

// We don't need to get these values when w.borderless is true
// because scaling is impossible without a border to grab to scale.
// Todo: except through programatic window resizing.

// The -leftOffset here is assuming that the bottom border has the same
// height as the left and right do width.

// On restore, resize to the previous saved rect size.

func (w *Window) SetFullScreen(fullscreen bool) error { _ = "STUB: not implemented"; return nil }

// Fullscreen impl copied from chromium
// https://src.chromium.org/viewvc/chrome/trunk/src/ui/views/win/fullscreen_handler.cc
// Save current window state if not already fullscreen.

// Save current window information.  We force the window into restored mode
// before going fullscreen because Windows doesn't seem to hide the
// taskbar if the window is in the maximized state.

// Set new window style and size.

// On expand, if we're given a window_rect, grow to it, otherwise do
// not resize.
// shiny cmt: Need to look into what this for_metro argument means,
// right now we don't use it
// if (!for_metro) {

// }

// Reset original window style and size.  The multiple window size/moves
// here are ugly, but if SetWindowPos() doesn't redraw, the taskbar won't be
// repainted.  Better-looking methods welcome.

// if !for_metro {
// On restore, resize to the previous saved rect size.

//}

// HideCursor turns the OS cursor into a 1x1 transparent image.
func (w *Window) HideCursor() error { _ = "STUB: not implemented"; return nil }

// SetIcon sets this window's taskbar (and top left corner) icon
func (w *Window) SetIcon(icon image.Image) error {
	_ = "STUB: not implemented"
	// windows supports four modes of setting icons:
	// 1. loading internal resources embedded into binaries in a windows-specific fashion
	// 2. loading from file
	// 3. using windows-os built in icons like question marks
	// 4. hand crafting black and white icons via combining AND and XOR masks
	//
	// note, none of these are 'use an icon held in application memory'
	//
	// 1 is not an option for a multiplatform app.
	// 3 is not an option because icons are usually not built in windows icons.
	// 4 is not an option because icons are usually colorful.
	//
	// so we're left with 2: take the image given, write it as an icon to a temporary
	// file, load that file, set it as the icon, delete that file.
	return nil
}

func isWindowsSuccessError(err error) bool { _ = "STUB: not implemented"; return false }

// we got a confusing 'this operation completed successfully'
// no, this does not actually mean the operation necessarily succeeded
// no, win32.GetLastError will not necessarily return a real error to clarify things

func (w *Window) MoveWindow(x, y, wd, ht int) error { _ = "STUB: not implemented"; return nil }

func drawWindow(dc win32.HDC, src2dst f64.Aff3, src interface{}, sr image.Rectangle, op draw.Op) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// general drawing

// copy bitmap

// scale bitmap

// TODO: check if this (and below) works when src2dst[0] < 0.

// TODO: check if this (and below) works when src2dst[4] < 0.

func (w *Window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Publish() { _ = "STUB: not implemented"; return }

func init() {
	send := func(hwnd win32.HWND, e interface{}) {
		windowLock.RLock()
		w := allWindows[hwnd]
		windowLock.RUnlock()

		w.Send(e)
	}
	win32.MouseEvent = func(hwnd win32.HWND, e mouse.Event) { send(hwnd, e) }
	win32.PaintEvent = func(hwnd win32.HWND, e paint.Event) { send(hwnd, e) }
	win32.KeyEvent = func(hwnd win32.HWND, e key.Event) { send(hwnd, e) }
	win32.LifecycleEvent = lifecycleEvent
	win32.SizeEvent = sizeEvent
}

func lifecycleEvent(hwnd win32.HWND, to lifecycle.Stage) { _ = "STUB: not implemented"; return }

func sizeEvent(hwnd win32.HWND, e size.Event) { _ = "STUB: not implemented"; return }

// cmd is used to carry parameters between user code
// and Windows message pump thread.
type cmd struct {
	id  int
	err error

	src2dst f64.Aff3
	sr      image.Rectangle
	dp      image.Point
	dr      image.Rectangle
	color   color.Color
	op      draw.Op
	texture syscall.Handle
	buffer  *bufferImpl
}

const (
	cmdDraw = iota
	cmdFill
	cmdUpload
	cmdDrawUniform
)

// msgCmd is the stored value for our handleCmd function for syscalls.
var msgCmd = win32.AddWindowMsg(handleCmd)

func (w *Window) execCmd(c *cmd) { _ = "STUB: not implemented"; return }

// TODO handle errors

func handleCmd(hwnd win32.HWND, uMsg uint32, wParam, lParam uintptr) {
	_ = "STUB: not implemented"
	return
}

// TODO: adjust if dp is outside dst bounds, or sr is outside buffer bounds.

func (w *Window) GetCursorPosition() (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }

func (w *Window) SetTopMost(topMost bool) error { _ = "STUB: not implemented"; return nil }

// Note: although you can change a window's ex style to include EX_TOPMOST
// this will not work after window creation. The following is what you need to
// do instead.

// TODO: extract and parse os error
