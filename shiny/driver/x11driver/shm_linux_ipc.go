// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || ppc64 || ppc64le || s390x)
// +build linux
// +build 386 ppc64 ppc64le s390x

package x11driver

import (
	"unsafe"
)

// These constants are from /usr/include/linux/ipc.h
const (
	ipcPrivate = 0
	ipcRmID    = 0

	shmAt  = 21
	shmDt  = 22
	shmGet = 23
	shmCtl = 24
)

func shmOpen(size int) (shmid uintptr, addr unsafe.Pointer, err error) {
	_ = "STUB: not implemented"
	return 0, *new(unsafe.Pointer), nil
}

func shmClose(p unsafe.Pointer) error { _ = "STUB: not implemented"; return nil }
