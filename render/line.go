package render

import (
	"image"
	"image/color"
)

// Todo:
// Our current concept of thickness expands out in both directions,
// so it's impossible to draw a even-pixel thick line. This is probably
// okay for an easy syntax like this but we might want to add in a
// "Line constructor" type object like our ray-casters
// so this behavior can be customized, i.e.-- if you take thickness as
// pixel thickness, do you expand out left or right, or center, and how
// are ties broken, etc. That would also help prevent the number of
// different functions for line-drawing from continually increasing.

// NewLine returns a line from x1,y1 to x2,y2 with the given color
func NewLine(x1, y1, x2, y2 float64, c color.Color) *Sprite { _ = "STUB: not implemented"; return nil }

// NewThickLine returns a Line that has some value of thickness
func NewThickLine(x1, y1, x2, y2 float64, c color.Color, thickness int) *Sprite {
	_ = "STUB: not implemented"
	return nil
}

// NewGradientLine returns a Line that has some value of thickness along with a start and end color
func NewGradientLine(x1, y1, x2, y2 float64, c1, c2 color.Color, thickness int) *Sprite {
	_ = "STUB: not implemented"
	return nil
}

// NewColoredLine returns a line with a custom function for how each pixel in that line should be colored.
func NewColoredLine(x1, y1, x2, y2 float64, colorer Colorer, thickness int) *Sprite {
	_ = "STUB: not implemented"
	return nil

	// We subtract the minimum from each side here
	// to normalize the new line segment toward the origin
}

// DrawLine draws a line onto an image rgba from one point to another
func DrawLine(rgba *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	_ = "STUB: not implemented"
	return
}

// DrawThickLine acts like DrawlineOnto, but takes in thickness of the given line
func DrawThickLine(rgba *image.RGBA, x1, y1, x2, y2 int, c color.Color, thickness int) {
	_ = "STUB: not implemented"
	return
}

// DrawGradientLine acts like DrawThickLine but also applies a gradient to the line
func DrawGradientLine(rgba *image.RGBA, x1, y1, x2, y2 int, c1, c2 color.Color, thickness int) {
	_ = "STUB: not implemented"
	return
}

// DrawLineColored acts like DrawThickLine, but takes in a custom colorer function for how it draws its line.
func DrawLineColored(rgba *image.RGBA, x1, y1, x2, y2, thickness int, colorer Colorer) {
	_ = "STUB: not implemented"
	return
}

func drawLineBetween(x1, y1, x2, y2 int, colorer Colorer, thickness int) *image.RGBA {
	_ = "STUB: not implemented"

	// Bresenham's line-drawing algorithm from wikipedia
	return nil
}

// Todo: document why we add one here
// It has something to do with zero-height rgbas, but is always useful
