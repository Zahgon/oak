// Package format provides audio file and format parsers
package format

import (
	"io"
	"sync"

	"github.com/oakmound/oak/v4/audio/pcm"
)

// A Loader can parse the data from an io.Reader and convert it into PCM encoded audio data with
// a known format.
type Loader func(r io.Reader) (pcm.Reader, error)

var fileLoadersLock sync.RWMutex
var fileLoaders = map[string]func(r io.Reader) (pcm.Reader, error){}

// Register registers a format by file extension (eg '.mp3') with its parsing function.
func Register(extension string, fn Loader) { _ = "STUB: not implemented"; return }

// LoaderForExtension returns a previously registered loader.
func LoaderForExtension(extension string) (Loader, bool) {
	_ = "STUB: not implemented"
	return *new(Loader), false
}
