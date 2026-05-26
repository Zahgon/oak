package floatgeom

import (
	"github.com/oakmound/oak/v4/alg/span"
)

// A Rect2 represents a span from one point in 2D space to another.
// If Min is less than max on any axis, it will return undefined results
// for methods.
type Rect2 struct {
	Min, Max Point2
}

// MaxDimensions reports that a Rect2 has only two dimensions of definition.
func (r Rect2) MaxDimensions() int {
	_ = "STUB: not implemented"

	// A Rect3 represents a span from one point in 3D space to another.
	// If Min is less than Max on any axis, it will return undefined results
	// for methods.
	return 0
}

type Rect3 struct {
	Min, Max Point3
}

// MaxDimensions reports that a Rect3 has only three dimensions of definition.
func (r Rect3) MaxDimensions() int {
	_ = "STUB: not implemented"

	// NewRect2 returns an (X,Y):(X2,Y2) rectangle. This enforces that
	// x < x2 and y < y2, and will swap the inputs if that is not true.
	// If that enforcement is not desired, construct the struct manually.
	return 0
}

func NewRect2(x, y, x2, y2 float64) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// NewRect2WH returns an (X,Y):(X+W,Y+H) rectangle. This enforces that
// w and h are positive, and will decrease x and y respectively if that is not true.
func NewRect2WH(x, y, w, h float64) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// NewBoundingRect2 will produce the minimal rectangle that contains all of
// the input points.
func NewBoundingRect2(pts ...Point2) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// NewRect3 returns an (X,Y,Z):(X2,Y2,Z2) rectangle. This enforces that
// x < x2, y < y2, and z < z2, and will swap the inputs if that is not true.
func NewRect3(x, y, z, x2, y2, z2 float64) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// NewRect3WH returns an (X,Y,Z):(X+W,Y+H,Z+D) rectangle. This enforces that
// w, h, and d and positive, and will decrease x, y, and z respectively if that
// is not true.
func NewRect3WH(x, y, z, w, h, d float64) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// NewBoundingRect3 will produce the minimal rectangle that contains all of
// the input points.
func NewBoundingRect3(pts ...Point3) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// Shift moves the rectangle by a point returns a new instance
func (r Rect2) Shift(p Point2) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// Shift moves the rectangle by a point returns a new instance
func (r Rect3) Shift(p Point3) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// Area returns W * H.
func (r Rect2) Area() float64 { _ = "STUB: not implemented"; return 0 }

// Span returns the span on this rectangle's ith axis.
func (r Rect2) Span(i int) float64 { _ = "STUB: not implemented"; return 0 }

// W returns the width of this rectangle.
func (r Rect2) W() float64 {
	_ = "STUB: not implemented"

	// H returns the height of this rectangle.
	return 0
}

func (r Rect2) H() float64 {
	_ = "STUB: not implemented"

	// Space returns W * H * D
	return 0
}

func (r Rect3) Space() float64 { _ = "STUB: not implemented"; return 0 }

// Span returns the span on this rectangle's ith axis.
func (r Rect3) Span(i int) float64 { _ = "STUB: not implemented"; return 0 }

// W returns the width of this rectangle.
func (r Rect3) W() float64 {
	_ = "STUB: not implemented"

	// H returns the height of this rectangle.
	return 0
}

func (r Rect3) H() float64 {
	_ = "STUB: not implemented"

	// D returns the depth of this rectangle.
	return 0
}

func (r Rect3) D() float64 {
	_ = "STUB: not implemented"

	// Midpoint returns the midpoint of this rectangle's span over a given dimension.
	return 0
}

func (r Rect2) Midpoint(i int) float64 { _ = "STUB: not implemented"; return 0 }

// Midpoint returns the midpoint of this rectangle's span over a given dimension.
func (r Rect3) Midpoint(i int) float64 { _ = "STUB: not implemented"; return 0 }

// Center returns the center of this rectangle
func (r Rect2) Center() Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Center returns the center of this rectangle
func (r Rect3) Center() Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Perimeter computes the sum of the edge lengths of a rectangle.
func (r Rect2) Perimeter() float64 {
	_ = "STUB: not implemented"
	// The number of edges in an n-dimensional rectangle is n * 2^(n-1)
	// (http://en.wikipedia.org/wiki/Hypercube_graph).  Thus the number
	// of edges of length (ai - bi), where the rectangle is determined
	// by p = (a1, a2, ..., an) and q = (b1, b2, ..., bn), is 2^(n-1).
	//
	// The margin of the rectangle, then, is given by the formula
	// 2^(n-1) * [(b1 - a1) + (b2 - a2) + ... + (bn - an)].
	return 0
}

// Margin computes the sum of the edge lengths of a rectangle.
func (r Rect3) Margin() float64 { _ = "STUB: not implemented"; return 0 }

// Contains tests whether p is located inside or on the boundary of r.
func (r Rect2) Contains(p Point2) bool { _ = "STUB: not implemented"; return false }

// Contains tests whether p is located inside or on the boundary of r.
func (r Rect3) Contains(p Point3) bool { _ = "STUB: not implemented"; return false }

// ContainsRect tests whether r2 is is located inside r1.
func (r Rect2) ContainsRect(r2 Rect2) bool { _ = "STUB: not implemented"; return false }

// ContainsRect tests whether r2 is is located inside r1.
func (r Rect3) ContainsRect(r2 Rect3) bool { _ = "STUB: not implemented"; return false }

// GreaterOf returns a rectangle formed of the lowest values on each
// dimension for Min, and the highest for Max.
func (r Rect2) GreaterOf(r2 Rect2) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// GreaterOf returns a rectangle formed of the lowest values on each
// dimension for Min, and the highest for Max.
func (r Rect3) GreaterOf(r2 Rect3) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// Intersects returns whether the two rectangles intersect.
func (r Rect3) Intersects(r2 Rect3) bool {
	_ = "STUB: not implemented"
	// There are four cases of overlap:
	//
	//  1. a1------------b1
	//     a2------------b2
	//     p--------q
	//
	//  2. a1------------b1
	//     a2------------b2
	//     p--------q
	//
	//  3. a1-----------------b1
	//     a2-------b2
	//     p--------q
	//
	//  4. a1-------b1
	//     a2-----------------b2
	//     p--------q
	//
	// Thus there are only two cases of non-overlap:
	//
	//  1. a1------b1
	//     a2------b2
	//
	//  2. a1------b1
	//     a2------b2
	//
	// Enforced by constructor: a1 <= b1 and a2 <= b2.  So we can just
	// check the endpoints.
	return false
}

// Intersects returns whether the two rectangles intersect.
func (r Rect2) Intersects(r2 Rect2) bool {
	_ = "STUB: not implemented"
	// There are four cases of overlap:
	//
	//  1. a1------------b1
	//     a2------------b2
	//     p--------q
	//
	//  2. a1------------b1
	//     a2------------b2
	//     p--------q
	//
	//  3. a1-----------------b1
	//     a2-------b2
	//     p--------q
	//
	//  4. a1-------b1
	//     a2-----------------b2
	//     p--------q
	//
	// Thus there are only two cases of non-overlap:
	//
	//  1. a1------b1
	//     a2------b2
	//
	//  2. a1------b1
	//     a2------b2
	//
	// Enforced by constructor: a1 <= b1 and a2 <= b2.  So we can just
	// check the endpoints.
	return false
}

// ProjectZ projects the Rect3 onto the z axis, removing it's
// z component and returning a Rect2
func (r Rect3) ProjectZ() Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// MulConst multiplies the boundary points of this rectangle by i.
func (r Rect2) MulConst(i float64) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// Poll returns a pseudorandom point from within this rectangle
func (r Rect2) Poll() Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Clamp returns a version of the provided point such that it is contained within r. If it was already contained in
// r, it will not be changed.
func (r Rect2) Clamp(pt Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Percentile returns a point within this rectangle along the vector from the top left to the bottom right of the
// rectangle, where for example, 0.0 will be r.Min, 1.0 will be r.Max, and 2.0 will be project the vector beyond r
// and return r.Min + {r.W()*2, r.H()*2}
func (r Rect2) Percentile(f float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// MulSpan returns this rectangle as a Point2 Span after multiplying the boundary points of the rectangle by f.
func (r Rect2) MulSpan(f float64) span.Span[Point2] { _ = "STUB: not implemented"; return nil }

// MulConst multiplies the boundary points of this rectangle by i.
func (r Rect3) MulConst(i float64) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// Poll returns a pseudorandom point from within this rectangle
func (r Rect3) Poll() Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Clamp returns a version of the provided point such that it is contained within r. If it was already contained in
// r, it will not be changed.
func (r Rect3) Clamp(pt Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Percentile returns a point within this rectangle along the vector from the top left to the bottom right of the
// rectangle, where for example, 0.0 will be r.Min, 1.0 will be r.Max, and 2.0 will be project the vector beyond r
// and return r.Min + {r.W()*2, r.H()*2, r.D()*2}
func (r Rect3) Percentile(f float64) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// MulConst multiplies the boundary points of this rectangle by i.
func (r Rect3) MulSpan(f float64) span.Span[Point3] { _ = "STUB: not implemented"; return nil }
