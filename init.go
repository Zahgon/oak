package oak

import (
	"image"
)

var (
	zeroPoint = image.Point{0, 0}
)

// Init initializes the oak engine.
// After the configuration options have been parsed and validated, this will run concurrent
// routines drawing to an OS window or app, forwarding OS inputs to this window's configured
// event handler, and running scenes: first the predefined 'loading' scene, then firstScene
// as provided here, then scenes following commands sent to the window or returned by ending
// scenes.
func (w *Window) Init(firstScene string, configOptions ...ConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

// This error cannot happen as it would surface in Parse above

// assume we are in focus on window creation

// seed math/rand with time.Now, useful for minimal examples
//that would tend to forget to do this.
