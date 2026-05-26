package debugtools

import (
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"

	"github.com/oakmound/oak/v4/collision"
)

// NewRTree creates a wrapper around a tree that supports coloring the spaces
func NewRTree(ctx *scene.Context, t *collision.Tree) *Rtree { _ = "STUB: not implemented"; return nil }

// NewThickRTree creates a wrapper around tree that colors spaces up to a thickness
func NewThickRTree(ctx *scene.Context, t *collision.Tree, thickness int) *Rtree {
	_ = "STUB: not implemented"
	return nil
}

// NewThickColoredRTree creates a wrapper around tree that colors spaces up to a thickness based on a coloring map
func NewThickColoredRTree(ctx *scene.Context, t *collision.Tree, thickness int, colorMapping map[collision.Label]color.RGBA) *Rtree {
	_ = "STUB: not implemented"
	return nil
}

// An Rtree wraps around a collision tree and can draw debug rectangles for every entity in
// the tree.
type Rtree struct {
	*collision.Tree
	Thickness int
	render.LayeredPoint
	OutlineColor color.RGBA
	ColorMap     map[collision.Label]color.RGBA
	DrawDisabled bool
	Context      *scene.Context
}

// GetDims returns the total possible area to draw this on.
func (r *Rtree) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Draw will draw the collision outlines
func (r *Rtree) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// Get all spaces on screen

// Draw spaces that are on screen (as outlines)
