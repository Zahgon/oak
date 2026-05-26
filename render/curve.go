package render

import (
	"image"
	"image/color"
)

// NewCircle creates a sprite and draws a circle onto it
func NewCircle(c color.Color, radius, thickness float64, offsets ...float64) *Sprite {
	_ = "STUB: not implemented"
	return nil
}

// DrawCircle draws a circle on the input rgba, of color c.
func DrawCircle(rgba *image.RGBA, c color.Color, radius, thickness float64, offsets ...float64) {
	_ = "STUB: not implemented"
	return
}

// DrawCurve draws a curve inward on the input rgba, of color c.
func DrawCurve(rgba *image.RGBA, c color.Color, radius, thickness,
	initialAngle, circlePercentage float64, offsets ...float64) {
	_ = "STUB: not implemented"
	return
}

// We add rVec to move from -1->1 to 0->2 in terms of radius scale

// this pixel is radius minus the delta, to move inward
