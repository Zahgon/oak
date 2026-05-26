package grid

import (
	"github.com/oakmound/oak/v4/entities/x/btn"
)

// An Option modifies a generator prior to grid generation
type Option func(Generator) Generator

// Content sets the button option to create for each
// button at x,y coordinates on this grid. This should
// not be used with Width and Height. Options set here act
// like And() when used with Defaults.
func Content(content [][]btn.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// ContentAt sets the button uption to create for a button
// at a given x,y coordinate on this grid. If the grid has
// already had content defined by Width, Height, or Content,
// and the x,y value given would not fall on the defined grid,
// the grids dimensions will be expanded so that it will.
// Negative x or y values will result in this option having
// no effect.
// ContentAt will overwrite options set by Content. It acts
// like And() when used with Defaults.
func ContentAt(x, y int, opts ...btn.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the number of buttons vertically that this grid will make.
func Height(h int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Width sets the number of buttons horizontally that this grid will make.
func Width(w int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Defaults sets the starting option used to create buttons in this grid.
func Defaults(defaults btn.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// YGap sets the gap between buttons vertically on this grid.
func YGap(gap float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// XGap sets the gap between buttons horizontally on this grid.
func XGap(gap float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Todo: combine grids?
// So could have a grid of color definitions,
// and it with a grid of something else...
// Todo also: row or grid permutations?
// Todo also: ContentAnd
