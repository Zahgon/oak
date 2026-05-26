// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package drawer provides functions that help implement screen.Drawer methods.
package drawer

import (
	"image"
	"image/draw"

	"github.com/oakmound/oak/v4/shiny/screen"
)

// Copy implements the Copy method of the screen.Drawer interface by calling
// the Draw method of that same interface.
func Copy(dst screen.SimpleDrawer, dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// Scale implements the Scale method of the screen.Drawer interface by calling
// the Draw method of that same interface.
func Scale(dst screen.SimpleDrawer, dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
