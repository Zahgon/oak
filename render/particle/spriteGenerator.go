package particle

import (
	"github.com/oakmound/oak/v4/alg/span"

	"github.com/oakmound/oak/v4/render"
)

// A SpriteGenerator generate SpriteParticles
type SpriteGenerator struct {
	BaseGenerator
	SpriteRotation span.Span[float64]
	Base           *render.Sprite
}

// NewSpriteGenerator creates a SpriteGenerator
func NewSpriteGenerator(options ...func(Generator)) Generator {
	_ = "STUB: not implemented"
	return *new(Generator)
}

func (sg *SpriteGenerator) setDefaults() { _ = "STUB: not implemented"; return }

// Generate creates a source using this generator
func (sg *SpriteGenerator) Generate(layer int) *Source {
	_ = "STUB: not implemented"
	// Convert rotation from degrees to radians
	return nil
}

// GenerateParticle creates a particle from a generator
func (sg *SpriteGenerator) GenerateParticle(bp *baseParticle) Particle {
	_ = "STUB: not implemented"
	return *new(Particle)
}

// A Sprited can have a sprite set to it
type Sprited interface {
	SetSprite(*render.Sprite)
	SetSpriteRotation(f span.Span[float64])
}

// Sprite sets a Sprited's sprite
func Sprite(s *render.Sprite) func(Generator) { _ = "STUB: not implemented"; return nil }

// SetSprite is the function on a sprite generator that satisfies
// Sprited
func (sg *SpriteGenerator) SetSprite(s *render.Sprite) {
	_ = "STUB: not implemented"

	// SpriteRotation sets a Sprited's rotation
	return
}

func SpriteRotation(f span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// SetSpriteRotation satisfied Sprited for SpriteGenerators
func (sg *SpriteGenerator) SetSpriteRotation(f span.Span[float64]) {
	_ = "STUB: not implemented"
	return

	// GetParticleSize returns the size of the sprite that the generator generates
}

func (sg *SpriteGenerator) GetParticleSize() (w float64, h float64, perParticle bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}
