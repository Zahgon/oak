package oak

import (
	"github.com/oakmound/oak/v4/scene"
)

// AddScene is shorthand for w.SceneMap.AddScene
func (w *Window) AddScene(name string, s scene.Scene) error { _ = "STUB: not implemented"; return nil }

func (w *Window) sceneTransition(result *scene.Result) { _ = "STUB: not implemented"; return }

// TODO: Transition should take in the amount of time passed, not number of frames,
// to account for however long the transition itself takes.
