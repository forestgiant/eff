package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type myShape struct{}

func (m *myShape) init(canvas eff.Canvas) {
	childCount := 0
	addChild := func(parent *eff.Shape) *eff.Shape {
		child := &eff.Shape{}
		child.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: int(float64(parent.Rect().W) / 1.1),
			H: int(float64(parent.Rect().H) / 1.1),
		})
		color := eff.White()
		if childCount%2 == 0 {
			color = eff.Black()
		}
		color = eff.RandomColor()
		child.SetBackgroundColor(color)
		parent.AddChild(child)
		vec := eff.Point{X: 1, Y: 1}
		child.SetUpdateHandler(func() {
			x := child.Rect().X + vec.X
			y := child.Rect().Y + vec.Y
			if x <= 0 || x >= (child.Parent().Rect().W-child.Rect().W) {
				vec.X *= -1
			}
			if y <= 0 || y >= (child.Parent().Rect().H-child.Rect().H) {
				vec.Y *= -1
			}

			child.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: child.Rect().W,
				H: child.Rect().H,
			})
		})

		childCount++

		return child
	}

	p := &eff.Shape{}
	p.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: 800,
		H: 600,
	})
	p.SetBackgroundColor(eff.White())
	canvas.AddChild(p)
	for i := 0; i < 50; i++ {
		p = addChild(p)
	}

}

func main() {
	canvas := sdl.NewCanvas("Children", 800, 600, eff.White(), 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		m := myShape{}
		m.init(canvas)
	})
}
