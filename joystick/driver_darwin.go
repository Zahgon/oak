package joystick

func osinit() error { _ = "STUB: not implemented"; return nil }

func newOsJoystick() osJoystick { _ = "STUB: not implemented"; return *new(osJoystick) }

type osJoystick struct {
}

func (j *Joystick) prepare() error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) getState() (*State, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *Joystick) vibrate(left, right uint16) error { _ = "STUB: not implemented"; return nil }

func (j *Joystick) close() error { _ = "STUB: not implemented"; return nil }

func getJoysticks() []*Joystick { _ = "STUB: not implemented"; return nil }
