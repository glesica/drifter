package geometry

import (
	"math"
)

type DEM struct {
	heights []float64
	width   int
	height  int
}

func NewDEM(w, h int) *DEM {
	return &DEM{
		heights: make([]float64, w*h),
		width:   w,
		height:  h,
	}
}

func (d *DEM) Set(x, y int, elev float64) {
	d.heights[y*d.width+x] = elev
}

// At produces the elevation in the given grid space.
func (d *DEM) At(x, y int) float64 {
	return d.heights[y*d.width+x]
}

// AtInter produces an interpolated elevation for the given point
// within the DEM bounds. Interpolation is linear.
// TODO: Need to handle region edges properly
func (d *DEM) AtInter(x, y float64) float64 {
	xBase := math.Floor(x)
	yBase := math.Floor(y)

	xOff := x - xBase
	yOff := y - yBase

	var xOther, yOther float64
	if xOff < 0.5 {
		xOther = xBase - 1
	} else {
		xOther = xBase + 1
	}
	if yOff < 0.5 {
		yOther = yBase - 1
	} else {
		yOther = yBase + 1
	}

	xBaseWt := 1.0 - math.Abs(0.5-xOff)
	yBaseWt := 1.0 - math.Abs(0.5-yOff)

	xOtherWt := 1.0 - xBaseWt
	yOtherWt := 1.0 - yBaseWt

	baseHt := d.At(int(xBase), int(yBase))
	xOtherHt := d.At(int(xOther), int(yBase))
	yOtherHt := d.At(int(xBase), int(yOther))

	return (baseHt*xBaseWt + baseHt*yBaseWt + xOtherHt*xOtherWt + yOtherHt*yOtherWt) / 2.0
}

// AccelAt returns the estimated acceleration at the given point
// within the region covered by the DEM, where x ranges within
// [0.0, d.width] and y ranges within [0.0, d.height].
func (d *DEM) AccelAt(x, y float64) (float64, float64) {
	centerX, centerY := d.discretize(x, y)

	localDX := x - centerX
	localDY := y - centerY

	// Return the slope for that triangle
	x0, y0 := int(centerX), int(centerY)

	var x1, y1, x2, y2 int
	switch {
	case localDX >= 0.5 && localDY >= 0.5:
		// Northeast
		x1, y1 = x0, y0+1
		x2, y2 = x0+1, y0
	case localDX >= 0.5 && localDY < 0.5:
		// Southeast
		x1, y1 = x0+1, y0
		x2, y2 = x0, y0-1
	case localDX < 0.5 && localDY < 0.5:
		// Southwest
		x1, y1 = x0, y0-1
		x2, y2 = x0-1, y0
	case localDX < 0.5 && localDY >= 0.5:
		// Northwest
		x1, y1 = x0-1, y0
		x2, y2 = x0, y0+1
	}

	// Account for edges

	x1, y1, x2, y2 = d.clamp(x1, y1, x2, y2)

	pt0 := Vec{X: float64(x0), Y: float64(y0), Z: d.At(x0, y0)}
	pt1 := Vec{X: float64(x1), Y: float64(y1), Z: d.At(x1, y1)}
	pt2 := Vec{X: float64(x2), Y: float64(y2), Z: d.At(x2, y2)}

	dx, dy := Slopes(pt0, pt1, pt2)
	return -dx, -dy
}

func (d *DEM) Size() (int, int) {
	return d.width, d.height
}

// clamp adjusts the 2nd and 3rd points of the triangle to be
// in the opposite direction if they go off the edge of the DEM.
func (d *DEM) clamp(x1, y1, x2, y2 int) (int, int, int, int) {
	if x1 < 0 {
		x1 += 2
	} else if x1 >= d.width {
		x1 -= 2
	}

	if y1 < 0 {
		y1 += 2
	} else if y1 >= d.height {
		y1 -= 2
	}

	if x2 < 0 {
		x2 += 2
	} else if x2 >= d.width {
		x2 -= 2
	}

	if y2 < 0 {
		y2 += 2
	} else if y2 >= d.height {
		y2 -= 2
	}

	return x1, y1, x2, y2
}

func (d *DEM) discretize(x, y float64) (xi float64, yi float64) {
	if x > 0.0 {
		xi = math.Ceil(x - 1.0)
	}
	if y > 0.0 {
		yi = math.Ceil(y - 1.0)
	}
	return
}
