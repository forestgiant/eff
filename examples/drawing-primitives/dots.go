package main

import (
	"math/rand"

	"github.com/forestgiant/eff/eff"
)

type dots struct {
	points []eff.Point
}

func (d *dots) randomPoints(count int, maxX int, maxY int) []eff.Point {
	var points []eff.Point
	for i := 0; i < count; i++ {
		points = append(points, eff.Point{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		})
	}
	return points
}

func (d *dots) Init(canvas eff.Canvas) {
	d.points = d.randomPoints(10000, canvas.Width(), canvas.Height())
}

func (d *dots) Draw(canvas eff.Canvas) {
	//Draw Points in a random color
	canvas.DrawPoints(d.points, eff.RandomColor())
}

func (d *dots) Update(canvas eff.Canvas) {
	updateRandomPoints := func() {
		for i := range d.points {
			d.points[i].X = rand.Intn(canvas.Width())
			d.points[i].Y = rand.Intn(canvas.Height())
		}
	}

	updateRandomPoints()
}
