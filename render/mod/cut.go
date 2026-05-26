package mod

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/shape"
)

// CutRound rounds the edges of the Modifiable with Bezier curves.
func CutRound(xOff, yOff float64) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// start off as a copy

// For each corner, define directions

// X, Y, xDir, yDir

// 3 point Bezier curve

// Progressing along the curve, whenever a new y value is
// intersected at a pixel delete all values
// from the image above(or below, for negative c[3])
// that pixel

// todo: non-arbitrary progress increment

// Could only redo this loop at new y values to save time,
// but because this is currently just a pre-processing Mod
// it should be okay

// CutShape unsets pixels that are not in the provided shape.
func CutShape(sh shape.Shape) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// todo: this should not be in this package
func pointBetween(p1, p2 floatgeom.Point2, f float64) floatgeom.Point2 {
	_ = "STUB: not implemented"
	return *new(floatgeom.Point2)
}

// CutFn  can reduce or add blank space to an input image.
// Each input function decides the starting location or offset of a cut.
func CutFn(xMod, yMod, wMod, hMod func(int) int) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// CutFromLeft acts like cut but removes from the left and top rather than the right and bottom
func CutFromLeft(newWidth, newHeight int) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// CutRel acts like Cut, but takes in a multiplier on the
// existing dimensions of the image.
func CutRel(relWidth, relHeight float64) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// Cut reduces (or increases, adding nothing)
// the dimensions of the input image, setting them to newWidth and
// newHeight.
func Cut(newWidth, newHeight int) Mod { _ = "STUB: not implemented"; return *new(Mod) }
