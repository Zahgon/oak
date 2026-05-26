package shape

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// BezierCurve will form a Bezier on the given coordinates, expected in (x,y)
// pairs. If the inputs have an odd length, an error noting so is returned, and
// the Bezier returned is nil.
func BezierCurve(coords ...float64) (Bezier, error) {
	_ = "STUB: not implemented"
	return *new(Bezier), nil
}

// A Bezier has a function indicating how far along a curve something is given
// some float64 progress between 0 and 1. This allows points, lines, and limitlessly complex
// bezier curves to be represented under this interface.
//
// Beziers should not break if given an input outside of 0 to 1, but the results
// shouldn't be relied upon.
type Bezier interface {
	Pos(progress float64) (x, y float64)
}

// A BezierNode ties together and find points between two other Beziers
type BezierNode struct {
	Left, Right Bezier
}

// Pos returns the a point progress percent between this node's left and
// right progress percent points.
func (bn BezierNode) Pos(progress float64) (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }

// A BezierPoint covers cases where only 1 point is supplied, and serve as roots.
type BezierPoint floatgeom.Point2

// Pos returns this point.
func (bp BezierPoint) Pos(float64) (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }
