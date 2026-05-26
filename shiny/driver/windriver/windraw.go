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
	"syscall"

	"github.com/oakmound/oak/v4/shiny/driver/internal/win32"
)

func mkbitmap(size image.Point) (syscall.Handle, *byte, error) {
	_ = "STUB: not implemented"
	return *new(syscall.Handle), nil, nil
}

// negative height to force top-down drawing

var blendOverFunc = _BLENDFUNCTION{
	BlendOp:             _AC_SRC_OVER,
	BlendFlags:          0,
	SourceConstantAlpha: 255,           // only use per-pixel alphas
	AlphaFormat:         _AC_SRC_ALPHA, // premultiplied
}

func copyBitmapToDC(dc win32.HDC, dr image.Rectangle, src syscall.Handle, sr image.Rectangle, op draw.Op) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// This output device does not support blending capabilities,
// so the subsequent output is incorrect, but is the best we
// can do on systems that do not support AlphaBlend.

func fill(dc win32.HDC, dr image.Rectangle, c color.Color, op draw.Op) error {
	_ = "STUB: not implemented"
	return nil
}

// AlphaBlend will stretch the input image (using StretchBlt's
// COLORONCOLOR mode) to fill the output rectangle. Testing
// this shows that the result appears to be the same as if we had
// used a MxN bitmap instead.

// TODO handle error?
