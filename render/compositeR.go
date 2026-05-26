package render

import (
	"image"
	"image/draw"
	"sync"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/alg/intgeom"
)

// A CompositeR is equivalent to a CompositeM for Renderables instead of
// Modifiables. CompositeRs also implements Stackable.
type CompositeR struct {
	LayeredPoint
	toPush      []Renderable
	toUndraw    []Renderable
	rs          []Renderable
	predrawLock sync.Mutex
}

// NewCompositeR creates a new CompositeR from a slice of renderables
func NewCompositeR(sl ...Renderable) *CompositeR { _ = "STUB: not implemented"; return nil }

// AppendOffset adds a new renderable to CompositeR with an offset
func (cs *CompositeR) AppendOffset(r Renderable, p floatgeom.Point2) {
	_ = "STUB: not implemented"
	return
}

// AddOffset adds an offset to a given renderable of the slice
func (cs *CompositeR) AddOffset(i int, p floatgeom.Point2) { _ = "STUB: not implemented"; return }

// Append adds a new renderable to the end of the CompositeR.
func (cs *CompositeR) Append(r Renderable) { _ = "STUB: not implemented"; return }

// Prepend adds a new renderable to the front of the CompositeR.
func (cs *CompositeR) Prepend(r Renderable) { _ = "STUB: not implemented"; return }

// Len returns the number of renderables in this composite.
func (cs *CompositeR) Len() int {
	_ = "STUB: not implemented"

	// SetIndex places a renderable at a certain point in the composites renderable slice
	return 0
}

func (cs *CompositeR) SetIndex(i int, r Renderable) {
	_ = "STUB: not implemented"

	// SetOffsets sets all renderables in CompositeR to the passed in Vector positions positions
	return
}

func (cs *CompositeR) SetOffsets(ps ...floatgeom.Point2) { _ = "STUB: not implemented"; return }

// Draw Draws the CompositeR with an offset from its logical location.
func (cs *CompositeR) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// Undraw undraws the CompositeR and its consituent renderables
func (cs *CompositeR) Undraw() { _ = "STUB: not implemented"; return }

// GetRGBA always returns nil from Composites
func (cs *CompositeR) GetRGBA() *image.RGBA {
	_ = "STUB: not implemented"

	// Get returns renderable from a given index in CompositeR
	return nil
}

func (cs *CompositeR) Get(i int) Renderable {
	_ = "STUB: not implemented"

	// Add stages a renderable to be added to the Composite at the next PreDraw
	return *new(Renderable)
}

func (cs *CompositeR) Add(r Renderable, _ ...int) Renderable {
	_ = "STUB: not implemented"
	return *new(Renderable)
}

// Replace updates a renderable in the CompositeR to the new Renderable
func (cs *CompositeR) Replace(old, new Renderable, i int) { _ = "STUB: not implemented"; return }

// PreDraw updates the CompositeR with the new renderables to add.
// This helps keep consistency and mitigates the threat of unsafe operations.
func (cs *CompositeR) PreDraw() { _ = "STUB: not implemented"; return }

// Copy returns a new composite with the same length slice of renderables but no actual renderables...
// CompositeRs cannot have their internal elements copied,
// as renderables cannot be copied.
func (cs *CompositeR) Copy() Stackable { _ = "STUB: not implemented"; return *new(Stackable) }

// DrawToScreen draws the elements in this composite to the given screen image.
func (cs *CompositeR) DrawToScreen(world draw.Image, viewPos *intgeom.Point2, screenW, screenH int) {
	_ = "STUB: not implemented"
	return
}

// Clear resets a composite to be empty.
func (cs *CompositeR) Clear() { _ = "STUB: not implemented"; return }

// ToSprite converts the composite into a sprite by drawing each layer in order
// and overwriting lower layered pixels
func (cs *CompositeR) ToSprite() *Sprite { _ = "STUB: not implemented"; return nil }
