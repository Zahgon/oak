//go:build js
// +build js

package oak

import (
	"syscall/js"
)

func overrideInit(w *Window) { _ = "STUB: not implemented"; return }

func (w *Window) requestFrame(this js.Value, args []js.Value) interface{} {
	_ = "STUB: not implemented"
	return nil
}
