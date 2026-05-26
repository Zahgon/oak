// Package shake provides methods for rapidly shifting graphical components' positions
package shake

import (
	"context"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/window"
)

// A Shaker knows how to shake something by a (or up to a) given magnitude.
// If Random is true, the Shaker will shake up to the (negative or positive)
// magnitude of each the X and Y axes. Otherwise, it will oscillate between
// negative magnitude and positive magnitude.
type Shaker struct {
	Magnitude floatgeom.Point2
	Delay     time.Duration
	Random    bool
	// ResetPosition determines whether the shaken entity will be reset back to its original position
	// after shaking is complete. True by default.
	ResetPosition bool
}

var (
	// DefaultShaker is the global default shaker, used when shake.Screen or shake.Shake are called.
	DefaultShaker = &Shaker{
		Random:        false,
		Magnitude:     floatgeom.Point2{3.0, 3.0},
		Delay:         30 * time.Millisecond,
		ResetPosition: true,
	}
)

// A ShiftPoser can have its position shifted by an x,y pair
type ShiftPoser interface {
	ShiftPos(x, y float64)
}

// Shake shakes a ShiftPoser for the given duration. It uses the settings
// in DefaultShaker to determine the quality of the shake.
func Shake(sp ShiftPoser, dur time.Duration) { _ = "STUB: not implemented"; return }

// Shake shakes a ShiftPoser for the given duration.
func (sk *Shaker) Shake(sp ShiftPoser, dur time.Duration) { _ = "STUB: not implemented"; return }

// ShakeContext shakes a ShiftPoser for the given duration or until the context is done,
// whichever comes first.
func (sk *Shaker) ShakeContext(ctx context.Context, sp ShiftPoser, dur time.Duration) {
	_ = "STUB: not implemented"
	return
}

type screenToPoser struct {
	window.App
}

func (stp screenToPoser) ShiftPos(x, y float64) { _ = "STUB: not implemented"; return }

// Screen shakes the screen that the context controls for the given duration.
// It uses the settings in DefaultShaker to determine the quality of the shake.
func Screen(ctx *scene.Context, dur time.Duration) { _ = "STUB: not implemented"; return }

// ShakeScreen shakes the screen that the context controls for the given duration.
func (sk *Shaker) ShakeScreen(ctx *scene.Context, dur time.Duration) {
	_ = "STUB: not implemented"
	return
}
