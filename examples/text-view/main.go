package main

import (
	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/examples/text-view/drawable"
	"github.com/forestgiant/eff/sdl"
)

func main() {
	canvas := sdl.NewCanvas("Text View", 800, 540, eff.White(), 60, true)
	canvas.Run(func() {
		d := drawable.TextViewer{}
		canvas.AddDrawable(&d)
	})
}
