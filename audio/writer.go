package audio

import (
	"github.com/oakmound/oak/v4/audio/pcm"
)

// NewWriter returns a writer which can accept audio streamed matching the given format
func NewWriter(f pcm.Format) (pcm.Writer, error) {
	_ = "STUB: not implemented"
	return *

	// MustNewWriter calls NewWriter and panics if an error is returned.
	new(pcm.Writer), nil
}

func MustNewWriter(f pcm.Format) pcm.Writer { _ = "STUB: not implemented"; return *new(pcm.Writer) }
