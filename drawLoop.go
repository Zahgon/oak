package oak

import (
	"image"
)

// A Background can be used as a background draw layer. Backgrounds will be drawn as the first
// element in each frame, and are expected to cover up data drawn on the previous frame.
type Background interface {
	GetRGBA() *image.RGBA
}

// DrawLoop
// Unless told to stop, the draw channel will repeatedly
// 1. draw the background color to a temporary buffer
// 2. draw all visible rendered elements onto the temporary buffer.
// 3. draw the buffer's data at the viewport's position to the screen.
// 4. publish the screen to display in window.
func (w *Window) drawLoop() { _ = "STUB: not implemented"; return }

// Publish what was drawn last frame to screen, then work on preparing the next frame.

// this code is duplicated as an optimization: it's much faster
// to have a 'default' case than to flood a channel, and we can't conditionally
// add a default case to a select.

func (w *Window) publish() { _ = "STUB: not implemented"; return }

// every frame, swap buffers. This enables drivers which might hold on to the rgba buffers we publish as if they
// were immutable.

// DoBetweenDraws will execute the given function in-between draw frames. It will prevent draws from happening until
// the provided function has terminated. DoBetweenDraws will block until the provided function is called within the
// draw loop's schedule, but will not wait for that function itself to terminate.
func (w *Window) DoBetweenDraws(f func()) { _ = "STUB: not implemented"; return }
