package fileutil

import (
	"io"
	"io/fs"
	"os"
)

var (
	// FS is the filesystem that Open, ReadFile and ReadDir will query.
	FS fs.FS = os.DirFS(".")
	// FixWindowsPaths will reset all file paths loaded to replace windows style slashes
	// with unix style slashes. This is important when using io/fs or embed, because the
	// path/filepath package will produce windows style paths on a windows system, but
	// these stdlib packages will reject all windows paths.
	FixWindowsPaths = true
	// OSFallback will fallback to loading via os.Open / io.ReadFile if loading otherwise fails.
	// This is necessary when reading system level fallback fonts. Fixed paths will not be applied
	// to this fallback route.
	OSFallback = true
)

// Open is a wrapper around os.Open that will also check FS to access
// embedded data. The intended use is to use the an embedding library to create an
// Asset function that matches this signature.
func Open(file string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// ReadFile replaces ioutil.ReadFile, trying to use FS.
func ReadFile(file string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadDir replaces ioutil.ReadDir, trying to use FS.
func ReadDir(file string) ([]fs.DirEntry, error) { _ = "STUB: not implemented"; return nil, nil }

func fixWindowsPath(file string) string { _ = "STUB: not implemented"; return "" }
