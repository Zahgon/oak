// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package windriver

import (
	"image"
	"image/color"
	"image/draw"
	"sync"
	"syscall"

	"github.com/oakmound/oak/v4/shiny/driver/internal/win32"
	"github.com/oakmound/oak/v4/shiny/screen"
)

type textureImpl struct {
	size   image.Point
	dc     syscall.Handle
	bitmap syscall.Handle

	mu       sync.Mutex
	released bool
}

type handleCreateTextureParams struct {
	size   image.Point
	dc     syscall.Handle
	bitmap syscall.Handle
	err    error
}

var msgCreateTexture = win32.AddScreenMsg(handleCreateTexture)

func newTexture(size image.Point, screenHWND win32.HWND) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

func handleCreateTexture(hwnd win32.HWND, uMsg uint32, wParam, lParam uintptr) {
	_ = "STUB: not implemented"
	// This code needs to run on Windows message pump thread.
	// Firstly, it calls GetDC(nil) and, according to Windows documentation
	// (https://msdn.microsoft.com/en-us/library/windows/desktop/dd144871(v=vs.85).aspx),
	// has to be released on the same thread.
	// Secondly, according to Windows documentation
	// (https://msdn.microsoft.com/en-us/library/windows/desktop/dd183489(v=vs.85).aspx),
	// ... thread that calls CreateCompatibleDC owns the HDC that is created.
	// When this thread is destroyed, the HDC is no longer valid. ...
	// So making Windows message pump thread own returned HDC makes DC
	// live as long as we want to.
	return
}

func (t *textureImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (t *textureImpl) Fill(r image.Rectangle, c color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// TODO handle error

func (t *textureImpl) Release() { _ = "STUB: not implemented"; return }

// TODO handle error

func (t *textureImpl) release() error { _ = "STUB: not implemented"; return nil }

func (t *textureImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (t *textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

// TODO handle error

// update prepares texture t for update and executes f over texture device
// context dc in a safe manner.
func (t *textureImpl) update(f func(dc syscall.Handle) error) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Select t.bitmap into t.dc, so our drawing gets recorded
// into t.bitmap and not into 1x1 default bitmap created
// during CreateCompatibleDC call.
