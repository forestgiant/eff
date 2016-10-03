package main

import (
	"github.com/forestgiant/eff/eff"
	"github.com/forestgiant/eff/sdl"
)

type textDrawable struct {
	initialized bool
	font        eff.Font
}

func (t *textDrawable) Init(canvas eff.Canvas) {
	t.font.Path = "../assets/vcr_osd_mono.ttf"
}

func (t *textDrawable) Draw(canvas eff.Canvas) {
	canvas.DrawText("hello, world!", 24, t.font, eff.RandomColor(), Point{X: 0, Y: 0})
}

func (t *textDrawable) Update(canvas eff.Canvas) {

}

func (t *textDrawable) Initialized() bool {
	return t.initialized
}

func main() {
	t := textDrawable{}
	canvas := sdl.NewCanvas("Animating Text", 800, 540, 60, true)

	canvas.AddDrawable(t)
	canvas.Run()
}
