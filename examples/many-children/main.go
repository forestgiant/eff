package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW = 1024
	windowH = 768
	cols    = 128
	rows    = 128
)

func main() {
	canvas := sdl.NewCanvas("Many Children", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < cols*rows; i++ {
			s := &eff.Shape{}
			s.SetRect(eff.Rect{
				X: (i % cols) * (windowW / cols),
				Y: (i / cols) * (windowH / rows),
				W: windowW / cols,
				H: windowH / rows,
			})
			s.SetBackgroundColor(eff.RandomColor())
			canvas.AddChild(s)
		}
	})
}
