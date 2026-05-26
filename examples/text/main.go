package main

import (
	"embed"

	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

//go:embed assets
var assets embed.FS

type floatStringer struct {
	f *float64
}

func (fs floatStringer) String() string { _ = "STUB: not implemented"; return "" }

func main() {
	oak.AddScene("demo",
		scene.Scene{Start: func(ctx *scene.Context) {
			render.Draw(render.NewDrawFPS(0.25, nil, 10, 10))
			drawFallbackFonts(ctx)
			drawColorChangingText(ctx)
		},
		})
	oak.SetFS(assets)
	oak.Init("demo")
}

func drawFallbackFonts(ctx *scene.Context) { _ = "STUB: not implemented"; return }

// TODO: support multi-color glyphs

func drawColorChangingText(ctx *scene.Context) { _ = "STUB: not implemented"; return }
