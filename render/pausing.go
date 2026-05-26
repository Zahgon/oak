package render

// CanPause types have pause functions to start and stop animation
type CanPause interface {
	Pause()
	Unpause()
}

type pauseBool struct {
	playing bool
}

func (p *pauseBool) Pause() { _ = "STUB: not implemented"; return }

func (p *pauseBool) Unpause() { _ = "STUB: not implemented"; return }
