package particle

import (
	"image/color"
	"image/draw"
)

// A GradientParticle has a gradient from one color to another
type GradientParticle struct {
	ColorParticle
	startColor2 color.Color
	endColor2   color.Color
}

// Draw redirects to DrawOffsetGen
func (gp *GradientParticle) Draw(buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// DrawOffsetGen draws a particle with it's generator's variables
func (gp *GradientParticle) DrawOffsetGen(generator Generator, buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}
