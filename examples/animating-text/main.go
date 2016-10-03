package main

import (
	"github.com/forestgiant/eff/eff"
	"github.com/forestgiant/eff/sdl"
)

type textDrawable struct {
	initialized bool
}

func (t *textDrawable) Init(canvas eff.Canvas) {
	font := eff.Font{
		Path: "../assets/vcr_osd_mono.ttf",
	}

	canvas.SetFont(font, 24)
	t.initialized = true
}

func (t *textDrawable) Draw(canvas eff.Canvas) {
	textColor := eff.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	canvas.DrawText("hello, world!", textColor, eff.Point{X: 0, Y: 0})
}

func (t *textDrawable) Update(canvas eff.Canvas) {

}

func (t *textDrawable) Initialized() bool {
	return t.initialized
}

func main() {
	t := textDrawable{}
	canvas := sdl.NewCanvas("Animating Text", 800, 540, 60, true)

	canvas.AddDrawable(&t)
	canvas.Run()
}
