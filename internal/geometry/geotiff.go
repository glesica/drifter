package geometry

import (
	"image"
	"image/color"
)

func ReadGeoTiff(img image.Image, dem *DEM) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.At(x, y).(color.Gray16)
			height := float64(pixel.Y)
			dem.Set(x, y, height)
		}
	}
}
