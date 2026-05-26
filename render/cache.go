package render

import (
	"image"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/oakmound/oak/v4/alg/intgeom"
)

// DefaultCache is the receiver for package level sprites, sheets, and font loading operations.
var DefaultCache = NewCache()

// Cache is a simple image data cache
type Cache struct {
	imageLock    sync.RWMutex
	loadedImages map[string]*image.RGBA

	sheetLock    sync.RWMutex
	loadedSheets map[string]*Sheet

	fontLock    sync.RWMutex
	loadedFonts map[string]*truetype.Font
}

// NewCache returns an empty Cache
func NewCache() *Cache { _ = "STUB: not implemented"; return nil }

// ClearAll will remove all elements from a Cache
func (c *Cache) ClearAll() { _ = "STUB: not implemented"; return }

// Clear will remove elements matching the given key from the Cache.
func (c *Cache) Clear(key string) { _ = "STUB: not implemented"; return }

// GetSprite calls GetSprite on the Default Cache.
func GetSprite(file string) (*Sprite, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadSprite calls LoadSprite on the Default Cache.
func LoadSprite(file string) (*Sprite, error) { _ = "STUB: not implemented"; return nil, nil }

// GetSheet calls GetSheet on the Default Cache.
func GetSheet(file string) (*Sheet, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadSheet calls LoadSheet on the Default Cache.
func LoadSheet(file string, cellSize intgeom.Point2) (*Sheet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFont calls GetFont on the Default Cache.
func GetFont(file string) (*truetype.Font, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadFont calls LoadFont on the Default Cache.
func LoadFont(file string) (*truetype.Font, error) { _ = "STUB: not implemented"; return nil, nil }
