package render

import (
	"image"
	"io"
)

// Decoder functions convert arbitrary readers to images.
// The input of a decoder in oak's loader will generally
// be an image file.
type Decoder func(io.Reader) (image.Image, error)

// CfgDecoder is an equivalent to Decoder that just exports
// the color model and dimensions of the image.
type CfgDecoder func(io.Reader) (image.Config, error)

var (
	fileDecoders = map[string]Decoder{}
	cfgDecoders  = map[string]CfgDecoder{}
)

// RegisterDecoder adds a decoder to the set of image decoders
// for file loading. If the extension string is already set,
// the existing decoder will not be overwritten.
func RegisterDecoder(ext string, decoder Decoder) error { _ = "STUB: not implemented"; return nil }

// RegisterCfgDecoder acts like RegisterDecoder for CfgDecoders
func RegisterCfgDecoder(ext string, decoder CfgDecoder) error {
	_ = "STUB: not implemented"
	return nil
}
