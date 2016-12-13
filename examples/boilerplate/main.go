package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW    = 800
	windowH    = 540
	squareSize = 100
)

type myShape struct {
	eff.Shape
}

func main() {
	canvas := sdl.NewCanvas("Boilerplate", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		m := &myShape{}
		m.SetRect(eff.Rect{
			X: (windowW - squareSize) / 2,
			Y: (windowH - squareSize) / 2,
			W: squareSize,
			H: squareSize,
		})
		minSpeed := 3
		maxSpeed := 10
		vec := eff.Point{X: rand.Intn(maxSpeed-minSpeed) + minSpeed, Y: rand.Intn(maxSpeed-minSpeed) + minSpeed}
		m.SetUpdateHandler(func() {
			x := m.Rect().X + vec.X
			y := m.Rect().Y + vec.Y
			if x <= 0 || x >= (canvas.Rect().W-m.Rect().W) {
				vec.X *= -1
			}

			if y <= 0 || y >= (canvas.Rect().H-m.Rect().H) {
				vec.Y *= -1
			}

			m.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: m.Rect().W,
				H: m.Rect().H,
			})
		})
		m.SetBackgroundColor(eff.RandomColor())
		canvas.AddChild(m)
	})
}
