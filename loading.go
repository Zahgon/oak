package oak

import (
	"io/fs"
)

func (w *Window) loadAssets(imageDir, audioDir string) { _ = "STUB: not implemented"; return }

func (w *Window) endLoad() { _ = "STUB: not implemented"; return }

// SetFS updates all calls oak or oak's subpackages will make to read from the given filesystem.
// By default, this is set to os.DirFS(".")
func SetFS(filesystem fs.FS) { _ = "STUB: not implemented"; return }
