// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x11driver

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

	"github.com/BurntSushi/xgb/render"
	"github.com/BurntSushi/xgb/xproto"

	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/image/math/f64"
)

const textureDepth = 32

type textureImpl struct {
	s *screenImpl

	size image.Point
	xm   xproto.Pixmap
	xp   render.Picture

	// renderMu is a mutex that enforces the atomicity of methods like
	// Window.Draw that are conceptually one operation but are implemented by
	// multiple X11/Render calls. X11/Render is a stateful API, so interleaving
	// X11/Render calls from separate higher-level operations causes
	// inconsistencies.
	renderMu sync.Mutex

	releasedMu sync.Mutex
	released   bool
}

func (t *textureImpl) degenerate() bool  { _ = "STUB: not implemented"; return false }
func (t *textureImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }
func (t *textureImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (t *textureImpl) Release() { _ = "STUB: not implemented"; return }

func (t *textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (t *textureImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// f64ToFixed converts from float64 to X11/Render's 16.16 fixed point.
func f64ToFixed(x float64) render.Fixed { _ = "STUB: not implemented"; return *new(render.Fixed) }

func inv(x *f64.Aff3) f64.Aff3 { _ = "STUB: not implemented"; return *new(f64.Aff3) }

func (t *textureImpl) draw(xp render.Picture, src2dst *f64.Aff3, sr image.Rectangle, op draw.Op) {
	_ = "STUB: not implemented"
	return
}

// For simple copies and scales, the inverse matrix is trivial to compute,
// and we do not need the "Src becomes OutReverse plus Over" dance (see
// below). Thus, draw can be one render.SetPictureTransform call and then
// one render.Composite call, regardless of whether or not op is Src.

// TODO: check if this (and below) works when src2dst[0] < 0.

// TODO: check if this (and below) works when src2dst[4] < 0.

// SrcX, SrcY,
// MaskX, MaskY,
// DstX, DstY,
// Width, Height,

// The X11/Render transform matrix maps from destination pixels to source
// pixels, so we invert src2dst.

// render.TriFan visits every dst-space pixel in the axis-aligned
// bounding box (AABB) containing the transformation of the sr
// rectangle in src-space to a quad in dst-space.
//
// render.TriFan is like render.Composite, except that the AABB is
// defined implicitly by the transformed triangle vertices instead of
// being passed explicitly as arguments. It implies the minimal AABB.
//
// In any case, for arbitrary src2dst affine transformations, which
// include rotations, this means that a naive render.TriFan call will
// affect those pixels inside the AABB but outside the quad. For the
// draw.Src operator, this means that pixels in that AABB can be
// incorrectly set to zero.
//
// Instead, we implement the draw.Src operator as two render.TriFan
// calls. The first one (using the PictOpOutReverse operator and a
// fully opaque source) clears the dst-space quad but leaves pixels
// outside that quad (but inside the AABB) untouched. The second one
// (using the PictOpOver operator and the texture t as source) fills in
// the quad and again does not touch the pixels outside.
//
// What X11/Render calls PictOpOutReverse is also known as dst-out. See
// http://www.w3.org/TR/SVGCompositing/examples/compop-porterduff-examples.png
// for a visualization.

func trifanPoints(src2dst *f64.Aff3, sr image.Rectangle) [4]render.Pointfix {
	_ = "STUB: not implemented"
	return nil
}

func renderOp(op draw.Op) byte { _ = "STUB: not implemented"; return 0 }
