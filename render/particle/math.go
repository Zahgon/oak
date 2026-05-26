package particle

import (
	"image/color"
)

// floatFromSpread returns a random value between
// 0 and a given float64 f
func floatFromSpread(f float64) float64 { _ = "STUB: not implemented"; return 0 }

// randColor returns a random color from two arguments:
// a base color and a color representing the maximum
// potential offset for each of R,G,B, and A.
func randColor(c, ra color.Color) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

// uint16Spread returns a random uint16 between
// n-r/2 and n+r/2, not higher than 2^16-1
func uint16Spread(n, r uint32) uint16 { _ = "STUB: not implemented"; return 0 }

// uint16OnScale returns a uint16, progress % between n and endN.
// At 0 progress, endN will be returned. At 1 progress, n will be returned.
func uint16OnScale(n, endN uint32, progress float64) uint16 { _ = "STUB: not implemented"; return 0 }
