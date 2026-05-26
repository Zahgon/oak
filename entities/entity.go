package entities

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)

type Generator struct {
	Position   floatgeom.Point2
	Dimensions floatgeom.Point2
	Speed      floatgeom.Point2

	Parent event.Caller

	Color      color.Color
	Renderable render.Renderable

	Mod mod.Mod

	Label collision.Label

	DrawLayers []int

	UseMouseTree     bool
	WithoutCollision bool

	Children         [][]Option
	ExplicitChildren []*Entity
}

func And(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithChild(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExplicitChild(e *Entity) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRect(v floatgeom.Rect2) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOffset(p floatgeom.Point2) Option { _ = "STUB: not implemented"; return *new(Option) }

var defaultGenerator = Generator{
	Dimensions: floatgeom.Point2{1, 1},
	DrawLayers: []int{0},
}

type Entity struct {
	event.CallerID

	ctx *scene.Context

	Rect  floatgeom.Rect2
	Speed floatgeom.Point2
	Delta floatgeom.Point2

	Renderable render.Renderable

	collision.Phase

	Space *collision.Space
	Tree  *collision.Tree

	metadata map[string]string

	Children []*Entity
}

func (e Entity) CID() event.CallerID { _ = "STUB: not implemented"; return *new(event.CallerID) }

func (e Entity) X() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) Y() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) W() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) H() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) Top() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) Bottom() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) Left() float64 { _ = "STUB: not implemented"; return 0 }

func (e Entity) Right() float64 { _ = "STUB: not implemented"; return 0 }

func (e *Entity) ShiftDelta() { _ = "STUB: not implemented"; return }

func (e *Entity) Shift(delta floatgeom.Point2) {
	_ = "STUB: not implemented"
	// TODO: attachment?
	// TODO: helper
	return
}

func (e *Entity) SetX(x float64) { _ = "STUB: not implemented"; return }

func (e *Entity) SetY(y float64) { _ = "STUB: not implemented"; return }

func (e *Entity) ShiftX(x float64) { _ = "STUB: not implemented"; return }

func (e *Entity) ShiftY(y float64) { _ = "STUB: not implemented"; return }

func (e *Entity) SetPos(p floatgeom.Point2) { _ = "STUB: not implemented"; return }

func (e *Entity) ShiftPos(x, y float64) { _ = "STUB: not implemented"; return }

func (e *Entity) HitLabel(label collision.Label) *collision.Space {
	_ = "STUB: not implemented"
	return nil
}

func (e *Entity) Destroy() { _ = "STUB: not implemented"; return }

// SetMetadata sets the metadata for some key to some value. Empty value strings
// will not be stored.
func (e *Entity) SetMetadata(k, v string) { _ = "STUB: not implemented"; return }

// Metadata accesses the value, and whether it existed, for a given metadata key
func (e *Entity) Metadata(k string) (v string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func New(ctx *scene.Context, opts ...Option) *Entity { _ = "STUB: not implemented"; return nil }
