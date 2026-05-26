// Package riff reads and umarshalls RIFF files.
package riff

import (
	"bytes"
	"reflect"
)

// A Reader is a bytes reader with some helper functions to read IDs, Lens, and Data
// from RIFF files.
type Reader struct {
	*bytes.Reader
}

// NewReader returns an initial Reader
func NewReader(data []byte) *Reader { _ = "STUB: not implemented"; return nil }

// Print prints a reader without any knowledge of the structure of the reader,
// so all values will be []bytes.
// It assumes the reader has not advanced at all. Todo: Change that
func (r *Reader) Print() { _ = "STUB: not implemented"; return }

func deepPrint(r *Reader, prefix string, readLimit int) { _ = "STUB: not implemented"; return }

// There will be a bogus byte at the end of some prints.

// Unmarshal is a mirror of json.Unmarshal, for RIFF files
func Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Reader) unmarshal(v interface{}) error {
	_ = "STUB: not implemented"
	// Mirrors json.unmarshal
	return nil
}

// The first ID in the riff should be RIFF

// The next ID identifies this file type. We don't want it.

// NextID returns the next four byte sof the reader as a string
func (r *Reader) NextID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// NextIDLen returns NextID and NextLen
func (r *Reader) NextIDLen() (string, uint32, bool, error) {
	_ = "STUB: not implemented"
	return "", 0, false, nil
}

// NextLen returns the next four bytes of the reader as a length.
func (r *Reader) NextLen() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) chunks(rv reflect.Value, inLength int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	// Find chunkId in rv
	// If it can't be found, ignore it as a value the user does not want
	return *new(reflect.Value), nil
}

func (r *Reader) sliceChunks(rv reflect.Value, inLength int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// structChunks reads chunks and matches them to fields on rv (which is a struct)
// structChunks sets the fields of rv to be the output it gets
func (r *Reader) structChunks(rv reflect.Value, inLength int) error {
	_ = "STUB: not implemented"
	return nil
}

//spew.Dump(fields[i])

// get contents from recursive call

// if length is odd read one more

// next id

// Skip this id
// if length is odd read one more

// next id

// Todo: the switch here should change to some separate functions, there's some
// repetition here that is not necessary.
func (r *Reader) fieldValue(rv reflect.Value, ln uint32) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Something on this struct has an undefined size
// Read each field in part by part.
