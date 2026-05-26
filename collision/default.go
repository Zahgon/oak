package collision

// DefaultTree is a collision tree intended to be used by default if no other
// is instantiated. Methods on a collision tree are duplicated as functions
// in this package, so `tree.Add(...)` can instead be `collision.Add(...)` if
// the codebase is coordinated to just use the default tree.
var (
	DefaultTree = NewTree()
)

// Clear resets the default tree's contents
func Clear() { _ = "STUB: not implemented"; return }

// Add adds a set of spaces to the rtree
func Add(sps ...*Space) { _ = "STUB: not implemented"; return }

// Remove removes a space from the rtree
func Remove(sps ...*Space) { _ = "STUB: not implemented"; return }

// UpdateSpace resets a space's location to a given rect.
func UpdateSpace(x, y, w, h float64, s *Space) error { _ = "STUB: not implemented"; return nil }

// ShiftSpace adds x and y to a space and updates its position
// in the collision rtree that should not be a package global
func ShiftSpace(x, y float64, s *Space) error { _ = "STUB: not implemented"; return nil }

// Hits returns the set of spaces which are colliding
// with the passed in space.
func Hits(sp *Space) []*Space { _ = "STUB: not implemented"; return nil }

// HitLabel acts like hits, but reutrns the first space within hits
// that matches one of the input labels
func HitLabel(sp *Space, labels ...Label) *Space { _ = "STUB: not implemented"; return nil }

// Update updates this space with the default rtree
func (s *Space) Update(x, y, w, h float64) error { _ = "STUB: not implemented"; return nil }

// SetDim sets the dimensions of the space in the default rtree
func (s *Space) SetDim(w, h float64) error { _ = "STUB: not implemented"; return nil }

// UpdateLabel changes the label behind this space and resets
// it in the default rtree
func (s *Space) UpdateLabel(classtype Label) { _ = "STUB: not implemented"; return }
