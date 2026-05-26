package debugstream

import (
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/window"
)

// AddDefaultsForScope for debugging.
func (sc *ScopedCommands) AddDefaultsForScope(scopeID int32, controller window.Window) {
	_ = "STUB: not implemented"
	return
}

// assume the scope for easy usage here

func moveWindow(w window.Window) func([]string) string { _ = "STUB: not implemented"; return nil }

const explainFullScreen = "specify off 'fullscreen off' to exit fullscreen"

func fullScreen(w window.Window) func([]string) string { _ = "STUB: not implemented"; return nil }

const explainMouseDetails = "the mext mouse click on the given window will print the cursor's location"

func mouseCommands(w window.Window) func([]string) string { _ = "STUB: not implemented"; return nil }

func mouseDetails(w window.Window) func(*mouse.Event) event.Response {
	_ = "STUB: not implemented"
	return nil
}

const explainQuit = "close the given window"

func quitCommands(w window.Window) func([]string) string { _ = "STUB: not implemented"; return nil }

func skipCommands(w window.Window) func([]string) string { _ = "STUB: not implemented"; return nil }

const explainFade = "fade the specified renderable by the given int if given. Renderable must be registered in debugtools"

func fadeCommands(tokenString []string) (out string) { _ = "STUB: not implemented"; return "" }

func parseTokenAsInt(tokenString []string, arrIndex int, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}
