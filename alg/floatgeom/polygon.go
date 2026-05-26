package floatgeom

// A Polygon2 is a series of points in 2D space.
type Polygon2 struct {
	// Bounding is a cached bounding box calculated from the input points
	// It is exported for convenience, but should be modified with care
	Bounding Rect2
	// The component points of the polygon. If modified, Bounding should
	// be updated with NewBoundingRect2.
	Points      []Point2
	rectangular bool
}

// NewPolygon2 is a helper method to construct a valid polygon. Polygons
// cannot contain less than 3 points.
func NewPolygon2(p1, p2, p3 Point2, pn ...Point2) Polygon2 {
	_ = "STUB: not implemented"
	return *new(Polygon2)
}

// Contains returns whether or not the current Polygon contains the passed in Point.
// If it is known that the polygon is convex, ConvexContains should be preferred for
// performance.
func (pg Polygon2) Contains(x, y float64) (contains bool) { _ = "STUB: not implemented"; return false }

// Three comparisons
// One Comparison, Four add/sub, Two mult/div

// ConvexContains returns whether the given point is contained by the input polygon.
// It assumes the polygon is convex.
func (pg Polygon2) ConvexContains(x, y float64) bool { _ = "STUB: not implemented"; return false }

// TODO: rename this to its real math name, export it
func getSide(a, b Point2) int { _ = "STUB: not implemented"; return 0 }

// OverlappingRectCollides returns whether a Rect2 intersects or is contained by this Polygon.
// This method differs from RectCollides because it assumes that we already know r overlaps with pg.Bounding.
// It is only valid for convex polygons.
func (pg Polygon2) OverlappingRectCollides(r Rect2) bool { _ = "STUB: not implemented"; return false }

// Checking line segment from last to next

// RectCollides returns whether a Rect2 intersects or is contained by this Polygon.
// It is only valid for convex polygons.
func (pg Polygon2) RectCollides(r Rect2) bool { _ = "STUB: not implemented"; return false }

func isRectangular(pts ...Point2) bool { _ = "STUB: not implemented"; return false }

// The last point needs to share an x or y value with this point

func orient(p1, p2, p3 Point2) int8 { _ = "STUB: not implemented"; return 0 }
