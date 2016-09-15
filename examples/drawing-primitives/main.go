package main

import (
	"math/rand"
	"os"

	"github.com/forestgiant/eff"
)

type dots struct {
	points []eff.Point
}

func (d *dots) randomPoints(count int, maxX int, maxY int) *[]eff.Point {
	points := make([]eff.Point, count)
	for i := 0; i < count; i++ {
		points[i] = eff.Point{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		}
	}
	return &points
}

func (d *dots) Init(canvas eff.Canvas) {
	d.points = *d.randomPoints(10000, canvas.Width(), canvas.Height())
}

func (d *dots) Draw(canvas eff.Canvas) {
	//Draw Points in a random color
	canvas.DrawPoints(&d.points, eff.Color{R: rand.Intn(255), G: rand.Intn(255), B: rand.Intn(255), A: rand.Intn(255)})
}

func (d *dots) Update(canvas eff.Canvas) {
	updateRandomPoints := func() {
		for i, _ := range d.points {
			d.points[i].X = rand.Intn(canvas.Width())
			d.points[i].Y = rand.Intn(canvas.Height())
		}
	}

	updateRandomPoints()
}

func main() {
	//Create drawables
	d := dots{}

	//Create Eff Canvas
	canvas := eff.SDLCanvas{}

	//Add drawables to canvas
	canvas.AddDrawable(&d)

	//Start the run loop
	os.Exit(canvas.Run())
}
