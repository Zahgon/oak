package collision

import (
	"sync"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/oakerr"
)

// A Tree provides a space for managing collisions between rectangles
type Tree struct {
	*Rtree
	sync.Mutex
}

const (
	defaultMinChildren = 20
	defaultMaxChildren = 40
)

// NewTree returns a new collision Tree. defaultMinChildren and defaultMaxChildren
// are used for node sizing.
func NewTree() *Tree { _ = "STUB: not implemented"; return nil }

// NewCustomTree returns a new collision Tree with custom node sizes.
// minChildren must be less than maxChildren.
func NewCustomTree(minChildren, maxChildren int) (*Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clear resets a tree's contents to be empty
func (t *Tree) Clear() { _ = "STUB: not implemented"; return }

// Add adds a set of spaces to the rtree
func (t *Tree) Add(sps ...*Space) { _ = "STUB: not implemented"; return }

// Remove removes spaces from the rtree and
// returns the number of spaces removed.
func (t *Tree) Remove(sps ...*Space) int { _ = "STUB: not implemented"; return 0 }

// UpdateLabel will set the input space's label. DEPRECATED. Just set
// the Label field on the Space pointer.
func (t *Tree) UpdateLabel(classtype Label, s *Space) { _ = "STUB: not implemented"; return }

// ErrNotExist is returned by methods on spaces
// when the space to update or act on did not exist
var ErrNotExist = oakerr.NotFound{InputName: "Space"}

// UpdateSpace is not an operation on a space because
// a space can exist in multiple trees.

// UpdateSpace resets a space's location to a given
// rect.
func (t *Tree) UpdateSpace(x, y, w, h float64, s *Space) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateSpaceRect acts as UpdateSpace, but takes in a rectangle instead
// of four distinct arguments.
func (t *Tree) UpdateSpaceRect(rect floatgeom.Rect3, s *Space) error {
	_ = "STUB: not implemented"
	return nil
}

// ShiftSpace adds x and y to a space and updates its position
func (t *Tree) ShiftSpace(x, y float64, s *Space) error { _ = "STUB: not implemented"; return nil }

// Hits returns the set of spaces which are colliding
// with the passed in space. All spaces collide with
// themselves, if they exist in the tree, but self-collision
// will not be reported by Hits.
func (t *Tree) Hits(sp *Space) []*Space { _ = "STUB: not implemented"; return nil }

// HitLabel acts like Hits, but returns the first space within hits
// that matches one of the input labels. HitLabel can return the same
// space that is passed into it, if that space has a label in the set of
// accepted labels.
func (t *Tree) HitLabel(sp *Space, labels ...Label) *Space { _ = "STUB: not implemented"; return nil }

// Hit is an experimental new syntax that probably has performance hits
// relative to Hits/HitLabel, see filters.go
func (t *Tree) Hit(sp *Space, fs ...Filter) []*Space { _ = "STUB: not implemented"; return nil }
