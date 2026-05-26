// Package noop provides a nonfunctional testing driver for accessing a screen.
package noop

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/screen"
)

func Main(f func(screen.Screen)) { _ = "STUB: not implemented"; return }

type screenImpl struct{}

func (screenImpl) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
	return *new(screen.Image), nil
}

func (screenImpl) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

func (screenImpl) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	return *new(screen.Window), nil
}

type imageImpl struct {
	size image.Point
	rgba *image.RGBA
}

func (ii imageImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (ii imageImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (imageImpl) Release() { _ = "STUB: not implemented"; return }

func (ii imageImpl) RGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }

type textureImpl struct {
	size image.Point
}

func (ti textureImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (ti textureImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}
func (textureImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
func (textureImpl) Release() { _ = "STUB: not implemented"; return }

type Window struct {
	event.Deque
}

func (*Window) Release() { _ = "STUB: not implemented"; return }
func (*Window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
func (*Window) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (*Window) Publish() { _ = "STUB: not implemented"; return }
