//go:build js
// +build js

package jsdriver

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/shiny/screen"
)

type textureImpl struct {
	screen *screenImpl
	size   image.Point
	rgba   *image.RGBA
}

func (ti *textureImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (ti *textureImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (ti *textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (*textureImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
func (*textureImpl) Release() { _ = "STUB: not implemented"; return }
