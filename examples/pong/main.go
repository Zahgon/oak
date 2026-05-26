package main

import (
	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var (
	score1 = 0
	score2 = 0
)

const (
	hitPaddle collision.Label = 1
)

func main() {
	oak.AddScene("pong",
		scene.Scene{Start: func(ctx *scene.Context) {
			newPaddle(ctx, 20, 200, 1)
			newPaddle(ctx, 600, 200, 2)
			newBall(ctx, 320, 240)
			ctx.Draw(render.NewIntText(&score2, 200, 20), 3)
			ctx.Draw(render.NewIntText(&score1, 400, 20), 3)
		}})
	oak.Init("pong")
}

func newBallDelta() floatgeom.Point2 { _ = "STUB: not implemented"; return *new(floatgeom.Point2) }

func newBall(ctx *scene.Context, x, y float64) { _ = "STUB: not implemented"; return }

func newPaddle(ctx *scene.Context, x, y float64, player int) { _ = "STUB: not implemented"; return }

func enterPaddle(up, down key.Code) func(*entities.Entity, event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	return nil
}
