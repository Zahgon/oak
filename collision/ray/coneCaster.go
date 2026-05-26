package ray

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/collision"
)

var (
	// DefaultConeCaster is a global caster that all NewConeCaster()
	// calls are built on before options are applied.
	DefaultConeCaster = &ConeCaster{
		Caster:     DefaultCaster,
		CenterCone: true,
		ConeSpread: 1,
		Rays:       1,
	}
)

// SetDefaultConeCaster is analogous to SetDefaultCaster, however
// is equivalent to setting the global variable.
func SetDefaultConeCaster(coneCaster *ConeCaster) { _ = "STUB: not implemented"; return }

// A ConeCaster will repeatedly Cast
// its underlying Caster in a cone shape.
type ConeCaster struct {
	*Caster
	CenterCone bool
	// ConeSpread is represented in radians
	ConeSpread float64
	Rays       float64
}

// A ConeCastOption represents a transformation on a ConeCaster.
type ConeCastOption func(*ConeCaster)

// NewConeCaster copies the DefaultConeCaster and modifies it with the input
// options, returning the modified Caster. Zero arguments is valid input.
func NewConeCaster(opts ...ConeCastOption) *ConeCaster { _ = "STUB: not implemented"; return nil }

// CastTo casts a ray from origin to target, and otherwise acts as Cast.
func (cc *ConeCaster) CastTo(origin, target floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// Cast creates a ray from origin pointing at the given angle and returns
// some spaces collided with at the point of collision, given the settings of
// this ConeCaster. By default, all spaces hit will be returned. ConeCasters in
// addition will recast at progressive angles until they have cast up to their
// Rays value. Angles progress in counter-clockwise order.
func (cc *ConeCaster) Cast(origin, angle floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// Copy copies a ConeCaster.
func (cc *ConeCaster) Copy() *ConeCaster { _ = "STUB: not implemented"; return nil }

// ConeCast calls DefaultConeCaster.Cast. See (*ConeCaster).Cast
func ConeCast(origin, angle floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// ConeCastTo calls DefaultConeCaster.CastTo. See (*ConeCaster).CastTo
func ConeCastTo(origin, target floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// CenterCone sets whether the caster should center its cones around the
// input angles or progress out from those input angles. True by default.
//
// Example:
// Casting from a to b:
// if True:
//
//	  .  b  .
//	 . . .
//	...
//
// a
// if False:
//
//	  b     .
//	 .   .
//	. .
//
// a
func CenterCone(on bool) ConeCastOption { _ = "STUB: not implemented"; return *new(ConeCastOption) }

// ConeSpread sets how far a ConeCaster should progress its angles in degrees.
func ConeSpread(degrees float64) ConeCastOption {
	_ = "STUB: not implemented"
	return *new(ConeCastOption)
}

// ConeSpreadRadians sets how far a ConeCaster should progress its angles in radians.
func ConeSpreadRadians(radians float64) ConeCastOption {
	_ = "STUB: not implemented"
	return *new(ConeCastOption)
}

// ConeRays sets how many rays a ConeCaster should divide its spread into.
func ConeRays(rays int) ConeCastOption { _ = "STUB: not implemented"; return *new(ConeCastOption) }
