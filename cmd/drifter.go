package main

import (
	"github.com/fogleman/gg"
	"github.com/glesica/drifter/internal/art"
	"github.com/glesica/drifter/internal/sim"
	terrain2 "github.com/glesica/drifter/internal/terrain"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(2342)

	field := terrain2.NewFuncMap()
	field.AccelFunc = func(x, y float64) (float64, float64) {
		ax := 1 + 1/(x+1)
		ay := 1 + 1/(y+1)
		return ax, ay
	}

	traces := terrain2.RandLayout(1, 500, 500)

	driver := sim.NewDriver(field, traces)
	for i := 0; i < 10000; i++ {
		driver.Update(0.01)
	}

	canvas := gg.NewContext(500, 500)
	drawer := art.LineDrawer(canvas)
	driver.Render(drawer)

	outFile, err := os.Create("output.png")
	if err != nil {
		panic("failed to open output file")
	}

	err = canvas.EncodePNG(outFile)
	if err != nil {
		panic("failed to encode png")
	}

	err = outFile.Close()
	if err != nil {
		panic("failed to close output file")
	}
}
