package shapes

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type Dots struct {
	eff.Shape

	points []eff.Point
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

func (d *Dots) Init(width int, height int) {
	numDots := (width * height) / 20
	d.points = d.randomPoints(numDots, width, height)
	d.SetUpdateHandler(func() {
		for i := range d.points {
			d.points[i].X = rand.Intn(width)
			d.points[i].Y = rand.Intn(height)
		}
		d.Clear()
		d.DrawPoints(d.points, eff.RandomColor())
	})
}
