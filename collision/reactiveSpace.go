package collision

import "sync"

// ReactiveSpace is a space that keeps track of a map of collision events
type ReactiveSpace struct {
	*Space
	Tree *Tree

	onHitsLock sync.Mutex
	onHits     map[Label]OnHit
}

// NewReactiveSpace creates a reactive space on the default collision tree
func NewReactiveSpace(s *Space, onHits map[Label]OnHit) *ReactiveSpace {
	_ = "STUB: not implemented"
	return nil
}

// CallOnHits calls CallOnHits on the underlying space of a reactive space
// with the reactive spaces' map of collision events, and returns the channel
// it will send the done signal from. It is not safe to call concurrently with
// add / remove / clear.
func (rs *ReactiveSpace) CallOnHits() chan bool { _ = "STUB: not implemented"; return nil }

// Add adds a mapping to a reactive spaces' onhit map
func (rs *ReactiveSpace) Add(i Label, oh OnHit) { _ = "STUB: not implemented"; return }

// Remove removes a mapping from a reactive spaces' onhit map
func (rs *ReactiveSpace) Remove(i Label) { _ = "STUB: not implemented"; return }

// Clear resets a reactive space's onhit map
func (rs *ReactiveSpace) Clear() { _ = "STUB: not implemented"; return }
