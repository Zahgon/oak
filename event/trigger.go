package event

// TriggerForCaller acts like Trigger, but will only trigger for the given caller.
func (bus *Bus) TriggerForCaller(callerID CallerID, eventID UnsafeEventID, data interface{}) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// Trigger will scan through the event bus and call all bindables found attached
// to the given event, with the passed in data.
func (bus *Bus) Trigger(eventID UnsafeEventID, data interface{}) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// TriggerOn calls Trigger with a strongly typed event.
func TriggerOn[T any](b Handler, ev EventID[T], data T) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// TriggerForCallerOn calls TriggerForCaller with a strongly typed event.
func TriggerForCallerOn[T any](b Handler, cid CallerID, ev EventID[T], data T) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}
