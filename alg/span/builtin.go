package span

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// A Spanable must be usable in basic arithmetic-- addition, subtraction, and multiplication.
type Spanable interface {
	constraints.Float | constraints.Integer
}

// NewConstant returns a span where the minimum and maximum are both i. Poll, Percentile, and Clamp will always return i.
func NewConstant[T Spanable](i T) Span[T] { _ = "STUB: not implemented"; return nil }

type constant[T Spanable] struct {
	val T
}

func (c constant[T]) Poll() T { _ = "STUB: not implemented"; return *new(T) }

func (c constant[T]) MulSpan(i float64) Span[T] { _ = "STUB: not implemented"; return nil }

func (c constant[T]) Clamp(T) T { _ = "STUB: not implemented"; return *new(T) }

func (c constant[T]) Percentile(float64) T {
	_ = "STUB: not implemented"

	// NewLinear returns a linear span between min and max. The linearity implies that no point in the span is preferred,
	// and Percentile will scale in a constant fashion from min to max.
	return *new(T)
}

func NewLinear[T Spanable](min, max T) Span[T] { _ = "STUB: not implemented"; return nil }

// NewSpread returns a linear span from base-spread to base+spread.
func NewSpread[T Spanable](base, spread T) Span[T] { _ = "STUB: not implemented"; return nil }

type linear[T Spanable] struct {
	min, max T
	rng      *rand.Rand
	flipped  bool
}

func (lir linear[T]) Poll() T { _ = "STUB: not implemented"; return *new(T) }

func (lir linear[T]) MulSpan(i float64) Span[T] { _ = "STUB: not implemented"; return nil }

func (lir linear[T]) Clamp(i T) T { _ = "STUB: not implemented"; return *new(T) }

func (lir linear[T]) Percentile(f float64) T { _ = "STUB: not implemented"; return *new(T) }

// 0 - 255 * .1 = -25 + 255 = 230 // 255 - 0 * .1 = 25
