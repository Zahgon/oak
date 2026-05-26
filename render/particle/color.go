package particle

import (
	"image/color"
)

// A Colorable can have colors set on it.
type Colorable interface {
	SetStartColor(color.Color, color.Color)
	SetEndColor(color.Color, color.Color)
}

// Color sets colors on a Colorable
func Color(start, startRand, end, endRand color.Color) func(Generator) {
	_ = "STUB: not implemented"
	return nil
}

// A Colorable2 can have more colors set on it
type Colorable2 interface {
	SetStartColor2(color.Color, color.Color)
	SetEndColor2(color.Color, color.Color)
}

// Color2 sets more colors on a Colorable2
func Color2(start, startRand, end, endRand color.Color) func(Generator) {
	_ = "STUB: not implemented"
	return nil
}
