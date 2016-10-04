package main

import (
	"github.com/forestgiant/eff/eff"
	"github.com/forestgiant/eff/sdl"
)

type textDrawable struct {
	initialized bool
	t           float64
	text        string
	textColor   eff.Color
}

func (t *textDrawable) Init(canvas eff.Canvas) {
	font := eff.Font{
		Path: "../assets/fonts/Jellee-Roman.ttf",
	}

	canvas.SetFont(font, 24)

	t.text = "Effulgent, Effulgent, Effulgent, Effulgent, Effulgent, Effulgent"
	t.initialized = true
}

func (t *textDrawable) Draw(canvas eff.Canvas) {
	var index int
	index = int(t.t*float64(len(t.text))) + 1
	textColor := eff.RandomColor()

	canvas.DrawText(t.text[:index], textColor, eff.Point{X: 0, Y: 0})
}

func (t *textDrawable) Update(canvas eff.Canvas) {
	t.t += 0.005
	if t.t > 1 {
		t.t = 0
	}
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
