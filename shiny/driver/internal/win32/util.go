package win32

func boolToBOOL(value bool) BOOL { _ = "STUB: not implemented"; return *new(BOOL) }

// MakeGUID allocates a GUID from a [16]byte array. If the array
// is uninitialized a random byte array will be used.
func MakeGUID(guid [16]byte) GUID { _ = "STUB: not implemented"; return *new(GUID) }

// Implementation from hallazzang/go-windows-programming. Would have
// reimplemented but it's so simple that its hard to do so.
