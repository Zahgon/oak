package audio

import (
	"time"

	"github.com/oakmound/oak/v4/audio/pcm"
)

// FadeIn wraps a reader such that it will linearly fade in over the given duration.
func FadeIn(dur time.Duration, in pcm.Reader) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

type fadeInReader struct {
	pcm.Reader
	toFadeIn, totalToFadeIn int
}

func (fir *fadeInReader) ReadPCM(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FadeOut wraps a reader such that it will linearly fade out over the given duration.
func FadeOut(dur time.Duration, in pcm.Reader) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

type fadeOutReader struct {
	pcm.Reader
	toFadeOut, totaltoFadeOut int
}

func (fir *fadeOutReader) ReadPCM(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ pcm.Reader = &fadeOutReader{}
var _ pcm.Reader = &fadeInReader{}
