package collision

import "github.com/oakmound/oak/v4/event"

// A Filter will take a set of collision spaces
// and return the subset that match some requirement
type Filter func([]*Space) []*Space

// FirstLabel returns the first space that has a label in the input, or nothing
func FirstLabel(ls ...Label) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// With will filter spaces so that only those returning true
// from the input keepFn will be in the output
func With(keepFn func(*Space) bool) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// Without will filter spaces so that no spaces returning true
// from the input tossFn will be in the output
func Without(tossFn func(*Space) bool) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// WithoutCIDs will return no spaces with a CID in the input
func WithoutCIDs(cids ...event.CallerID) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// WithLabels will only return spaces with a label in the input
func WithLabels(ls ...Label) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// WithoutLabels will return no spaces with a label in the input
func WithoutLabels(ls ...Label) Filter { _ = "STUB: not implemented"; return *new(Filter) }
