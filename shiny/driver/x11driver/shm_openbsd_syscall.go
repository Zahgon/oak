// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build openbsd && (i386 || amd64)
// +build openbsd
// +build i386 amd64

package x11driver

import (
	"unsafe"
)

// These constants are from /usr/include/sys/ipc.h
const (
	ipcPrivate = 0
	ipcRmID    = 0
)

func shmOpen(size int) (shmid uintptr, addr unsafe.Pointer, err error) {
	_ = "STUB: not implemented"
	return 0, *new(unsafe.Pointer), nil
}

func shmClose(p unsafe.Pointer) error { _ = "STUB: not implemented"; return nil }
