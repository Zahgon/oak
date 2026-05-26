package shape

import (
	"math"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

// In functions return whether the given coordinate lies
// in a shape.
type In func(x, y int, sizes ...int) bool

// AndIn will combine multiple In functions into one, where
// if any of the shapes are false the result is false.
func AndIn(is ...In) In { _ = "STUB: not implemented"; return *new(In) }

// OrIn will combine multiple In functions into one, where
// if any of the shapes are true the result is true.
func OrIn(is ...In) In { _ = "STUB: not implemented"; return *new(In) }

// NotIn returns the opposite of a given In function for any query.
func NotIn(i In) In { _ = "STUB: not implemented"; return *new(In) }

// A JustIn lets an In function serve as a shape by automatically
// wrapping it in assistant functions for other utilites.
type JustIn In

// In acts as the underlying In function
func (ji JustIn) In(x, y int, sizes ...int) bool { _ = "STUB: not implemented"; return false }

// Rect calls InToRect on a JustIn's In
func (ji JustIn) Rect(sizes ...int) [][]bool { _ = "STUB: not implemented"; return nil }

// Outline calls ToOutline on a JustIn
func (ji JustIn) Outline(sizes ...int) ([]intgeom.Point2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	// Square will return true for any [x][y]
	Square = JustIn(func(x, y int, sizes ...int) bool {
		return true
	})

	// Rectangle will return true for any [x][y] in w, h
	Rectangle = JustIn(func(x, y int, sizes ...int) bool {
		w := sizes[0]
		h := sizes[0]
		if len(sizes) > 1 {
			h = sizes[1]
		}
		if x < w && y < h && x >= 0 && y >= 0 {
			return true
		}
		return false
	})
	// Diamond has a shape like the following:
	// . . t . .
	// . t t t .
	// t t t t t
	// . t t t .
	// . . t . .
	Diamond = JustIn(func(x, y int, sizes ...int) bool {
		radius := sizes[0] / 2
		return math.Abs(float64(x-radius))+math.Abs(float64(y-radius)) < float64(radius)
	})
	// Circle has a shape like the following:
	// . . . . . . .
	// . . t t t . .
	// . t t t t t .
	// . t t t t t .
	// . t t t t t .
	// . . t t t . .
	// . . . . . . .
	Circle = JustIn(func(x, y int, sizes ...int) bool {
		radius := sizes[0] / 2
		dx := math.Abs(float64(x - radius))
		dy := math.Abs(float64(y - radius))
		radiusf64 := float64(radius)
		if dx+dy <= radiusf64 {
			return true
		}
		return math.Pow(dx, 2)+math.Pow(dy, 2) < math.Pow(radiusf64, 2)
	})
	// Checkered has a shape like the following:
	// t . t . t .
	// . t . t . t
	// t . t . t .
	// . t . t . t
	// t . t . t .
	// . t . t . t
	Checkered = JustIn(func(x, y int, sizes ...int) bool {
		return (x+y)%2 == 0
	})
)

// XRange is an example In utility which returns values within a given
// relative range (where 0 = 0 and 1 = size).
func XRange(a, b float64) In { _ = "STUB: not implemented"; return *new(In) }
