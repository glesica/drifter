package drifter

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type EbitenRenderer struct {
	drawer         Drawer
	history        *History
	nextTimestamp  float64
	renderViewport *Viewport
	simViewport    *Viewport
	timestamp      float64
	timestampDelta float64
}

func NewEbitenRenderer() *EbitenRenderer {
	simViewport := NewViewport(-10, -10, 10, 10)
	// This will be set the first time it is drawn to the screen
	renderViewport := NewViewport(0, 0, 0, 0)

	return &EbitenRenderer{
		history:        &History{},
		renderViewport: renderViewport,
		simViewport:    simViewport,
		timestampDelta: 1.0 / 60.0,
	}
}

func (e *EbitenRenderer) Update() error {
	if e.timestamp+e.timestampDelta < e.history.LastTimestamp {
		e.nextTimestamp = e.timestamp + e.timestampDelta
	} else {
		e.nextTimestamp = e.history.LastTimestamp
	}
	return nil
}

func (e *EbitenRenderer) Draw(screen *ebiten.Image) {
	canvas := gg.NewContextForImage(screen)
	canvas.Push()
	canvas.Translate(e.renderViewport.WidthFloat()/2, e.renderViewport.HeightFloat()/2)

	widthScale := e.renderViewport.WidthFloat() / e.simViewport.WidthFloat()
	heightScale := e.renderViewport.HeightFloat() / e.simViewport.HeightFloat()
	canvas.Scale(widthScale, -heightScale)

	for _, t := range e.history.Traces {
		e.drawer.Trace(t, e.nextTimestamp, canvas)
	}

	canvas.Pop()

	screenImage := ebiten.NewImageFromImage(canvas.Image())
	screen.Fill(color.Black)
	screen.DrawImage(screenImage, nil)

	e.timestamp = e.nextTimestamp
}

func (e *EbitenRenderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	e.renderViewport.Y1 = outsideHeight
	e.renderViewport.X2 = outsideWidth

	return outsideWidth, outsideHeight
}
