package oak

import (
	"github.com/oakmound/oak/v4/alg/intgeom"
)

type Viewport struct {
	Position       intgeom.Point2
	Bounds         intgeom.Rect2
	BoundsEnforced bool
}

// ShiftViewport shifts the viewport by x,y
func (w *Window) ShiftViewport(delta intgeom.Point2) { _ = "STUB: not implemented"; return }

// SetViewport positions the viewport to be at x,y
func (w *Window) SetViewport(pt intgeom.Point2) { _ = "STUB: not implemented"; return }

// ViewportBounds returns the boundary of this window's viewport, or the rectangle
// that the viewport is not allowed to exit as it moves around. It often represents
// the total size of the world within a given scene. If bounds are not enforced, ok will
// be false.
func (w *Window) ViewportBounds() (rect intgeom.Rect2, ok bool) {
	_ = "STUB: not implemented"
	return *new(intgeom.Rect2), false
}

// RemoveViewportBounds removes restrictions on the viewport's movement. It will not
// cause the viewport to update immediately.
func (w *Window) RemoveViewportBounds() { _ = "STUB: not implemented"; return }

// SetViewportBounds sets the minimum and maximum position of the viewport, including
// screen dimensions
func (w *Window) SetViewportBounds(rect intgeom.Rect2) { _ = "STUB: not implemented"; return }

// Viewport returns the viewport's position. Its width and height are the window's
// width and height. This position plus width/height cannot exceed ViewportBounds.
func (w *Window) Viewport() intgeom.Point2 { _ = "STUB: not implemented"; return *new(intgeom.Point2) }
