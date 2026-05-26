package floatgeom

// Point2 represents a 2D point on a plane.
type Point2 [2]float64

// Point3 represents a 3D point in space.
type Point3 [3]float64

// Point4 represents a 4D point, in space + some additional dimension.
type Point4 [4]float64

// AnglePoint creates a unit vector from the given angle in degrees as a Point2.
func AnglePoint(angle float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// RadianPoint creates a unit vector from the given angle in radians as a Point2.
func RadianPoint(radians float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Dim returns the value of p in the ith dimension.
// Panics if i > 1.
func (p Point2) Dim(i int) float64 {
	_ = "STUB: not implemented"

	// Dim returns the value of p in the ith dimension.
	// Panics if i > 2.
	return 0
}

func (p Point3) Dim(i int) float64 {
	_ = "STUB: not implemented"

	// Dim returns the value of p in the ith dimension.
	// Panics if i > 3.
	return 0
}

func (p Point4) Dim(i int) float64 {
	_ = "STUB: not implemented"

	// X returns p's value on the X axis.
	return 0
}

func (p Point2) X() float64 {
	_ = "STUB: not implemented"

	// Y returns p's value on the Y axis.
	return 0
}

func (p Point2) Y() float64 {
	_ = "STUB: not implemented"

	// X returns p's value on the X axis.
	return 0
}

func (p Point3) X() float64 {
	_ = "STUB: not implemented"

	// Y returns p's value on the Y axis.
	return 0
}

func (p Point3) Y() float64 {
	_ = "STUB: not implemented"

	// Z returns p's value on the Z axis.
	return 0
}

func (p Point3) Z() float64 {
	_ = "STUB: not implemented"

	// W returns p's value on the W axis.
	return 0
}

func (p Point4) W() float64 {
	_ = "STUB: not implemented"

	// X returns p's value on the X axis.
	return 0
}

func (p Point4) X() float64 {
	_ = "STUB: not implemented"

	// Y returns p's value on the Y axis.
	return 0
}

func (p Point4) Y() float64 {
	_ = "STUB: not implemented"

	// Z returns p's value on the Z axis.
	return 0
}

func (p Point4) Z() float64 {
	_ = "STUB: not implemented"

	// Distance calculates the distance between this Point2 and another.
	return 0
}

func (p Point2) Distance(p2 Point2) float64 { _ = "STUB: not implemented"; return 0 }

// Distance calculates the distance between this Point3 and another.
func (p Point3) Distance(p2 Point3) float64 { _ = "STUB: not implemented"; return 0 }

// Distance2 calculates the euclidean distance between two points, as two (x,y) pairs
func Distance2(x1, y1, x2, y2 float64) float64 { _ = "STUB: not implemented"; return 0 }

// Distance3 calculates the euclidean distance between two points, as two (x,y,z) triplets
func Distance3(x1, y1, z1, x2, y2, z2 float64) float64 { _ = "STUB: not implemented"; return 0 }

// LesserOf returns the lowest values on each axis of the input points as a point.
func (p Point2) LesserOf(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// LesserOf returns the lowest values on each axis of the input points as a point.
func (p Point3) LesserOf(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// GreaterOf returns the highest values on each axis of the input points as a point.
func (p Point2) GreaterOf(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// GreaterOf returns the highest values on each axis of the input points as a point.
func (p Point3) GreaterOf(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Add combines the input points via addition.
func (p Point2) Add(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Sub combines the input points via subtraction.
func (p Point2) Sub(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Mul combines in the input points via multiplication.
func (p Point2) Mul(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// MulConst multiplies all elements of a point by the input floats
func (p Point2) MulConst(fs ...float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Cross gets the cross product of two Point 3s
func (p Point3) Cross(p2 Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Div combines the input points via division.
// Div does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point2) Div(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// DivConst divides all elements of a point by the input floats
// DivConst does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point2) DivConst(fs ...float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Add combines the input points via addition.
func (p Point3) Add(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Sub combines the input points via subtraction.
func (p Point3) Sub(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Mul combines in the input points via multiplication.
func (p Point3) Mul(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// MulConst multiplies all elements of a point by the input floats
func (p Point3) MulConst(fs ...float64) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Div combines the input points via division.
// Div does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point3) Div(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// DivConst divides all elements of a point by the input floats
// DivConst does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point3) DivConst(fs ...float64) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// MulConst multiplies all elements of a point by the input floats
func (p Point4) MulConst(fs ...float64) Point4 { _ = "STUB: not implemented"; return *new(Point4) }

// DivConst divides all elements of a point by the input floats
// DivConst does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point4) DivConst(fs ...float64) Point4 { _ = "STUB: not implemented"; return *new(Point4) }

// Dot returns the dot product of the input points
func (p Point2) Dot(p2 Point2) float64 { _ = "STUB: not implemented"; return 0 }

// Dot returns the dot product of the input points
func (p Point3) Dot(p2 Point3) float64 { _ = "STUB: not implemented"; return 0 }

// Dot returns the dot product of the input points
func (p Point4) Dot(p2 Point4) float64 { _ = "STUB: not implemented"; return 0 }

// Magnitude returns the magnitude of the combined components of a Point
func (p Point2) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// Magnitude returns the magnitude of the combined components of a Point
func (p Point3) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// Magnitude returns the magnitude of the combined components of a Point
func (p Point4) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// Normalize converts this point into a unit vector.
func (p Point2) Normalize() Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Normalize converts this point into a unit vector.
func (p Point3) Normalize() Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Normalize converts this point into a unit vector.
func (p Point4) Normalize() Point4 { _ = "STUB: not implemented"; return *new(Point4) }

// Rotate takes in a set of angles and rotates v by their sum
// the input angles are expected to be in degrees.
func (p Point2) Rotate(fs ...float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// RotateRadians takes in a set of angles and rotates v by their sum
// the input angles are expected to be in radians.
func (p Point2) RotateRadians(fs ...float64) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// ToRect converts this point into a rectangle spanning span distance
// in each axis.
func (p Point2) ToRect(span float64) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// ToRect converts this point into a rectangle spanning span distance
// in each axis.
func (p Point3) ToRect(span float64) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

// ProjectX projects the Point3 onto the x axis, removing it's
// x component and returning a Point2
func (p Point3) ProjectX() Point2 {
	_ = "STUB: not implemented"
	return *

	// ProjectY projects the Point3 onto the y axis, removing it's
	// y component and returning a Point2
	new(Point2)
}

func (p Point3) ProjectY() Point2 {
	_ = "STUB: not implemented"
	return *

	// ProjectZ projects the Point3 onto the z axis, removing it's
	// z component and returning a Point2
	new(Point2)
}

func (p Point3) ProjectZ() Point2 {
	_ = "STUB: not implemented"
	return *

	// ToAngle returns this point as an angle in degrees.
	new(Point2)
}

func (p Point2) ToAngle() float64 { _ = "STUB: not implemented"; return 0 }

// ToRadians returns this point as an angle in radians.
func (p Point2) ToRadians() float64 { _ = "STUB: not implemented"; return 0 }

// AngleTo returns the angle from p to p2 in degrees.
func (p Point2) AngleTo(p2 Point2) float64 { _ = "STUB: not implemented"; return 0 }

// RadiansTo returns the angle from p to p2 in radians.
func (p Point2) RadiansTo(p2 Point2) float64 { _ = "STUB: not implemented"; return 0 }

// Conjugate returns a value of a Point4 often obtained to calculate the inverse
func (p Point4) Conjugate() Point4 { _ = "STUB: not implemented"; return *new(Point4) }

// Inverse of a Point4, often used to get the inverse rotation of a quaternion
func (p Point4) Inverse() Point4 { _ = "STUB: not implemented"; return *new(Point4) }

// MulQuat multiplies two quaternions to generate a combined quarternion that represents both rotations
// ref: https://www.mathworks.com/help/aeroblks/quaternionmultiplication.html
func (p Point4) MulQuat(p2 Point4) Point4 { _ = "STUB: not implemented"; return *new(Point4) }
