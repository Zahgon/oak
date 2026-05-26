package collision

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/physics"
)

// ID Types constant
const (
	NONE = iota
	IDTypeCID
	IDTypePID
)

// A Space is a rectangle
// with a couple of ways of identifying
// an underlying object.
type Space struct {
	Location floatgeom.Rect3
	// A label can store type information.
	// Recommended to use with an enum.
	Label Label
	// A CID can be used to get the exact
	// entity which this rectangle belongs to.
	CID event.CallerID
	// Type represents which ID space the above ID
	// corresponds to.
	Type int
}

// Bounds satisfies the rtreego.Spatial interface.
func (s *Space) Bounds() floatgeom.Rect3 {
	_ = "STUB: not implemented"

	// X returns a space's x position (leftmost)
	return *new(floatgeom.Rect3)
}

func (s *Space) X() float64 { _ = "STUB: not implemented"; return 0 }

// Y returns a space's y position (upmost)
func (s *Space) Y() float64 { _ = "STUB: not implemented"; return 0 }

// GetW returns a space's width (rightmost x - leftmost x)
// Deprecated: Use W instead
func (s *Space) GetW() float64 { _ = "STUB: not implemented"; return 0 }

// GetH returns a space's height (upper y - lower y)
// Deprecated: Use H instead
func (s *Space) GetH() float64 { _ = "STUB: not implemented"; return 0 }

// W returns a space's width (rightmost x - leftmost x)
func (s *Space) W() float64 { _ = "STUB: not implemented"; return 0 }

// H returns a space's height (upper y - lower y)
func (s *Space) H() float64 { _ = "STUB: not implemented"; return 0 }

// GetCenter returns the center point of the space
func (s *Space) GetCenter() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// GetPos returns both y and x
func (s *Space) GetPos() (float64, float64) {
	_ = "STUB: not implemented"
	return 0,

		// Above returns how much above this space another space is
		// Important note: (10,10) is Above (10,20), because in oak's
		// display, lower y values are higher than higher y values.
		0
}

func (s *Space) Above(other *Space) float64 { _ = "STUB: not implemented"; return 0 }

// Below returns how much below this space another space is,
// Equivalent to -1 * Above
func (s *Space) Below(other *Space) float64 { _ = "STUB: not implemented"; return 0 }

// Contains returns whether this space contains another
func (s *Space) Contains(other *Space) bool {
	_ = "STUB: not implemented"
	// You contain another space if it is fully inside your space
	// If you are the same size and location as the space you are checking then you both contain eachother
	return false
}

// LeftOf returns how far to the left other is of this space
func (s *Space) LeftOf(other *Space) float64 { _ = "STUB: not implemented"; return 0 }

// RightOf returns how far to the right other is of this space.
// Equivalent to -1 * LeftOf
func (s *Space) RightOf(other *Space) float64 { _ = "STUB: not implemented"; return 0 }

// Overlap returns how much this space overlaps with another space
func (s *Space) Overlap(other *Space) (xOver, yOver float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// OverlapVector returns Overlap as a vector
func (s *Space) OverlapVector(other *Space) physics.Vector {
	_ = "STUB: not implemented"
	return *new(physics.Vector)
}

// SubtractRect removes a subrectangle from this rectangle and
// returns the rectangles remaining after the portion has been
// removed. The input x,y is relative to the original space:
// Example: removing 1,1 from 10,10 -> 12,12 is OK, but removing
// 11,11 from 10,10 -> 12,12 will not act as expected.
func (s *Space) SubtractRect(x2, y2, w2, h2 float64) []*Space {
	_ = "STUB: not implemented"
	return nil
}

// Left, Top, Right, Bottom
// X, Y, W, H

// Todo: these spaces overlap on the corners. We could remove that.

// NewUnassignedSpace returns a space that just has a rectangle
func NewUnassignedSpace(x, y, w, h float64) *Space { _ = "STUB: not implemented"; return nil }

// NewSpace returns a space with an associated caller id
func NewSpace(x, y, w, h float64, cID event.CallerID) *Space { _ = "STUB: not implemented"; return nil }

// NewLabeledSpace returns a space with an associated integer label
func NewLabeledSpace(x, y, w, h float64, l Label) *Space { _ = "STUB: not implemented"; return nil }

// NewFullSpace returns a space with both a label and a caller id
func NewFullSpace(x, y, w, h float64, l Label, cID event.CallerID) *Space {
	_ = "STUB: not implemented"
	return nil
}

// NewRect2Space returns a space with an associated caller id from a rect2
func NewRect2Space(rect floatgeom.Rect2, cID event.CallerID) *Space {
	_ = "STUB: not implemented"
	return nil
}

// NewRectSpace creates a colliison space with the specified 3D rectangle
func NewRectSpace(rect floatgeom.Rect3, l Label, cID event.CallerID) *Space {
	_ = "STUB: not implemented"
	return nil
}

// NewRect is a wrapper around rtreego.NewRect,
// casting the given x,y to an rtreego.Point.
// Used to not expose rtreego.Point to the user.
// Invalid widths and heights are converted to be valid.
// If zero width or height is given, it is replaced with 1.
// If a negative width or height is given, the rectangle is
// shifted to the left or up by that negative dimension and
// the dimension is made positive.
func NewRect(x, y, w, h float64) floatgeom.Rect3 {
	_ = "STUB: not implemented"
	return *new(floatgeom.Rect3)
}

// SetZLayer sets a space's z layer.
func (s *Space) SetZLayer(z float64) { _ = "STUB: not implemented"; return }
