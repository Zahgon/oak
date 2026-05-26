package collision

import (
	"github.com/oakmound/oak/v4/event"
)

// A Phase is a struct that other structs who want to use PhaseCollision
// should be composed of
type Phase struct {
	OnCollisionS *Space
	tree         *Tree
	bus          event.Handler
	// If allocating maps becomes an issue
	// we can have two constant maps that we
	// switch between on alternating frames
	Touching map[Label]bool
}

func (cp *Phase) getCollisionPhase() *Phase { _ = "STUB: not implemented"; return nil }

func (cp *Phase) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

type collisionPhase interface {
	getCollisionPhase() *Phase
}

// PhaseCollision binds to the entity behind the space's CID so that it will
// receive CollisionStart and CollisionStop events, appropriately when
// entities begin to collide or stop colliding with the space.
// If tree is nil, it uses DefTree
func PhaseCollision(s *Space, tree *Tree) error { _ = "STUB: not implemented"; return nil }

// PhaseCollisionWithBus allows for a non-default bus in a phase collision binding.
func PhaseCollisionWithBus(s *Space, tree *Tree, bus event.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

// CollisionStart/Stop: when a PhaseCollision entity starts/stops touching some label.
var (
	Start = event.RegisterEvent[Label]()
	Stop  = event.RegisterEvent[Label]()
)

func phaseCollisionEnter(id event.CallerID, handler event.Handler, _ interface{}) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}

// check hits

// if any are new, trigger on collision

// if we lost any, trigger off collision
