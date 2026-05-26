package mouse

import (
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/event"
)

// CollisionPhase is a component that can be placed into another struct to
// enable PhaseCollision on the struct. See PhaseCollision.
type CollisionPhase struct {
	OnCollisionS *collision.Space
	CallerMap    *event.CallerMap
	LastEvent    *Event

	wasTouching bool
}

func (cp *CollisionPhase) getCollisionPhase() *CollisionPhase {
	_ = "STUB: not implemented"
	return nil
}

func (cp *CollisionPhase) CID() event.CallerID {
	_ = "STUB: not implemented"
	return *new(event.CallerID)
}

type collisionPhase interface {
	getCollisionPhase() *CollisionPhase
}

// PhaseCollision binds to the entity behind the space's CID so that it will
// receive MouseCollisionStart and MouseCollisionStop events, appropriately when
// the mouse begins to hover or stops hovering over the input space.
func PhaseCollision(s *collision.Space, handler event.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

// MouseCollisionStart/Stop: see collision Start/Stop, for mouse collision
var (
	Start = event.RegisterEvent[*Event]()
	Stop  = event.RegisterEvent[*Event]()
)

func phaseCollisionEnter(id event.CallerID, handler event.Handler, _ interface{}) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}

// TODO: think about how this can more cleanly work with multiple windows
