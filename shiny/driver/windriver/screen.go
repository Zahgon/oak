// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package windriver

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/driver/internal/win32"
	"github.com/oakmound/oak/v4/shiny/screen"
)

type screenImpl struct {
	screenHWND win32.HWND
}

func newScreen(hwnd win32.HWND) *screenImpl { _ = "STUB: not implemented"; return nil }

func (*screenImpl) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
	// Buffer length must fit in BITMAPINFO.Header.SizeImage (uint32), as
	// well as in Go slice length (int). It's easiest to be consistent
	// between 32-bit and 64-bit, so we just use int32.
	return *new(screen.Image), nil
}

func (s *screenImpl) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

func (s *screenImpl) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	return *new(screen.Window), nil
}
