package main

import (
	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/examples/text-view/shapes"
	"github.com/forestgiant/eff/sdl"
)

func main() {
	canvas := sdl.NewCanvas("Text View", 800, 540, eff.White(), 60, true)
	canvas.Run(func() {
		s := &shapes.TextViewer{}
		s.SetRect(canvas.Rect())
		canvas.AddChild(s)
		s.Init(canvas)
	})
}
