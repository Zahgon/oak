package floatgeom

// Tri3 is a triangle of Point3s
type Tri3 [3]Point3

// Barycentric finds the barycentric coordinates of the given x,y cartesian
// coordinates within this triangle. If the point (x,y) is outside of the
// triangle, one of the output values will be negative.
// Credit goes to github.com/yellingintothefan for their work in gel
func (t Tri3) Barycentric(x, y float64) Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Normal calculates the surface normal of a triangle
func (t Tri3) Normal() Point3 { _ = "STUB: not implemented"; return *new(Point3) }

// Check that the triangle is defined in a clockwise fashion.
