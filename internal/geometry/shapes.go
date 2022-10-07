package geometry

import "math"

func MakeBowl(w, h int) *DEM {
	midX := w / 2
	midY := h / 2

	dem := NewDEM(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dx := math.Abs(float64(x - midX))
			dy := math.Abs(float64(y - midY))

			dist := math.Sqrt(dx*dx + dy*dy)

			dem.Set(x, y, math.Pow(dist, 1.25))
		}
	}

	return dem
}
