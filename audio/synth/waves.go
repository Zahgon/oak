// Package synth provides functions and types to support waveform synthesis.
package synth

import (
	"github.com/oakmound/oak/v4/audio/pcm"
)

// Wave functions take a set of options and return an audio
type Wave func(opts ...Option) pcm.Reader

// Sourced from https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables
func phase(freq Pitch, i int, sampleRate uint32) float64 { _ = "STUB: not implemented"; return 0 }

// Sin produces a Sin wave
//
//	         __
//	       --  --
//	      /      \
//	--__--        --__--
func (s Source) Sin(opts ...Option) pcm.Reader { _ = "STUB: not implemented"; return *new(pcm.Reader) }

func (s Source) SinWave(idx int) float64 { _ = "STUB: not implemented"; return 0 }

func (s Source) Square(opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

// Pulse acts like Square when given a pulse of 2, when given any lesser
// pulse the time up and down will change so that 1/pulse time the wave will
// be up.
//
//	    __    __
//	    ||    ||
//	____||____||____
func (s Source) Pulse(pulse float64) func(opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return nil
}

func PulseWave(pulse float64) Waveform { _ = "STUB: not implemented"; return *new(Waveform) }

// Saw produces a saw wave
//
//	  ^   ^   ^
//	 / | / | /
//	/  |/  |/
func (s Source) Saw(opts ...Option) pcm.Reader { _ = "STUB: not implemented"; return *new(pcm.Reader) }

func (s Source) SawWave(idx int) float64 { _ = "STUB: not implemented"; return 0 }

// Triangle produces a Triangle wave
//
//	  ^   ^
//	 / \ / \
//	v   v   v
func (s Source) Triangle(opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

func (s Source) TriangleWave(idx int) float64 { _ = "STUB: not implemented"; return 0 }

// Noise produces random audio data.
func (s Source) Noise(opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

var _ Waveform = Source.NoiseWave

// NoiseWave returns noise pcm data bounded by this source's volume.
func (s Source) NoiseWave(idx int) float64 { _ = "STUB: not implemented"; return 0 }

func (s Source) modPhase(idx int) float64 { _ = "STUB: not implemented"; return 0 }

// A Waveform is a function that can report a point of audio data given some source parameters for generating the audio
// and an index of where in the generated waveform the requested point lies
type Waveform func(s Source, idx int) float64

// Wave converts a waveform function into a pcm.Reader
func (s Source) Wave(waveFn Waveform, opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

// MultiWave converts a series of waveform functions into a combined reader, outputting the average
// of all of the source waveforms at any given index
func (s Source) MultiWave(waveFns []Waveform, opts ...Option) pcm.Reader {
	_ = "STUB: not implemented"
	return *new(pcm.Reader)
}

type wave8Reader struct {
	Source
	lastIndex int
	waveFunc  func(s Source, idx int) int8
}

func (pr *wave8Reader) ReadPCM(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type wave16Reader struct {
	Source
	lastIndex int
	waveFunc  func(s Source, idx int) int16
}

func (pr *wave16Reader) ReadPCM(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type wave32Reader struct {
	Source
	lastIndex int
	waveFunc  func(s Source, idx int) int32
}

func (pr *wave32Reader) ReadPCM(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
