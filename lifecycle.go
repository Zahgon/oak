package oak

import (
	"github.com/oakmound/oak/v4/shiny/screen"
)

func (w *Window) lifecycleLoop(s screen.Screen) { _ = "STUB: not implemented"; return }

// Right here, query the backing scale factor of the physical screen
// Apply that factor to the scale

// Quit sends a signal to the window to close itself, closing the window and
// any spun up resources. It should not be called before Init. After it is called,
// it must not be called again.
func (w *Window) Quit() {
	_ = "STUB: not implemented"
	// We could have hit this before the window was created
	return
}

func (w *Window) newWindow(x, y, width, height int) error {
	_ = "STUB: not implemented"
	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.
	return nil
}

// SetAspectRatio will enforce that the displayed window does not distort the
// input screen away from the given x:y ratio. The screen will not use these
// settings until a new size event is received from the OS.
func (w *Window) SetAspectRatio(xToY float64) { _ = "STUB: not implemented"; return }

// ChangeWindow sets the width and height of the game window. Although exported,
// calling it without a size event will probably not act as expected.
func (w *Window) ChangeWindow(width, height int) error {
	_ = "STUB: not implemented"
	// Draw the background to cover up smears
	return nil
}

// UpdateViewSize updates the size of this window's viewport. If the window has yet
// to be initialized, it will update ScreenWidth and ScreenHeight, and then exit.
func (w *Window) UpdateViewSize(width, height int) error { _ = "STUB: not implemented"; return nil }

// this is being called before Init
