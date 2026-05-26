package render

import (
	"time"

	"github.com/oakmound/oak/v4/event"
)

// LogicFPS is a Stackable that will draw the logical fps onto the screen when a part
// of the draw stack.
type LogicFPS struct {
	event.CallerID
	*Text
	fps       int
	lastTime  time.Time
	Smoothing float64
}

func (lf LogicFPS) CID() event.CallerID {
	_ = "STUB: not implemented"

	// NewLogicFPS returns a LogicFPS, which will render a counter of how fast it receives event.Enter events.
	// If font is not provided, DefaultFont is used. If smoothing is 0, a reasonable default is used.
	return *new(event.CallerID)
}

func NewLogicFPS(smoothing float64, font *Font, x, y float64) *LogicFPS {
	_ = "STUB: not implemented"
	return nil
}

// TODO: not default bus

func logicFPSBind(lf *LogicFPS, _ event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return *new(event.Response)
}
