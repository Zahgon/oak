// Package timing provides utilities for time.
package timing

import (
	"time"
)

const (
	nanoPerSecond = 1000000000

	maximumFPS = 1200
)

// FPS returns the number of frames being processed per second,
// supposing a time interval from lastTime to now.
func FPS(lastTime, now time.Time) float64 { _ = "STUB: not implemented"; return 0 }

// This indicates that we recorded two times within
// the inaccuracy of the OS's system clock, so the values
// were the same. 1200 is chosen because on windows,
// fps will be 1200 instead of a negative value.

// FPSToNano converts a framesPerSecond value to the number of
// nanoseconds that should take place for each frame.
func FPSToNano(fps float64) int64 { _ = "STUB: not implemented"; return 0 }

// FPSToFrameDelay converts a frameRate like 60fps into a delay time between frames
func FPSToFrameDelay(frameRate int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// FrameDelayToFPS converts a duration of delay between frames into a frames per second count
func FrameDelayToFPS(dur time.Duration) float64 { _ = "STUB: not implemented"; return 0 }
