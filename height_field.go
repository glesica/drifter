package drifter

import "math"

const aScale = 0.02

type HeightField struct {
	heights     []float64
	width       int
	widthFloat  float64
	height      int
	heightFloat float64
}

func NewHeightField(heights []float64, width int) *HeightField {
	if len(heights)%width != 0 {
		panic("invalid height field size")
	}

	height := len(heights) / width
	return &HeightField{
		heights:     heights,
		width:       width,
		widthFloat:  float64(width),
		height:      height,
		heightFloat: float64(height),
	}
}

func (h *HeightField) Acceleration(x, y float64) (float64, float64) {
	if !h.Valid(x, y) {
		panic("coordinates out of bounds")
	}

	x, y = h.Offset(x, y)
	dxIn, dyIn := x-math.Floor(x), y-math.Floor(y)
	ix, iy := int(x), int(y)

	hSelf := h.Height(ix, iy)
	hEast := h.Height(ix+1, iy)
	hWest := h.Height(ix-1, iy)
	hNorth := h.Height(ix, iy+1)
	hSouth := h.Height(ix, iy-1)

	dxAvg := ((hEast-hSelf)*dxIn + (hSelf-hWest)*(1-dxIn)) / 2
	dyAvg := ((hNorth-hSelf)*dyIn + (hSelf-hSouth)*(1-dyIn)) / 2

	return -dxAvg * aScale, -dyAvg * aScale
}

func (h *HeightField) Damping(x, y float64) float64 {
	return 0.1
}

func (h *HeightField) Height(x, y int) float64 {
	// Adjust y for the cartesian plane.
	y = h.height - 1 - y

	// Clamp to the area since we assume that the border around the field
	// has the same height as the adjacent cells.
	if x >= h.width {
		x = h.width - 1
	}
	if x < 0 {
		x = 0
	}
	if y >= h.height {
		y = h.height - 1
	}
	if y < 0 {
		y = 0
	}

	return h.heights[y*h.width+x]
}

// Offset adjusts the point so that we're in the first quadrant to simplify
// indexing into the underlying height array.
func (h *HeightField) Offset(x, y float64) (float64, float64) {
	x = x + (h.widthFloat-1)/2
	y = y + (h.heightFloat-1)/2
	return x, y
}

func (h *HeightField) Valid(x, y float64) bool {
	x, y = h.Offset(x, y)
	if x < 0.0 || x > h.widthFloat {
		return false
	}
	if y < 0.0 || y > h.heightFloat {
		return false
	}
	return true
}

func (h *HeightField) Wrap(x, y float64) (float64, float64) {
	return x, y
}
