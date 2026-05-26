// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

package mtldriver

import "image"

// bufferImpl implements screen.Buffer.
type bufferImpl struct {
	rgba *image.RGBA
}

func (*bufferImpl) Release()            { _ = "STUB: not implemented"; return }
func (b *bufferImpl) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }
func (b *bufferImpl) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}
func (b *bufferImpl) RGBA() *image.RGBA { _ = "STUB: not implemented"; return nil }
