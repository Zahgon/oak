package render

import (
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render/mod"
)

// The Reverting structure lets modifications be made to a Modifiable and then
// reverted, up to arbitrary history limits.
type Reverting struct {
	Modifiable
	rs []Modifiable
}

// NewReverting returns a Reverting type wrapped around the given modifiable
func NewReverting(m Modifiable) *Reverting { _ = "STUB: not implemented"; return nil }

// Revert goes back n steps in this Reverting's history and displays that Modifiable
func (rv *Reverting) Revert(n int) { _ = "STUB: not implemented"; return }

// RevertAll resets this reverting to its original Modifiable
func (rv *Reverting) RevertAll() { _ = "STUB: not implemented"; return }

// RevertAndModify reverts n steps and then modifies this reverting. This
// is a separate function from Revert followed by Modify to prevent skipped
// draw frames.
func (rv *Reverting) RevertAndModify(n int, ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// RevertAndFilter acts as RevertAndModify, but with Filters.
func (rv *Reverting) RevertAndFilter(n int, fs ...mod.Filter) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Modify alters this reverting by the given modifications, appending the new
// modified renderable to it's list of modified versions and displaying it.
func (rv *Reverting) Modify(ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Filter alters this reverting by the given filters, appending the new
// modified renderable to it's list of modified versions and displaying it.
func (rv *Reverting) Filter(ms ...mod.Filter) { _ = "STUB: not implemented"; return }

// Copy returns a copy of this Reverting
func (rv *Reverting) Copy() Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

// This might not ever be called?
func (rv *Reverting) update() { _ = "STUB: not implemented"; return }

// SetTriggerID sets the ID AnimationEnd will trigger on for animating subtypes.
func (rv *Reverting) SetTriggerID(cid event.CallerID) { _ = "STUB: not implemented"; return }

// Pause ceases animating any renderable types that animate underneath this
func (rv *Reverting) Pause() { _ = "STUB: not implemented"; return }

// Unpause resumes animating any renderable types that animate underneath this
func (rv *Reverting) Unpause() { _ = "STUB: not implemented"; return }

// IsInterruptable returns if whatever this reverting is currently dispalying is interruptable.
func (rv *Reverting) IsInterruptable() bool { _ = "STUB: not implemented"; return false }

// IsStatic returns if whatever this reverting is currently displaying is static.
func (rv *Reverting) IsStatic() bool { _ = "STUB: not implemented"; return false }

// Get calls Get on the active renderable below this Reverting. If nothing has a Get
// method, it returns the empty string.
func (rv *Reverting) Get() string { _ = "STUB: not implemented"; return "" }

// Set calls Set on underlying types below this Reverting that can be Set
// Todo: if Set becomes used by more types, this should use an interface like
// CanPause
func (rv *Reverting) Set(k string) error { _ = "STUB: not implemented"; return nil }
