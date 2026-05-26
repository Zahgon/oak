package span

import "image/color"

type linearColor struct {
	r, g, b, a Span[uint32]
}

// NewLinearColor returns a linear color distribution between min and maxColor
func NewLinearColor(minColor, maxColor color.Color) Span[color.Color] {
	_ = "STUB: not implemented"
	return nil
}

func (l linearColor) Clamp(c color.Color) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (l linearColor) MulSpan(i float64) Span[color.Color] { _ = "STUB: not implemented"; return nil }

func (l linearColor) Poll() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (l linearColor) Percentile(f float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func rgbaFromInts(r, g, b, a uint32) color.RGBA { _ = "STUB: not implemented"; return *new(color.RGBA) }
