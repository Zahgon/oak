package mod

import (
	"image"
	"image/color"
)

// A Filter modifies an input image in place. This is useful notably for modifying
// a screen buffer, as they will refuse to be modified in any other way. This cannot
// change the dimensions of the underlying image.
type Filter func(*image.RGBA)

// AndFilter combines multiple filters into one.
func AndFilter(fs ...Filter) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// ConformToPalette( is not a modification, but acts like ConformToPalette(
// without allocating a new *image.RGBA
func ConformToPalette(p color.Model) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// Fade reduces the alpha of an image. It takes an alpha from 0-255.
func Fade(alpha int) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// ApplyMask mixes the rgba values of two images, according to
// their alpha levels, and returns that as a new rgba.
func ApplyMask(img image.RGBA) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// ApplyColor mixes a color into the rgba values of an image
// and returns that new rgba.
func ApplyColor(c color.Color) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// FillMask replaces alpha 0 pixels in an RGBA with corresponding
// pixels in a second RGBA.
func FillMask(img image.RGBA) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// InPlace converts a Mod to a Filter.
func InPlace(m Mod) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// StripOuterAlpha from the image given a source image and a alpha level to denote stripping.
// Note that this was implemented for ease of implementation but not speed.
// We could use image lib or a real depth first search to do fewer checks but this is easier...
func StripOuterAlpha(m *image.RGBA, level int) Filter {
	_ = "STUB: not implemented"
	return *new(Filter)
}

//get an image

// check downwards for the given level.x

// Treating as transparent

// check left to right for the given level.

// Treating as transparent

// check bottom up for the given level.

// Treating as transparent

// check right to left for the given level.

// Treating as transparent

// There is no function to convert a Filter to a Mod, to promote not doing so.
// Mods are significantly less efficient than Filters.
