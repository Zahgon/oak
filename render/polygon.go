package render

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// A Polygon is a renderable that is represented by a set of in order points
// on a plane.
type Polygon struct {
	*Sprite
	floatgeom.Polygon2
}

// NewPointsPolygon is a helper function for `NewPolygon(floatgeom.NewPolygon2(p1, p2, p3, pn...))`
func NewPointsPolygon(p1, p2, p3 floatgeom.Point2, pn ...floatgeom.Point2) *Polygon {
	_ = "STUB: not implemented"
	return nil
}

// NewPolygon constructs a renderable polygon. It will display nothing until
// Fill or FillInverse is called on it.
func NewPolygon(poly floatgeom.Polygon2) *Polygon { _ = "STUB: not implemented"; return nil }

// GetOutline returns a set of lines of the given color along this polygon's outline
func (pg *Polygon) GetOutline(c color.Color) *CompositeM { _ = "STUB: not implemented"; return nil }

// GetThickOutline returns a set of lines of the given color along this polygon's outline,
// at the given thickness
func (pg *Polygon) GetThickOutline(c color.Color, thickness int) *CompositeM {
	_ = "STUB: not implemented"
	return nil
}

// GetGradientOutline returns a set of lines of the given color along this polygon's outline,
// at the given thickness, ranging from c1 to c2 in color
func (pg *Polygon) GetGradientOutline(c1, c2 color.Color, thickness int) *CompositeM {
	_ = "STUB: not implemented"
	return nil
}

// GetColoredOutline returns a set of lines of the given color along this polygon's outline
func (pg *Polygon) GetColoredOutline(colorer Colorer, thickness int) *CompositeM {
	_ = "STUB: not implemented"
	return nil
}

// FillInverse colors this polygon's exterior the given color
func (pg *Polygon) FillInverse(c color.Color) { _ = "STUB: not implemented"; return }

// Fill fills the inside of this polygon with the input color
func (pg *Polygon) Fill(c color.Color) {
	_ = "STUB: not implemented"
	// Reset the rgba of the polygon
	return
}
