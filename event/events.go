package event

import (
	"time"
)

// An UnsafeEventID is a non-typed eventID. EventIDs are just these, with type information attached.
type UnsafeEventID int64

// A EventID represents an event associated with a given payload type.
type EventID[T any] struct {
	UnsafeEventID
}

var (
	nextEventID int64
)

// RegisterEvent returns a unique ID to associate an event with. EventIDs not created through RegisterEvent are
// not valid for use in type-safe bindings.
func RegisterEvent[T any]() EventID[T] { _ = "STUB: not implemented"; return nil }

// EnterPayload is the payload sent down to Enter bindings
type EnterPayload struct {
	FramesElapsed  int
	SinceLastFrame time.Duration
	TickPercent    float64
}

var (
	// Enter: the beginning of every logical frame.
	Enter = RegisterEvent[EnterPayload]()
)
