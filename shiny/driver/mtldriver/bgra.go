//go:build arm64 && darwin
// +build arm64,darwin

package mtldriver

import (
	"image"
	"image/color"

	"golang.org/x/image/math/f64"
)

// This file is a copy of much of x/image/draw and draw/image
// To enable fast conversions from RGBA (which Oak uses everywhere internally)
// and BGRA (which metal refuses not to use for windows)

var _ image.Image = &BGRA{}

// BGRA is an in-memory image whose At method returns BGRA values.
type BGRA struct {
	// Pix holds the image's pixels, in B, G, R, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *BGRA) ColorModel() color.Model { _ = "STUB: not implemented"; return *new(color.Model) }

func (p *BGRA) Bounds() image.Rectangle { _ = "STUB: not implemented"; return *new(image.Rectangle) }

func (p *BGRA) At(x, y int) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (p *BGRA) RGBA64At(x, y int) color.RGBA64 {
	_ = "STUB: not implemented"
	return *new(color.RGBA64)
}

// Small cap improves performance, see https://golang.org/issue/27857

func (p *BGRA) RGBAAt(x, y int) color.RGBA { _ = "STUB: not implemented"; return *new(color.RGBA) }

// Small cap improves performance, see https://golang.org/issue/27857

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *BGRA) PixOffset(x, y int) int { _ = "STUB: not implemented"; return 0 }

func (p *BGRA) Set(x, y int, c color.Color) { _ = "STUB: not implemented"; return }

// Small cap improves performance, see https://golang.org/issue/27857

func (p *BGRA) SetRGBA64(x, y int, c color.RGBA64) { _ = "STUB: not implemented"; return }

// Small cap improves performance, see https://golang.org/issue/27857

func (p *BGRA) SetRGBA(x, y int, c color.RGBA) { _ = "STUB: not implemented"; return }

// Small cap improves performance, see https://golang.org/issue/27857

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *BGRA) SubImage(r image.Rectangle) image.Image {
	_ = "STUB: not implemented"
	return *

	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	new(image.Image)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *BGRA) Opaque() bool { _ = "STUB: not implemented"; return false }

// NewBGRA returns a new RGBA image with the given bounds.
func NewBGRA(r image.Rectangle) *BGRA { _ = "STUB: not implemented"; return nil }

// pixelBufferLength returns the length of the []uint8 typed Pix slice field
// for the NewXxx functions. Conceptually, this is just (bpp * width * height),
// but this function panics if at least one of those is negative or if the
// computation would overflow the int type.
//
// This panics instead of returning an error because of backwards
// compatibility. The NewXxx functions do not return an error.
func pixelBufferLength(bytesPerPixel int, r image.Rectangle, imageTypeName string) int {
	_ = "STUB: not implemented"
	return 0
}

// mul3NonNeg returns (x * y * z), unless at least one argument is negative or
// if the computation overflows the int type, in which case it returns -1.
func mul3NonNeg(x int, y int, z int) int { _ = "STUB: not implemented"; return 0 }

// clip clips r against each image's bounds (after translating into the
// destination image's coordinate space) and shifts the points sp and mp by
// the same amount as the change in r.Min.
func clip(dst *BGRA, r *image.Rectangle, src *image.RGBA, sp *image.Point, mask image.Image, mp *image.Point) {
	_ = "STUB: not implemented"
	return
}

type nnInterpolator struct{}

func (z nnInterpolator) Transform(dst *BGRA, s2d f64.Aff3, src *image.RGBA, sr image.Rectangle) {
	_ = "STUB: not implemented"
	// Try to simplify a Transform to a Copy.
	// if s2d[0] == 1 && s2d[1] == 0 && s2d[3] == 0 && s2d[4] == 1 {
	// 	dx := int(s2d[2])
	// 	dy := int(s2d[5])
	// 	if float64(dx) == s2d[2] && float64(dy) == s2d[5] {
	// 		Copy(dst, image.Point{X: sr.Min.X + dx, Y: sr.Min.X + dy}, src, sr, op, opts)
	// 		return
	// 	}
	// }
	return
}

// adr is the affected destination pixels.

// bias is a translation of the mapping from dst coordinates to src
// coordinates such that the latter temporarily have non-negative X
// and Y coordinates. This allows us to write int(f) instead of
// int(math.Floor(f)), since "round to zero" and "round down" are
// equivalent when f >= 0, but the former is much cheaper. The X--
// and Y-- are because the TransformLeaf methods have a "sx -= 0.5"
// adjustment.

// Make adr relative to dr.Min.

// sr is the source pixels. If it extends beyond the src bounds,
// we cannot use the type-specific fast paths, as they access
// the Pix fields directly without bounds checking.
//
// Similarly, the fast paths assume that the masks are nil.

func (nnInterpolator) transform_BGRA_RGBA_Over(dst *BGRA, dr, adr image.Rectangle, d2s *f64.Aff3, src *image.RGBA, sr image.Rectangle, bias image.Point) {
	_ = "STUB: not implemented"
	return
}

// transformRect returns a rectangle dr that contains sr transformed by s2d.
func transformRect(s2d *f64.Aff3, sr *image.Rectangle) (dr image.Rectangle) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// The +1 adjustments below are because an image.Rectangle is inclusive
// on the low end but exclusive on the high end.

func clipAffectedDestRect(adr image.Rectangle, dstMask image.Image, dstMaskP image.Point) (image.Rectangle, image.Image) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Image)
}

// TODO: clip to dstMask.Bounds() if the color model implies that out-of-bounds means 0 alpha?

func invert(m *f64.Aff3) f64.Aff3 { _ = "STUB: not implemented"; return *new(f64.Aff3) }
