package audio

import (
	"github.com/oakmound/oak/v4/audio/pcm"
)

var _ pcm.Reader = &LoopingReader{}
var _ pcm.Reader = &BytesReader{}

// LoopReader will cache read bytes as they are read and resend them after the reader returns EOF.
func LoopReader(r pcm.Reader) pcm.Reader { _ = "STUB: not implemented"; return *new(pcm.Reader) }

// A LoopingReader will read from Reader continually, even after it has been fully consumed. The data read
// from the reader will be cached after read within the LoopingReader structure, potentially inflating memory
// if provided a large stream.
type LoopingReader struct {
	pcm.Reader
	buffer     []byte
	bufferPos  int
	eofReached bool
}

func (l *LoopingReader) ReadPCM(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// Note a quirk of this implementation: read calls in succession
	// will not return buffers of a similar size, instead of fully populating
	// the requested p we assert the caller will recall Read resuming from
	// the front of our buffer
	return 0, nil
}

// A BytesReader acts like a bytes.Buffer for converting raw []bytes into pcm Readers.
type BytesReader struct {
	pcm.Format
	Buffer []byte
	Offset int
}

func (b *BytesReader) ReadPCM(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *BytesReader) Copy() *BytesReader { _ = "STUB: not implemented"; return nil }

// ReadAll will read all of the content within a reader and convert it into a BytesReader. Use carefully; use on
// a LoopingReader or reader which generates its data (e.g. synth types) will likely read until OOM.
func ReadAll(r pcm.Reader) *BytesReader { _ = "STUB: not implemented"; return nil }

// Add more capacity (let append pick how much).

// ReadFull acts like io.ReadFull with a pcm Reader. It will read until the provided buffer
// is competely populated by the reader.
func ReadFull(r pcm.Reader, buf []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadAtLeast acts like io.ReadAtLeast with a pcm Reader. It will read until at least min
// bytes have been read into the provided buffer.
func ReadAtLeast(r pcm.Reader, buf []byte, min int) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
