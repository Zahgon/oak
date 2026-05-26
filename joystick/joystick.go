// Package joystick provides utilities for querying and reacting to joystick or
// gamepad inputs.
package joystick

import (
	"sync"
	"time"

	"github.com/oakmound/oak/v4/event"
)

type Input string

// Button or input names
const (
	InputA             Input = "A"
	InputB             Input = "B"
	InputX             Input = "X"
	InputY             Input = "Y"
	InputUp            Input = "Up"
	InputDown          Input = "Down"
	InputLeft          Input = "Left"
	InputRight         Input = "Right"
	InputBack          Input = "Back"
	InputStart         Input = "Start"
	InputLeftShoulder  Input = "LeftShoulder"
	InputRightShoulder Input = "RightShoulder"
	InputLeftStick     Input = "LeftStick"
	InputRightStick    Input = "RightStick"
)

// Events. All events but Disconnected include a *State payload.
var (
	Change          = event.RegisterEvent[*State]()
	ButtonDown      = event.RegisterEvent[*State]()
	ButtonUp        = event.RegisterEvent[*State]()
	RtTriggerChange = event.RegisterEvent[*State]()
	LtTriggerChange = event.RegisterEvent[*State]()
	RtStickChange   = event.RegisterEvent[*State]()
	LtStickChange   = event.RegisterEvent[*State]()
	// Disconnected includes the ID of the joystick that disconnected.
	Disconnected = event.RegisterEvent[uint32]()
)

// Init calls any os functions necessary to detect joysticks
func Init() error {
	_ = "STUB: not implemented"

	// A Triggerer can either be an event bus or event CID, allowing
	// joystick triggers to be listened to globally or sent to particular entities.
	return nil
}

type Triggerer interface {
	Trigger(eventID event.UnsafeEventID, data interface{}) <-chan struct{}
}

// A Joystick represents a (usually) physical controller connected to the machine.
// It can be listened to to obtain button / trigger states of the controller
// and propagate changes through an event handler.
type Joystick struct {
	Handler  Triggerer
	PollRate time.Duration
	id       uint32
	osJoystick
}

// State represents a snapshot of a joystick's inputs
type State struct {
	// Frame advances every time the State changes,
	// Making it easier to tell if the state has not changed
	// since the last poll.
	Frame uint32
	// ID is the joystick ID this state is associated with
	ID uint32
	// There isn't a canonical list of buttons (yet), because
	// this can be so dependent on controller type. To get a list of expected
	// buttons for a joystick, this map can be checked after a single GetState().
	// The package will ensure that the keyset of this map will never change for
	// a given joystick variable. True = Pressed/Down. False = Released/Up.
	// Todo: ensure the above guarantee
	Buttons  map[string]bool
	TriggerL uint8
	TriggerR uint8
	StickLX  int16
	StickLY  int16
	StickRX  int16
	StickRY  int16
}

// ListenOptions can be passed into a joystick's Listen method to
// change what events will be propagated out.
type ListenOptions struct {
	// Each boolean dictates given event type(s) will be sent during Listen
	// "JoystickChange": *State
	JoystickChanges bool
	// Whenever any button is changed:
	// "ButtonDown": *State
	// "ButtonUp": *State
	GenericButtonPresses bool
	// "($buttonName)Up/Down": *State
	// e.g.
	// "AButtonUp": *State
	// "XButtonDown:" *State
	ButtonPresses bool
	// "RtTriggerChange": *State
	// "LtTriggerChange": *State
	TriggerChanges bool
	// "RtStickChange": *State
	// "LtStickChange": *State
	StickChanges bool
	// StickDeadzones enable preventing movement near the center of
	// the analog control being sent to a logic handler. A
	// StickDeadzone value will be treated as an absolute threshold.
	StickDeadzoneLX int16
	StickDeadzoneLY int16
	StickDeadzoneRX int16
	StickDeadzoneRY int16
}

func (lo *ListenOptions) sendFn() func(Triggerer, *State, *State) {
	_ = "STUB: not implemented"
	// Todo: benchmark that this is an effective way of reducing time spent
	// sending unwanted events
	return nil
}

var upEventsLock sync.Mutex
var upEvents = map[string]event.EventID[*State]{}

func Up(s string) event.EventID[*State] { _ = "STUB: not implemented"; return nil }

var downEventsLock sync.Mutex
var downEvents = map[string]event.EventID[*State]{}

func Down(s string) event.EventID[*State] { _ = "STUB: not implemented"; return nil }

func deltaExceedsThreshold(old, new, threshold int16) bool { _ = "STUB: not implemented"; return false }

func intAbs(x int16) (positiveX int16) { _ = "STUB: not implemented"; return 0 }

// Listen causes the joystick to send its inputs to its Handler, by regularly
// querying GetState. The type of events returned can be specified by options.
// If the options are nil, only JoystickChange events will be sent.
func (j *Joystick) Listen(opts *ListenOptions) (cancel func()) {
	_ = "STUB: not implemented"
	return nil
}

// Perform required initialization to receive inputs from OS

// Wait on inputs

// ID returns which player this joystick is associated with
func (j *Joystick) ID() uint32 {
	_ = "STUB: not implemented"

	// Vibrate triggers vibration on a joystick (if it is supported).
	return 0
}

func (j *Joystick) Vibrate(left, right uint16) error { _ = "STUB: not implemented"; return nil }

// Prepare allocates any operating system resources needed to read signals
// for the joystick.
func (j *Joystick) Prepare() error {
	_ = "STUB: not implemented"

	// GetState returns the current button, trigger, and analog stick
	// state of the joystick
	return nil
}

func (j *Joystick) GetState() (*State, error) {
	_ = "STUB: not implemented"
	return nil,

		// Close closes any operating system resources backing this joystick's signals
		nil
}

func (j *Joystick) Close() error {
	_ = "STUB: not implemented"

	// GetJoysticks returns all known active joysticks.
	return nil
}

func GetJoysticks() []*Joystick { _ = "STUB: not implemented"; return nil }

// WaitForJoysticks will regularly call GetJoysticks so to send signals
// on the output channel when new joysticks are connected to the system.
// Call `cancel' to close the channel and stop polling.
func WaitForJoysticks(pollRate time.Duration) (joyCh <-chan *Joystick, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send all existing joysticks

// j is new
