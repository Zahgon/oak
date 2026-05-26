//go:build linux || darwin

package audio

import (
	"bytes"
	"sync"

	"github.com/jfreymuth/pulse"

	"github.com/oakmound/oak/v4/audio/pcm"
)

// This mutex may be unneeded
var newWriterMutex sync.Mutex

func newPulseWriter(f pcm.Format) (pcm.Writer, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Writer), nil
}

// TODO:
// 1. Volume scales with pitch-- lower pitches are quieter -- only happens for sin32 and tri32, so probably us

type eofFReader struct {
	*bytes.Buffer
	eof bool
}

func (m *eofFReader) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// if we've closed, make sure nothing more is read
	return 0, nil
}

// This enables the pulse library to continually read from
// the buffer we are handing data over to in WritePCM
// TODO: to resolve glitches

type pulseWriter struct {
	sync.Mutex
	pcm.Format
	handOver      *eofFReader
	playBack      *pulse.PlaybackStream
	client        *pulse.Client
	startComplete chan (struct{})
	playing       bool
}

func (dsw *pulseWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (dsw *pulseWriter) WritePCM(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// this function blocks if it does not have enough data yet.
// we don't want to structure our API around start/stop so far,
// so just start and wait for it to unblock once we've written
// enough for the os to be happy. Signal to the rest of the system
// that this process has completed by closing startComplete after.
// (without this channel, close may be called while start is being called,
// causing a panic)
