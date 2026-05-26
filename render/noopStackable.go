package render

import (
	"image/draw"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

// NoopStackable is a Stackable element where all methods are no-ops.
// Use for tests to disable rendering.
type NoopStackable struct{}

// PreDraw on a NoopStackable does nothing.
func (ns NoopStackable) PreDraw() {
	_ = "STUB: not implemented"

	// Add on a NoopStackable does nothing. The input Renderable is still returned.
	return
}

func (ns NoopStackable) Add(r Renderable, _ ...int) Renderable {
	_ = "STUB: not implemented"

	// Replace on a NoopStackable does nothing.
	return *new(Renderable)
}

func (ns NoopStackable) Replace(Renderable, Renderable, int) {
	_ = "STUB: not implemented"

	// Copy on a NoopStackable returns itself.
	return
}

func (ns NoopStackable) Copy() Stackable {
	_ = "STUB: not implemented"

	// DrawToScreen on a NoopStackable does nothing.
	return *new(Stackable)
}

func (ns NoopStackable) DrawToScreen(draw.Image, *intgeom.Point2, int, int) {
	_ = "STUB: not implemented"

	// Clear on a NoopStackable does nothing.
	return
}

func (ns NoopStackable) Clear() { _ = "STUB: not implemented"; return }
