//go:build darwin

package audio

import (
	"fmt"

	"github.com/oakmound/oak/v4/audio/pcm"
)

func initOS(driver Driver) error { _ = "STUB: not implemented"; return nil }

// Sanity check that pulse is installed and a sink is defined

// osx: brew install pulseaudio
// linux: sudo apt install pulseaudio

var newWriter = func(f pcm.Format) (pcm.Writer, error) {
	return nil, fmt.Errorf("this package has not been initialized")
}
