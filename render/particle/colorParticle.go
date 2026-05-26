package particle

import (
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v4/physics"
)

// A ColorParticle is a particle with a defined color and size
type ColorParticle struct {
	*baseParticle
	startColor color.Color
	endColor   color.Color
	size       float64
	endSize    float64
}

// Draw redirects to DrawOffsetGen
func (cp *ColorParticle) Draw(buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// DrawOffsetGen draws a particle with it's generator's variables
func (cp *ColorParticle) DrawOffsetGen(generator Generator, buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// Hmm. this is expensive.
// This work should be done by the Source because if the draw rate is faster
// than the enter frame rate than this is doing duplicate work
// does that mean every particle is the same struct (
//	baseParticle + image
//)
// and different particle types are just different update functions?
// -No- because we still need to keep track of variable things on these particles
// but it -does- mean that particles should track an image that they all have a function
// to create instead of these Draw functions which should just be provided by
// baseParticle

// GetLayer returns baseParticle GetLayer. This is a safety check against auto-generated
// code which would not contain the nil check here
func (cp *ColorParticle) GetLayer() int { _ = "STUB: not implemented"; return 0 }

// GetPos returns the middle of a color particle
func (cp *ColorParticle) GetPos() physics.Vector {
	_ = "STUB: not implemented"
	return *new(physics.Vector)
}

// GetDims returns the color particle's size, twice
func (cp *ColorParticle) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }
