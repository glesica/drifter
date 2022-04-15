package drifter

import "math"

// A Viewport represents a window into a coordinate space, which could be
// either the logical Sim space or the real Renderer space.
//
// X1 - lower-left X
// Y1 - lower-left Y
// X2 - upper-right X
// Y2 - upper-right Y
type Viewport struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func NewViewport(x1, y1, x2, y2 int) *Viewport {
	return &Viewport{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

func (v *Viewport) Resize(x1, y1, x2, y2 int) {
	v.X1 = x1
	v.Y1 = y1
	v.X2 = x2
	v.Y2 = y2
}

func (v *Viewport) Width() int {
	return v.X2 - v.X1
}

func (v *Viewport) WidthFloat() float64 {
	return float64(v.X2 - v.X1)
}

func (v *Viewport) Height() int {
	return v.Y2 - v.Y1
}

func (v *Viewport) HeightFloat() float64 {
	return math.Abs(float64(v.Y2 - v.Y1))
}

func (v *Viewport) Mirror() bool {
	return v.Y1 > v.Y2
}
