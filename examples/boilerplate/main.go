package main

import (
	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/examples/boilerplate/drawable"
	"github.com/forestgiant/eff/sdl"
)

func main() {
	canvas := sdl.NewCanvas("Boilerplate", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		d := drawable.MyDrawable{}
		canvas.AddDrawable(&d)
	})
}
