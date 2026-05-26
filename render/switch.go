package render

import (
	"image"
	"image/draw"
	"sync"

	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/physics"
	"github.com/oakmound/oak/v4/render/mod"
)

// The Switch type is intended for use to easily swap between multiple
// renderables that are drawn at the same position on the same layer.
// A common use case for this would be a character entitiy who switches
// their animation based on how they are moving or what they are doing,
// or a button that has changes state when selected or hovered over.
type Switch struct {
	LayeredPoint
	subRenderables map[string]Modifiable
	curRenderable  string
	lock           sync.RWMutex
}

// NewSwitch creates a new Switch from a map of names to modifiables
func NewSwitch(start string, m map[string]Modifiable) *Switch {
	_ = "STUB: not implemented"
	return nil
}

// Add makes a new entry in the Switch's map. If the key already
// existed, it will be overwritten and an error will be returned.
func (c *Switch) Add(k string, v Modifiable) (err error) { _ = "STUB: not implemented"; return nil }

// Set sets the current renderable to the one specified
func (c *Switch) Set(k string) error { _ = "STUB: not implemented"; return nil }

// GetSub returns a keyed Modifiable from this Switch's map
func (c *Switch) GetSub(s string) Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

// Get returns the Switch's current key
func (c *Switch) Get() string { _ = "STUB: not implemented"; return "" }

// SetOffsets sets the logical offset for the specified key
func (c *Switch) SetOffsets(k string, offsets physics.Vector) { _ = "STUB: not implemented"; return }

// Copy creates a copy of the Switch
func (c *Switch) Copy() Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

// GetRGBA returns the current renderables rgba
func (c *Switch) GetRGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }

// Modify performs the input modifications on all elements of the Switch
func (c *Switch) Modify(ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Filter filters all elements of the Switch with fs
func (c *Switch) Filter(fs ...mod.Filter) { _ = "STUB: not implemented"; return }

// Draw draws the Switch at an offset from its logical location
func (c *Switch) Draw(buff draw.Image, xOff float64, yOff float64) {
	_ = "STUB: not implemented"
	return
}

// ShiftPos shifts the Switch's logical position
func (c *Switch) ShiftPos(x, y float64) { _ = "STUB: not implemented"; return }

// GetDims gets the current Renderables dimensions
func (c *Switch) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Pause stops the current Renderable if possible
func (c *Switch) Pause() { _ = "STUB: not implemented"; return }

// Unpause tries to unpause the current Renderable if possible
func (c *Switch) Unpause() { _ = "STUB: not implemented"; return }

// IsInterruptable returns whether the current renderable is interruptable
func (c *Switch) IsInterruptable() bool { _ = "STUB: not implemented"; return false }

// IsStatic returns whether the current renderable is static
func (c *Switch) IsStatic() bool { _ = "STUB: not implemented"; return false }

// SetTriggerID sets the ID AnimationEnd will trigger on for animating subtypes.
// Todo: standardize this with the other interface Set functions so that it
// also only acts on the current subRenderable, or the other way around, or
// somehow offer both options
func (c *Switch) SetTriggerID(cid event.CallerID) { _ = "STUB: not implemented"; return }

// Revert will revert all parts of this Switch that can be reverted
func (c *Switch) Revert(mod int) { _ = "STUB: not implemented"; return }

// RevertAll will revert all parts of this Switch that can be reverted, back
// to their original state.
func (c *Switch) RevertAll() { _ = "STUB: not implemented"; return }

func (c *Switch) update() { _ = "STUB: not implemented"; return }
