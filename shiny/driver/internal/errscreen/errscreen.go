// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errscreen provides a stub Screen implementation.
package errscreen

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/screen"
)

// Stub returns a Screen whose methods all return the given error.
func Stub(err error) screen.Screen { _ = "STUB: not implemented"; return *new(screen.Screen) }

type stub struct {
	err error
}

func (s stub) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
	return *new(screen.Image), nil
}
func (s stub) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}
func (s stub) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	return *new(screen.Window), nil
}
