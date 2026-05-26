package mod

import (
	"image"
	"image/color"
)

// A Mod takes an image and returns that image transformed in some way.
type Mod func(image.Image) *image.RGBA

// A Transform is a longer name for writing Mod
type Transform = Mod

// And chains together multiple Mods into a single Mod
func And(ms ...Mod) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// SafeAnd removes any nil mods before passing the resultant set to the And function.
// It will also return a functional no-op if the mods passed in are all nil.
func SafeAnd(ms ...Mod) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// TrimColor will trim inputs so that any rows or columns where each pixel is
// less than or equal to the input color are removed. This will change the dimensions
// of the image.
func TrimColor(trimUnder color.Color) Mod { _ = "STUB: not implemented"; return *new(Mod) }

func colorLess(r, r2, g, g2, b, b2, a, a2 uint32) bool { _ = "STUB: not implemented"; return false }

// Zoom zooms into a position on the input image.
// The position is determined by the input percentages, and how far the zoom
// is deep depends on the input zoom level-- 2.0 would quarter the number of
// unique pixels from the input to the output.
func Zoom(xPerc, yPerc, zoom float64) func(rgba image.Image) *image.RGBA {
	_ = "STUB: not implemented"
	return nil
}
