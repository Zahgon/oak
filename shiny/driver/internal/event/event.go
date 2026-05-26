// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package event provides an infinitely buffered double-ended queue of events.
package event

import (
	"sync"
)

// Deque is an infinitely buffered double-ended queue of events. The zero value
// is usable, but a Deque value must not be copied.
type Deque struct {
	mu    sync.Mutex
	cond  sync.Cond     // cond.L is lazily initialized to &Deque.mu.
	back  []interface{} // FIFO.
	front []interface{} // LIFO.
}

func (q *Deque) lockAndInit() { _ = "STUB: not implemented"; return }

// NextEvent implements the screen.EventDeque interface.
func (q *Deque) NextEvent() interface{} { _ = "STUB: not implemented"; return nil }

// Send implements the screen.EventDeque interface.
func (q *Deque) Send(event interface{}) { _ = "STUB: not implemented"; return }
