package main

import (
	"image/color"
	"log"

	oak "github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg"
	"github.com/oakmound/oak/v4/alg/span"
	"github.com/oakmound/oak/v4/debugstream"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/oakerr"
	"github.com/oakmound/oak/v4/physics"
	"github.com/oakmound/oak/v4/render"
	pt "github.com/oakmound/oak/v4/render/particle"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/shape"
)

var (
	startColor     color.Color
	startColorRand color.Color
	endColor       color.Color
	endColorRand   color.Color
	src            *pt.Source
)

func parseShape(args []string) shape.Shape { _ = "STUB: not implemented"; return *new(shape.Shape) }

func main() {

	debugstream.AddCommand(debugstream.Command{Name: "followMouse", Operation: func(args []string) string {
		event.GlobalBind(event.DefaultBus, event.Enter, func(ev event.EnterPayload) event.Response {
			// It'd be interesting to attach to the mouse position
			src.SetPos(float64(mouse.LastEvent.X()), float64(mouse.LastEvent.Y()))
			return 0
		})
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "shape", Operation: func(args []string) string {
		if len(args) > 0 {
			sh := parseShape(args)
			if sh != nil {
				src.Generator.(pt.Shapeable).SetShape(sh)
			}
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "size", Operation: func(args []string) string {
		f1, f2, two, err := parseInts(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.(pt.Sizeable).SetSize(span.NewConstant(f1))
		} else {
			src.Generator.(pt.Sizeable).SetSize(span.NewLinear(f1, f2))
		}

		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "endsize", Operation: func(args []string) string {
		f1, f2, two, err := parseInts(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.(pt.Sizeable).SetEndSize(span.NewConstant(f1))
		} else {
			src.Generator.(pt.Sizeable).SetEndSize(span.NewLinear(f1, f2))
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "count", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.GetBaseGenerator().NewPerFrame = span.NewConstant(npf)
		} else {
			src.Generator.GetBaseGenerator().NewPerFrame = span.NewLinear(npf, npf2)
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "life", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.GetBaseGenerator().LifeSpan = span.NewConstant(npf)
		} else {
			src.Generator.GetBaseGenerator().LifeSpan = span.NewLinear(npf, npf2)
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "rotation", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.GetBaseGenerator().Rotation = span.NewConstant(npf)
		} else {
			src.Generator.GetBaseGenerator().Rotation = span.NewLinear(npf, npf2)
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "angle", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.GetBaseGenerator().Angle = span.NewConstant(npf * alg.DegToRad)
		} else {
			src.Generator.GetBaseGenerator().Angle = span.NewLinear(npf*alg.DegToRad, npf2*alg.DegToRad)
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "speed", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			src.Generator.GetBaseGenerator().Speed = span.NewConstant(npf)
		} else {
			src.Generator.GetBaseGenerator().Speed = span.NewLinear(npf, npf2)
		}
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "spread", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			return oakerr.InsufficientInputs{AtLeast: 2, InputName: "speeds"}.Error()
		}
		src.Generator.GetBaseGenerator().Spread.SetPos(npf, npf2)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "gravity", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			return oakerr.InsufficientInputs{AtLeast: 2, InputName: "speeds"}.Error()
		}
		src.Generator.GetBaseGenerator().Gravity.SetPos(npf, npf2)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "speeddecay", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			return oakerr.InsufficientInputs{AtLeast: 2, InputName: "speeds"}.Error()
		}
		src.Generator.GetBaseGenerator().SpeedDecay.SetPos(npf, npf2)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "pos", Operation: func(args []string) string {
		npf, npf2, two, err := parseFloats(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		if !two {
			return oakerr.InsufficientInputs{AtLeast: 2, InputName: "positions"}.Error()
		}
		src.Generator.SetPos(npf, npf2)

		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "startcolor", Operation: func(args []string) string {
		if len(args) < 3 {
			return oakerr.InsufficientInputs{AtLeast: 3, InputName: "colorvalues"}.Error()
		}
		r, g, b, a, err := parseRGBA(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		startColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		src.Generator.(pt.Colorable).SetStartColor(startColor, startColorRand)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "startrand", Operation: func(args []string) string {
		if len(args) < 3 {
			return oakerr.InsufficientInputs{AtLeast: 3, InputName: "colorvalues"}.Error()
		}
		r, g, b, a, err := parseRGBA(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		startColorRand = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		src.Generator.(pt.Colorable).SetStartColor(startColor, startColorRand)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "endcolor", Operation: func(args []string) string {
		if len(args) < 3 {
			return oakerr.InsufficientInputs{AtLeast: 3, InputName: "colorvalues"}.Error()
		}
		r, g, b, a, err := parseRGBA(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		endColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		src.Generator.(pt.Colorable).SetEndColor(endColor, endColorRand)
		return ""
	}})

	debugstream.AddCommand(debugstream.Command{Name: "endrand", Operation: func(args []string) string {
		if len(args) < 3 {
			return oakerr.InsufficientInputs{AtLeast: 3, InputName: "colorvalues"}.Error()
		}
		r, g, b, a, err := parseRGBA(args)
		if err != nil {
			return oakerr.UnsupportedFormat{Format: err.Error()}.Error()
		}
		endColorRand = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		src.Generator.(pt.Colorable).SetEndColor(endColor, endColorRand)
		return ""
	}})

	oak.AddScene("demo", scene.Scene{Start: func(*scene.Context) {
		render.Draw(render.NewDrawFPS(0, nil, 10, 10))
		x := 320.0
		y := 240.0
		newPf := span.NewLinear(1.0, 2.0)
		life := span.NewLinear(100.0, 120.0)
		angle := span.NewLinear(0.0, 360.0)
		speed := span.NewLinear(1.0, 5.0)
		size := span.NewConstant(1)
		layerFn := func(v physics.Vector) int {
			return 1
		}
		startColor = color.RGBA{255, 255, 255, 255}
		startColorRand = color.RGBA{0, 0, 0, 0}
		endColor = color.RGBA{255, 255, 255, 255}
		endColorRand = color.RGBA{0, 0, 0, 0}
		shape := shape.Square

		src = pt.NewColorGenerator(
			pt.Pos(x, y),
			pt.Duration(pt.Inf),
			pt.LifeSpan(life),
			pt.Angle(angle),
			pt.Speed(speed),
			pt.Layer(layerFn),
			pt.Shape(shape),
			pt.Size(size),
			pt.Color(startColor, startColorRand, endColor, endColorRand),
			pt.NewPerFrame(newPf)).Generate(0)
	}})

	render.SetDrawStack(
		render.NewCompositeR(),
	)

	err := oak.Init("demo", func(c oak.Config) (oak.Config, error) {
		c.Debug.Level = "VERBOSE"
		c.DrawFrameRate = 1200
		c.FrameRate = 60
		c.Title = "Particle Demo"
		c.EnableDebugConsole = true
		return c, nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func parseRGBA(args []string) (r, g, b, a int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func parseFloats(args []string) (f1, f2 float64, two bool, err error) {
	_ = "STUB: not implemented"
	return 0, 0, false, nil
}

func parseInts(args []string) (i1, i2 int, two bool, err error) {
	_ = "STUB: not implemented"
	return 0, 0, false, nil
}
