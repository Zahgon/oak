package particle

import (
	"time"

	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/physics"
)

const (
	//IgnoreEnd refers to the life value given to particles that want to skip their source's end function.
	IgnoreEnd = -2000 / 2
)

// A Source is used to store and control a set of particles.
type Source struct {
	Generator Generator
	*Allocator

	rotateBinding event.Binding

	particles [blockSize]Particle
	nextPID   int
	event.CallerID
	pIDBlock     int
	stackLevel   int
	EndFunc      func()
	stopRotateAt time.Time
	paused       bool
	started      bool
	stopped      bool
}

// NewDefaultSource creates a new sourceattached to the default event bus.
func NewDefaultSource(g Generator, stackLevel int) *Source { _ = "STUB: not implemented"; return nil }

// NewSource for particles constructed from a generator with specifications on how the particles should be handled.
func NewSource(handler event.Handler, g Generator, stackLevel int) *Source {
	_ = "STUB: not implemented"
	return nil
}

// cid must be set before the following bind call

// CID of our particle source
func (ps *Source) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

func (ps *Source) cycleParticles() bool { _ = "STUB: not implemented"; return false }

// Ignore dead particles

// Apply rotational acceleration

// Layer is shorthand for getting the base generator behind a source's layer
func (ps *Source) Layer(v physics.Vector) int { _ = "STUB: not implemented"; return 0 }

func (ps *Source) addParticles() { _ = "STUB: not implemented"; return }

// Regularly create particles (up until max particles)

// If this particle has not been allocated yet

// If this is a 'recycled' particle waiting to be redrawn

// rotateParticles updates particles over time as long
// as a Source is active.
func rotateParticles(ps *Source, _ event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}

// clearParticles is used after a Source has been stopped
// to continue moving old particles for as long as they exist.
func clearParticles(ps *Source, _ event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}

// TODO: not default

// Stop manually stops a Source, if its duration is infinite
// or if it should be stopped before expiring naturally.
func (ps *Source) Stop() { _ = "STUB: not implemented"; return }

// Pause on a Source just stops the repetition
// of its rotation function, which moves, destroys,
// ages and generates particles. Existing particles will
// stay in place.
func (ps *Source) Pause() {
	_ = "STUB: not implemented"

	// UnPause on a source a Source rebinds it's rotate function.
	return
}

func (ps *Source) UnPause() {
	_ = "STUB: not implemented"

	// IsPaused checks for whether the source is currently in a paused state.
	// It probably would have made more sense to export paused but this way if a lock is needed here in the future...
	// Then it wont change the api.
	return
}

func (ps *Source) IsPaused() bool {
	_ = "STUB: not implemented"

	// ShiftX shift's a source's underlying generator
	return false
}

func (ps *Source) ShiftX(x float64) { _ = "STUB: not implemented"; return }

// ShiftY shift's a source's underlying generator
func (ps *Source) ShiftY(y float64) { _ = "STUB: not implemented"; return }

// SetPos sets a source's underlying generator
func (ps *Source) SetPos(x, y float64) { _ = "STUB: not implemented"; return }
