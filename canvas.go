package drifter

import (
	"image/color"
)

// A Canvas represents a drawing surface with helpful methods for creating
// shapes and other graphical elements.
//
// The interface is based on gg.Context in order to provide a limited,
// potentially replaceable, subset of that functionality.
type Canvas interface {
	DrawCircle(x, y, r float64)
	DrawLine(x1, y1, x2, y2 float64)
	DrawPoint(x, y, r float64)
	Fill()
	LineTo(x, y float64)
	MoveTo(x, y float64)
	Push()
	Pop()
	SetColor(color color.Color)
	SetLineWidth(width float64)
	Stroke()
}
