// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

package mtldriver

import (
	"image"
	"image/color"

	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/draw"
)

// textureImpl implements screen.Texture.
type textureImpl struct {
	rgba *image.RGBA
}

func (*textureImpl) Release()            { _ = "STUB: not implemented"; return }
func (t *textureImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }
func (t *textureImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (t *textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (t *textureImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
