package main

import (
	"fmt"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type textDrawable struct {
	initialized bool
	t           float64
	text        string
}

func (t *textDrawable) Init(canvas eff.Canvas) {
	font := eff.Font{
		Path: "../assets/fonts/Jellee-Roman.ttf",
	}

	err := canvas.SetFont(font, 24)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t.text = "Effulgent, Effulgent, Effulgent, Effulgent, Effulgent, Effulgent"
	t.initialized = true
}

func (t *textDrawable) Draw(canvas eff.Canvas) {
	var index int
	index = int(t.t*float64(len(t.text))) + 1
	textColor := eff.RandomColor()

	err := canvas.DrawText(t.text[:index], textColor, eff.Point{X: 0, Y: 0})

	if err != nil {
		fmt.Println(err)
	}
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
	canvas := sdl.NewCanvas("Animating Text", 800, 540, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.AddDrawable(&textDrawable{})
	})
}
