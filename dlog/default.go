package dlog

import (
	"io"
	"sync"
)

var (
	_ Logger = &logger{}
)

type logger struct {
	bytPool     sync.Pool
	debugLevel  Level
	debugFilter func(string) bool
	writer      io.Writer
}

// NewLogger returns an instance of the default logger with no filter,
// no file, and level set to ERROR
func NewLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// dLog, the primary function of the package,
// prints out and writes to file a string
// containing the logged data separated by spaces,
// prepended with file and line information.
// It only includes logs which pass the current filters.
func (l *logger) dLog(level Level, in ...interface{}) {
	_ = "STUB: not implemented"
	// (pc uintptr, file string, line int, ok bool)
	// TODO: restructure so dlog functions work like t.Helper,
	// incrementing this traceBack value for us
	return
}

// Note on errors: these functions all return
// errors, but they are always nil.

// This can error, but we can't do anything about it if it does.

func truncateFileName(f string) string { _ = "STUB: not implemented"; return "" }

func (l *logger) checkFilter(f string, in ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// SetFilter defines a custom filter function. Log lines that
// return false when passed to this function will not be output.
func (l *logger) SetFilter(filter func(string) bool) { _ = "STUB: not implemented"; return }

// SetLogLevel sets what message levels of debug
// will be printed.
func (l *logger) SetLogLevel(level Level) error { _ = "STUB: not implemented"; return nil }

// Error will write a dlog if the debug level is not NONE
func (l *logger) Error(in ...interface{}) { _ = "STUB: not implemented"; return }

// Info will write a dLog if the debug level is higher than WARN
func (l *logger) Info(in ...interface{}) { _ = "STUB: not implemented"; return }

// Verb will write a dLog if the debug level is higher than INFO
func (l *logger) Verb(in ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }
