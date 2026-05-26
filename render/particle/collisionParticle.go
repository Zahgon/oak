package particle

import (
	"image/draw"

	"github.com/oakmound/oak/v4/collision"
)

// A CollisionParticle is a wrapper around other particles that also
// has a collision space and can functionally react with the environment
// on collision
type CollisionParticle struct {
	Particle
	s *collision.ReactiveSpace
}

// Draw redirects to DrawOffsetGen
func (cp *CollisionParticle) Draw(buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// DrawOffsetGen draws a particle with it's generator's variables
func (cp *CollisionParticle) DrawOffsetGen(generator Generator, buff draw.Image, xOff, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// Cycle updates the collision particles variables once per rotation
func (cp *CollisionParticle) Cycle(generator Generator) { _ = "STUB: not implemented"; return }

// GetDims returns the dimensions of the space of the particle
func (cp *CollisionParticle) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }
