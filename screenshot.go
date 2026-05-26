package oak

import (
	"image"
	"image/gif"
)

// ScreenShot takes a snap shot of the window's image content.
// ScreenShot is not safe to call while an existing ScreenShot call has
// yet to finish executing. This could change in the future.
func (w *Window) ScreenShot() *image.RGBA { _ = "STUB: not implemented"; return nil }

// We need to take the shot when the screen is not being redrawn
// We know the screen has everything drawn on it when it is published

// Copy the buffer

// gifShot is internally used by RecordGIF
func (w *Window) gifShot() *image.Paletted { _ = "STUB: not implemented"; return nil }

// We need to take the shot when the screen is not being redrawn
// We know the screen has everything drawn on it when it is published

// Copy the buffer

// RecordGIF will start recording frames via screen shots with the given
// time delay (in 1/100ths of a second) between frames. When the returned
// stop function is called, the frames will be compiled into a gif.
func (w *Window) RecordGIF(hundredths int) (stop func() *gif.GIF) {
	_ = "STUB: not implemented"
	return nil
}
