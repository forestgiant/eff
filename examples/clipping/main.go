package main

import (
	"fmt"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW         = 800
	windowH         = 600
	parentSize      = 300
	biggerChildSize = 500
)

type myShape struct {
	eff.Shape
	biggerChild *eff.Shape
}

func (m *myShape) init() {
	m.SetRect(eff.Rect{
		X: (windowW - parentSize) / 2,
		Y: (windowH - parentSize) / 2,
		W: parentSize,
		H: parentSize,
	})
	m.SetBackgroundColor(eff.Black())

	m.biggerChild = &eff.Shape{}
	m.biggerChild.SetRect(eff.Rect{
		X: (parentSize - biggerChildSize) / 2,
		Y: (parentSize - windowH) / 2,
		W: biggerChildSize,
		H: windowH,
	})
	m.biggerChild.SetBackgroundColor(eff.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0x66})
	m.AddChild(m.biggerChild)

	dot := &eff.Shape{}
	dot.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: 5,
		H: 5,
	})
	dot.SetBackgroundColor(eff.Black())
	m.biggerChild.AddChild(dot)
	vec := eff.Point{X: 10, Y: 10}
	dot.SetUpdateHandler(func() {
		x := dot.Rect().X + vec.X
		y := dot.Rect().Y + vec.Y
		if x <= 0 || x >= (dot.Parent().Rect().W-dot.Rect().W) {
			vec.X *= -1
		}
		if y <= 0 || y >= (dot.Parent().Rect().H-dot.Rect().H) {
			vec.Y *= -1
		}

		dot.SetRect(eff.Rect{
			X: x,
			Y: y,
			W: dot.Rect().W,
			H: dot.Rect().H,
		})
	})
}

func main() {
	canvas := sdl.NewCanvas("Clipping", windowW, windowH, eff.White(), 60, true)
	canvas.Run(func() {
		fmt.Println("Press space to toggle between clip and no clip")
		m := &myShape{}
		m.init()
		canvas.AddChild(m)
		canvas.AddKeyUpEnumHandler(func(keyCode sdl.Keycode) {
			if keyCode == sdl.KeySpace {
				m.biggerChild.SetShouldClip(!m.biggerChild.ShouldClip())
				if m.biggerChild.ShouldClip() {
					fmt.Println("Clipping")
				} else {
					fmt.Println("No Clip")
				}
			}
		})
	})
}
