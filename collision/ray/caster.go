package ray

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/collision"
)

var (
	// DefaultCaster is a global caster that all
	// NewCaster() calls are built on before options are applied.
	DefaultCaster = &Caster{
		PointSize: floatgeom.Point2{.1, .1},
		PointSpan: 1.0,
		// CastDistance needs to be defined, but
		// there isn't a reasonable default.
		// Consider: Cast() could take in distance as well.
		CastDistance: 200,
		Tree:         collision.DefaultTree,
	}
)

// SetDefaultCaster sets the global caster to be the input, and
// sets the caster behind the global cone caster as well.
func SetDefaultCaster(caster *Caster) { _ = "STUB: not implemented"; return }

// A Caster can cast rays and return the colliding collision points
// of rays cast from points at angles. This behavior is customizable
// through CastOptions.
type Caster struct {
	Filters      []CastFilter
	Limits       []CastLimit
	PointSize    floatgeom.Point2
	PointSpan    float64
	CastDistance float64
	Tree         *collision.Tree
	CenterPoints bool
}

// A CastOption represents a transformation to a ray caster.
type CastOption func(*Caster)

// NewCaster will copy and modify the DefaultCaster by the input options
// and return the modified Caster. Giving no inputs is valid.
func NewCaster(opts ...CastOption) *Caster { _ = "STUB: not implemented"; return nil }

// CastTo casts a ray from origin to target, and otherwise acts as Cast.
func (c *Caster) CastTo(origin, target floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// Cast creates a ray from origin pointing at the given angle and returns
// some spaces collided with at the point of collision, given the settings of
// this Caster. By default, all spaces hit will be returned.
func (c *Caster) Cast(origin, angle floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// Copy copies a Caster.
func (c *Caster) Copy() *Caster { _ = "STUB: not implemented"; return nil }

// Cast calls DefaultCaster.Cast. See (*Caster).Cast
func Cast(origin, angle floatgeom.Point2) []collision.Point { _ = "STUB: not implemented"; return nil }

// CastTo calls DefaultCaster.CastTo. See (*Caster).CastTo
func CastTo(origin, target floatgeom.Point2) []collision.Point {
	_ = "STUB: not implemented"
	return nil
}

// Tree sets the collision tree of a Caster.
func Tree(t *collision.Tree) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// CenterPoints sets whether a Caster should center its collision points that
// form its ray. This is by default false, and is only significant if said
// points' dimensions are significantly large.
func CenterPoints(on bool) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// Distance determines how far a caster will project rays before stopping
func Distance(dist float64) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// PointSize determines the size of a caster's collision checks
func PointSize(ps floatgeom.Point2) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// PointSpan determines the distance between collision check points
func PointSpan(span float64) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }
