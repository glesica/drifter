package main

import (
	"github.com/glesica/drifter"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

const screenWidth = 320
const screenHeight = 320

const viewportWidth = 40
const viewportHeight = 40

func random10() float64 {
	return (rand.Float64() - 0.5) * 20
}

func perturbed(value float64) float64 {
	return value + (rand.Float64()-0.5)*value*0.2
}

func main() {
	rand.Seed(98723)

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Drifter Demo")

	field := drifter.NewField(drifter.Valley)

	sim := drifter.NewSim(
		field,
	)
	sim.AddTrace(&drifter.Trace{
		ID: 0,
		X:  perturbed(10.0),
		Y:  perturbed(10.0),
		VX: random10(),
		VY: random10(),
	})
	sim.AddTrace(&drifter.Trace{
		ID: 1,
		X:  perturbed(-10.0),
		Y:  perturbed(10.0),
		VX: random10(),
		VY: random10(),
	})
	sim.AddTrace(&drifter.Trace{
		ID: 2,
		X:  perturbed(-10.0),
		Y:  perturbed(-10.0),
		VX: random10(),
		VY: random10(),
	})
	sim.AddTrace(&drifter.Trace{
		ID: 3,
		X:  perturbed(10.0),
		Y:  perturbed(-10.0),
		VX: random10(),
		VY: random10(),
	})

	renderer := drifter.NewRenderer(sim)
	//renderer.DrawField = true

	minX := -viewportWidth / 2
	minY := -viewportHeight / 2
	maxX := viewportWidth / 2
	maxY := viewportHeight / 2
	renderer.Viewport().Resize(minX, minY, maxX, maxY)

	if err := ebiten.RunGame(renderer); err != nil {
		log.Fatal(err)
	}
}
