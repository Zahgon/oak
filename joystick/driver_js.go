package joystick

import (
	"syscall/js"
)

func osinit() error {
	_ = "STUB: not implemented"
	// TODO: listen to joystick connected and joystick disconnected? We'd still need to
	// list from getGamepads every frame, it seems, to get new button presses.
	return nil
}

func newJoystick(gp js.Value, id uint32) *Joystick { _ = "STUB: not implemented"; return nil }

func newOsJoystick(gp js.Value) osJoystick { _ = "STUB: not implemented"; return *new(osJoystick) }

func refreshGamepadState(j *Joystick, gp js.Value) { _ = "STUB: not implemented"; return }

type jsGamepadState struct {
	axes      []float64
	buttons   []jsButton
	connected bool
	// osID      string
	// index     int
	mapping string
}

type jsButton struct {
	value float64
	//touched bool
	pressed bool
}

type osJoystick struct {
	cache      State
	jsState    jsGamepadState
	newJSState jsGamepadState
	newButtons map[string]bool
}

var (
	standardMappingButtons = []string{
		0: "A",
		1: "B",
		2: "X",
		3: "Y",
		4: "LeftShoulder",
		5: "RightShoulder",
		//6: LeftTrigger
		//7: RightTrigger
		8:  "Back",
		9:  "Start",
		10: "LeftStick",
		11: "RightStick",
		12: "Up",
		13: "Down",
		14: "Left",
		15: "Right",
	}
)

func (j *Joystick) prepare() error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) getState() (*State, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *Joystick) vibrate(left, right uint16) error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) close() error { _ = "STUB: not implemented"; return nil }

func getJoysticks() []*Joystick { _ = "STUB: not implemented"; return nil }
