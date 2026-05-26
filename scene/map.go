package scene

import (
	"sync"
)

// A Map lets scenes be accessed via associated names.
type Map struct {
	CurrentScene string
	scenes       map[string]Scene
	// This could be a RWMutex, but it isn't anticipated that
	// reads will be more common than writes.
	lock sync.Mutex
}

// NewMap creates a scene map
func NewMap() *Map { _ = "STUB: not implemented"; return nil }

// Get returns the scene associated with the given name, if it exists. If it
// does not exist, it returns a zero value and false.
func (m *Map) Get(name string) (Scene, bool) { _ = "STUB: not implemented"; return *new(Scene), false }

// GetCurrent returns the current scene, as defined by map.CurrentScene
func (m *Map) GetCurrent() (Scene, bool) { _ = "STUB: not implemented"; return *new(Scene), false }

// AddScene takes a scene struct, checks that its assigned name does not
// conflict with an existing name in the map, and then adds it to the map.
// If a conflict occurs, the scene will not be overwritten.
// Checks if the Scene's start is nil, sets to noop if so.
// Checks if the Scene's end is nil, sets to loop to this scene if so.
func (m *Map) AddScene(name string, s Scene) error { _ = "STUB: not implemented"; return nil }
