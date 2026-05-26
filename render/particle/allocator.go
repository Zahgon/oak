package particle

import (
	"github.com/oakmound/oak/v4/event"
)

const (
	blockSize = 2048
)

// An Allocator can allocate ids for particles
type Allocator struct {
	particleBlocks map[int]event.CallerID
	nextOpenCh     chan int
	freeCh         chan int
	allocCh        chan event.CallerID
	requestCh      chan int
	responseCh     chan event.CallerID
	stopCh         chan struct{}
}

// NewAllocator creates a new allocator
func NewAllocator() *Allocator { _ = "STUB: not implemented"; return nil }

// Run spins up an allocator to accept allocation requests. It will run until
// Stop is called. This is a blocking call.
func (a *Allocator) Run() { _ = "STUB: not implemented"; return }

// DefaultAllocator is an allocator that starts running as soon as this package is imported.
var DefaultAllocator = NewAllocator()

// This is an always-called init instead of Init because oak does not import this
// package by default. If this package is not used, it will not run this goroutine.
func init() {
	go DefaultAllocator.Run()
}

func (a *Allocator) freereceive(i int) int { _ = "STUB: not implemented"; return 0 }

// Allocate requests a new block in the particle space for the given cid
func (a *Allocator) Allocate(id event.CallerID) int { _ = "STUB: not implemented"; return 0 }

// Deallocate requests that the given block be removed from the particle space
func (a *Allocator) Deallocate(block int) {
	_ = "STUB: not implemented"

	// LookupSource requests the source that generated a pid
	return
}

func (a *Allocator) LookupSource(id int) *Source { _ = "STUB: not implemented"; return nil }

// TODO: not default?

// Lookup requests a specific particle in the particle space
func (a *Allocator) Lookup(id int) Particle { _ = "STUB: not implemented"; return *new(Particle) }

// Stop stops the allocator's ongoing Run. Once stopped, allocator may not be reused.
// Stop must not be called more than once.
func (a *Allocator) Stop() { _ = "STUB: not implemented"; return }
