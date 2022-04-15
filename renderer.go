package drifter

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// A Renderer knows how to manage and draw a Sim.
//
// `BlockCount` - the number of blocks to divide the viewport into for rendering vectors
// `DrawField` - whether to draw the vector field itself
// `sim` - the sim to be managed and rendered
// `logicalViewport` - the viewport relative to the Sim
// `realViewport` - the viewport relative to the current output canvas
type Renderer struct {
	BlockCount      int
	DrawField       bool
	Drawer          Drawer
	sim             *Sim
	logicalViewport *Viewport
	realViewport    *Viewport
}

func NewRenderer(sim *Sim) *Renderer {
	logicalViewport := NewViewport(-10, -10, 10, 10)
	imageViewport := NewViewport(0, 10, 10, 0)

	return &Renderer{
		BlockCount:      10,
		Drawer:          NewSimplerDrawer(),
		sim:             sim,
		logicalViewport: logicalViewport,
		realViewport:    imageViewport,
	}
}

func (r *Renderer) Viewport() *Viewport {
	return r.logicalViewport
}

func (r *Renderer) Update() error {
	r.sim.Advance(1.0 / 60.0)
	return nil
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	screen.Clear()

	canvas := gg.NewContextForImage(screen)
	canvas.Push()
	canvas.Translate(r.realViewport.WidthFloat()/2, r.realViewport.HeightFloat()/2)
	canvas.Scale(24, -24)

	for _, trace := range r.sim.traces {
		r.Drawer.Trace(trace, canvas)
	}

	canvas.Pop()

	screenImage := ebiten.NewImageFromImage(canvas.Image())
	screen.Fill(color.Black)
	screen.DrawImage(screenImage, nil)
}

func (r *Renderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	r.realViewport.Y1 = outsideHeight
	r.realViewport.X2 = outsideWidth

	return outsideWidth, outsideHeight
}
