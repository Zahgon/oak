//go:build js
// +build js

package jsdriver

import "image"

type imageImpl struct {
	screen *screenImpl
	size   image.Point
	rgba   *image.RGBA
}

func (ii imageImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (ii imageImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (imageImpl) Release() { _ = "STUB: not implemented"; return }

func (ii imageImpl) RGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }
