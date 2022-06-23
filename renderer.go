package drifter

type Renderer interface {
	SetData(h *History)
	SetDelta(d float64)
	SetDrawer(d Drawer)
	SetViewport(v *Viewport)
}

func (e *EbitenRenderer) SetData(h *History) {
	e.history = h
}

func (e *EbitenRenderer) SetDelta(d float64) {
	e.timestampDelta = d
}

func (e *EbitenRenderer) SetDrawer(d Drawer) {
	e.drawer = d
}

func (e *EbitenRenderer) SetViewport(v *Viewport) {
	e.simViewport = v
}
