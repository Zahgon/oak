package synth

import (
	"time"

	"github.com/oakmound/oak/v4/audio/pcm"
)

// A Source stores necessary information for generating waveform data
type Source struct {
	pcm.Format
	Pitch Pitch
	// Volume, between 0.0 -> 1.0
	Volume  float64
	Seconds float64
}

// PlayLength returns the time it will take before audio generated from this
// source will stop.
func (s Source) PlayLength() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Phase is shorthand for phase(s.Pitch, i, s.SampleRate).
// Some sources might have custom phase functions in the future, however.
func (s Source) Phase(i int) float64 { _ = "STUB: not implemented"; return 0 }

// Update is shorthand for applying a set of options to a source
func (s Source) Update(opts ...Option) Source { _ = "STUB: not implemented"; return *new(Source) }

var (
	// Int16 is a default source for building 16-bit audio
	Int16 = Source{
		Format: pcm.Format{
			SampleRate: 44100,
			Channels:   2,
			// within a source, if Bits is not specified, it'll default to 16.
			Bits: 16,
		},
		Pitch:   A4,
		Volume:  .25,
		Seconds: 1,
	}
)
