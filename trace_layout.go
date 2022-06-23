package drifter

import (
	"math"
	"math/rand"
)

func RingLayout(count int, radius float64) []*Trace {
	traces := make([]*Trace, count)
	for i := 0; i < count; i++ {
		theta := rand.Float64() * 2 * math.Pi
		x, y := math.Cos(theta)*radius, math.Sin(theta)*radius
		traces[i] = NewTrace(x, y)
	}
	return traces
}

func RandomLayout(count int, viewport *Viewport) []*Trace {
	traces := make([]*Trace, count)
	for i := 0; i < count; i++ {
		x := rand.Float64()*viewport.WidthFloat() + float64(viewport.X1)
		y := rand.Float64()*viewport.HeightFloat() + float64(viewport.Y1)
		traces[i] = NewTrace(x, y)
	}
	return traces
}

func GridLayout(across, down int, viewport *Viewport) []*Trace {
	dx := viewport.WidthFloat() / (float64(across) + 1)
	dy := viewport.HeightFloat() / (float64(down) + 1)
	count := across * down

	index := 0
	traces := make([]*Trace, count)
	for row := 0; row < down; row++ {
		for col := 0; col < across; col++ {
			x := float64(viewport.X1) + dx*float64(col) + dx
			y := float64(viewport.Y1) + dy*float64(row) + dy
			traces[index] = NewTrace(x, y)
			index++
		}
	}

	return traces
}
