package windriver

import (
	"image"
	"io"
)

// adapted from https://github.com/Kodeworks/golang-image-ico

type icondir struct {
	reserved  uint16
	imageType uint16
	numImages uint16
}

type icondirentry struct {
	imageWidth   uint8
	imageHeight  uint8
	numColors    uint8
	reserved     uint8
	colorPlanes  uint16
	bitsPerPixel uint16
	sizeInBytes  uint32
	offset       uint32
}

func newIcondir() icondir { _ = "STUB: not implemented"; return *new(icondir) }

func newIcondirentry() icondirentry { _ = "STUB: not implemented"; return *new(icondirentry) }

// windows is supposed to not mind 0 or 1, but other icon files seem to have 1 here
// can be 24 for bitmap or 24/32 for png. Set to 32 for now
//6 icondir + 16 icondirentry, next image will be this image size + 16 icondirentry, etc

func encodeIco(w io.Writer, im image.Image) error { _ = "STUB: not implemented"; return nil }
