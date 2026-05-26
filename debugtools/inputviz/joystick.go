package inputviz

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/dlog"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/joystick"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

//go:embed controllerOutline.png
var controllerOutline []byte

var pngOutline image.Image

func init() {
	var err error
	pngOutline, err = png.Decode(bytes.NewBuffer(controllerOutline))
	if err != nil {
		dlog.Error("failed to decode background data: %w", err)
	}
}

// Joystick visualizes the inputs sent to a controller
type Joystick struct {
	// Rect is the rect this joystick should be drawn to.
	// Defaults to (0,0)->(320,240)
	Rect floatgeom.Rect2

	// StickDeadzone is the lowest value of stick movements that should
	// be rendered.
	StickDeadzone int16

	// BaseLayer is the base layer to render the resulting renderables to
	// if -1, it will render only to the layer provided to RenderAndListen.
	BaseLayer int

	ctx *scene.Context
	event.CallerID
	joy          *joystick.Joystick
	rs           map[string]render.Modifiable
	lastState    *joystick.State
	triggerY     float64
	lStickCenter floatgeom.Point2
	rStickCenter floatgeom.Point2
	cancel       func()

	bindings []event.Binding
}

func (j *Joystick) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

func (j *Joystick) RenderAndListen(ctx *scene.Context, joy *joystick.Joystick, layer int) error {
	_ = "STUB: not implemented"
	return nil
}

// Draw the triggers behind the outline to simulate pressing down

// adjust all offsets

// TODO: it is bad that you need to import two 'key' packages

func (j *Joystick) Destroy() { _ = "STUB: not implemented"; return }
