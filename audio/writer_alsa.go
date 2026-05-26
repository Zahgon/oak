//go:build linux
// +build linux

package audio

import (
	"sync"

	"github.com/oakmound/alsa"
	"github.com/oakmound/oak/v4/audio/pcm"
)

func newALSAWriter(f pcm.Format) (pcm.Writer, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Writer), nil
}

// Todo: annotate these errors with more info

// Default value at recommendation of library

type alsaWriter struct {
	sync.Mutex
	pcm.Format
	*alsa.Device
	playing bool
	period  int
}

var (
	// Todo: support more customized audio device usage
	openDeviceLock sync.Mutex
	openedDevice   *alsa.Device
)

func openDevice() (*alsa.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// We've a found a device we can hypothetically use
// don't close this card

func alsaFormat(bits uint16) (alsa.FormatType, error) {
	_ = "STUB: not implemented"
	return *new(alsa.FormatType), nil
}

func (aw *alsaWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (aw *alsaWriter) WritePCM(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
