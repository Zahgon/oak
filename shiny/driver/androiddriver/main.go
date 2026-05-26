//go:build android
// +build android

package androiddriver

import (
	"github.com/oakmound/oak/v4/shiny/screen"
)

func Main(f func(screen.Screen)) { _ = "STUB: not implemented"; return }

// TODO: expose touch events in a way an oak program can
// differentiate them from clicks
