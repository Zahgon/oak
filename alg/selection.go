package alg

// A Float64Generator must be able to generate a float64. This is generally used to implement randomness
// See rand.Float64 for an example
type Float64Generator interface {
	Float64() float64
}

// UniqueChooseX returns n indices from the input weights at a count
// relative to the weight of each index. This will never return duplicate indices.
// if n > len(weights), it will return -1 after depleting the n elements from
// weights.
func UniqueChooseX(weights []float64, n int) []int { _ = "STUB: not implemented"; return nil }

// UniqueChooseXSeeded returns n indices from the input weights at a count
// relative to the weight of each index. This will never return duplicate indices.
// if n > len(weights), it will return -1 after depleting the n elements from
// weights. If you do not want to use your own rand
// use UniqueChooseX.
func UniqueChooseXSeeded(weights []float64, n int, rng Float64Generator) []int {
	_ = "STUB: not implemented"
	return nil
}

func uniqueChooseX(weights []float64, n int, rngFxn func() float64) []int {
	_ = "STUB: not implemented"
	return nil
}

// ChooseX - also known as Roulette Search.
// This returns n indices from the input weights at a count
// relative to the weight of each index. It can return the same index
// multiple times.
func ChooseX(weights []float64, n int) []int { _ = "STUB: not implemented"; return nil }

// WeightedMapChoice converts the input map into a set where keys are
// indices and values are weights for WeightedChooseOne, then returns
// the key for WeightedChooseOne of the weights.
func WeightedMapChoice(weightMap map[int]float64) int { _ = "STUB: not implemented"; return 0 }

// WeightedMapChoiceSeeded converts the input map into a set where keys are
// indices and values are weights for WeightedChooseOne, then returns
// the key for WeightedChooseOne of the weight. If you do not want to use your own rand
// use WeightedMapChoice.
func WeightedMapChoiceSeeded(weightMap map[int]float64, rng Float64Generator) int {
	_ = "STUB: not implemented"
	return 0
}

func weightedMapChoice(weightMap map[int]float64, rngFxn func() float64) int {
	_ = "STUB: not implemented"
	return 0
}

// CumulativeWeights converts a slice of weights into
// a slice of cumulative weights, where each index
// is the sum of all weights following that index in
// the original slice
func CumulativeWeights(weights []float64) []float64 { _ = "STUB: not implemented"; return nil }

// WeightedChooseOne returns a single index from the weights given
// at a rate relative to the magnitude of each weight. It expects
// the input to be in the form of CumulativeWeights, cumulative with
// the total at index 0.
func WeightedChooseOne(cumulative []float64) int { _ = "STUB: not implemented"; return 0 }

// WeightedChooseOneSeeded returns a single index from the weights given
// at a rate relative to the magnitude of each weight. It expects
// the input to be in the form of CumulativeWeights, cumulative with
// the total at index 0. If you do not want to use your own rand
// use WeightedChooseOne.
func WeightedChooseOneSeeded(cumulative []float64, rng Float64Generator) int {
	_ = "STUB: not implemented"
	return 0
}

// weightedChooseOne returns a single index from the weights given
// at a rate relative to the magnitude of each weight. It expects
// the input to be in the form of CumulativeWeights, cumulative with
// the total at index 0.
func weightedChooseOne(cumulative []float64, rngFxn func() float64) int {
	_ = "STUB: not implemented"
	return 0
}
