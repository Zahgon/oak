package render

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/shape"
)

// BezierLine converts a bezier into a line sprite.
func BezierLine(b shape.Bezier, c color.Color) *Sprite { _ = "STUB: not implemented"; return nil }

// BezierThickLine draws a BezierLine wrapping each colored pixel in
// a square of width and height = thickness
func BezierThickLine(b shape.Bezier, c color.Color, thickness int) *Sprite {
	_ = "STUB: not implemented"
	return nil
}

func roundToIntPoint(x, y float64) intgeom.Point2 {
	_ = "STUB: not implemented"
	return *new(intgeom.Point2)
}

func bezierDraw(b shape.Bezier, pts *[]intgeom.Point2, low, high float64, lowPt, highPt intgeom.Point2) {
	_ = "STUB: not implemented"
	return
}

// If we haven't yet added this point at this low or high value
// I.E. if the low and high values are sufficiently close together
