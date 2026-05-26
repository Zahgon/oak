//go:build js
// +build js

package jsdriver

import (
	"syscall/js"
)

// Adapted from Mark Farnan's go-canvas library (github.com/markfarnan/go-canvas)
type Canvas2D struct {
	// DOM properties
	window js.Value
	doc    js.Value
	body   js.Value

	// Canvas properties
	canvas  js.Value
	ctx     js.Value
	imgData js.Value

	copybuff js.Value
}

func NewCanvas2d(width int, height int) *Canvas2D { _ = "STUB: not implemented"; return nil }

// TODO: screen position

// Setup the 2D Drawing context

// Note Width, then Height

// Static JS buffer for copying data out to JS. Defined once and re-used to save on un-needed allocations
