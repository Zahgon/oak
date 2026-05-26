package oak

import (
	"github.com/oakmound/oak/v4/event"
)

// InputType expresses some form of input to the engine to represent a player
type InputType int32

var trackingJoystickChange = event.RegisterEvent[struct{}]()

// The following constants define valid types of input sent via the InputChange event.
const (
	InputNone InputType = iota
	InputKeyboard
	InputMouse
	InputJoystick
)

func (w *Window) trackInputChanges() { _ = "STUB: not implemented"; return }

type joyHandler struct {
	handler event.Handler
}

func (jh *joyHandler) Trigger(eventID event.UnsafeEventID, data interface{}) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func trackJoystickChanges(handler event.Handler) { _ = "STUB: not implemented"; return }
