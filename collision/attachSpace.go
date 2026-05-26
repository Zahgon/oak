package collision

import (
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/physics"
)

// An AttachSpace is a composable struct that provides attachment
// functionality for entities. An entity with AttachSpace can have its
// associated space passed into Attach with the vector the space should
// be attached to.
// Example usage: Any moving character with a collision space. When
// moving the character around by the vector passed in to Attach, the space
// will move with it.
type AttachSpace struct {
	follow     physics.Vector
	aSpace     **Space
	tree       *Tree
	offX, offY float64
	binding    event.Binding
}

func (as *AttachSpace) getAttachSpace() *AttachSpace { _ = "STUB: not implemented"; return nil }

func (as *AttachSpace) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

var _ attachSpace = &AttachSpace{}

type attachSpace interface {
	event.Caller
	getAttachSpace() *AttachSpace
}

// Attach attaches v to the given space with optional x,y offsets. See AttachSpace.
func Attach(v physics.Vector, s *Space, tree *Tree, offsets ...float64) error {
	_ = "STUB: not implemented"
	return nil
}

func AttachWithBus(v physics.Vector, s *Space, tree *Tree, bus event.Handler, offsets ...float64) error {
	_ = "STUB: not implemented"
	return nil
}

// Detach removes the attachSpaceEnter binding from an entity composed with
// AttachSpace
func Detach(s *Space) error { _ = "STUB: not implemented"; return nil }

func DetachWithBus(s *Space, bus event.Handler) error { _ = "STUB: not implemented"; return nil }

func attachSpaceEnter(asIface attachSpace, _ event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}
