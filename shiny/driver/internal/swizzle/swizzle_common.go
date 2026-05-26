// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package swizzle provides functions for converting between RGBA pixel
// formats.
package swizzle

// BGRA converts a pixel buffer between Go's RGBA and other systems' BGRA byte
// orders.
//
// It panics if the input slice length is not a multiple of 4.
func BGRA(p []byte) { _ = "STUB: not implemented"; return }

// Use asm code for 16- or 4-byte chunks, if supported.
