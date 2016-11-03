package main

import (
	"fmt"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type button struct {
	rect             eff.Rect
	defaultBGColor   eff.Color
	defaultTextColor eff.Color
	overBGColor      eff.Color
	overTextColor    eff.Color
	downBGColor      eff.Color
	downTextColor    eff.Color
	bgColor          eff.Color
	textColor        eff.Color
	text             string
	mouseDown        bool
	mouseOver        bool
	clickHandler     func(b *button)
}

func (b *button) Hitbox() eff.Rect {
	return b.rect
}

func (b *button) MouseDown(leftState bool, middleState bool, rightState bool) {
	b.mouseDown = true
	if leftState {
		b.bgColor = b.downBGColor
		b.textColor = b.downTextColor
	}
}

func (b *button) MouseUp(leftState bool, middleState bool, rightState bool) {
	if b.mouseDown {
		b.mouseDown = false
		b.clickHandler(b)
	}

	b.bgColor = b.overBGColor
	b.textColor = b.overTextColor
}

func (b *button) MouseOver() {
	b.mouseOver = true
	b.bgColor = b.overBGColor
	b.textColor = b.overTextColor
}

func (b *button) MouseOut() {
	b.mouseOver = false
	b.mouseDown = false
	b.bgColor = b.defaultBGColor
	b.textColor = b.defaultTextColor
}

func (b *button) IsMouseOver() bool { return b.mouseOver }

func newButton(text string, rect eff.Rect) button {
	defaultBG := eff.Color{R: 0xDD, G: 0xDD, B: 0xDD, A: 0xFF}
	defaultText := eff.Black()

	overBG := eff.Color{R: 0x99, G: 0x9F, B: 0xAD, A: 0xFF}
	overText := eff.White()

	downBG := eff.Color{R: 0x3F, G: 0x54, B: 0x7F, A: 0xFF}
	downText := eff.Color{R: 0xF5, G: 0x87, B: 0x35, A: 0xFF}

	b := button{
		defaultBGColor:   defaultBG,
		defaultTextColor: defaultText,
		overBGColor:      overBG,
		overTextColor:    overText,
		downBGColor:      downBG,
		downTextColor:    downText,
		rect:             rect,
		bgColor:          defaultBG,
		textColor:        defaultText,
		text:             text,
	}

	return b
}

type buttonTest struct {
	initialized bool
	buttons     []*button
	middleText  string
}

func (b *buttonTest) Init(c eff.Canvas) {
	font := eff.Font{
		Path: "../assets/fonts/roboto/Roboto-Bold.ttf",
	}
	err := c.SetFont(font, 15)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	clickHandler := func(button *button) {
		b.middleText = button.text
	}

	padding := 20
	buttonWidth := 100
	buttonHeight := 30
	topLeftButton := newButton("NW", eff.Rect{X: padding, Y: padding, W: buttonWidth, H: buttonHeight})
	topLeftButton.clickHandler = clickHandler
	b.buttons = append(b.buttons, &topLeftButton)
	c.AddClickable(&topLeftButton)

	bottomLeftButton := newButton("SW", eff.Rect{X: padding, Y: c.Height() - padding - buttonHeight, W: buttonWidth, H: buttonHeight})
	bottomLeftButton.clickHandler = clickHandler
	b.buttons = append(b.buttons, &bottomLeftButton)
	c.AddClickable(&bottomLeftButton)

	topRightButton := newButton("NE", eff.Rect{X: c.Width() - padding - buttonWidth, Y: padding, W: buttonWidth, H: buttonHeight})
	topRightButton.clickHandler = clickHandler
	b.buttons = append(b.buttons, &topRightButton)
	c.AddClickable(&topRightButton)

	bottonRightButton := newButton("SE", eff.Rect{X: c.Width() - padding - buttonWidth, Y: c.Height() - padding - buttonHeight, W: buttonWidth, H: buttonHeight})
	bottonRightButton.clickHandler = clickHandler
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
		c.FillRect(button.rect, button.bgColor)
		tW, tH, err := c.GetTextSize(button.text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		textPoint := eff.Point{
			X: button.rect.X + ((button.rect.W - tW) / 2),
			Y: button.rect.Y + ((button.rect.H - tH) / 2),
		}

		c.DrawText(button.text, button.textColor, textPoint)
	}

	tW, tH, err := c.GetTextSize(b.middleText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	textPoint := eff.Point{
		X: (c.Width() - tW) / 2,
		Y: (c.Height() - tH) / 2,
	}
	c.DrawText(b.middleText, eff.Black(), textPoint)
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
