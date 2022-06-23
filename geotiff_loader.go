package drifter

import (
	"image"
	"image/color"
	"math"
)

func NewGeotiffField(img image.Image, viewport *Viewport) *HeightField {
	//img, err := tiff.Decode(bufio.NewReader(file))
	//if err != nil {
	//	panic(err.Error())
	//}

	width := viewport.Width()
	height := viewport.Height()

	heights := make([]float64, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.At(x, y).(color.Gray16)
			height := float64(pixel.Y)
			heights[y*width+x] = height
		}
	}

	return NewHeightField(heights, width)
}

func MakeViewport(bounds image.Rectangle, mapRegion *Region, simRegion *Region) *Viewport {
	xPixelRatio := float64(bounds.Dx()) / mapRegion.Width()
	yPixelRatio := float64(bounds.Dy()) / mapRegion.Height()

	leftX := int(simRegion.LeftLon * xPixelRatio)
	rightX := int(math.Round(simRegion.RightLon * xPixelRatio))
	topY := int(math.Round(simRegion.TopLat * yPixelRatio))
	bottomY := int(simRegion.BottomLat * yPixelRatio)

	return NewViewport(leftX, bottomY, rightX, topY)
}

type Region struct {
	LeftLon   float64
	RightLon  float64
	TopLat    float64
	BottomLat float64
}

func (r *Region) Width() float64 {
	return r.RightLon - r.LeftLon
}

func (r *Region) Height() float64 {
	return r.TopLat - r.BottomLat
}
