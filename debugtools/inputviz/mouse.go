// Package inputviz provides components that enable visualization of user input (e.g. mouse, keyboard) for debugging
package inputviz

import (
	"sync"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type Mouse struct {
	Rect      floatgeom.Rect2
	BaseLayer int

	event.CallerID
	ctx *scene.Context

	rs map[mouse.Button]*render.Switch

	lastMousePos *posStringer
	posText      *render.Text

	stateIncLock sync.RWMutex
	stateInc     map[mouse.Button]int

	bindings []event.Binding
}

func (m *Mouse) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

func (m *Mouse) RenderAndListen(ctx *scene.Context, layer int) error {
	_ = "STUB: not implemented"
	return nil
}

type posStringer struct {
	floatgeom.Point2
}

func (ps *posStringer) String() string { _ = "STUB: not implemented"; return "" }

func (m *Mouse) Destroy() {
	_ = "STUB: not implemented"
	// TODO: this is a lot of code to write to track and unbind all of an entity's bindings
	return
}
