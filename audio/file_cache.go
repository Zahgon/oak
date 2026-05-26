package audio

import (
	"sync"

	"github.com/oakmound/oak/v4/audio/pcm"
)

// DefaultCache is the receiver for package level loading operations.
var DefaultCache = NewCache()

// Cache is a simple audio data cache
type Cache struct {
	mu   sync.RWMutex
	data map[string]*BytesReader
}

// NewCache returns an empty Cache
func NewCache() *Cache { _ = "STUB: not implemented"; return nil }

// ClearAll will remove all elements from a Cache
func (c *Cache) ClearAll() { _ = "STUB: not implemented"; return }

// Clear will remove elements matching the given key from the Cache.
func (c *Cache) Clear(key string) { _ = "STUB: not implemented"; return }

func (c *Cache) setLoaded(file string, r pcm.Reader) {
	_ = "STUB: not implemented"
	// This ReadAll and .Copy() on Cache.Read ensure that multiple loads from the cache do not
	// change the data that will be read on future reads.
	return
}

// Load calls Load on the Default Cache.
func Load(file string) (pcm.Reader, error) { _ = "STUB: not implemented"; return *new(pcm.Reader), nil }

// Get calls Get on the Default Cache.
func Get(file string) (pcm.Reader, error) { _ = "STUB: not implemented"; return *new(pcm.Reader), nil }
