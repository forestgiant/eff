package button

import (
	"log"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/util"
)

func defaultBGColor() eff.Color {
	return eff.Color{R: 219, G: 217, B: 217, A: 255}
}

func defaultTextColor() eff.Color {
	return eff.Black()
}

func overBGColor() eff.Color {
	return eff.Color{R: 238, G: 238, B: 238, A: 255}
}

func overTextColor() eff.Color {
	return eff.Black()
}

func downBGColor() eff.Color {
	return eff.Color{R: 246, G: 246, B: 246, A: 255}
}

func downTextColor() eff.Color {
	return eff.Black()
}

// Click function that is called when the button is clicked
type Click func(*Button)

// Button defines an eff.Drawable that maintains the state of a button
type Button struct {
	eff.Shape

	mouseDown        bool
	mouseOver        bool
	Text             string
	ClickHandler     func(b *Button)
	font             eff.Font
	defaultBGColor   eff.Color
	defaultTextColor eff.Color
	overBGColor      eff.Color
	overTextColor    eff.Color
	downBGColor      eff.Color
	downTextColor    eff.Color
	bgColor          eff.Color
	textColor        eff.Color
}

// Hitbox returns the hitbox rect of the button, this is the same as Button.Rect
func (b *Button) Hitbox() eff.Rect {
	return b.Rect()
}

// MouseDown function that is called when any mouse button is pressed down while the cursor is inside the hitbox
func (b *Button) MouseDown(leftState bool, middleState bool, rightState bool) {
	b.mouseDown = true

	b.drawButton()
}

// MouseUp function that is called when any mouse button is released while the cursor is inside the hitbox
func (b *Button) MouseUp(leftState bool, middleState bool, rightState bool) {
	if b.mouseDown {
		b.mouseDown = false
		b.ClickHandler(b)
	}

	b.drawButton()
}

// MouseOver function that is called when the mouse moves into the hitbox
func (b *Button) MouseOver() {
	b.mouseOver = true

	b.drawButton()
}

// MouseOut function that is called when the mouse moves out of the hitbox
func (b *Button) MouseOut() {
	b.mouseOver = false
	b.mouseDown = false

	b.drawButton()
}

func (b *Button) drawButton() {
	b.Clear()
	text, err := util.EllipseText(b.font, b.Text, b.Rect().W, b.Graphics())
	if err != nil {
		log.Fatal(err)
	}

	textW, textH, err := b.Graphics().GetTextSize(b.font, text)
	if err != nil {
		log.Fatal(err)
	}

	textPoint := eff.Point{
		X: (b.Rect().W - textW) / 2,
		Y: (b.Rect().H - textH) / 2,
	}

	bgColor := b.defaultBGColor
	textColor := b.defaultTextColor
	if b.mouseDown {
		bgColor = b.downBGColor
		textColor = b.downTextColor
	} else if b.mouseOver {
		bgColor = b.overBGColor
		textColor = b.overTextColor
	}

	b.SetBackgroundColor(bgColor)
	b.DrawText(b.font, text, textColor, textPoint)
}

// IsMouseOver function that returns true if the mouse cursor is currently inside the hitbox
func (b *Button) IsMouseOver() bool { return b.mouseOver }

func (b *Button) SetDefaultBGColor(c eff.Color) {
	b.defaultBGColor = c
	b.drawButton()
}

func (b *Button) DefaultBGColor() eff.Color {
	return b.defaultBGColor
}

func (b *Button) SetDefaultTextColor(c eff.Color) {
	b.defaultTextColor = c
	b.drawButton()
}

func (b *Button) DefaultTextColor() eff.Color {
	return b.defaultTextColor
}

func (b *Button) SetOverBGColor(c eff.Color) {
	b.overBGColor = c
	b.drawButton()
}

func (b *Button) OverBGColor() eff.Color {
	return b.overBGColor
}

func (b *Button) SetOverTextColor(c eff.Color) {
	b.overTextColor = c
	b.drawButton()
}

func (b *Button) OverTextColor() eff.Color {
	return b.overTextColor
}

func (b *Button) SetDownBGColor(c eff.Color) {
	b.downBGColor = c
	b.drawButton()
}

func (b *Button) DownBGColor() eff.Color {
	return b.downBGColor
}

func (b *Button) SetDownTextColor(c eff.Color) {
	b.downTextColor = c
	b.drawButton()
}

func (b *Button) DownTextColor() eff.Color {
	return b.downTextColor
}

// NewButton function that creates an instance of the component button
func NewButton(font eff.Font, text string, rect eff.Rect, clickhandler Click) *Button {
	b := &Button{
		font:         font,
		Text:         text,
		ClickHandler: clickhandler,
	}

	b.SetRect(rect)
	b.defaultBGColor = defaultBGColor()
	b.defaultTextColor = defaultTextColor()
	b.downBGColor = downBGColor()
	b.downTextColor = downTextColor()
	b.overBGColor = overBGColor()
	b.overTextColor = overTextColor()
	b.bgColor = defaultBGColor()
	b.textColor = defaultTextColor()
	b.SetGraphicsReadyHandler(func() {
		b.drawButton()
	})

	return b
}
