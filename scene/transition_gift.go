//go:build !nogift
// +build !nogift

package scene

// Fade is a scene transition that fades to black at a given rate for
// a total of 'frames' frames
func Fade(rate float32, frames int) Transition { _ = "STUB: not implemented"; return *new(Transition) }
