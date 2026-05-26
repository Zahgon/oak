package scene

import (
	"context"
	"time"

	"github.com/oakmound/oak/v4/render"
)

// DoAfter will execute the given function after some duration. When the scene
// ends, DoAfter will exit without calling f. This call blocks until one of those
// conditions is reached.
func (c *Context) DoAfter(d time.Duration, f func()) { _ = "STUB: not implemented"; return }

// DoAfterContext will execute the given function once the passed in context is closed.
// When the scene ends, DoAfterContext will exit without calling f. This call blocks until
// one of those conditions is reached.
func (c *Context) DoAfterContext(ctx context.Context, f func()) { _ = "STUB: not implemented"; return }

// DrawForTime draws, and after d, undraws an element
func (c *Context) DrawForTime(r render.Renderable, d time.Duration, layers ...int) error {
	_ = "STUB: not implemented"
	return nil
}
