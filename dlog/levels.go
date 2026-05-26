package dlog

// Level represents the levels a debug message can have
type Level int

// Level values const
const (
	NONE Level = iota
	ERROR
	INFO
	VERBOSE
)

var logLevels = map[Level]string{
	NONE:    "NONE",
	ERROR:   "ERROR",
	INFO:    "INFO",
	VERBOSE: "VERBOSE",
}

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

// ParseDebugLevel parses the input string as a known debug levels
func ParseDebugLevel(level string) (Level, error) {
	_ = "STUB: not implemented"
	return *new(Level), nil
}
