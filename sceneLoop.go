package oak

// the oak loading scene is a reserved scene
// for preloading assets
const oakLoadingScene = "oak:loading"

func (w *Window) sceneLoop(first string, trackingInputs bool) { _ = "STUB: not implemented"; return }

// kick start the draw loop

// Post transition, begin loading animation

// Send a signal to resume (or begin) drawing

// We don't want enterFrames going off between scenes

// Send a signal to stop drawing

// Reset transient portions of the engine
// We start by clearing the event bus to
// remove most ongoing code

// We follow by clearing collision areas
// because otherwise collision function calls
// on non-entities (i.e. particles) can still
// be triggered and attempt to access an entity

// Todo: Add in customizable loading scene between regular scenes,
// In addition to the existing customizable loading renderable?

// For convenience, we allow the user to return nil
// but it gets translated to an empty result
