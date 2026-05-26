package btn

import (
	"fmt"

	"github.com/oakmound/oak/v4/render"
)

// Text sets the text of the button to be generated
func Text(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// TextPtr sets the text of the button to be generated
// to a string pointer.
func TextPtr(s *string) Option { _ = "STUB: not implemented"; return *new(Option) }

// TextStringer sets the text of the generated button to
// use a fmt.Stringer String call
func TextStringer(s fmt.Stringer) Option { _ = "STUB: not implemented"; return *new(Option) }

// Font sets the font for the text of the button to be generated
func Font(f *render.Font) Option { _ = "STUB: not implemented"; return *new(Option) }

// TxtOff sets the text offset  of the button generator from the bottom left
func TxtOff(x, y float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// FitText adjusts a btn's width, given it has text and font defined, to
// be large enough for the given text plus the provided buffer
func FitText(buffer int) Option { _ = "STUB: not implemented"; return *new(Option) }
