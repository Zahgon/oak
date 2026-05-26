package collision

// An OnHit is a function which takes in two spaces
type OnHit func(s, s2 *Space)

// CallOnHits will send a signal to the passed in channel
// when it has completed all collision functions in the hitmap.
func CallOnHits(s *Space, onHits map[Label]OnHit, doneCh chan bool) {
	_ = "STUB: not implemented"
	return
}

// CallOnHits will send a signal to the passed in channel
// when it has completed all collision functions in the hitmap.
func (t *Tree) CallOnHits(s *Space, onHits map[Label]OnHit, doneCh chan bool) {
	_ = "STUB: not implemented"
	return
}

// This waits to send our signal that we've
// finished until we've counted signals for
// each collision entity

// OnIDs converts a function on two CIDs to an OnHit
func OnIDs(fn func(int, int)) func(s, s2 *Space) { _ = "STUB: not implemented"; return nil }
