package drifter

import (
	"github.com/fogleman/gg"
	"image/color"
	"io"
)

type PNGRenderer struct {
	delta    float64
	drawer   Drawer
	history  *History
	viewport *Viewport
}

func NewPNGRenderer() *PNGRenderer {
	return &PNGRenderer{
		delta:    1.0 / 60.0,
		history:  &History{},
		viewport: NewViewportSym(10),
	}
}

func (p *PNGRenderer) SetData(h *History) {
	p.history = h
}

func (p *PNGRenderer) SetDelta(d float64) {
	p.delta = d
}

func (p *PNGRenderer) SetDrawer(d Drawer) {
	p.drawer = d
}

func (p *PNGRenderer) SetViewport(v *Viewport) {
	p.viewport = v
}

func (p *PNGRenderer) WriteFinal(w io.Writer, width, height int) error {
	widthFloat, heightFloat := float64(width), float64(height)
	widthScale := widthFloat / p.viewport.WidthFloat()
	heightScale := heightFloat / p.viewport.HeightFloat()

	canvas := gg.NewContext(width, height)

	canvas.Push()
	canvas.SetColor(color.Black)
	canvas.DrawRectangle(0, 0, widthFloat, heightFloat)
	canvas.Fill()
	canvas.Pop()

	canvas.Push()
	canvas.Translate(widthFloat/2, heightFloat/2)
	canvas.Scale(widthScale, -heightScale)
	for _, t := range p.history.Traces {
		p.drawer.Trace(t, p.history.LastTimestamp, canvas)
	}
	canvas.Pop()

	return canvas.EncodePNG(w)
}
