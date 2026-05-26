//go:build android
// +build android

package androiddriver

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/shiny/screen"
)

type textureImpl struct {
	screen *Screen
	size   image.Point
	img    *imageImpl
}

func NewTexture(s *Screen, size image.Point) *textureImpl { _ = "STUB: not implemented"; return nil }

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
