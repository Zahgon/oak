package shape

import (
	"github.com/oakmound/oak/v4/alg/intgeom"
)

const (
	top = iota
	topright
	right
	bottomright
	bottom
	bottomleft
	left
	topleft
	lastdirection
)

var (
	xyMods = []int{
		0, -1,
		1, -1,
		1, 0,
		1, 1,
		0, 1,
		-1, 1,
		-1, 0,
		-1, -1,
	}
	pointDeltas = []int{
		1, 0,
		0, 1,
		0, 1,
		-1, 0,
		-1, 0,
		0, -1,
		0, -1,
		1, 0,
	}
)

// ToOutline returns the set of points along the input shape's outline, if
// one exists.
func ToOutline(shape Shape) func(...int) ([]intgeom.Point2, error) {
	_ = "STUB: not implemented"
	return nil
}

func parseSizes(sizes []int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// this is a hack to support 4 and 8 directional outlines
func toOutline(shape Shape, dirInc int, sizes ...int) ([]intgeom.Point2, error) {
	_ = "STUB: not implemented"
	return nil,

		//First decrement on diagonal to find start of outline
		nil
}

//Here we have found a point on the outline

func followOutline(shape Shape, dirInc, x, y, sx, sy, w, h, direction int, outline []intgeom.Point2) []intgeom.Point2 {
	_ = "STUB: not implemented"
	//Follow the outline point by point
	return nil
}

//From a point on the outline look clockwise around for next direction

// ToOutline4 returns the set of points along the input shape's outline, if
// one exists, but will move only up, left, right, or down to form this outline.
func ToOutline4(shape Shape) func(...int) ([]intgeom.Point2, error) {
	_ = "STUB: not implemented"
	return nil
}

func inOutline(s Shape, x, y, w, h int) bool { _ = "STUB: not implemented"; return false }
