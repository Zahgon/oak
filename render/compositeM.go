package render

import (
	"image"
	"image/draw"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/render/mod"
)

// CompositeM Types display all of their parts at the same time,
// and respect the positions of their parts as relative to the
// position of the composite itself
type CompositeM struct {
	LayeredPoint
	rs []Modifiable
}

// NewCompositeM creates a CompositeM
func NewCompositeM(sl ...Modifiable) *CompositeM { _ = "STUB: not implemented"; return nil }

// AppendOffset adds a new offset modifiable to the CompositeM
func (cs *CompositeM) AppendOffset(r Modifiable, p floatgeom.Point2) {
	_ = "STUB: not implemented"
	return
}

// Append adds a renderable as is to the CompositeM
func (cs *CompositeM) Append(r Modifiable) { _ = "STUB: not implemented"; return }

// Prepend adds a new renderable to the front of the CompositeMR.
func (cs *CompositeM) Prepend(r Modifiable) { _ = "STUB: not implemented"; return }

// SetIndex places a renderable at a certain point in the CompositeMs renderable slice
func (cs *CompositeM) SetIndex(i int, r Modifiable) {
	_ = "STUB: not implemented"

	// Slice creates a new CompositeM as a subslice of the existing CompositeM.
	// No Modifiables will be copied, and the original will not be modified.
	return
}

func (cs *CompositeM) Slice(start, end int) *CompositeM { _ = "STUB: not implemented"; return nil }

// Len returns the number of renderables in this CompositeM.
func (cs *CompositeM) Len() int {
	_ = "STUB: not implemented"

	// AddOffset offsets all renderables in the CompositeM by a vector
	return 0
}

func (cs *CompositeM) AddOffset(i int, p floatgeom.Point2) { _ = "STUB: not implemented"; return }

// SetOffsets applies the initial offsets to the entire CompositeM
func (cs *CompositeM) SetOffsets(vs ...floatgeom.Point2) { _ = "STUB: not implemented"; return }

// Get returns a renderable at the given index within the CompositeM
func (cs *CompositeM) Get(i int) Modifiable {
	_ = "STUB: not implemented"

	// Draw draws the CompositeM with some offset from its logical position
	// (and therefore sub renderables logical positions).
	return *new(Modifiable)
}

func (cs *CompositeM) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// Undraw stops the CompositeM from being drawn
func (cs *CompositeM) Undraw() { _ = "STUB: not implemented"; return }

// GetRGBA always returns nil from Composites
func (cs *CompositeM) GetRGBA() *image.RGBA {
	_ = "STUB: not implemented"

	// Modify applies mods to the CompositeM
	return nil
}

func (cs *CompositeM) Modify(ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Filter filters each component part of this CompositeM by all of the inputs.
func (cs *CompositeM) Filter(fs ...mod.Filter) { _ = "STUB: not implemented"; return }

// ToSprite converts the composite into a sprite by drawing each layer in order
// and overwriting lower layered pixels
func (cs *CompositeM) ToSprite() *Sprite { _ = "STUB: not implemented"; return nil }

// Copy makes a new CompositeM with the same renderables
func (cs *CompositeM) Copy() Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }
