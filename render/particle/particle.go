// Package particle provides options for generating renderable
// particle sources.
package particle

import (
	"image/draw"

	"github.com/oakmound/oak/v4/physics"
	"github.com/oakmound/oak/v4/render"
)

// A Particle is a renderable that is spawned by a generator, usually very fast,
// usually very small, for visual effects
type Particle interface {
	render.Renderable
	GetBaseParticle() *baseParticle
	GetPos() physics.Vector
	DrawOffsetGen(gen Generator, buff draw.Image, xOff, yOff float64)
	Cycle(gen Generator)
	setPID(int)
}

type baseParticle struct {
	render.LayeredPoint
	Src       *Source
	Vel       physics.Vector
	Life      float64
	totalLife float64
	pID       int
}

func (bp *baseParticle) GetLayer() int { _ = "STUB: not implemented"; return 0 }

func (bp *baseParticle) GetBaseParticle() *baseParticle { _ = "STUB: not implemented"; return nil }

func (bp *baseParticle) GetPos() physics.Vector {
	_ = "STUB: not implemented"
	return *new(physics.Vector)
}

func (bp *baseParticle) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (bp *baseParticle) Cycle(gen Generator) { _ = "STUB: not implemented"; return }

func (bp *baseParticle) setPID(pid int) { _ = "STUB: not implemented"; return }
