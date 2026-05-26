//go:build js
// +build js

// Package jsdriver provides a WASM/JS driver for accessing a screen.
package jsdriver

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/screen"
)

func Main(f func(screen.Screen)) { _ = "STUB: not implemented"; return }

type screenImpl struct {
	windows []*Window
}

func (s *screenImpl) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
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
