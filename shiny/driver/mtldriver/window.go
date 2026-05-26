// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

package mtldriver

import (
	"image"

	"dmitri.shuralyov.com/gpu/mtl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/driver/internal/lifecycler"
	"github.com/oakmound/oak/v4/shiny/driver/mtldriver/internal/coreanim"
)

// Window implements screen.Window.
type Window struct {
	device mtl.Device
	window *glfw.Window
	chans  windowRequestChannels
	ml     coreanim.MetalLayer
	cq     mtl.CommandQueue

	event.Deque
	lifecycler lifecycler.State

	bgra    *BGRA
	texture mtl.Texture // Used in Publish.

	title      string
	fullscreen bool
	borderless bool

	w, h int
	x, y int
}

func (w *Window) HideCursor() error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetBorderless(borderless bool) error { _ = "STUB: not implemented"; return nil }

func (w *Window) SetFullScreen(full bool) error { _ = "STUB: not implemented"; return nil }

func (w *Window) MoveWindow(x, y, width, height int) error { _ = "STUB: not implemented"; return nil }

func (w *Window) GetCursorPosition() (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }

func (w *Window) Release() { _ = "STUB: not implemented"; return }

// Break main loop out of glfw.WaitEvents so it can receive on releaseWindowCh.

func (w *Window) SetTitle(title string) error { _ = "STUB: not implemented"; return nil }

// Break main loop out of glfw.WaitEvents so it can receive on releaseWindowCh.

type attribPair struct {
	key glfw.Hint
	val int
}

func (w *Window) SetTopMost(topMost bool) error { _ = "STUB: not implemented"; return nil }

// Break main loop out of glfw.WaitEvents so it can receive on releaseWindowCh.

// BUG: this doesn't work, and it doesn't error either
func (w *Window) SetIcon(img image.Image) error { _ = "STUB: not implemented"; return nil }

func (w *Window) NextEvent() interface{} { _ = "STUB: not implemented"; return nil }

// TODO(dmitshur): this is the best place/time/frequency to do this
//                 I've found so far, but see if it can be even better

// Set drawable size, create backing image and texture.

func (w *Window) Publish() {
	_ = "STUB: not implemented"
	// Copy w.rgba pixels into a texture.
	return
}

// Copy the texture into the drawable.
