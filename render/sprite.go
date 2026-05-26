package render

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/render/mod"
)

// A Sprite is a basic wrapper around image data and a point. The most basic Renderable.
type Sprite struct {
	LayeredPoint
	r *image.RGBA
}

// NewEmptySprite returns a sprite of the given dimensions with a blank RGBA
func NewEmptySprite(x, y float64, w, h int) *Sprite { _ = "STUB: not implemented"; return nil }

// NewSprite creates a new sprite
func NewSprite(x, y float64, r *image.RGBA) *Sprite { _ = "STUB: not implemented"; return nil }

// GetRGBA returns the rgba behind this sprite
func (s *Sprite) GetRGBA() *image.RGBA {
	_ = "STUB: not implemented"

	// GetDims returns the dimensions of this sprite, or if this sprite has no
	// defined RGBA returns default values.
	return nil
}

func (s *Sprite) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// SetRGBA will replace the rgba behind this sprite
func (s *Sprite) SetRGBA(r *image.RGBA) {
	_ = "STUB: not implemented"

	// Bounds is an alternative to GetDims that alows a sprite
	// to satisfy draw.Image.
	return
}

func (s *Sprite) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *

	// ColorModel allows sprites to satisfy draw.Image. Returns
	// color.RGBAModel.
	new(image.Rectangle)
}

func (s *Sprite) ColorModel() color.Model {
	_ = "STUB: not implemented"
	return *

	// At returns the color of a given pixel location
	new(color.Model)
}

func (s *Sprite) At(x, y int) color.Color {
	_ = "STUB: not implemented"
	return *

	// Set sets a color of a given pixel location
	new(color.Color)
}

func (s *Sprite) Set(x, y int, c color.Color) {
	_ = "STUB: not implemented"

	// Draw draws this sprite at +xOff, +yOff
	return
}

func (s *Sprite) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// Copy returns a copy of this Sprite
func (s *Sprite) Copy() Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

func rgbaCopy(r *image.RGBA) *image.RGBA { _ = "STUB: not implemented"; return nil }

// Modify takes in modifications (modify.go) and alters this sprite accordingly
func (s *Sprite) Modify(ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Filter filters this sprite's rgba on all the input filters
func (s *Sprite) Filter(fs ...mod.Filter) { _ = "STUB: not implemented"; return }

// OverlaySprites combines sprites together through masking to form a single sprite
func OverlaySprites(sps []*Sprite) *Sprite { _ = "STUB: not implemented"; return nil }
