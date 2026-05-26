package render

import (
	"image"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

// LoadSheet loads a file in some directory with sheets of (w,h) sized sprites.
// This will blow away any cached sheet with the same fileName.
func (c *Cache) LoadSheet(file string, cellSize intgeom.Point2) (*Sheet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeSheet converts an image into a sheet with cellSize sized sprites
func MakeSheet(rgba *image.RGBA, cellSize intgeom.Point2) (*Sheet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSheet tries to find the given file in the set of loaded sheets.
// If SheetIsLoaded(filename) is not true, this returns an error.
// Otherwise it will return the sheet as a 2d array of sprites
func (c *Cache) GetSheet(fileName string) (*Sheet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func subImage(rgba *image.RGBA, x, y, w, h int) *image.RGBA { _ = "STUB: not implemented"; return nil }
