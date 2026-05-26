package physics

// An Attachable can be attached to static or moving vectors.
type Attachable interface {
	Detach()
	Attach(Vecer, float64, float64)
	AttachX(Vecer, float64)
	AttachY(Vecer, float64)
	Vecer
}

// A Vecer can be converted into a Vector
type Vecer interface {
	Vec() Vector
}

// Vec returns a vector itself
func (v Vector) Vec() Vector {
	_ = "STUB: not implemented"

	// Attach takes in something for this vector to attach to and a set of
	// offsets.
	return *new(Vector)
}

func (v *Vector) Attach(a Vecer, offX, offY float64) { _ = "STUB: not implemented"; return }

// AttachX performs an attachment that only attaches on the X axis.
func (v *Vector) AttachX(a Vecer, offX float64) { _ = "STUB: not implemented"; return }

// AttachY performs an attachment that only attaches on the Y axis.
func (v *Vector) AttachY(a Vecer, offY float64) { _ = "STUB: not implemented"; return }

// Detach modifies a vector to no longer be attached to anything.
func (v *Vector) Detach() { _ = "STUB: not implemented"; return }

// DetachX modifies a vector to no longer be attached on the X Axis.
func (v *Vector) DetachX() { _ = "STUB: not implemented"; return }

// DetachY modifies a vector to no longer be attached on the Y Axis.
func (v *Vector) DetachY() { _ = "STUB: not implemented"; return }
