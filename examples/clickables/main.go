package main

import (
	"fmt"
	"log"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/button"
	"github.com/forestgiant/eff/sdl"
)

type buttonTest struct {
	initialized bool
	buttons     []*button.Button
	middleText  string
	font        eff.Font
}

func (b *buttonTest) Init(c eff.Canvas) {
	font, err := c.OpenFont("../assets/fonts/roboto/Roboto-Bold.ttf", 15)
	if err != nil {
		log.Fatal(err)
	}

	b.font = font

	clickHandler := func(button *button.Button) {
		b.middleText = button.Text
	}

	drawButton := func(text string, rect eff.Rect, bgColor eff.Color, textColor eff.Color, c eff.Canvas) {
		tW, tH, err := c.GetTextSize(b.font, text)
		if err != nil {
			log.Fatal(err)
		}

		textPoint := eff.Point{
			X: rect.X + ((rect.W - tW) / 2),
			Y: rect.Y + ((rect.H - tH) / 2),
		}
		c.FillRect(rect, bgColor)
		c.DrawText(b.font, text, textColor, textPoint)
	}

	drawDefault := func(button *button.Button, c eff.Canvas) {
		bgColor := eff.Color{R: 0xDD, G: 0xDD, B: 0xDD, A: 0xFF}
		textColor := eff.Black()
		drawButton(button.Text, button.Rect, bgColor, textColor, c)
	}

	drawDown := func(button *button.Button, c eff.Canvas) {
		bgColor := eff.Color{R: 0x3F, G: 0x54, B: 0x7F, A: 0xFF}
		textColor := eff.Color{R: 0xF5, G: 0x87, B: 0x35, A: 0xFF}
		drawButton(button.Text, button.Rect, bgColor, textColor, c)
	}

	drawOver := func(button *button.Button, c eff.Canvas) {
		bgColor := eff.Color{R: 0x99, G: 0x9F, B: 0xAD, A: 0xFF}
		textColor := eff.White()
		drawButton(button.Text, button.Rect, bgColor, textColor, c)
	}

	padding := 20
	buttonWidth := 100
	buttonHeight := 30
	topLeftButton := button.NewButton("NW", eff.Rect{X: padding, Y: padding, W: buttonWidth, H: buttonHeight}, drawDefault, drawDown, drawOver, clickHandler)
	b.buttons = append(b.buttons, &topLeftButton)
	c.AddClickable(&topLeftButton)

	bottomLeftButton := button.NewButton("SW", eff.Rect{X: padding, Y: c.Height() - padding - buttonHeight, W: buttonWidth, H: buttonHeight}, drawDefault, drawDown, drawOver, clickHandler)
	b.buttons = append(b.buttons, &bottomLeftButton)
	c.AddClickable(&bottomLeftButton)

	topRightButton := button.NewButton("NE", eff.Rect{X: c.Width() - padding - buttonWidth, Y: padding, W: buttonWidth, H: buttonHeight}, drawDefault, drawDown, drawOver, clickHandler)
	b.buttons = append(b.buttons, &topRightButton)
	c.AddClickable(&topRightButton)

	bottonRightButton := button.NewButton("SE", eff.Rect{X: c.Width() - padding - buttonWidth, Y: c.Height() - padding - buttonHeight, W: buttonWidth, H: buttonHeight}, drawDefault, drawDown, drawOver, clickHandler)
	b.buttons = append(b.buttons, &bottonRightButton)
	c.AddClickable(&bottonRightButton)

	b.middleText = "NW"
	b.initialized = true
}

func (b *buttonTest) Initialized() bool {
	return b.initialized
}

func (b *buttonTest) Draw(c eff.Canvas) {
	for _, button := range b.buttons {
		button.Draw(c)
	}

	tW, tH, err := c.GetTextSize(b.font, b.middleText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	textPoint := eff.Point{
		X: (c.Width() - tW) / 2,
		Y: (c.Height() - tH) / 2,
	}
	c.DrawText(b.font, b.middleText, eff.Black(), textPoint)
}

func (b *buttonTest) Update(c eff.Canvas) {

}

func main() {
	canvas := sdl.NewCanvas("Clickables", 800, 540, eff.White(), 60, true)

	canvas.Run(func() {
		bt := buttonTest{}
		canvas.AddDrawable(&bt)
	})
}
