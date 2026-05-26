package shape

import "math"

// Eq represents a basic equation-- a mapping of x values to
// y values. This equation is expected to represent some part or all
// of a shape from -1 to 1. This range is chosen because it's often
// easier to write shape equations around the center of a graph.
type Eq func(x float64) (y float64)

// Below returns an In which reports true for all x,y coordinates below
// the equation curve.
func (eq Eq) Below() In { _ = "STUB: not implemented"; return *new(In) }

// shift from 0 to size to -1 to 1

// Above returns an In which reports true for all x,y coordinates above
// the equation curve.
func (eq Eq) Above() In { _ = "STUB: not implemented"; return *new(In) }

var (
	// Top half of heart
	hf1 Eq = func(x float64) (y float64) {
		return -2.2*math.Pow(.4+x, 2) + 1
	}
	hf2 Eq = func(x float64) (y float64) {
		return hf1(-x)
	}
	// Bottom half of heart
	hf3 Eq = func(x float64) (y float64) {
		return -math.Sqrt((x + 1)) + .2
	}
	hf4 Eq = func(x float64) (y float64) {
		return hf3(-x)
	}

	// Heart has an shape like the following:
	// . . t . t . .
	// . t t t t t .
	// t t t t t t t
	// t t t t t t t
	// . t t t t t .
	// . . t t t . .
	// . . . . . . .
	Heart = JustIn(OrIn(
		AndIn(
			XRange(0, 0.5), hf1.Below(), hf3.Above()),
		AndIn(
			XRange(0.5, 1), hf2.Below(), hf4.Above()),
	))
)
