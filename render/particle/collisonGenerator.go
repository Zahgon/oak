package particle

import (
	"github.com/oakmound/oak/v4/collision"
)

// A CollisionGenerator generates collision particles
type CollisionGenerator struct {
	Generator
	Fragile bool
	HitMap  map[collision.Label]collision.OnHit
}

// NewCollisionGenerator creates a new collision generator
func NewCollisionGenerator(g Generator, options ...func(*CollisionGenerator)) Generator {
	_ = "STUB: not implemented"
	return *new(Generator)
}

func (cg *CollisionGenerator) setDefault() { _ = "STUB: not implemented"; return }

// Generate creates a source using this generator
func (cg *CollisionGenerator) Generate(layer int) *Source { _ = "STUB: not implemented"; return nil }

// GenerateParticle creates a particle from a generator
func (cg *CollisionGenerator) GenerateParticle(bp *baseParticle) Particle {
	_ = "STUB: not implemented"
	return *new(Particle)
}

// GetParticleSize on a CollisionGenerator tells the caller that the particle size
// is per-particle specific
func (cg *CollisionGenerator) GetParticleSize() (w float64, h float64, perParticle bool) {
	_ = "STUB: not implemented"

	// Fragile sets whether the particles from this collisionGenerator are destroyed
	// on contact
	return 0, 0, false
}

func Fragile(f bool) func(*CollisionGenerator) { _ = "STUB: not implemented"; return nil }

// HitMap sets functions to be called when particles from this generator collide
// with other spaces
func HitMap(hm map[collision.Label]collision.OnHit) func(*CollisionGenerator) {
	_ = "STUB: not implemented"
	return nil
}
