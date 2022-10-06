package terrain

import "math"

// Valley produces a kind of uniform bowl-shaped landscape.
func Valley(x, y float64) (float64, float64) {
	absX := math.Abs(x)
	absY := math.Abs(y)

	d := math.Sqrt(x*x + y*y)
	sd := (d / 10) * (d / 10)

	return sd * (-x / absX), sd * (-y / absY)
}
