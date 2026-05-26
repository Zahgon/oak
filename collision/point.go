package collision

import "github.com/oakmound/oak/v4/alg/floatgeom"

// A Point is a specific point where
// collision occurred and a zone to identify
// what was collided with.
type Point struct {
	floatgeom.Point3
	Zone *Space
}

// NewPoint creates a new point
func NewPoint(s *Space, x, y float64) Point { _ = "STUB: not implemented"; return *new(Point) }

// IsNil returns whether the underlying zone of a Point is nil
func (cp Point) IsNil() bool { _ = "STUB: not implemented"; return false }
