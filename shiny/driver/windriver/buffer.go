// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package windriver

import (
	"image"
	"sync"
	"syscall"
)

type bufferImpl struct {
	hbitmap syscall.Handle
	buf     []byte
	rgba    image.RGBA
	size    image.Point

	mu        sync.Mutex
	nUpload   uint32
	released  bool
	cleanedUp bool
}

func (b *bufferImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }
func (b *bufferImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}
func (b *bufferImpl) RGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }

func (b *bufferImpl) preUpload() {
	_ = "STUB: not implemented"
	// Check that the program hasn't tried to modify the rgba field via the
	// pointer returned by the bufferImpl.RGBA method. This check doesn't catch
	// 100% of all cases; it simply tries to detect some invalid uses of a
	// screen.Image such as:
	//
	//	*buffer.RGBA() = anotherImageRGBA
	return
}

func (b *bufferImpl) postUpload() { _ = "STUB: not implemented"; return }

func (b *bufferImpl) Release() { _ = "STUB: not implemented"; return }

func (b *bufferImpl) cleanUp() { _ = "STUB: not implemented"; return }

func (b *bufferImpl) blitToDC(dc syscall.Handle, dp image.Point, sr image.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}
