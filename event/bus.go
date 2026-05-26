package event

import (
	"sync"
	"time"
)

// A Bus stores bindables to be triggered by events.
type Bus struct {
	// nextBindID is an atomically incrementing value to track bindings within this structure
	nextBindID *int64

	// resetCount increments every time the bus is reset. bindings and unbindings make sure that
	// they are called on a bus with an unchanged reset count, and become NOPs if performed on
	// a bus with a different reset count to ensure they do not interfere with a bus using different
	// bind IDs.
	resetCount         int64
	bindingMap         map[UnsafeEventID]map[CallerID]bindableList
	persistentBindings []persistentBinding

	callerMap *CallerMap

	mutex sync.RWMutex
}

// a persistentBinding is rebound every time the bus is reset.
type persistentBinding struct {
	eventID  UnsafeEventID
	callerID CallerID
	fn       UnsafeBindable
}

// NewBus returns an empty event bus with an assigned caller map. If nil
// is provided, the caller map used will be DefaultCallerMap
func NewBus(callerMap *CallerMap) *Bus { _ = "STUB: not implemented"; return nil }

// SetCallerMap updates a bus to use a specific set of callers.
func (bus *Bus) SetCallerMap(cm *CallerMap) {
	_ = "STUB: not implemented"

	// GetCallerMap returns this bus's caller map.
	return
}

func (b *Bus) GetCallerMap() *CallerMap {
	_ = "STUB: not implemented"

	// ClearPersistentBindings removes all persistent bindings. It will not unbind them
	// from the bus, but they will not be bound following the next bus reset.
	return nil
}

func (eb *Bus) ClearPersistentBindings() { _ = "STUB: not implemented"; return }

// Reset unbinds all present, non-persistent bindings on the bus. It will block until
// persistent bindings are in place.
func (bus *Bus) Reset() { _ = "STUB: not implemented"; return }

// EnterLoop triggers Enter events at the specified rate until the returned cancel is called.
func EnterLoop(bus Handler, frameDelay time.Duration) (cancel func()) {
	_ = "STUB: not implemented"
	return nil
}

// Q: why send here as well as close
// A: to ensure that no more ticks are sent, the above goroutine has to
//    acknowledge that it should stop and return-- just closing would
//    enable code following this cancel function to assume no enters were
//    being triggered when they still are.
