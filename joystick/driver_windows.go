package joystick

import (
	"sync"

	"github.com/oakmound/w32"
)

func newJoystick(id uint32) *Joystick { _ = "STUB: not implemented"; return nil }

type osJoystick struct {
	// Todo: mutex these values?
	wstate    *w32.XInputState
	vibration *w32.XInputVibration
}

// The windows driver currently uses the xinput api.
// We should consider providing alternatives.

var once sync.Once

func osinit() error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) prepare() error { _ = "STUB: not implemented"; return nil }

type buttonName struct {
	name      Input
	xinputVal uint16
}

var (
	chkButtons = []buttonName{
		{InputUp, w32.XINPUT_GAMEPAD_DPAD_UP},
		{InputDown, w32.XINPUT_GAMEPAD_DPAD_DOWN},
		{InputLeft, w32.XINPUT_GAMEPAD_DPAD_LEFT},
		{InputRight, w32.XINPUT_GAMEPAD_DPAD_RIGHT},
		{InputStart, w32.XINPUT_GAMEPAD_START},
		{InputBack, w32.XINPUT_GAMEPAD_BACK},
		{InputLeftStick, w32.XINPUT_GAMEPAD_LEFT_THUMB},
		{InputRightStick, w32.XINPUT_GAMEPAD_RIGHT_THUMB},
		{InputLeftShoulder, w32.XINPUT_GAMEPAD_LEFT_SHOULDER},
		{InputRightShoulder, w32.XINPUT_GAMEPAD_RIGHT_SHOULDER},
		{InputA, w32.XINPUT_GAMEPAD_A},
		{InputB, w32.XINPUT_GAMEPAD_B},
		{InputX, w32.XINPUT_GAMEPAD_X},
		{InputY, w32.XINPUT_GAMEPAD_Y},
	}
)

func (j *Joystick) getState() (*State, error) { _ = "STUB: not implemented"; return nil, nil }

// Convert windows state into os-regular state

func (j *Joystick) vibrate(left, right uint16) error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) close() error {
	_ = "STUB: not implemented"
	// It seemingly makes sense to do this, but doing this disables
	// detection of future joysticks
	// return w32.XInputEnable(false)
	return nil
}

func getJoysticks() []*Joystick {
	_ = "STUB: not implemented"
	// With xinput there are explicitly up to 4 controllers
	return nil
}
