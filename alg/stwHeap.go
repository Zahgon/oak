package alg

type stwHeap struct {
	bh           []float64
	weightsBelow []float64
}

// Select Total Weight Heap
// This name was chosen relatively arbitrarily, if there
// is a canonical academic name for this structure we'd gladly
// use that instead
func newSTWHeap(f []float64) *stwHeap { _ = "STUB: not implemented"; return nil }

// The order of elements literally does not
// matter, so 'heap' is a misnomer.

func (stwh *stwHeap) Pop(rng float64) int { _ = "STUB: not implemented"; return 0 }

// With the >= here, we don't accept 0 weights

// Propagate to left child

// Switch to right child

// Instead of removing a node we set its weight to 0.

// All parents of the index need to be reduced
// in total weight.
