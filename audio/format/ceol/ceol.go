// Package ceol provides functionality to handle .ceol files and .ceol encoded data (Bosca Ceoil files).
package ceol

import (
	"io"
	"time"
)

// Raw Ceol types, holds all information in ceol file

// Ceol represents a complete .ceol file
type Ceol struct {
	Version       int
	Swing         int
	Effect        int
	EffectValue   int
	Bpm           int
	PatternLength int
	BarLength     int
	Instruments   []Instrument
	Patterns      []Pattern
	LoopStart     int
	LoopEnd       int
	Arrangement   [][8]int
}

// Instrument represents a single entry in a .ceol's instrument block
type Instrument struct {
	Index        int
	IsDrumkit    int
	Palette      int
	LPFCutoff    int
	LPFResonance int
	Volume       int
}

// Pattern represents a single entry in a .ceol's pattern block
type Pattern struct {
	Key        int
	Scale      int
	Instrument int
	Palette    int
	Notes      []Note
	Filters    []Filter
}

// Note represents a single entry in a .ceol's pattern's note block
type Note struct {
	PitchIndex int // C4 = 60
	Length     int
	Offset     int
}

// Filter represents a single entry in a .ceol's pattern's filter block
type Filter struct {
	Volume       int
	LPFCutoff    int
	LPFResonance int
}

// DurationFromQuarters should not be here, should be in a package
// managing bpm and time
// Duration from quarters expects four quarters to occur per beat,
// (direct complaints at terry cavanagh), and returns a time.Duration
// for n quarters in the given bpm.
func DurationFromQuarters(bpm, quarters int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Open returns a Ceol from an io.Reader
func Open(r io.Reader) (Ceol, error) { _ = "STUB: not implemented"; return *new(Ceol), nil }

// Dummy value here
