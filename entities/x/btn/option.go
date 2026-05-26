package btn

import (
	"image/color"

	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/mouse"

	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
)

// And combines a variadic number of options
func And(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Clear resets the button to be empty
func Clear() Option { _ = "STUB: not implemented"; return *new(Option) }

// Width sets the Width of the button to be generated
func Width(w float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the Height of the button to be generated
func Height(h float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Pos sets the position of the button  to be generated
func Pos(x, y float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Offset increments the position of the button to be generated
func Offset(x, y float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// CID sets the starting CID of the button to be generated
func CID(c event.CallerID) Option { _ = "STUB: not implemented"; return *new(Option) }

// Color sets the colorboxes color for the button to be generated
func Color(c color.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// VGradient creates a vertical color gradient for the btn
func VGradient(c1, c2 color.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// Mod sets the modifications to apply to the initial color box for the button to be generated
func Mod(m mod.Transform) Option { _ = "STUB: not implemented"; return *new(Option) }

// AndMod combines the input modification with whatever existing modifications
// exist for the generator, as opposed to Mod which resets previous modifications.
func AndMod(m mod.Transform) Option { _ = "STUB: not implemented"; return *new(Option) }

// Layers sets the layer of the button to be generated
func Layers(ls ...int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Renderable sets a renderable to use as a base image for the button.
// Not compatible with Color / Toggle.
func Renderable(r render.Modifiable) Option { _ = "STUB: not implemented"; return *new(Option) }

// Binding appends a function to be called when a specific event
// is triggered.
func Binding[Payload any](ev event.EventID[Payload], bnd event.Bindable[*entities.Entity, Payload]) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// TODO: not default

// Click appends a function to be called when the button is clicked on.
func Click(bnd event.Bindable[*entities.Entity, *mouse.Event]) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Label(l collision.Label) Option { _ = "STUB: not implemented"; return *new(Option) }
