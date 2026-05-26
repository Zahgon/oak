package event

// Q: Why do Bind / Unbind / etc not immediately take effect?
// A: For concurrent safety, most operations on a bus lock the bus. Triggers acquire a read lock on the bus,
//    as they iterate over internal bus components. Most logic within an event bus will happen from within
//    a Trigger call-- when an entity is destroyed by some collision, for example, all of its bindings should
//    be unregistered. If one were to call Unbind from within a call to Trigger, the trigger would never release
//    its lock-- so the unbind would never be able to take the lock-- so the bus would be unrecoverably stuck.

// Q: Why not trust users to call Bind / Unbind / etc with `go`, to allow the caller to decide when to use
//    concurrency?
// A: It is almost never correct to not call such a function with `go`, and it is a bad user experience for
//    the engine to deadlock unexpectedly because you forgot to begin some call with a goroutine.

// A Binding, returned from calls to Bind, references the details of a binding and where that binding is
// stored within a handler. The common use case for this structure would involve a system that wanted to
// keep track of its bindings for later remote unbinding. This structure can also be used to construct
// and unbind a known reference.
type Binding struct {
	Handler  Handler
	EventID  UnsafeEventID
	CallerID CallerID
	BindID   BindID

	busResetCount int64

	// Bound is closed once the binding has been applied. Wait on this condition carefully; bindings
	// will not take effect while an event is being triggered (e.g. in a event callback's returning thread)
	Bound <-chan struct{}
}

// Unbind unbinds the callback associated with this binding from it's own event handler. If this binding
// does not belong to its handler or has already been unbound, this will do nothing.
func (b Binding) Unbind() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// A BindID is a unique identifier for a binding within a bus.
type BindID int64

// UnsafeBind registers a callback function to be called whenever the provided event is triggered
// against this bus. The binding is concurrently bound, and therefore may not be immediately
// available to be triggered. When Reset is called on a Bus, all prior bindings are unbound and any
// concurrent calls to UnsafeBind will not take effect. This call is 'unsafe' because UnsafeBindables
// use bare interface{} types.
func (bus *Bus) UnsafeBind(eventID UnsafeEventID, callerID CallerID, fn UnsafeBindable) Binding {
	_ = "STUB: not implemented"
	return *new(Binding)
}

// The event bus has reset while we we were waiting to bind this

// PersistentBind calls UnsafeBind, and causes UnsafeBind to be called with these inputs when a Bus is Reset, i.e.
// persisting the binding through bus resets. Unbinding this will not stop it from being rebound on the next
// Bus Reset-- ClearPersistentBindings will. If called concurrently during a bus Reset, the request may not be
// bound until the next bus Reset.
func (bus *Bus) PersistentBind(eventID UnsafeEventID, callerID CallerID, fn UnsafeBindable) Binding {
	_ = "STUB: not implemented"
	return *new(Binding)
}

// Unbind unregisters a binding from a bus concurrently. Once complete, triggers that would
// have previously caused the Bindable callback to execute will no longer do so.
func (bus *Bus) Unbind(loc Binding) <-chan struct{} { _ = "STUB: not implemented"; return nil }

// This binding is not valid for this bus (in this state)

// A Bindable is a strongly typed callback function to be executed on Trigger. It must be paired
// with an event registered via RegisterEvent.
type Bindable[C any, Payload any] func(C, Payload) Response

// Bind will cause the function fn to be called whenever the event ev is triggered on the given event handler. The function
// will be called with the provided caller as its first argument, and will also be called when the provided event is specifically
// triggered on the caller's ID.
func Bind[C Caller, Payload any](h Handler, ev EventID[Payload], caller C, fn Bindable[C, Payload]) Binding {
	_ = "STUB: not implemented"
	return *new(Binding)
}

// A GlobalBindable is a bindable that is not bound to a specific caller.
type GlobalBindable[Payload any] func(Payload) Response

// GlobalBind will cause the function fn to be called whenever the event ev is triggered on the given event handler.
func GlobalBind[Payload any](h Handler, ev EventID[Payload], fn GlobalBindable[Payload]) Binding {
	_ = "STUB: not implemented"
	return *new(Binding)
}

// UnsafeBindable defines the underlying signature of all bindings.
type UnsafeBindable func(CallerID, Handler, interface{}) Response

// UnbindAllFrom unbinds all bindings currently bound to the provided caller via ID.
func (bus *Bus) UnbindAllFrom(c CallerID) <-chan struct{} { _ = "STUB: not implemented"; return nil }
