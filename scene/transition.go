package scene

import (
	"image"
)

// Transition functions can be set to occur at the end of a scene.
type Transition func(*image.RGBA, int) bool

// Zoom transitions by performing a simplistic zoom each frame towards some
// percentage-based part of the screen.
func Zoom(xPerc, yPerc float64, frames int, zoomRate float64) Transition {
	_ = "STUB: not implemented"
	return *new(Transition)
}
