// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !dragonfly && !openbsd
// +build !linux,!dragonfly,!openbsd

package x11driver

import (
	"unsafe"
)

func shmOpen(size int) (shmid uintptr, addr unsafe.Pointer, err error) {
	_ = "STUB: not implemented"
	return 0, *new(unsafe.Pointer), nil
}

func shmClose(p unsafe.Pointer) error { _ = "STUB: not implemented"; return nil }
