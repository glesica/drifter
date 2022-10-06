package art

import (
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

var AllColors = []color.Color{
	colornames.Darkcyan,
	colornames.Aquamarine,
	colornames.Cadetblue,
	colornames.Darkgoldenrod,
	colornames.Darkmagenta,
	colornames.Darkorchid,
	colornames.Darkorange,
	colornames.Gainsboro,
	colornames.Crimson,
	colornames.Lawngreen,
	colornames.Firebrick,
	colornames.Lightseagreen,
	colornames.Lightcoral,
	colornames.Lightgreen,
	colornames.Lightcyan,
	colornames.Lightskyblue,
	colornames.Lightyellow,
}

type ColorPicker func(x, y, vx, vy float64) color.Color

func VelocityColorPicker(slow, fast colorful.Color, minSpeed, maxSpeed float64) ColorPicker {
	return func(x, y, vx, vy float64) color.Color {
		speed := math.Sqrt(vx*vx + vy*vy)
		scale := (speed - minSpeed) / (maxSpeed - minSpeed)
		return slow.BlendRgb(fast, scale)
	}
}
