package shape

import (
	"github.com/oakmound/oak/v4/alg/intgeom"
)

// GetHoles finds sets of points which are not In this shape that
// are adjacent.
func GetHoles(sh Shape, w, h int) [][]intgeom.Point2 { _ = "STUB: not implemented"; return nil }

// GetBorderHoles finds sets of points which are not In this shape that
// are adjacent in addition to the space around the shape
// (ie points that border the shape)
func GetBorderHoles(sh Shape, w, h int) [][]intgeom.Point2 { _ = "STUB: not implemented"; return nil }

// getHoles is an internal function that finds sets of points which are not In this shape that
// are adjacent.
func getHoles(sh Shape, w, h int, includeBorder bool) [][]intgeom.Point2 {
	_ = "STUB: not implemented"
	return nil
}

// flooding is now a map of holes, points which are false
// but not on the border.

func borderPoints(w, h int) []intgeom.Point2 { _ = "STUB: not implemented"; return nil }

func bfsFlood(m map[intgeom.Point2]bool, start intgeom.Point2) []intgeom.Point2 {
	_ = "STUB: not implemented"
	return nil
}

// literally adjacent points for adjacency
