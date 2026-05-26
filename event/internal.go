package event

type bindableList map[BindID]UnsafeBindable

func (eb *Bus) getBindableList(eventID UnsafeEventID, callerID CallerID) bindableList {
	_ = "STUB: not implemented"
	return *new(bindableList)
}

func (bus *Bus) trigger(binds bindableList, eventID UnsafeEventID, callerID CallerID, data interface{}) {
	_ = "STUB: not implemented"
	return
}

// Q: Why does this call bus.Unbind when it already has the event index to delete?
// A: This goroutine does not own a write lock on the bus, and should therefore
//    not modify its contents. We do not have a simple way of promoting our read lock
//    to a write lock.
