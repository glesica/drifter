package drifter

import (
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type ColorPicker interface {
	TraceColor(t *TraceFrame) color.Color
}

var allColors = []color.Color{
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

type IDColorPicker struct{}

func (i *IDColorPicker) TraceColor(t *TraceFrame) color.Color {
	return allColors[t.ID%len(allColors)]
}

type VelocityColorPicker struct {
	slowColor colorful.Color
	fastColor colorful.Color
	minSpeed  float64
	maxSpeed  float64
	rangeSize float64
}

func NewVelocityColorPicker(slowColor, fastColor color.Color, minSpeed, maxSpeed float64) *VelocityColorPicker {
	slow, _ := colorful.MakeColor(slowColor)
	fast, _ := colorful.MakeColor(fastColor)
	return &VelocityColorPicker{
		slowColor: slow,
		fastColor: fast,
		minSpeed:  minSpeed,
		maxSpeed:  maxSpeed,
		rangeSize: maxSpeed - minSpeed,
	}
}

func (v *VelocityColorPicker) TraceColor(t *TraceFrame) color.Color {
	speed := math.Sqrt(t.VX*t.VX + t.VY*t.VY)
	scale := (speed - v.minSpeed) / v.rangeSize
	return v.slowColor.BlendRgb(v.fastColor, scale)
}
