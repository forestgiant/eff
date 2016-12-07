package main

import (
	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type myShape struct {
	eff.Shape
}

const (
	windowW         = 800
	windowH         = 600
	parentSize      = 300
	biggerChildSize = 500
)

func (m *myShape) init() {
	m.SetRect(eff.Rect{
		X: (windowW - parentSize) / 2,
		Y: (windowH - parentSize) / 2,
		W: parentSize,
		H: parentSize,
	})
	m.SetBackgroundColor(eff.Black())

	biggerChild := &eff.Shape{}
	biggerChild.SetRect(eff.Rect{
		X: (parentSize - biggerChildSize) / 2,
		Y: (parentSize - biggerChildSize) / 2,
		W: biggerChildSize,
		H: biggerChildSize,
	})
	biggerChild.SetBackgroundColor(eff.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0x66})
	m.AddChild(biggerChild)
}

func main() {
	canvas := sdl.NewCanvas("Clipping", windowW, windowH, eff.White(), 60, true)
	canvas.Run(func() {
		m := &myShape{}
		m.init()
		canvas.AddChild(m)
	})
}
