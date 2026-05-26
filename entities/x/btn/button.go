package btn

import (
	"fmt"
	"image/color"

	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/shape"
)

// A Generator defines the variables used to create buttons from optional arguments
type Generator struct {
	X, Y         float64
	W, H         float64
	TxtX, TxtY   float64
	Color        color.Color
	Color2       color.Color
	ProgressFunc func(x, y, w, h int) float64
	Mod          mod.Transform
	R            render.Modifiable
	R1           render.Modifiable
	R2           render.Modifiable
	RS           []render.Modifiable
	Cid          event.CallerID
	Font         *render.Font
	Layers       []int
	Text         string
	TextPtr      *string
	TextStringer fmt.Stringer
	Children     []Generator
	Bindings     []func(ctx *scene.Context, caller *entities.Entity) event.Binding
	Trigger      string
	Shape        shape.Shape
	Label        collision.Label
}

func defGenerator() Generator {
	_ = "STUB: not implemented"
	// A number of these fields could be removed, because they are the zero
	// value, but are left for documentation
	return *new(Generator)
}

// Generate creates a Button from a generator.
func (g Generator) Generate(ctx *scene.Context) *entities.Entity {
	_ = "STUB: not implemented"
	return nil

	// handle different renderable options that could be passed to the generator
}

// An Option is used to populate generator fields prior to generation of a button
type Option func(Generator) Generator

// New creates a button with the given options and defaults for all variables not set.
func New(ctx *scene.Context, opts ...Option) *entities.Entity {
	_ = "STUB: not implemented"
	return nil
}
