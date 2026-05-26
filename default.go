package oak

import (
	"image"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var defaultWindow *Window

var initDefaultWindowOnce sync.Once

func initDefaultWindow() { _ = "STUB: not implemented"; return }

// Init calls Init on the default window. The default window
// will be set to use render.GlobalDrawStack and event.DefaultBus.
func Init(scene string, configOptions ...ConfigOption) error { _ = "STUB: not implemented"; return nil }

// AddScene calls AddScene on the default window.
func AddScene(name string, sc scene.Scene) error { _ = "STUB: not implemented"; return nil }

// IsDown calls IsDown on the default window.
func IsDown(k key.Code) bool { _ = "STUB: not implemented"; return false }

// IsHeld calls IsHeld on the default window.
func IsHeld(k key.Code) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

// SetViewportBounds calls SetViewportBounds on the default window.
func SetViewportBounds(rect intgeom.Rect2) { _ = "STUB: not implemented"; return }

// ShiftViewport calls ShiftViewport on the default window.
func ShiftViewport(pt intgeom.Point2) { _ = "STUB: not implemented"; return }

// SetViewport calls SetViewport on the default window.
func SetViewport(pt intgeom.Point2) { _ = "STUB: not implemented"; return }

// UpdateViewSize calls UpdateViewSize on the default window.
func UpdateViewSize(w, h int) error { _ = "STUB: not implemented"; return nil }

// ScreenShot calls ScreenShot on the default window.
func ScreenShot() *image.RGBA { _ = "STUB: not implemented"; return nil }

// SetLoadingRenderable calls SetLoadingRenderable on the default window.
func SetLoadingRenderable(r render.Renderable) { _ = "STUB: not implemented"; return }

// SetBackground calls SetBackground on the default window.
func SetBackground(b Background) { _ = "STUB: not implemented"; return }

// SetColorBackground calls SetColorBackground on the default window.
func SetColorBackground(img image.Image) { _ = "STUB: not implemented"; return }

// Bounds returns the default window's boundary.
func Bounds() intgeom.Point2 { _ = "STUB: not implemented"; return *new(intgeom.Point2) }
