package button

import "github.com/forestgiant/eff"

// Draw function that is called to draw a button at a particular state
type Draw func(*Button, eff.Canvas)

// Click function that is called when the button is clicked
type Click func(*Button)

// Button defines an eff.Drawable that maintains the state of a button
type Button struct {
	Rect         eff.Rect
	mouseDown    bool
	mouseOver    bool
	Text         string
	ClickHandler func(b *Button)
	DrawDefault  Draw
	DrawDown     Draw
	DrawOver     Draw
}

// Hitbox returns the hitbox rect of the button, this is the same as Button.Rect
func (b *Button) Hitbox() eff.Rect {
	return b.Rect
}

// MouseDown function that is called when any mouse button is pressed down while the cursor is inside the hitbox
func (b *Button) MouseDown(leftState bool, middleState bool, rightState bool) {
	b.mouseDown = true
}

// MouseUp function that is called when any mouse button is released while the cursor is inside the hitbox
func (b *Button) MouseUp(leftState bool, middleState bool, rightState bool) {
	if b.mouseDown {
		b.mouseDown = false
		b.ClickHandler(b)
	}
}

// MouseOver function that is called when the mouse moves into the hitbox
func (b *Button) MouseOver() {
	b.mouseOver = true
}

// MouseOut function that is called when the mouse moves out of the hitbox
func (b *Button) MouseOut() {
	b.mouseOver = false
	b.mouseDown = false
}

// IsMouseOver function that returns true if the mouse cursor is currently inside the hitbox
func (b *Button) IsMouseOver() bool { return b.mouseOver }

// Draw calls the appropriate draw function based on the button state
func (b *Button) Draw(c eff.Canvas) {
	var drawFunc Draw
	if b.mouseDown {
		drawFunc = b.DrawDown
	} else if b.mouseOver {
		drawFunc = b.DrawOver
	} else {
		drawFunc = b.DrawDefault
	}

	if drawFunc == nil {
		drawFunc = b.DrawDefault
	}

	if drawFunc != nil {
		drawFunc(b, c)
	}
}

// NewButton function that creates an instance of the component button
func NewButton(text string, rect eff.Rect, drawDefault Draw, drawDown Draw, drawOver Draw, clickhandler Click) Button {
	return Button{
		Rect:         rect,
		Text:         text,
		ClickHandler: clickhandler,
		DrawDefault:  drawDefault,
		DrawDown:     drawDown,
		DrawOver:     drawOver,
	}
}
