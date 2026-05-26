//go:build darwin && amd64
// +build darwin,amd64

package mtldriver

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

func (w *Window) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

func (w *Window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

type BGRA = image.RGBA

var NewBGRA = image.NewRGBA
