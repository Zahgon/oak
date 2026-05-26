package mod

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

func HighlightOff(c color.Color, thickness, xOff, yOff int) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

func InnerHighlightOff(c color.Color, thickness, xOff, yOff int) Mod {
	_ = "STUB: not implemented"
	return *new(Mod)
}

// todo overlay instead

func InnerHighlight(c color.Color, thickness int) Mod { _ = "STUB: not implemented"; return *new(Mod) }

func Highlight(c color.Color, thickness int) Mod { _ = "STUB: not implemented"; return *new(Mod) }

type InsetFilter func(color.Color) color.Color

func Inset(fn InsetFilter, dir intgeom.Dir2) Mod { _ = "STUB: not implemented"; return *new(Mod) }

// todo: depth

// Darker produces a darker color by f percentage (0 to 1) difference
func Darker(c color.Color, f float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

// Don't touch alpha

// Lighter produces a lighter color by f percentage (0 to 1) difference
func Lighter(c color.Color, f float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

// FadeColor produces a color with more transparency by f percentage (0 to 1)
func FadeColor(c color.Color, f float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}
