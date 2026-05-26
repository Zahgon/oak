//go:build windows
// +build windows

package win32

import "sync"

var emptyCursor HCURSOR
var emptyCursorOnce sync.Once

// Create a custom cursor at run time.
func GetEmptyCursor() HCURSOR { _ = "STUB: not implemented"; return *new(HCURSOR) }

// app. instance
// horizontal position of hot spot
// vertical position of hot spot
// 0 width/height is unsupported in testing
// cursor width
// cursor height

// TODO: Add image.Image to cursor conversion and setting functionality
// this can currently be done in oak by having a image follow the cursor around,
// but that will inherently not be as smooth as setting the OS cursor. (but more portable)
