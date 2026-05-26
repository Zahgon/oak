package entities

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/key"
)

// WASD moves the given mover based on its speed as W,A,S, and D are pressed
func WASD(mvr *Entity) { _ = "STUB: not implemented"; return }

// Arrows moves the given mover based on its speed as the arrow keys are pressed
func Arrows(mvr *Entity) { _ = "STUB: not implemented"; return }

// TopDown moves the given mover based on its speed as the given keys are pressed
func TopDown(mvr *Entity, up, down, left, right key.Code) { _ = "STUB: not implemented"; return }

// CenterScreenOn will cause the screen to center on the given mover, obeying
// viewport limits if they have been set previously
func CenterScreenOn(mvr *Entity) { _ = "STUB: not implemented"; return }

// Limit restricts the movement of the mover to stay within a given rectangle
func Limit(mvr *Entity, rect floatgeom.Rect2) { _ = "STUB: not implemented"; return }
