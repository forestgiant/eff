package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type myShape struct {
	eff.Shape
}

func (m *myShape) init() {
	makeChild := func() *eff.Shape {
		parent := &eff.Shape{}
		parent.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: 60,
			H: 60,
		})
		parent.SetBackgroundColor(eff.Black())
		pVec := eff.Point{X: rand.Intn(9) + 1, Y: rand.Intn(9) + 1}

		parent.SetUpdateHandler(func() {
			x := parent.Rect().X + pVec.X
			y := parent.Rect().Y + pVec.Y

			if x <= 0 || x >= (parent.Parent().Rect().W-parent.Rect().W) {
				pVec.X *= -1
			}
			if y <= 0 || y >= (parent.Parent().Rect().H-parent.Rect().H) {
				pVec.Y *= -1
			}

			parent.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: parent.Rect().W,
				H: parent.Rect().H,
			})
		})

		child := &eff.Shape{}
		child.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: 5,
			H: 5,
		})
		child.SetBackgroundColor(eff.White())
		parent.AddChild(child)
		cVec := eff.Point{X: rand.Intn(4) + 1, Y: rand.Intn(4) + 1}
		child.SetUpdateHandler(func() {
			x := child.Rect().X + cVec.X
			y := child.Rect().Y + cVec.Y
			if x <= 0 || x >= (child.Parent().Rect().W-child.Rect().W) {
				cVec.X *= -1
			}
			if y <= 0 || y >= (child.Parent().Rect().H-child.Rect().H) {
				cVec.Y *= -1
			}

			child.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: child.Rect().W,
				H: child.Rect().H,
			})
		})

		return parent
	}
	m.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: 800,
		H: 600,
	})

	for i := 0; i < 20; i++ {
		m.AddChild(makeChild())
	}

}

func main() {
	canvas := sdl.NewCanvas("Children", 800, 600, eff.White(), 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		m := myShape{}
		m.init()
		canvas.AddChild(&m)
	})
}
