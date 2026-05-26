// Package pcm provides a interface for interacting with PCM audio streams
package pcm

import (
	"fmt"
	"io"
)

var _ Reader = &IOReader{}

// A Reader mimics io.Reader for pcm data.
type Reader interface {
	Formatted
	ReadPCM(b []byte) (n int, err error)
}

// An IOReader converts an io.Reader into a pcm.Reader
type IOReader struct {
	Format
	io.Reader
}

func (ior *IOReader) ReadPCM(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// A Writer can have PCM formatted audio data written to it. It mimics io.Writer.
	return 0, nil
}

type Writer interface {
	io.Closer
	Formatted
	// WritePCM expects PCM bytes matching this Writer's format.
	// WritePCM will block until all of the bytes are consumed.
	WritePCM([]byte) (n int, err error)
}

// The Formatted interface represents types that are aware of a PCM Format they expect or provide.
type Formatted interface {
	// PCMFormat will return the Format used by an encoded audio or expected by an audio consumer.
	// Implementations can embed a Format struct to simplify this.
	PCMFormat() Format
}

// Format is a PCM format; it defines how binary audio data should be converted into real audio.
type Format struct {
	// SampleRate defines how many times per second a consumer should read a single value. An example
	// of a common value for this is 44100 or 44.1khz.
	SampleRate uint32
	// Channels defines how many concurrent audio channels are present in audio data. Common values are
	// 1 for mono and 2 for stereo.
	Channels uint16
	// Bits determines how many bits a single sample value takes up. 8, 16, and 32 are common values.
	// TODO: Do we need LE vs BE, float vs int representation?
	Bits uint16
}

// PCMFormat returns this format.
func (f Format) PCMFormat() Format {
	_ = "STUB: not implemented"

	// BytesPerSecond returns how many bytes this format would be encoded into per second in an audio stream.
	return *new(Format)
}

func (f Format) BytesPerSecond() uint32 { _ = "STUB: not implemented"; return 0 }

func (f Format) SampleSize() int { _ = "STUB: not implemented"; return 0 }

// ReadFloat reads a single sample from an audio stream, respecting bits and channels:
// f.Bits / 8 bytes * f.Channels bytes will be read from b, and this count will be returned as 'read'.
// the length of values will be equal to f.Channels, if no error is returned. If an error is returned,
// it will be io.ErrUnexpectedEOF or ErrUnsupportedBits
func (f Format) SampleFloat(b []byte) (values []float64, read int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// ErrUnsupportedBits represents that the Bits value for a Format was not supported for some operation.
var ErrUnsupportedBits = fmt.Errorf("unsupported bits in pcm format")
