package particle

import (
	"github.com/oakmound/oak/v4/alg/span"
	"github.com/oakmound/oak/v4/physics"
	"github.com/oakmound/oak/v4/render"
)

// And chains together particle options into a single option
// for prebaking option sets
func And(as ...func(Generator)) func(Generator) { _ = "STUB: not implemented"; return nil }

// NewPerFrame sets how many particles should be produced per frame
func NewPerFrame(npf span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// Pos sets the initial position of spawned particles
func Pos(x, y float64) func(Generator) { _ = "STUB: not implemented"; return nil }

// LifeSpan sets how long a particle should last before dying
func LifeSpan(ls span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// InfiniteLifeSpan will set particles to never die over time.
func InfiniteLifeSpan() func(Generator) { _ = "STUB: not implemented"; return nil }

// Angle sets the initial angle of a particle in degrees
func Angle(a span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// Speed sets the initial speed of a particle
func Speed(s span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// Spread sets how far from a generator's position a particle can spawn
func Spread(x, y float64) func(Generator) { _ = "STUB: not implemented"; return nil }

// Duration sets how long a generator should produce particles for
func Duration(i span.Span[int]) func(Generator) { _ = "STUB: not implemented"; return nil }

// Rotation rotates particles by a variable amount per frame
func Rotation(a span.Span[float64]) func(Generator) { _ = "STUB: not implemented"; return nil }

// Gravity sets how a particle should be shifted over time in either dimension
func Gravity(x, y float64) func(Generator) { _ = "STUB: not implemented"; return nil }

// SpeedDecay sets how the speed of a particle should decay
func SpeedDecay(x, y float64) func(Generator) { _ = "STUB: not implemented"; return nil }

// End sets what function should happen when a particle dies
func End(ef func(Particle)) func(Generator) { _ = "STUB: not implemented"; return nil }

// Layer sets a function to determine what draw layer a particle should exist on
func Layer(l func(physics.Vector) int) func(Generator) { _ = "STUB: not implemented"; return nil }

// Limit limits the total number of particles a particle generator can have
// active at once.
func Limit(limit int) func(Generator) { _ = "STUB: not implemented"; return nil }

// DrawStack sets the current drawstack so that we dont use the globaldrawstack
func DrawStack(drawStack *render.DrawStack) func(Generator) { _ = "STUB: not implemented"; return nil }
