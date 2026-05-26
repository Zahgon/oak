//go:build !nogift
// +build !nogift

package mod

import (
	"image"
	"image/color"

	"github.com/disintegration/gift"
)

// GiftTransform converts any set of gift.Filters into a Mod.
func GiftTransform(fs ...gift.Filter) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// GiftFilter converts any set of gift.Filters into a Filter.
// if a filter is internally a transformation in gift, this will
// not work and GiftTransform should be used instead.
func GiftFilter(fs ...gift.Filter) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// Brighten brightens an image between -100 and 100. 100 will be solid white,
// -100 will be solid black, for all colors not zero before filtering.
func Brighten(brightenBy float32) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// Saturate saturates the input between -100 and 500 percent.
func Saturate(saturateBy float32) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// ColorBalance takes in 3 numbers between -100 and 500 and applies it to the given image
func ColorBalance(r, g, b float32) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// Crop will return the given rectangle portion of transformed images. See gift.Crop
func Crop(rect image.Rectangle) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// CropToSize applies crop with an optional anchor. See gift.CropToSize
func CropToSize(width, height int, anchor gift.Anchor) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// FlipX returns a new rgba which is flipped
// over the horizontal axis.
var FlipX = GiftTransform(gift.FlipHorizontal())

// FlipY returns a new rgba which is flipped
// over the vertical axis.
var FlipY = GiftTransform(gift.FlipVertical())

// A Resampling is a strategy for image resizing.
type Resampling = gift.Resampling

// NearestNeighborResampling is a nearest neighbor resampling filter.
var NearestNeighborResampling = gift.NearestNeighborResampling

// BoxResampling is a box resampling filter (average of surrounding pixels).
var BoxResampling = gift.BoxResampling

// LinearResampling is a bilinear resampling filter.
var LinearResampling = gift.LinearResampling

// CubicResampling is a bicubic resampling filter (Catmull-Rom).
var CubicResampling = gift.CubicResampling

// LanczosResampling is a Lanczos resampling filter (3 lobes).
var LanczosResampling = gift.LanczosResampling

// Resize will transform images to match the input dimensions. See gift.Resize.
func Resize(width, height int, resampling Resampling) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// ResizeToFill will resize to fit and then crop using the given anchor. See gift.ResizeToFill.
func ResizeToFill(width, height int, resampling Resampling, anchor gift.Anchor) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// ResizeToFit will resize while preserving aspect ratio. See gift.ResizeToFit.
func ResizeToFit(width, height int, resampling Resampling) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// Rotate returns a rotated rgba.
func Rotate(degrees float32) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// RotateInterpolated acts as Rotate, but accepts an interpolation argument.
// standard rotation does this with Cubic Interpolation.
func RotateInterpolated(degrees float32, interpolation gift.Interpolation) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// RotateBackground acts as RotateInterpolated, but allows for supplying a specific
// background color to the rotation.
func RotateBackground(degrees float32, bckgrnd color.Color, interpolation gift.Interpolation) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// Rotate180 performs a specialized rotation for 180 degrees.
var Rotate180 = GiftTransform(gift.Rotate180())

// Rotate270 performs a specialized rotation for 270 degrees.
var Rotate270 = GiftTransform(gift.Rotate270())

// Rotate90 performs a specialized rotation for 360 degrees.
var Rotate90 = GiftTransform(gift.Rotate90())

// Transpose flips horizontally and rotates 90 degrees counter clockwise.
var Transpose = GiftTransform(gift.Transpose())

// Transverse flips vertically and rotates 90 degrees counter clockwise.
var Transverse = GiftTransform(gift.Transverse())

// Scale returns a scaled rgba.
func Scale(xRatio, yRatio float64) Mod { _ = "STUB: not implemented"; return *new(Mod) }
