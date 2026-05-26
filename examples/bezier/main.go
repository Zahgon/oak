package main

import (
	"fmt"
	"strconv"

	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/debugstream"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/shape"
)

var (
	cmp *render.CompositeM
)

func renderCurve(floats []float64) { _ = "STUB: not implemented"; return }

func main() {

	// bezier X Y X Y X Y ...
	// for defining custom points without using the mouse.
	// does not interact with the mouse points tracked through left clicks.
	debugstream.AddCommand(debugstream.Command{Name: "bezier", Operation: func(tokens []string) string {
		if len(tokens) < 4 {
			return ""
		}
		tokens = tokens[1:]
		var err error
		floats := make([]float64, len(tokens))
		for i, s := range tokens {
			floats[i], err = strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Println(err)
				return ""
			}
		}
		renderCurve(floats)
		return ""
	}})

	oak.AddScene("bezier", scene.Scene{Start: func(ctx *scene.Context) {
		mouseFloats := []float64{}
		event.GlobalBind(ctx,
			mouse.Press, func(me *mouse.Event) event.Response {
				// Left click to add a point to the curve
				if me.Button == mouse.ButtonLeft {
					mouseFloats = append(mouseFloats, float64(me.X()), float64(me.Y()))
					renderCurve(mouseFloats)
					// Perform any other click to reset the drawn curve
				} else {
					mouseFloats = []float64{}
					if cmp != nil {
						cmp.Undraw()
					}
				}
				return 0
			})
	}})
	oak.Init("bezier", func(c oak.Config) (oak.Config, error) {
		c.EnableDebugConsole = true
		return c, nil
	})
}

func bezierDraw(b shape.Bezier) *render.CompositeM { _ = "STUB: not implemented"; return nil }

func bezierDrawRec(b shape.Bezier, list *render.CompositeM, alpha uint8) {
	_ = "STUB: not implemented"
	return
}
