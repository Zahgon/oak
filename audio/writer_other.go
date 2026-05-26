//go:build !windows && !linux && !darwin

package audio

import (
	"github.com/oakmound/oak/v4/audio/pcm"
)

func initOS(driver Driver) error { _ = "STUB: not implemented"; return nil }

func newWriter(f pcm.Format) (pcm.Writer, error) {
	_ = "STUB: not implemented"
	return *new(pcm.Writer), nil
}
