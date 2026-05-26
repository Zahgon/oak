package particle

import (
	"image/color"
)

// A GradientGenerator is a ColorGenerator with a patterned gradient
// on its particles
type GradientGenerator struct {
	ColorGenerator
	StartColor2, StartColor2Rand color.Color
	EndColor2, EndColor2Rand     color.Color
	ProgressFunction             func(x, y, w, h int) float64
}

// NewGradientGenerator returns a new GradientGenerator
func NewGradientGenerator(options ...func(Generator)) Generator {
	_ = "STUB: not implemented"
	return *new(Generator)
}

func (gg *GradientGenerator) setDefaults() { _ = "STUB: not implemented"; return }

// Generate takes a generator and converts it into a source,
// drawing particles and binding functions for particle generation
// and rotation.
func (gg *GradientGenerator) Generate(layer int) *Source {
	_ = "STUB: not implemented"
	// Convert rotation from degrees to radians
	return nil
}

// GenerateParticle creates a particle from a generator
func (gg *GradientGenerator) GenerateParticle(bp *baseParticle) Particle {
	_ = "STUB: not implemented"
	return *new(Particle)
}

// Gradient Coloration
//

// SetStartColor2 satisfies Colorable2
func (gg *GradientGenerator) SetStartColor2(sc, scr color.Color) { _ = "STUB: not implemented"; return }

// SetEndColor2 satisfies Colorable2
func (gg *GradientGenerator) SetEndColor2(ec, ecr color.Color) { _ = "STUB: not implemented"; return }

// A Progresses has a SetProgress function where a progress function
// returns how far between two colors a given coordinate in a space is
type Progresses interface {
	SetProgress(func(x, y, w, h int) float64)
}

// Progress sets a Progresses' Progress Function
func Progress(pf func(x, y, w, h int) float64) func(Generator) {
	_ = "STUB: not implemented"
	return nil
}

// SetProgress satisfies Progresses
func (gg *GradientGenerator) SetProgress(pf func(x, y, w, h int) float64) {
	_ = "STUB: not implemented"
	return
}
