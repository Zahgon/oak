package shape

import (
	"github.com/oakmound/oak/v4/alg/intgeom"
)

// A Rect is a function that returns a 2d boolean array
// of booleans for a given size, where true represents
// that the bounded shape contains the point [x][y].
type Rect func(sizes ...int) [][]bool

// InToRect converts an In function into a Rect function.
// Know that, if you are planning on looping over this only
// once, it's better to just use the In function. The use
// case for this is if the same size rect will be queried
// on some function multiple times, and just having the booleans
// to re-access is needed.
func InToRect(i In) Rect { _ = "STUB: not implemented"; return *new(Rect) }

// A StrictRect is a shape that ignores input width and height given to it.
type StrictRect [][]bool

// NewStrictRect returns a StrictRect with the given strict dimensions, all
// values set to false.
func NewStrictRect(w, h int) StrictRect { _ = "STUB: not implemented"; return *new(StrictRect) }

// In returns whether the input x and y are within this StrictRect's shape.
// If the shape is undefined for the input values, it returns false.
func (sr StrictRect) In(x, y int, sizes ...int) bool { _ = "STUB: not implemented"; return false }

// Outline returns this StrictRect's outline, ignoring the input dimensions.
func (sr StrictRect) Outline(sizes ...int) ([]intgeom.Point2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rect returns the StrictRect itself.
func (sr StrictRect) Rect(sizes ...int) [][]bool { _ = "STUB: not implemented"; return nil }
