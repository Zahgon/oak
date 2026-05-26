package render

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

// NewColorBox returns a Sprite full of a given color with the given dimensions
// Deprecated: Use NewColorboxM (for a Modifiable) or NewColorBoxR.
func NewColorBox(w, h int, c color.Color) *Sprite { _ = "STUB: not implemented"; return nil }

// NewColorBoxM returns a modifiable Color Box (as a Sprite)
func NewColorBoxM(w, h int, c color.Color) *Sprite { _ = "STUB: not implemented"; return nil }

// ColorBoxR is a renderable color box. It is a smaller structure and should
// render faster than a ColorBoxM.
type ColorBoxR struct {
	LayeredPoint
	Dims  intgeom.Point2
	Color *image.Uniform
}

// NewColorBoxR creates a color box. Colorboxes made without using this constructor
// may not function.
func NewColorBoxR(w, h int, c color.Color) *ColorBoxR { _ = "STUB: not implemented"; return nil }

// GetDims returns the dimensiosn of this colorbox
func (cb *ColorBoxR) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Draw renders this colorbox to screen.
func (cb *ColorBoxR) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }
