package render

import (
	"regexp"

	"github.com/oakmound/oak/v4/alg/intgeom"
)

// BatchLoad loads subdirectories from the given base folder and imports all files,
// using alias rules to automatically determine the size of sprites and sheets in
// subfolders.
func BatchLoad(baseFolder string) error { _ = "STUB: not implemented"; return nil }

// BlankBatchLoad acts like BatchLoad, but will not load and instead return a blank image
// of the appropriate dimensions for anything above maxFileSize.
func BlankBatchLoad(baseFolder string, maxFileSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

// We've been given a bad base directory

// Ignore files we know we can't parse

var (
	sheetFileRegex      = regexp.MustCompile(`^[^\d]*(\d+)x(\d+)\..*$`)
	sheetDirectoryRegex = regexp.MustCompile(`^[^\d]*(\d+)x(\d+)$`)
)

func shouldLoadSheet(file string) (intgeom.Point2, bool) {
	_ = "STUB: not implemented"
	// when should we determine a file should be loaded as a sheet?
	// 1. If the file itself ends in that syntax: image_%dx%d.png
	// 2. If the file's final directory ends in special syntax: (%dx%d)
	// ... preferring the former
	return *new(intgeom.Point2), false
}

// if len matches != 3, or if these fail to parse, our regex is wrong
