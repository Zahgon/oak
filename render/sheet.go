package render

import (
	"image"
)

// Sheet is a 2D array of image rgbas
type Sheet [][]*image.RGBA

// SubSprite gets a sprite from a sheet at the given location
func (sh *Sheet) SubSprite(x, y int) *Sprite { _ = "STUB: not implemented"; return nil }

// ToSprites returns this sheet as a 2D array of Sprites
func (sh *Sheet) ToSprites() [][]*Sprite { _ = "STUB: not implemented"; return nil }

// NewSheetSequence creates a Sequence from a sheet and a list of x,y frame coordinates.
// A sequence will be created by getting the sheet's [i][i+1]th elements incrementally
// from the input frames. If the number of input frames is uneven, an error is returned.
func NewSheetSequence(sheet *Sheet, fps float64, frames ...int) (*Sequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
