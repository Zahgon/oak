package render

import (
	"fmt"
	"image/draw"
)

// A Text is a renderable that represents some text to print on screen
type Text struct {
	LayeredPoint
	text fmt.Stringer
	d    *Font
}

// NewStringerText creates a renderable text component that will draw the string
// provided by the given stringer each frame.
func (f *Font) NewStringerText(str fmt.Stringer, x, y float64) *Text {
	_ = "STUB: not implemented"
	return nil
}

type stringerIntPointer struct {
	v *int
}

func (sip stringerIntPointer) String() string { _ = "STUB: not implemented"; return "" }

// NewIntText wraps the given int pointer in a stringer interface
func (f *Font) NewIntText(str *int, x, y float64) *Text { _ = "STUB: not implemented"; return nil }

type stringStringer string

func (ss stringStringer) String() string {
	_ = "STUB: not implemented"

	// NewText creates a renderable text component with the given string body
	return ""
}

func (f *Font) NewText(str string, x, y float64) *Text { _ = "STUB: not implemented"; return nil }

type stringPtrStringer struct {
	s *string
}

func (sp stringPtrStringer) String() string { _ = "STUB: not implemented"; return "" }

// NewStrPtrText creates a renderable text component with a body matching
// and updating to match the content behind the provided string pointer
func (f *Font) NewStrPtrText(str *string, x, y float64) *Text {
	_ = "STUB: not implemented"
	return nil
}

func (t *Text) drawWithFont(buff draw.Image, xOff, yOff float64, fnt *Font) {
	_ = "STUB: not implemented"
	return
}

// Draw for a text draws the text at its layeredPoint position
func (t *Text) Draw(buff draw.Image, xOff, yOff float64) { _ = "STUB: not implemented"; return }

// SetFont sets the drawer which renders the text each frame
func (t *Text) SetFont(f *Font) {
	_ = "STUB: not implemented"

	// GetDims reports the width and height of a text renderable
	return
}

func (t *Text) GetDims() (int, int) {
	_ = "STUB: not implemented"
	// BUG: reported height is too low, test this impl:
	// bounds, adv := t.d.BoundString(t.text.String())
	// return adv.Round(), bounds.Max.Y.Round()
	return 0, 0
}

// Center will shift the text so that the existing leftmost point
// where the text sits becomes the center of the new text.
func (t *Text) Center() { _ = "STUB: not implemented"; return }

// SetString accepts a string itself as the stringer to be written
func (t *Text) SetString(str string) { _ = "STUB: not implemented"; return }

// SetStringPtr accepts a string pointer as the stringer to be written
func (t *Text) SetStringPtr(str *string) { _ = "STUB: not implemented"; return }

// SetStringer accepts an fmt.Stringer to write
func (t *Text) SetStringer(s fmt.Stringer) {
	_ = "STUB: not implemented"

	// SetInt takes and converts the input integer to a string to write
	return
}

func (t *Text) SetInt(i int) { _ = "STUB: not implemented"; return }

// SetIntPtr takes in an integer pointer that will draw the integer
// behind the pointer, in base 10, each frame
func (t *Text) SetIntPtr(i *int) { _ = "STUB: not implemented"; return }

// StringLiteral returns what text is currently rendering.
func (t *Text) StringLiteral() string { _ = "STUB: not implemented"; return "" }

// Wrap returns the input text split into a list of texts
// spread vertically, splitting after each charLimit is reached.
// the input vertInc is how much each text in the slice will differ by
// in y value
func (t *Text) Wrap(charLimit int, vertInc float64) []*Text { _ = "STUB: not implemented"; return nil }

// ToSprite converts this text into a sprite, so that it is no longer
// modifiable in terms of its text content, but is modifiable in terms
// of mod.Transform or mod.Filter.
func (t *Text) ToSprite() *Sprite { _ = "STUB: not implemented"; return nil }
