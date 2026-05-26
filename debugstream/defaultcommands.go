package debugstream

import (
	"context"
	"io"
	"sync"
)

var (
	// DefaultCommands to attach to.
	DefaultCommands *ScopedCommands
	defaultsOnce    sync.Once
)

func checkOrCreateDefaults() { _ = "STUB: not implemented"; return }

// AddCommand to the default command set.
// See ScopedCommands' AddComand.
func AddCommand(c Command) error { _ = "STUB: not implemented"; return nil }

// AttachToStream if possible to start consuming the stream
// and executing commands per the stored information in the ScopeCommands.
func AttachToStream(ctx context.Context, input io.Reader, output io.Writer) {
	_ = "STUB: not implemented"
	return
}

// AddDefaultsForScope for debugging.
func AddDefaultsForScope(scopeID int32, controller interface{}) { _ = "STUB: not implemented"; return }
