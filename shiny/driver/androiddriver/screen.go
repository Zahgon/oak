//go:build android
// +build android

// Package androiddriver provides a Android driver for accessing a screen.
package androiddriver

import (
	"image"

	"github.com/oakmound/oak/v4/shiny/driver/internal/event"
	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var _ screen.Screen = &Screen{}

type Screen struct {
	event.Deque

	app   app.App
	glctx gl.Context

	images       *glutil.Images
	activeImages []*imageImpl

	lastSz size.Event
}

func (s *Screen) NewImage(size image.Point) (screen.Image, error) {
	_ = "STUB: not implemented"
	return *new(screen.Image), nil
}

func (s *Screen) NewTexture(size image.Point) (screen.Texture, error) {
	_ = "STUB: not implemented"
	return *new(screen.Texture), nil
}

var _ screen.Window = &Screen{}

func (s *Screen) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	_ = "STUB: not implemented"
	// android does not support multiple windows
	return *new(screen.Window), nil
}

func (w *Screen) Publish() { _ = "STUB: not implemented"; return }

func (w *Screen) Release() { _ = "STUB: not implemented"; return }
func (w *Screen) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}
func (w *Screen) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}
