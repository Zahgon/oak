// Package flac provides functionality to handle .flac files and .flac encoded data.
//
// This package may be imported solely to register flacs as a parseable file type within oak:
//
//	import (
//	    _ "github.com/oakmound/oak/v4/audio/format/flac"
//	)
package flac

import (
	"io"

	"github.com/eaburns/flac"
	"github.com/oakmound/oak/v4/audio/format"
	"github.com/oakmound/oak/v4/audio/pcm"
)

func init() {
	format.Register(".flac", Load)
}

// Load reads a FLAC header from a reader, parsing it's PCM format and returning
// a pcm Reader for the data following the header. It will error if the reader
// does not contain enough data to fill a FLAC header or if the header does not
// look like a FLAC header.
func Load(r io.Reader) (pcm.Reader, error) { _ = "STUB: not implemented"; return *new(pcm.Reader), nil }

type reader struct {
	d         *flac.Decoder
	readAhead []byte
}

func (r *reader) Read(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
