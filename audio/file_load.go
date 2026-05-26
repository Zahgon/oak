package audio

import (
	"github.com/oakmound/oak/v4/audio/pcm"
)

// Get will read cached audio data from Load, or error if the given
// file is not in the cache.
func (c *Cache) Get(file string) (pcm.Reader, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Reader), nil
}

// Load loads the given file and caches it by two keys:
// the full file name given and the final element of the file's
// path. If the file cannot be found or if its extension is not
// supported an error will be returned.
func (c *Cache) Load(file string) (pcm.Reader, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Reader), nil
}

// provide an error message suggesting a missing import for cases where we know about a
// common provider

// BatchLoad attempts to load all audio files within a given directory
// should their file ending match a registered audio file parser
func BatchLoad(baseFolder string) error { _ = "STUB: not implemented"; return nil }

// BlankBatchLoad acts like BatchLoad, but replaces all loaded assets
// with empty audio constructs. This is intended to reduce start-up
// times in development.
func BlankBatchLoad(baseFolder string) error { _ = "STUB: not implemented"; return nil }

func batchLoad(baseFolder string, blankOut bool) error { _ = "STUB: not implemented"; return nil }

func blankLoad(filename string) { _ = "STUB: not implemented"; return }
