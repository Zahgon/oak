//go:build android
// +build android

package androiddriver

import (
	"image"
	"sync"

	"golang.org/x/mobile/exp/gl/glutil"
)

type imageImpl struct {
	screen   *Screen
	size     image.Point
	img      *glutil.Image
	deadLock sync.Mutex
	dead     bool
}

func (ii *imageImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (ii *imageImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (ii *imageImpl) Release() { _ = "STUB: not implemented"; return }

func (ii *imageImpl) RGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }
