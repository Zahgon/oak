package physics

// A Vector is a two-dimensional point or vector used throughout oak
// to maintain functionality between packages.
type Vector struct {
	x, y       *float64
	offX, offY float64
}

const (
	// CUTOFF is used for rounding after floating point operations to
	// zero out vector values that are sufficiently close to zero
	CUTOFF = 0.001
)

// NewVector returns a vector with the given x and y components
func NewVector(x, y float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// NewVector32 accepts float32s (and will cast them to float64s), to create a vector
func NewVector32(x, y float32) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// PtrVector takes in pointers as opposed to float values-- these
// same pointers will be used in the returned vector and by attachments
// to that vector.
func PtrVector(x, y *float64) Vector {
	_ = "STUB: not implemented"
	return *

	// AngleVector creates a unit vector by the cosine and sine of the given
	// angle in degrees
	new(Vector)
}

func AngleVector(angle float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// MaxVector returns whichever vector has a greater magnitude
func MaxVector(a, b Vector) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Copy copies a Vector
func (v Vector) Copy() Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Magnitude returns the magnitude of the combined components of a Vector
func (v Vector) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// Normalize divides both components in a vector by the vector's magnitude
func (v Vector) Normalize() Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Zero is shorthand for NewVector(0,0), but uses the input vector
func (v Vector) Zero() Vector {
	_ = "STUB: not implemented"
	return *

	// Add combines a set of vectors through addition
	new(Vector)
}

func (v Vector) Add(vs ...Vector) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Sub combines a set of vectors through subtraction
func (v Vector) Sub(vs ...Vector) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Scale scales a vector by a set of floating points
// Scale(f1,f2,f3) is equivalent to Scale(f1*f2*f3)
func (v Vector) Scale(fs ...float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Rotate takes in a set of angles and rotates v by their sum
// the input angles are assumed to be in degrees.
func (v Vector) Rotate(fs ...float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Angle returns this vector as an angle in degrees
func (v Vector) Angle() float64 { _ = "STUB: not implemented"; return 0 }

// Dot returns the dot product of the vectors
func (v Vector) Dot(v2 Vector) float64 { _ = "STUB: not implemented"; return 0 }

// Distance on two vectors returns the euclidean distance
// from v to v2
func (v Vector) Distance(v2 Vector) float64 { _ = "STUB: not implemented"; return 0 }

func (v Vector) round() Vector { _ = "STUB: not implemented"; return *new(Vector) }

// ShiftX is equivalent to v.X() += x
func (v Vector) ShiftX(x float64) Vector {
	_ = "STUB: not implemented"
	return *

	// ShiftY is equivalent to v.Y() += y
	new(Vector)
}

func (v Vector) ShiftY(y float64) Vector {
	_ = "STUB: not implemented"
	return *

	// X returns this vector's x component
	new(Vector)
}

func (v Vector) X() float64 { _ = "STUB: not implemented"; return 0 }

// Y returns this vector's x component
func (v Vector) Y() float64 { _ = "STUB: not implemented"; return 0 }

// SetX returns a vector with its x component set to x
func (v Vector) SetX(x float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// SetY returns a vector with its y component set to y
func (v Vector) SetY(y float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// Xp returns the real pointer behind this vector's x component
func (v Vector) Xp() *float64 {
	_ = "STUB: not implemented"

	// Yp returns the real pointer behind  this vector's y component
	return nil
}

func (v Vector) Yp() *float64 {
	_ = "STUB: not implemented"

	// SetPos is equivalent to NewVector(x,y)
	return nil
}

func (v Vector) SetPos(x, y float64) Vector { _ = "STUB: not implemented"; return *new(Vector) }

// GetPos returns both v.X() and v.Y()
func (v Vector) GetPos() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }
