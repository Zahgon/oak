package debugtools

import (
	"github.com/oakmound/oak/v4/render"
	"golang.org/x/sync/syncmap"
)

var (
	debugMap syncmap.Map
)

// SetDebugRenderable stores a renderable under a name in a package global map.
// this is used by some built in debugConsole helper functions.
func SetDebugRenderable(rName string, r render.Renderable) { _ = "STUB: not implemented"; return }

// GetDebugRenderable returns whatever renderable is stored under the input
// string, if any.
func GetDebugRenderable(rName string) (render.Renderable, bool) {
	_ = "STUB: not implemented"
	return *new(render.Renderable), false
}

// EnumerateDebugRenderableKeys lists all registered renderables by key.
// It does not check to see if the associated renderables are still valid in any respect.
func EnumerateDebugRenderableKeys() []string { _ = "STUB: not implemented"; return nil }
