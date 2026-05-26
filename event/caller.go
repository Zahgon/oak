package event

import (
	"sync"
)

// A CallerID is a caller ID that Callers use to bind themselves to receive callback
// signals when given events are triggered
type CallerID int64

func (c CallerID) CID() CallerID {
	_ = "STUB: not implemented"

	// Global is the CallerID associated with global bindings. A caller must not be assigned
	// this ID. Global may be used to manually create bindings scoped to no callers, but the GlobalBind function
	// should be preferred when possible for type safety.
	return *new(CallerID)
}

const Global CallerID = 0

type Caller interface {
	CID() CallerID
}

// A CallerMap tracks CallerID mappings to Entities.
// This is an alternative to passing in the entity via closure scoping,
// and allows for more general bindings as simple top level functions.
type CallerMap struct {
	highestID   CallerID
	callersLock sync.RWMutex
	callers     map[CallerID]Caller
}

// NewCallerMap creates a caller map. A CallerMap
// is not valid for use if not created via this function.
func NewCallerMap() *CallerMap { _ = "STUB: not implemented"; return nil }

// NextID finds the next available caller id
// and returns it, after adding the given entity to
// the caller map.
func (cm *CallerMap) Register(e Caller) CallerID { _ = "STUB: not implemented"; return *new(CallerID) }

// Q: Why not use atomic?
// A: We're in a mutex and therefore it is not needed.
// A2: We need the mutex to safely assign to the map.
// A3: We cannot atomically increment outside of the map, consider:
//     - GR1 calls Clear, waits on Lock
//     - GR2 calls Register, gets id 100, waits on lock
//     - GR1 claims lock, resets highestID to 0, exits
//     - GR2 claims lock, inserts id 100 in the map
//     - ... later, register silently overwrites entity 100, its
//       bindings will now panic on a bad type assertion
//
// Increment before assigning to preserve Global == caller 0

// Get returns the entity corresponding to the given ID within
// the caller map. If no entity is found, it returns nil.
func (cm *CallerMap) GetEntity(id CallerID) Caller { _ = "STUB: not implemented"; return *new(Caller) }

// Has returns whether the given caller id is an initialized entity
// within the caller map.
func (cm *CallerMap) HasEntity(id CallerID) bool { _ = "STUB: not implemented"; return false }

// Remove removes an entity from the caller map.
func (cm *CallerMap) RemoveEntity(id CallerID) { _ = "STUB: not implemented"; return }

// Clear clears the caller map to forget all registered callers.
func (cm *CallerMap) Clear() { _ = "STUB: not implemented"; return }
