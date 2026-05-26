package joystick

import (
	"os"
	"sync"

	"regexp"
)

// This has all been tested with wired xbox 360 controllers.
// Todo: get more controllers, test with more controllers.

func newJoystick(devName string, id uint32) *Joystick { _ = "STUB: not implemented"; return nil }

type osJoystick struct {
	devName string
	fh      *os.File
	cache   State
	sync.Mutex
	quit         chan struct{}
	disconnected bool
}

func osinit() error { _ = "STUB: not implemented"; return nil }

type jevent struct {
	Time   uint32
	Value  int16
	Type   uint8
	Number uint8
}

const (
	axisType   = 2
	buttonType = 1
)

var (
	buttons = []string{
		0: "A",
		1: "B",
		2: "X",
		3: "Y",
		4: "LeftShoulder",
		5: "RightShoulder",
		6: "Back",
		7: "Start",
		// 8 is the "Xbox" button in the center
		9:  "LeftStick",
		10: "RightStick",
	}
)

func (j *Joystick) prepare() error { _ = "STUB: not implemented"; return nil }

// Read events continually

// The controller offers int16 fidelity of the
// triggers. We're lowering it to Xinput's uint8
// Todo: Flip that around?

// No mutex here could cause a frame delay on inputs

func (j *Joystick) getState() (*State, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *Joystick) vibrate(left, right uint16) error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) close() error { _ = "STUB: not implemented"; return nil }

// Joysticks contain "js%d"
var joystickRegex = regexp.MustCompile("js(\\d)+")

func getJoysticks() []*Joystick { _ = "STUB: not implemented"; return nil }

// Find joysticks

// Ignore mice

// Todo: what else do we ignore?
