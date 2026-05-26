package main

import (
	"time"

	"github.com/oakmound/oak/v4/alg/span"

	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var (
	score int
)

// label pillars with a known constant, so when we hit them, we can restart the scene
const (
	pillar collision.Label = iota
)

func main() {
	oak.AddScene("flappy", scene.Scene{Start: func(ctx *scene.Context) {
		render.Draw(render.NewDrawFPS(0, nil, 10, 10), 2, 0)
		render.Draw(render.NewLogicFPS(0, nil, 10, 20), 2, 0)

		score = 0
		// 1. Make Player
		newFlappy(ctx, 90, 140)
		// 2. Make scrolling repeating pillars
		pillarFreq := span.NewLinear(1.0, 5.0)
		var pillarLoop func()
		pillarLoop = func() {
			newPillarPair(ctx)
			ctx.DoAfter(time.Duration(pillarFreq.Poll()*float64(time.Second)), pillarLoop)
		}
		go ctx.DoAfter(time.Duration(pillarFreq.Poll()*float64(time.Second)), pillarLoop)

		// 3. Make Score
		t := render.DefaultFont().NewIntText(&score, 200, 30)
		render.Draw(t, 0)
	}})
	oak.Init("flappy")
}

func newFlappy(ctx *scene.Context, x, y float64) { _ = "STUB: not implemented"; return }

// Gravity

var (
	gapPosition = span.NewLinear(10.0, 370.0)
	gapSpan     = span.NewLinear(100.0, 250.0)
)

func newPillarPair(ctx *scene.Context) { _ = "STUB: not implemented"; return }

func newPillar(ctx *scene.Context, x, y, h float64, isAbove bool) {
	_ = "STUB: not implemented"
	return
}

func enterPillar(isAbove bool) func(p *entities.Entity, ev event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return nil
}

// don't score one out of each two pillars
