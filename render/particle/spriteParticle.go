package particle

import (
	"image/draw"
)

// A SpriteParticle is a particle that has an amount of sprite rotation
type SpriteParticle struct {
	*baseParticle
	rotation float32
}

// Draw redirects to DrawOffsetGen
func (sp *SpriteParticle) Draw(buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// DrawOffsetGen draws a particle with it's generator's variables
func (sp *SpriteParticle) DrawOffsetGen(generator Generator, buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}
