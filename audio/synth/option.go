package synth

import "time"

// Option types modify waveform sources before they generate a waveform
type Option func(Source) Source

// Duration sets the duration of a generated waveform
func Duration(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// Volume sets the volume of a generated waveform. It guarantees that 0 <= v <= 1
// (silent <= v <= max volume)
func Volume(v float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// AtPitch sets the pitch of a generated waveform.
func AtPitch(p Pitch) Option { _ = "STUB: not implemented"; return *new(Option) }

// Mono sets a synth source to play mono audio.
func Mono() Option { _ = "STUB: not implemented"; return *new(Option) }

// Stereo sets a synth source to play stereo audio.
func Stereo() Option { _ = "STUB: not implemented"; return *new(Option) }

// Detune detunes between -1.0 and 1.0, 1.0 representing a half step up.
// Q: What is detuning? A: It's taking the pitch of the audio and adjusting it less than
// a single tone up or down. If you detune too far, you've just made the next pitch,
// but if you detune a little, you get a resonant sound.
func Detune(percent float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// TODO: does pitch need to be a float?
