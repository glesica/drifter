package drifter

import (
	"image/color"
	"math"
)

type Drawer interface {
	Trace(trace *TraceHistory, ts float64, canvas Canvas)
}

type SimpleDrawer struct {
	colorPicker ColorPicker
}

func NewSimplerDrawer(colorPicker ColorPicker) *SimpleDrawer {
	return &SimpleDrawer{
		colorPicker: colorPicker,
	}
}

func (s *SimpleDrawer) Trace(t *TraceHistory, ts float64, canvas Canvas) {
	var traceColor color.Color
	prevFrame := t.Frames[0]

	canvas.SetLineWidth(2)
	canvas.MoveTo(t.Frames[0].X, t.Frames[0].Y)

	for i := 1; i < t.Length(); i++ {
		frame := t.Frames[i]
		if frame.Timestamp > ts {
			break
		}

		d := math.Sqrt(math.Pow(frame.X-prevFrame.X, 2) + math.Pow(frame.Y-prevFrame.Y, 2))
		if d < 1.0 {
			continue
		}

		traceColor = s.colorPicker.TraceColor(frame)
		canvas.SetColor(traceColor)

		canvas.LineTo(frame.X, frame.Y)

		canvas.Stroke()

		canvas.MoveTo(frame.X, frame.Y)
		prevFrame = frame
	}

	//canvas.Push()
	//canvas.SetColor(colornames.Aqua)
	//canvas.DrawPoint(t.FirstX(), t.FirstY(), 5)
	//canvas.Fill()
	//canvas.Pop()
}
