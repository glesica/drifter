package drifter

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type Drawer interface {
	Trace(trace *Trace, canvas Canvas)
}

type SimpleDrawer struct{}

func NewSimplerDrawer() *SimpleDrawer {
	return &SimpleDrawer{}
}

func (s *SimpleDrawer) Trace(trace *Trace, canvas Canvas) {
	traceColorName := colornames.Names[trace.ID%len(colornames.Names)]
	traceColor := colornames.Map[traceColorName]

	canvas.Push()

	canvas.SetColor(traceColor)
	canvas.SetLineWidth(2)
	canvas.MoveTo(trace.X, trace.Y)
	for i := len(trace.PathX) - 1; i >= 0; i-- {
		canvas.LineTo(trace.PathX[i], trace.PathY[i])
	}
	canvas.Stroke()

	canvas.SetColor(color.White)
	canvas.DrawPoint(trace.X, trace.Y, 5)
	canvas.Fill()

	canvas.Pop()
}
