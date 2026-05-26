package render

import (
	"image"
	"image/draw"
	"time"

	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render/mod"
)

// A Sequence is a series of modifiables drawn as an animation. It is more
// primitive than animation, but less efficient.
type Sequence struct {
	LayeredPoint
	pauseBool
	InterruptBool
	rs         []Modifiable
	lastChange time.Time
	sheetPos   int
	frameTime  int64
	event.CallerID
}

// NewSequence returns a new sequence from the input modifiables, playing at
// the given fps rate.
func NewSequence(fps float64, mods ...Modifiable) *Sequence { _ = "STUB: not implemented"; return nil }

// SetFPS sets the number of frames that should advance per second to be
// the input fps
func (sq *Sequence) SetFPS(fps float64) { _ = "STUB: not implemented"; return }

// GetDims of a Sequence returns the dims of the current Renderable for the sequence
func (sq *Sequence) GetDims() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Copy copies each modifiable inside this sequence in order to produce a new
// copied sequence
func (sq *Sequence) Copy() Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

var AnimationEnd = event.RegisterEvent[struct{}]()

// SetTriggerID sets the ID that AnimationEnd will be triggered on when this
// sequence loops over from its last frame to its first
func (sq *Sequence) SetTriggerID(id event.CallerID) { _ = "STUB: not implemented"; return }

func (sq *Sequence) update() { _ = "STUB: not implemented"; return }

// TODO: not default bus

// Get returns the Modifiable stored at this sequence's ith index. If the sequence
// does not have an ith index this returns nil
func (sq *Sequence) Get(i int) Modifiable { _ = "STUB: not implemented"; return *new(Modifiable) }

// Draw draws this sequence at +xOff, +yOff
func (sq *Sequence) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// GetRGBA returns the RGBA of the currently showing frame of this sequence
func (sq *Sequence) GetRGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }

// Modify alters each renderable in this sequence by the given
// modifications
func (sq *Sequence) Modify(ms ...mod.Mod) Modifiable {
	_ = "STUB: not implemented"
	return *new(Modifiable)
}

// Filter filters each element in the sequence by the inputs
func (sq *Sequence) Filter(fs ...mod.Filter) { _ = "STUB: not implemented"; return }

// IsStatic returns false for sequences
func (sq *Sequence) IsStatic() bool {
	_ = "STUB: not implemented"

	// TweenSequence returns a sequence that is the tweening between the input images
	// at the given frame rate over the given frame count.
	return false
}

func TweenSequence(a, b image.Image, frames int, fps float64) *Sequence {
	_ = "STUB: not implemented"
	return nil
}
