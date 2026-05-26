package physics

const frozen = -64

// ForceVector is a vector that has some force and can operate on entites with mass
type ForceVector struct {
	Vector
	Force *float64
}

// NewForceVector returns a force vector
func NewForceVector(direction Vector, force float64) ForceVector {
	_ = "STUB: not implemented"
	return *new(ForceVector)
}

// DefaultForceVector returns a force vector that converts the mass given
// into a force float
func DefaultForceVector(delta Vector, mass float64) ForceVector {
	_ = "STUB: not implemented"
	return *new(ForceVector)
}

// GetForce is a self-returning call
func (f ForceVector) GetForce() ForceVector {
	_ = "STUB: not implemented"

	// GetForce on a non-force vector returns a zero-value for force
	return *new(ForceVector)
}

func (v Vector) GetForce() ForceVector { _ = "STUB: not implemented"; return *new(ForceVector) }

// A Mass can have forces applied against it
type Mass struct {
	mass float64
}

// SetMass of an object
func (m *Mass) SetMass(inMass float64) error { _ = "STUB: not implemented"; return nil }

// GetMass returns the mass of an object
func (m *Mass) GetMass() float64 {
	_ = "STUB: not implemented"

	// Freeze changes a pushables mass such that it can no longer be pushed.
	return 0
}

func (m *Mass) Freeze() {
	_ = "STUB: not implemented"

	// Pushable is implemented by anything that has mass and directional movement,
	// and therefore can be pushed.
	return
}

type Pushable interface {
	GetDelta() Vector
	GetMass() float64
}

// A Pushes can push Pushable things through its associated ForceVector, or
// how hard the Pushable should move in a given direction
type Pushes interface {
	GetForce() ForceVector
}

// Push applies the force from the pushing object its target
func Push(a Pushes, b Pushable) error { _ = "STUB: not implemented"; return nil }

//Copy a's force so that we dont change the original when we scale it later
