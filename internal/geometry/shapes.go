package geometry

import "math"

func MakeBowl(w, h int, depth float64) *DEM {
	radius := math.Sqrt(0.25*float64(w*w) + 0.25*float64(h*h))

	midX := w / 2
	midY := h / 2

	dem := NewDEM(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dx := math.Abs(float64(x - midX))
			dy := math.Abs(float64(y - midY))

			dist := math.Sqrt(dx*dx + dy*dy)

			dem.Set(x, y, dist/radius*depth)
		}
	}

	return dem
}
