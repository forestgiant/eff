package main

import (
	"log"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/button"
	"github.com/forestgiant/eff/sdl"
)

type buttonTest struct {
	eff.Shape

	buttons    []*button.Button
	middleText string
	font       eff.Font
}

func (b *buttonTest) Init(c eff.Canvas) {

	font, err := c.OpenFont("../assets/fonts/roboto/Roboto-Medium.ttf", 15)
	if err != nil {
		log.Fatal(err)
	}

	b.font = font
	drawMiddleText := func() {
		textW, textH, err := b.Graphics().GetTextSize(b.font, b.middleText)
		if err != nil {
			log.Fatal(err)
		}
		textPoint := eff.Point{
			X: (b.Rect().W - textW) / 2,
			Y: (b.Rect().H - textH) / 2,
		}
		b.DrawText(b.font, b.middleText, eff.Black(), textPoint)
	}

	clickHandler := func(button *button.Button) {
		b.Clear()
		b.middleText = button.Text
		drawMiddleText()
	}

	padding := 20
	buttonWidth := 100
	buttonHeight := 30
	topLeftButton := button.NewButton(b.font, "NW", eff.Rect{X: padding, Y: padding, W: buttonWidth, H: buttonHeight}, clickHandler)
	b.buttons = append(b.buttons, topLeftButton)
	c.AddClickable(topLeftButton)
	b.AddChild(topLeftButton)

	bottomLeftButton := button.NewButton(b.font, "SW", eff.Rect{X: padding, Y: c.Rect().H - padding - buttonHeight, W: buttonWidth, H: buttonHeight}, clickHandler)
	b.buttons = append(b.buttons, bottomLeftButton)
	c.AddClickable(bottomLeftButton)
	b.AddChild(bottomLeftButton)

	topRightButton := button.NewButton(b.font, "NE", eff.Rect{X: c.Rect().W - padding - buttonWidth, Y: padding, W: buttonWidth, H: buttonHeight}, clickHandler)
	b.buttons = append(b.buttons, topRightButton)
	c.AddClickable(topRightButton)
	b.AddChild(topRightButton)

	bottomRightButton := button.NewButton(b.font, "SE", eff.Rect{X: c.Rect().W - padding - buttonWidth, Y: c.Rect().H - padding - buttonHeight, W: buttonWidth, H: buttonHeight}, clickHandler)
	b.buttons = append(b.buttons, bottomRightButton)
	c.AddClickable(bottomRightButton)
	b.AddChild(bottomRightButton)

	b.middleText = "NW"
	drawMiddleText()
}

func main() {
	canvas := sdl.NewCanvas("Clickables", 800, 540, eff.White(), 1000, false)

	canvas.Run(func() {
		bt := buttonTest{}
		bt.SetRect(eff.Rect{X: 0, Y: 0, W: 800, H: 540})
		canvas.AddChild(&bt)
		bt.Init(canvas)
	})
}
