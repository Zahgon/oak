// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

// Package coreanim provides access to Apple's Core Animation API
// (https://developer.apple.com/documentation/quartzcore).
//
// This package is in very early stages of development.
// It's a minimal implementation with scope limited to
// supporting mtldriver.
//
// It was copied from dmitri.shuralyov.com/gpu/mtl/example/movingtriangle/internal/coreanim.
package coreanim

import (
	"unsafe"

	"dmitri.shuralyov.com/gpu/mtl"
)

/*
#cgo LDFLAGS: -framework QuartzCore -framework Foundation
#include <stdbool.h>
#include "coreanim.h"
*/
import "C"

// Layer is an object that manages image-based content and
// allows you to perform animations on that content.
//
// Reference: https://developer.apple.com/documentation/quartzcore/calayer.
type Layer interface {
	// Layer returns the underlying CALayer * pointer.
	Layer() unsafe.Pointer
}

// MetalLayer is a Core Animation Metal layer, a layer that manages a pool of Metal drawables.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer.
type MetalLayer struct {
	metalLayer unsafe.Pointer
}

// MakeMetalLayer creates a new Core Animation Metal layer.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer.
func MakeMetalLayer() MetalLayer { _ = "STUB: not implemented"; return *new(MetalLayer) }

// Layer implements the Layer interface.
func (ml MetalLayer) Layer() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *

	// PixelFormat returns the pixel format of textures for rendering layer content.
	//
	// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/1478155-pixelformat.
	new(unsafe.Pointer)
}

func (ml MetalLayer) PixelFormat() mtl.PixelFormat {
	_ = "STUB: not implemented"
	return *new(mtl.PixelFormat)
}

// SetDevice sets the Metal device responsible for the layer's drawable resources.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/1478163-device.
func (ml MetalLayer) SetDevice(device mtl.Device) { _ = "STUB: not implemented"; return }

// SetPixelFormat controls the pixel format of textures for rendering layer content.
//
// The pixel format for a Metal layer must be PixelFormatBGRA8UNorm, PixelFormatBGRA8UNormSRGB,
// PixelFormatRGBA16Float, PixelFormatBGRA10XR, or PixelFormatBGRA10XRSRGB.
// SetPixelFormat panics for other values.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/1478155-pixelformat.
func (ml MetalLayer) SetPixelFormat(pf mtl.PixelFormat) { _ = "STUB: not implemented"; return }

// SetMaximumDrawableCount controls the number of Metal drawables in the resource pool
// managed by Core Animation.
//
// It can set to 2 or 3 only. SetMaximumDrawableCount panics for other values.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/2938720-maximumdrawablecount.
func (ml MetalLayer) SetMaximumDrawableCount(count int) { _ = "STUB: not implemented"; return }

// SetDisplaySyncEnabled controls whether the Metal layer and its drawables
// are synchronized with the display's refresh rate.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/2887087-displaysyncenabled.
func (ml MetalLayer) SetDisplaySyncEnabled(enabled bool) { _ = "STUB: not implemented"; return }

// SetDrawableSize sets the size, in pixels, of textures for rendering layer content.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/1478174-drawablesize.
func (ml MetalLayer) SetDrawableSize(width, height int) { _ = "STUB: not implemented"; return }

// NextDrawable returns a Metal drawable.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametallayer/1478172-nextdrawable.
func (ml MetalLayer) NextDrawable() (MetalDrawable, error) {
	_ = "STUB: not implemented"
	return *new(MetalDrawable), nil
}

// MetalDrawable is a displayable resource that can be rendered or written to by Metal.
//
// Reference: https://developer.apple.com/documentation/quartzcore/cametaldrawable.
type MetalDrawable struct {
	metalDrawable unsafe.Pointer
}

// Drawable implements the mtl.Drawable interface.
func (md MetalDrawable) Drawable() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *

	// Texture returns a Metal texture object representing the drawable object's content.
	//
	// Reference: https://developer.apple.com/documentation/quartzcore/cametaldrawable/1478159-texture.
	new(unsafe.Pointer)
}

func (md MetalDrawable) Texture() mtl.Texture { _ = "STUB: not implemented"; return *new(mtl.Texture) }
