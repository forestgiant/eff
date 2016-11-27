package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
	"github.com/forestgiant/eff/sdl"
)

type textDrawable struct {
	initialized bool
	text        string
	tweener     tween.Tweener
	index       int
	font        eff.Font
}

func (t *textDrawable) Init(canvas eff.Canvas) {
	font, err := canvas.OpenFont("../assets/fonts/Jellee-Roman.ttf", 24)
	if err != nil {
		log.Fatal(err)
	}
	t.font = font

	t.index = 1
	t.text = "Effulgent, Effulgent, Effulgent, Effulgent, Effulgent, Effulgent"
	t.tweener = tween.NewTweener(time.Second*5, func(progress float64) {
		t.index = int(progress * float64(len(t.text)))
		t.index = int(math.Max(1, float64(t.index)))
		t.index = int(math.Min(float64(len(t.text)), float64(t.index)))
	}, true, false, nil, nil)
	t.initialized = true
}

func (t *textDrawable) Draw(canvas eff.Canvas) {
	textColor := eff.RandomColor()
	err := canvas.DrawText(t.font, t.text[:t.index], textColor, eff.Point{X: 0, Y: 0})
	if err != nil {
		fmt.Println(err)
	}
}

func (t *textDrawable) Update(canvas eff.Canvas) {

	t.tweener.Tween()
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
