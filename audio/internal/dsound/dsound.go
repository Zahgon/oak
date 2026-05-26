//go:build windows

package dsound

import (
	"sync"

	"github.com/oov/directsound-go/dsound"
)

// A Config contains the API interfaces initalized by this package
type Config struct {
	Interface *dsound.IDirectSound
	Devices   map[*dsound.GUID]string
}

var cfg Config

var initLock sync.Mutex

// Init initializes directsound or returns an already intialized direct sound instance.
func Init() (Config, error) { _ = "STUB: not implemented"; return *new(Config), nil }

// TODO: providing a GUID which is not nil appears to not succeed, even if those
// GUIDs were returned by Enumerate above

// Call() can return "The operation was completed successfully" as an error
