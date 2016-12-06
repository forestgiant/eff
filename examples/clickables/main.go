package main

import (
	"log"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/button"
	"github.com/forestgiant/eff/sdl"
	"github.com/forestgiant/eff/util"
)

type buttonTest struct {
	eff.Shape

	buttons    []*button.Button
	middleText string
	font       eff.Font
}

func (b *buttonTest) Init(c eff.Canvas) {
	font, err := c.OpenFont("../assets/fonts/roboto/Roboto-Bold.ttf", 15)
	if err != nil {
		log.Fatal(err)
	}

	b.font = font

	clickHandler := func(button *button.Button) {
		b.Clear()
		b.middleText = button.Text
		textPoint, err := util.CenterTextInRect(b.font, b.middleText, b.Rect(), b.Graphics())
		if err != nil {
			log.Fatal(err)
		}

		b.DrawText(b.font, b.middleText, eff.Black(), textPoint)
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
}

func main() {
	canvas := sdl.NewCanvas("Clickables", 800, 540, eff.White(), 60, true)

	canvas.Run(func() {
		bt := buttonTest{}
		canvas.AddChild(&bt)
		bt.Init(canvas)
	})
}
