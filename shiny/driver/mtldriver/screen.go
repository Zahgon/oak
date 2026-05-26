// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

package mtldriver

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/screen"
)

// screenImpl implements screen.Screen.
type screenImpl struct {
	newWindowCh chan newWindowReq
}

func (*screenImpl) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
	return *new(screen.Image), nil
}

func (*screenImpl) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

func (s *screenImpl) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	return *new(screen.Window), nil
}

// Break main loop out of glfw.WaitEvents so it can receive on newWindowCh.
