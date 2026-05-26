//go:build windows

package audio

import (
	"sync"

	"github.com/oakmound/oak/v4/audio/pcm"
	"github.com/oov/directsound-go/dsound"
)

func initOS(driver Driver) error { _ = "STUB: not implemented"; return nil }

// OK

var directSoundInterface *dsound.IDirectSound

func newWriter(f pcm.Format) (pcm.Writer, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Writer), nil
}

// These flags cover everything we should ever want to do

type directSoundWriter struct {
	sync.Mutex
	pcm.Format
	buff         *dsound.IDirectSoundBuffer
	lockedOffset uint32
	bufferSize   uint32
	playing      bool
}

func (dsw *directSoundWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (dsw *directSoundWriter) Seek(offset int64, whence int) (position int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (dsw *directSoundWriter) WritePCM(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Always loop-- these buffers are small, and are continually reused even for
// larger audio sources
