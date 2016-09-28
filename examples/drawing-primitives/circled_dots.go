package main

import (
	"math"
	"math/rand"

	"github.com/forestgiant/eff/eff"
)

type twoPositionDot struct {
	eff.Color
	p1 eff.Point
	p2 eff.Point
}

func (dot *twoPositionDot) linearInterpolate(normalizedPercentage float64) eff.Point {
	xDiff := dot.p2.X - dot.p1.X
	yDiff := dot.p2.Y - dot.p1.Y

	return eff.Point{
		X: dot.p1.X + int(math.Ceil((float64(xDiff) * normalizedPercentage))),
		Y: dot.p1.Y + int(math.Ceil((float64(yDiff) * normalizedPercentage))),
	}
}

type circleDots struct {
	t           float64
	tDir        float64
	dots        []twoPositionDot
	initialized bool
}

func (dot *circleDots) Init(canvas eff.Canvas) {
	dotCount := (canvas.Width() * canvas.Height()) / 100

	pointOnCirlce := func(radius int, index int, totalPoints int, w int, h int) eff.Point {
		cx := w / 2
		cy := h / 2
		angle := math.Pi * 2 * (float64(index) / float64(totalPoints))
		return eff.Point{
			X: cx + int(float64(radius)*math.Cos(angle)),
			Y: cy + int(float64(radius)*math.Sin(angle)),
		}
	}

	for i := 0; i < dotCount; i++ {
		d := twoPositionDot{
			Color: eff.RandomColor(),
			p1: eff.Point{
				X: rand.Intn(canvas.Width()),
				Y: rand.Intn(canvas.Height()),
			},
			p2: pointOnCirlce(100, i, dotCount, canvas.Width(), canvas.Height()),
		}

		dot.dots = append(dot.dots, d)
	}

	dot.tDir = 1
	dot.initialized = true
}

func (dot *circleDots) Draw(canvas eff.Canvas) {
	var colorPoints []eff.ColorPoint
	for _, d := range dot.dots {
		point := d.linearInterpolate(dot.t)
		colorPoints = append(colorPoints,
			eff.ColorPoint{
				Point: point,
				Color: d.Color,
			},
		)
	}

	canvas.DrawColorPoints(colorPoints)
}

func (dot *circleDots) Update(canvas eff.Canvas) {
	diff := 0.005 * dot.tDir
	dot.t += diff

	if dot.t > 1 || dot.t < 0 {
		dot.tDir *= -1
	}
}

func (dot *circleDots) Initialized() bool {
	return dot.initialized
}
