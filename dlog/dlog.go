package dlog

import (
	"io"
)

// A Logger is a minimal log interface for the content oak wants to log:
// four levels of logging.
type Logger interface {
	Error(...interface{})
	Info(...interface{})
	Verb(...interface{})
	SetFilter(func(string) bool)
	SetLogLevel(l Level) error
	SetOutput(io.Writer)
}

// DefaultLogger is the Logger which all oak log messages are passed through.
var DefaultLogger = NewLogger()

// ErrorCheck checks that the input is not nil, then calls Error on it if it is
// not. Otherwise it does nothing.
// Emits the input error as is for additional processing if desired.
func ErrorCheck(in error) error { _ = "STUB: not implemented"; return nil }

// Error will write a log if the debug level is not NONE
func Error(vs ...interface{}) { _ = "STUB: not implemented"; return }

// Info will write a log if the debug level is higher than ERROR
func Info(vs ...interface{}) { _ = "STUB: not implemented"; return }

// Verb will write a log if the debug level is higher than INFO
func Verb(vs ...interface{}) { _ = "STUB: not implemented"; return }

// SetFilter defines a custom filter function. Log lines that
// return false when passed to this function will not be output.
func SetFilter(filter func(string) bool) { _ = "STUB: not implemented"; return }

// SetLogLevel sets the log level of the default logger.
func SetLogLevel(l Level) error { _ = "STUB: not implemented"; return nil }

// SetOutput will output logs on the default logger to be written to
// the given writer.
func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }
