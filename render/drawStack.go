package render

import (
	"image/draw"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

var (
	// GlobalDrawStack is the stack that all draw calls are sent through.
	GlobalDrawStack = NewDrawStack(NewDynamicHeap())
)

// The DrawStack is a stack with a safe adding mechanism that creates isolation between draw steps via predraw
type DrawStack struct {
	as     []Stackable
	toPush []Stackable
	toPop  int
}

// A Stackable can be put onto a draw stack. It usually manages how a subset of renderables
// are drawn.
type Stackable interface {
	PreDraw()
	Add(Renderable, ...int) Renderable
	Replace(Renderable, Renderable, int)
	Copy() Stackable
	DrawToScreen(draw.Image, *intgeom.Point2, int, int)
	Clear()
}

// NewDrawStack creates a DrawStack with the given stackable items, drawn in descending index order.
func NewDrawStack(stack ...Stackable) *DrawStack { _ = "STUB: not implemented"; return nil }

// SetDrawStack takes in a set of Stackables which act as the Drawstack available
// and resets how calls to Draw will act. If this is called mid scene,
// all elements on the existing draw stack will be lost.
func SetDrawStack(stackLayers ...Stackable) { _ = "STUB: not implemented"; return }

// Clear clears all stackables in a draw stack. This should revert the stack to contain
// no renderable components.
func (ds *DrawStack) Clear() { _ = "STUB: not implemented"; return }

// DrawToScreen on a stack will render its contents to the input buffer, for a screen
// of w,h dimensions, from a view point of view.
func (ds *DrawStack) DrawToScreen(world draw.Image, view *intgeom.Point2, w, h int) {
	_ = "STUB: not implemented"
	return
}

// If we had concurrent operations, we'd do it here
// in that case each draw call would return to us something
// to composite onto the window / world

// Draw adds the given renderable to the global draw stack.
//
// If the draw stack has only one stackable, the item will be added to that
// stackable with the input layers as its argument. Otherwise, the item will be added
// to the layers[0]th stackable, with remaining layers supplied to the stackable
// as arguments.
//
// If zero layers are provided, it will add to the zeroth stack layer and
// give nothing to the stackable's argument.
func Draw(r Renderable, layers ...int) (Renderable, error) {
	_ = "STUB: not implemented"
	return *new(Renderable), nil
}

// Draw adds the given renderable to the draw stack at the appropriate position based
// on the input layers. See render.Draw.
func (ds *DrawStack) Draw(r Renderable, layers ...int) (Renderable, error) {
	_ = "STUB: not implemented"
	return *new(Renderable), nil
}

// Push appends a Stackable to the draw stack during the next PreDraw.
func (ds *DrawStack) Push(a Stackable) { _ = "STUB: not implemented"; return }

// Pop pops an element from the stack at the next PreDraw call.
func (ds *DrawStack) Pop() {
	_ = "STUB: not implemented"

	// PreDraw performs whatever processes need to occur before this can be
	// drawn. In the case of the stack, it enacts previous Push and Pop calls,
	// and signals to elements on the stack to also prepare to be drawn.
	return
}

func (ds *DrawStack) PreDraw() { _ = "STUB: not implemented"; return }

// Should use two toPush lists, for this and
// draw heaps, so this call won't ever drop anything

// Copy creates a new deep copy of a Drawstack
func (ds *DrawStack) Copy() *DrawStack { _ = "STUB: not implemented"; return nil }
