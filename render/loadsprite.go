package render

import (
	"image"
)

func loadSpriteNoCache(file string, maxFileSize int64) (*image.RGBA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This can't reasonably error as we already loaded the file above

// construct a blank image of the correct dimensions

// Todo: we internally just use *image.RGBA, but that choice
// of image encoding was arbitrary. If using the image.Image
// interface would not hurt performance considerably, we should
// just use that.

func (c *Cache) loadSprite(file string, maxFileSize int64) (*image.RGBA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSprite tries to find the given file in a private set of
// loaded sprites. If that file isn't cached, it will return an error.
func (c *Cache) GetSprite(file string) (*Sprite, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadSprite will load the given file as an image by combining directory and fileName.
// The resulting image, if found, will be cached under its last path element for
// later access through GetSprite.
func (c *Cache) LoadSprite(file string) (*Sprite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
