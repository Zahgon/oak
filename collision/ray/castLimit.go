package ray

import (
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/event"
)

// A CastLimit is a function that can be applied to
// a Caster's points to return after it adds each one.
// If a Limit returns false, that Caster will immediately
// cease casting.
//
// If a Caster's ray collides with multiple spaces at the same
// point, and some of them would pass a CastLimit, but others would
// not, a Caster will not reliably return those that would pass the
// limit.
type CastLimit func([]collision.Point) bool

// AddLimit is a helper for converting a CastLimit into a CastOption.
func AddLimit(cl CastLimit) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// LimitResults will cause a Caster to return a limited number of
// collision points.
func LimitResults(limit int) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }

// StopAtLabel will cause a caster to cease casting as soon as it
// hits one of the input labels.
func StopAtLabel(ls ...collision.Label) CastOption {
	_ = "STUB: not implemented"
	return *new(CastOption)
}

// StopAtID will cause a caster to cease casting as soon as it
// hits one of the input CIDs.
func StopAtID(ids ...event.CallerID) CastOption { _ = "STUB: not implemented"; return *new(CastOption) }
