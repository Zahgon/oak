package particle

import (
	"image/color"

	"github.com/oakmound/oak/v4/shape"

	"github.com/oakmound/oak/v4/alg/span"
)

// A ColorGenerator generates ColorParticles
type ColorGenerator struct {
	BaseGenerator
	StartColor, StartColorRand color.Color
	EndColor, EndColorRand     color.Color
	// The size, in pixel radius, of spawned particles
	Size    span.Span[int]
	EndSize span.Span[int]
	//
	// Some sort of particle type, for rendering triangles or squares or circles...
	Shape shape.Shape
}

// NewColorGenerator returns a new color generator with some applied options.
func NewColorGenerator(options ...func(Generator)) Generator {
	_ = "STUB: not implemented"
	return *new(Generator)
}

func (cg *ColorGenerator) setDefaults() { _ = "STUB: not implemented"; return }

// Generate creates a source using this generator
func (cg *ColorGenerator) Generate(layer int) *Source {
	_ = "STUB: not implemented"
	// Convert rotation from degrees to radians
	return nil
}

// GenerateParticle creates a particle from a generator
func (cg *ColorGenerator) GenerateParticle(bp *baseParticle) Particle {
	_ = "STUB: not implemented"
	return *new(Particle)
}

// GetParticleSize on a color generator returns that the particles
// are per-particle specifically sized
func (cg *ColorGenerator) GetParticleSize() (w float64, h float64, perParticle bool) {
	_ = "STUB: not implemented"

	// Coloration
	//
	return 0, 0, false
}

// SetStartColor lets cg have its color be set
func (cg *ColorGenerator) SetStartColor(sc, scr color.Color) { _ = "STUB: not implemented"; return }

// SetEndColor lets cg have its end color be set
func (cg *ColorGenerator) SetEndColor(ec, ecr color.Color) { _ = "STUB: not implemented"; return }

//
// Sizing
//

// A Sizeable is a generator that can have some size set to it
type Sizeable interface {
	SetSize(i span.Span[int])
	SetEndSize(i span.Span[int])
}

// Size is an option to set a Sizeable size
func Size(i span.Span[int]) func(Generator) { _ = "STUB: not implemented"; return nil }

// EndSize sets the end size of a Sizeable
func EndSize(i span.Span[int]) func(Generator) { _ = "STUB: not implemented"; return nil }

// SetSize satisfies Sizeable
func (cg *ColorGenerator) SetSize(i span.Span[int]) {
	_ = "STUB: not implemented"

	// SetEndSize stasfies Sizeable
	return
}

func (cg *ColorGenerator) SetEndSize(i span.Span[int]) {
	_ = "STUB: not implemented"

	// Shaping
	return
}

// SetShape satisfies Shapeable
func (cg *ColorGenerator) SetShape(sf shape.Shape) { _ = "STUB: not implemented"; return }
