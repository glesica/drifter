package main

import (
	"bufio"
	"flag"
	"github.com/chai2010/tiff"
	"github.com/fogleman/gg"
	"github.com/glesica/drifter"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"os"
)

const screenWidth = 320
const screenHeight = 320

const viewportWidth = 40
const viewportHeight = 40

func random1() float64 {
	return (rand.Float64() - 0.5) * 2
}

func perturbed(value float64) float64 {
	return value + (rand.Float64()-0.5)*value*0.2
}

func main() {
	//file, err := os.Open("test/ASTGTMV003_N45W115_dem.tif")
	//if err != nil {
	//	panic(err.Error())
	//}

	// TODO: parse the offsets from the file name
	//latOffset := float64(50)
	//lngOffset := float64(28)

	//outImg := gg.NewContext(3601, 3601)
	//outImg.SetColor(color.Black)
	//outImg.DrawRectangle(0, 0, 3601, 3601)
	//outImg.Fill()
	//
	//img, _ := tiff.Decode(bufio.NewReader(file))
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//bounds := img.Bounds()
	//fmt.Printf("(%d, %d) to (%d, %d)\n", bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)

	//upperLeftLat := 46.0001389
	//upperLeftLon := -115.0001389
	//
	//lowerRightLat := 44.9998611
	//lowerRightLon := -113.9998611
	//
	//trapperLat := 45.88979
	//trapperLon := -114.29758

	//minHeight := int16(math.MaxInt16)
	//zeroCount := 0
	//maxHeight := int16(0)
	//for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
	//	for x := bounds.Min.X; x < bounds.Max.X; x++ {
	//		// Normalize X and Y to a 0 to 1 space based on the size of the image.
	//		// Add the offsets to get coordinates.
	//		//lng := lngOffset + (float64(x) / float64(bounds.Max.X))
	//		//lat := latOffset + (float64(y) / float64(bounds.Max.Y))
	//
	//		height := int16(img.At(x, y).(color.Gray16).Y)
	//
	//		if height < minHeight {
	//			minHeight = height
	//		}
	//		if height > maxHeight {
	//			maxHeight = height
	//		}
	//
	//		if height < 927 {
	//			zeroCount++
	//			outImg.SetColor(colornames.Red)
	//			outImg.DrawPoint(float64(x), float64(y), 1)
	//			outImg.Fill()
	//		}
	//	}
	//}
	//
	//fmt.Printf("[%d, %d]\n", minHeight, maxHeight)
	//fmt.Printf("found %d zero-ish values\n", zeroCount)
	//fmt.Printf("channel count %d\n", tiff.ChannelsOf(img))
	//
	//outImg.SavePNG("test.png")
	//
	//return

	duration := flag.Float64("duration", defaultDuration, "simulation duration")
	delta := flag.Float64("delta", defaultDelta, "simulation time step")
	pngPath := flag.String("png", defaultPNGPath, "output PNG path")
	//jsonPath := flag.String("json", defaultJSONPath, "output JSON path")
	seed := flag.Int64("seed", defaultSeed, "default random seed")

	flag.Parse()

	//rand.Seed(98723)
	rand.Seed(*seed)

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Drifter Demo")

	file, err := os.Open("test/ASTGTMV003_N46W114_dem.tif")
	if err != nil {
		panic(err.Error())
	}

	outImg := gg.NewContext(3601, 3601)
	outImg.SetColor(color.Black)
	outImg.DrawRectangle(0, 0, 3601, 3601)
	outImg.Fill()

	img, _ := tiff.Decode(bufio.NewReader(file))
	if err != nil {
		panic(err.Error())
	}

	mapRegion := &drifter.Region{
		LeftLon:   -115.0001389,
		RightLon:  -113.9998611,
		TopLat:    47.0001389,
		BottomLat: 45.9998611,
	}
	simRegion := &drifter.Region{
		LeftLon:   -114.26919,
		RightLon:  -114.22487,
		TopLat:    46.68690,
		BottomLat: 46.66382,
	}
	mapViewport := drifter.MakeViewport(img.Bounds(), mapRegion, simRegion)
	field := drifter.NewGeotiffField(img, mapViewport)

	//field := drifter.NewHeightField([]float64{
	//	1, 1, 1, 1, 1,
	//	1, 2, 2, 2, 1,
	//	1, 2, 3, 2, 1,
	//	1, 2, 2, 2, 1,
	//	1, 1, 1, 1, 1,
	//}, 5)

	//field := drifter.NewFuncField(drifter.Valley)
	//field.DampFunc = func(x, y float64) float64 {
	//	return 0.1
	//}

	sim := drifter.NewSim(
		field,
	)

	simViewport := drifter.NewViewportWH(mapViewport.Width(), mapViewport.Height())
	//traces := drifter.RingLayout(25, 0.25)
	//traces := drifter.RandomLayout(100, simViewport)
	traces := drifter.GridLayout(5, 5, simViewport)
	for _, t := range traces {
		sim.AddTrace(t)
	}

	sim.AdvanceTo(*delta, *duration)

	outFile, err := os.Create(*pngPath)
	if err != nil {
		panic("failed to open output path")
	}

	history := sim.History()
	maxSpeed, minSpeed := history.SpeedBounds()
	drawer := drifter.NewSimplerDrawer(drifter.NewVelocityColorPicker(colornames.Green, colornames.Red, minSpeed, maxSpeed))

	renderer := drifter.NewPNGRenderer()
	renderer.SetData(sim.History())
	renderer.SetDrawer(drawer)
	renderer.SetViewport(simViewport)
	err = renderer.WriteFinal(outFile, 1000, 1000)
	if err != nil {
		panic("failed to write output file")
	}

	//sim.AddTrace(&drifter.Trace{
	//	ID: 0,
	//	X:  perturbed(10.0),
	//	Y:  perturbed(10.0),
	//	VX: random1(),
	//	VY: random1(),
	//})
	//sim.AddTrace(&drifter.Trace{
	//	ID: 1,
	//	X:  perturbed(-10.0),
	//	Y:  perturbed(10.0),
	//	VX: random1(),
	//	VY: random1(),
	//})
	//sim.AddTrace(&drifter.Trace{
	//	ID: 2,
	//	X:  perturbed(-10.0),
	//	Y:  perturbed(-10.0),
	//	VX: random1(),
	//	VY: random1(),
	//})
	//sim.AddTrace(&drifter.Trace{
	//	ID: 3,
	//	X:  perturbed(10.0),
	//	Y:  perturbed(-10.0),
	//	VX: random1(),
	//	VY: random1(),
	//})

	//renderer := drifter.NewEbitenRenderer()
	//renderer.SetData(sim.History())
	//minX := -viewportWidth / 2
	//minY := -viewportHeight / 2
	//maxX := viewportWidth / 2
	//maxY := viewportHeight / 2
	//renderer.SetViewport(drifter.NewViewport(minX, minY, maxX, maxY))
	//if err := ebiten.RunGame(renderer); err != nil {
	//	log.Fatal(err)
	//}
}
