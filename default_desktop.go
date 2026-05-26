//go:build (windows || linux || osx) && !js && !android && !nooswindow
// +build windows linux osx
// +build !js
// +build !android
// +build !nooswindow

package oak

import (
	"image"
)

// MoveWindow calls MoveWindow on the default window.
func MoveWindow(x, y, w, h int) error { _ = "STUB: not implemented"; return nil }

// SetFullScreen calls SetFullScreen on the default window.
func SetFullScreen(fs bool) error { _ = "STUB: not implemented"; return nil }

// SetBorderless calls SetBorderless on the default window.
func SetBorderless(bs bool) error { _ = "STUB: not implemented"; return nil }

// SetTopMost calls SetTopMost on the default window.
func SetTopMost(on bool) error { _ = "STUB: not implemented"; return nil }

// SetTitle calls SetTitle on the default window.
func SetTitle(title string) error { _ = "STUB: not implemented"; return nil }

// SetIcon calls SetIcon on the default window.
func SetIcon(icon image.Image) error { _ = "STUB: not implemented"; return nil }

// HideCursor calls HideCursor on the default window.
func HideCursor() error { _ = "STUB: not implemented"; return nil }

// GetCursorPosition calls GetCursorPosition on the default window.
func GetCursorPosition() (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }
