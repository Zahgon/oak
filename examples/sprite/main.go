package main

import (
	"embed"
	"fmt"
	"image"
	"path/filepath"

	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)

const (
	minX = 0
	minY = 0
	maxX = 578
	maxY = 416
)

var cache = [360]*image.RGBA{}

func main() {
	oak.AddScene(
		"demo",
		scene.Scene{Start: func(ctx *scene.Context) {
			render.Draw(render.NewDrawFPS(0.03, nil, 10, 10))
			render.Draw(render.NewLogicFPS(0.03, nil, 10, 20))

			layer := 0
			layerTxt := render.DefaultFont().NewIntText(&layer, 30, 20)
			layerTxt.SetLayer(100000000)
			render.Draw(layerTxt, 0)
			NewGopher(ctx, layer)
			layer++
			event.GlobalBind(ctx, event.Enter, func(ev event.EnterPayload) event.Response {
				if oak.IsDown(key.K) {
					NewGopher(ctx, layer)
					layer++
				}
				return 0
			})
			// Generate a rotation cache for comparison
			// Compare the use of the cache against the use of a reverting type below
			for i := 0; i < 360; i++ {
				s, err := render.LoadSprite(filepath.Join("assets", "images", "raw", "gopher11.png"))
				if err != nil {
					fmt.Println(err)
					return
				}
				s.Modify(mod.Rotate(float32(i)))
				cache[i] = s.GetRGBA()
			}
		},
		})

	render.SetDrawStack(
		render.NewCompositeR(),
	)

	oak.SetFS(assets)
	oak.Init("demo")
}

//go:embed assets
var assets embed.FS

// Gopher is a basic bouncing renderable
type Gopher struct {
	*render.Switch
	event.CallerID
	deltaX, deltaY float64
	rotation       int
}

// NewGopher creates a gopher sprite to bounce around
func NewGopher(ctx *scene.Context, layer int) { _ = "STUB: not implemented"; return }

func gophEnter(goph *Gopher, ev event.EnterPayload) event.Response {
	_ = "STUB: not implemented"
	// Compare against this version of rotation
	// (also swap the comments on lines in goph.Doodad's renderable)
	//goph.R.(*render.Reverting).RevertAndModify(1, render.Rotate(goph.rotation))
	return *new(event.Response)
}
