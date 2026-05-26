package intgeom

// Point2 represents a 2D point in space.
type Point2 [2]int

// Point3 represents a 3D point in space.
type Point3 [3]int

// Dim returns the value of p in the ith dimension.
// Panics if i > 1. No check is made for efficiency's sake, pending benchmarks,
// but adding an error here would significantly worsen the API.
func (p Point2) Dim(i int) int {
	_ = "STUB: not implemented"

	// Dim returns the value of p in the ith dimension.
	// Panics if i > 2. No check is made for efficiency's sake, pending benchmarks,
	// but adding an error here would significantly worsen the API.
	return 0
}

func (p Point3) Dim(i int) int {
	_ = "STUB: not implemented"

	// X returns the point's value on the X axis.
	return 0
}

func (p Point2) X() int {
	_ = "STUB: not implemented"

	// Y returns the point's value on the Y axis.
	return 0
}

func (p Point2) Y() int {
	_ = "STUB: not implemented"

	// X returns the point's value on the X axis.
	return 0
}

func (p Point3) X() int {
	_ = "STUB: not implemented"

	// Y returns the point's value on the Y axis.
	return 0
}

func (p Point3) Y() int {
	_ = "STUB: not implemented"

	// Z returns the point's value on the Z axis.
	return 0
}

func (p Point3) Z() int {
	_ = "STUB: not implemented"

	// Distance calculates the distance between this Point2 and another.
	return 0
}

func (p Point2) Distance(p2 Point2) float64 { _ = "STUB: not implemented"; return 0 }

// Distance calculates the distance between this Point3 and another.
func (p Point3) Distance(p2 Point3) float64 { _ = "STUB: not implemented"; return 0 }

// Distance2 calculates the euclidean distance between two points, as two (x,y) pairs
func Distance2(x1, y1, x2, y2 int) float64 { _ = "STUB: not implemented"; return 0 }

// Distance3 calculates the euclidean distance between two points, as two (x,y,z) triplets
func Distance3(x1, y1, z1, x2, y2, z2 int) float64 { _ = "STUB: not implemented"; return 0 }

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
func (p Point2) MulConst(fs ...int) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Div combines the input points via division.
// Div does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point2) Div(ps ...Point2) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// DivConst divides all elements of a point by the input floats
// DivConst does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point2) DivConst(fs ...int) Point2 { _ = "STUB: not implemented"; return *new(Point2) }

// Add combines the input points via addition.
func (p Point3) Add(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Sub combines the input points via subtraction.
func (p Point3) Sub(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Mul combines in the input points via multiplication.
func (p Point3) Mul(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// MulConst multiplies all elements of a point by the input floats
func (p Point3) MulConst(fs ...int) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Div combines the input points via division.
// Div does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point3) Div(ps ...Point3) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// DivConst divides all elements of a point by the input floats
// DivConst does not check that the inputs are non zero before operating,
// and can panic if that is not true.
func (p Point3) DivConst(fs ...int) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Dot returns the dot product of the input points
func (p Point2) Dot(p2 Point2) int { _ = "STUB: not implemented"; return 0 }

// Dot returns the dot product of the input points
func (p Point3) Dot(p2 Point3) int { _ = "STUB: not implemented"; return 0 }

// Magnitude returns the magnitude of the combined components of a Point
func (p Point2) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// Magnitude returns the magnitude of the combined components of a Point
func (p Point3) Magnitude() float64 { _ = "STUB: not implemented"; return 0 }

// ToRect converts this point into a rectangle spanning span distance
// in each axis.
func (p Point2) ToRect(span int) Rect2 { _ = "STUB: not implemented"; return *new(Rect2) }

// ToRect converts this point into a rectangle spanning span distance
// in each axis.
func (p Point3) ToRect(span int) Rect3 { _ = "STUB: not implemented"; return *new(Rect3) }

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
