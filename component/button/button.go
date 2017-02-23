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

func selectedBGColor() eff.Color {
	return eff.Black()
}

func selectedTextColor() eff.Color {
	return eff.White()
}

// Click function that is called when the button is clicked
type Click func(*Button)

// Button defines an eff.Drawable that maintains the state of a button
type Button struct {
	eff.Shape

	mouseDown         bool
	mouseOver         bool
	selected          bool
	Text              string
	ClickHandler      func(b *Button)
	font              eff.Font
	defaultBGColor    eff.Color
	defaultTextColor  eff.Color
	overBGColor       eff.Color
	overTextColor     eff.Color
	downBGColor       eff.Color
	downTextColor     eff.Color
	selectedBGColor   eff.Color
	selectedTextColor eff.Color
	bgColor           eff.Color
	textColor         eff.Color
	img               eff.Image
	imagePadding      int
	imgShape          *eff.Shape
}

// Hitbox returns the hitbox rect of the button, this is the same as Button.Rect
func (b *Button) Hitbox() eff.Rect {
	return b.ParentOffsetRect()
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
		if b.ClickHandler != nil {
			b.ClickHandler(b)
		}

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
	if b.Graphics() == nil {
		return
	}

	b.Clear()

	bgColor := b.defaultBGColor
	textColor := b.defaultTextColor
	if b.selected {
		bgColor = b.selectedBGColor
		textColor = b.selectedTextColor
	} else if b.mouseDown {
		bgColor = b.downBGColor
		textColor = b.downTextColor
	} else if b.mouseOver {
		bgColor = b.overBGColor
		textColor = b.overTextColor
	}
	b.SetBackgroundColor(bgColor)

	if len(b.Text) > 0 && b.font != nil {
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
		b.DrawText(b.font, text, textColor, textPoint)
	}

	if b.img != nil {
		if b.imgShape != nil && b.imgShape.Parent() != nil {
			b.RemoveChild(b.imgShape)
		}
		b.imgShape = &eff.Shape{}
		b.imgShape.SetBackgroundColor(eff.Color{R: 0x00, G: 0x00, B: 0x00, A: 0x00})
		b.imgShape.SetRect(eff.Rect{
			X: b.imagePadding,
			Y: b.imagePadding,
			W: b.Rect().W - (2 * b.imagePadding),
			H: b.Rect().H - (2 * b.imagePadding),
		})
		b.AddChild(b.imgShape)

		aspect := float64(b.img.Height()) / float64(b.img.Width())
		h := util.RoundToInt(float64(b.imgShape.Rect().W) * aspect)
		y := util.RoundToInt(float64(b.imgShape.Rect().H-h) / 2)
		b.Clear()
		b.imgShape.Clear()
		b.imgShape.DrawImage(b.img, eff.Rect{
			X: 0,
			Y: y,
			W: b.imgShape.Rect().W,
			H: h,
		})
	}
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

func (b *Button) SetSelectedBGColor(c eff.Color) {
	b.selectedBGColor = c
	b.drawButton()
}

func (b *Button) SelectedBGColor() eff.Color {
	return b.selectedBGColor
}

func (b *Button) SetSelectedTextColor(c eff.Color) {
	b.selectedTextColor = c
	b.drawButton()
}

func (b *Button) SelectedTextColor() eff.Color {
	return b.selectedTextColor
}

func (b *Button) Selected() bool {
	return b.selected
}

func (b *Button) SetSelected(selected bool) {
	b.selected = selected
	b.drawButton()
}

func (b *Button) SetImage(img eff.Image) {
	b.img = img
}

func (b *Button) SetImagePadding(padding int) {
	b.imagePadding = padding
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
	b.selectedBGColor = selectedBGColor()
	b.selectedTextColor = selectedTextColor()
	b.bgColor = defaultBGColor()
	b.textColor = defaultTextColor()

	b.SetGraphicsReadyHandler(func() {
		b.drawButton()
	})
	b.SetResizeHandler(func() {
		b.drawButton()
	})

	return b
}

func NewImageButton(imgPath string, rect eff.Rect, canvas eff.Canvas, clickhandler Click) (*Button, error) {
	img, err := canvas.OpenImage(imgPath)
	if err != nil {
		return nil, err
	}

	b := &Button{
		img:          img,
		ClickHandler: clickhandler,
	}

	b.SetRect(rect)
	b.defaultBGColor = defaultBGColor()
	b.defaultTextColor = defaultTextColor()
	b.downBGColor = downBGColor()
	b.downTextColor = downTextColor()
	b.overBGColor = overBGColor()
	b.overTextColor = overTextColor()
	b.selectedBGColor = selectedBGColor()
	b.selectedTextColor = selectedTextColor()
	b.bgColor = defaultBGColor()
	b.textColor = defaultTextColor()

	b.SetGraphicsReadyHandler(func() {
		b.drawButton()
	})
	b.SetResizeHandler(func() {
		b.drawButton()
	})

	return b, nil
}
