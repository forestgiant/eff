package drawables

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type Dots struct {
	points      []eff.Point
	initialized bool
}

func (d *Dots) randomPoints(count int, maxX int, maxY int) []eff.Point {
	var points []eff.Point
	for i := 0; i < count; i++ {
		points = append(points, eff.Point{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		})
	}
	return points
}

func (d *Dots) Init(canvas eff.Canvas) {
	numDots := (canvas.Width() * canvas.Height()) / 20
	d.points = d.randomPoints(numDots, canvas.Width(), canvas.Height())
	d.initialized = true
}

func (d *Dots) Draw(canvas eff.Canvas) {
	//Draw Points in a random color
	canvas.DrawPoints(d.points, eff.RandomColor())
}

func (d *Dots) Update(canvas eff.Canvas) {
	updateRandomPoints := func() {
		for i := range d.points {
			d.points[i].X = rand.Intn(canvas.Width())
			d.points[i].Y = rand.Intn(canvas.Height())
		}
	}

	updateRandomPoints()
}

func (d *Dots) Initialized() bool {
	return d.initialized
}
