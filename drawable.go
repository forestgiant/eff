package eff

import (
	"errors"
	"fmt"
	"sync"
)

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Draw(canvas Canvas)

	SetRect(Rect)
	Rect() Rect
	ParentOffsetRect() Rect

	SetParent(Drawable)
	Parent() Drawable

	SetGraphics(Graphics)
	Graphics() Graphics

	SetUpdateHandler(func())
	HandleUpdate()

	SetGraphicsReadyHandler(func())

	SetResizeHandler(func())

	AddChild(Drawable) error
	RemoveChild(Drawable) error
	Children() []Drawable

	ShouldDraw() bool
	SetShouldDraw(bool)
	RedrawChildren()

	TextureInvalid() bool
	SetTextureInvalid(bool)
	InvalidateChildTextures()

	IsVisible() bool
}

type drawable struct {
	mu                   sync.RWMutex
	rect                 Rect
	parent               Drawable
	graphics             Graphics
	drawNeeded           bool
	textureInvalid       bool
	children             []Drawable
	updateHandler        func()
	graphicsReadyHandler func()
	resizeHandler        func()
}

func (d *drawable) init() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.children == nil {
		d.children = []Drawable{}
	}
}

func (d *drawable) SetRect(r Rect) {
	resized := false
	if r.W != d.rect.W || r.H != d.rect.H {
		resized = true
	}

	d.rect = r
	if d.Parent() != nil {
		d.Parent().SetShouldDraw(true)
	}

	if resized {
		d.SetTextureInvalid(true)
		if d.resizeHandler != nil {
			d.resizeHandler()
		}
	}
}

func (d *drawable) Rect() Rect {
	return d.rect
}

func (d *drawable) SetParent(p Drawable) {
	d.parent = p
}

func (d *drawable) Parent() Drawable {
	return d.parent
}

func (d *drawable) SetGraphics(g Graphics) {
	d.graphics = g
	if d.graphics != nil && d.graphicsReadyHandler != nil {
		d.graphicsReadyHandler()
	}

	for _, child := range d.children {
		child.SetGraphics(g)
	}
}

func (d *drawable) Graphics() Graphics {
	return d.graphics
}

func (d *drawable) Draw(c Canvas) {}

func (d *drawable) AddChild(c Drawable) error {
	if d == nil {
		return errors.New("parent is nil")
	}

	if c == nil {
		return errors.New("child was nil")
	}

	d.init()
	d.mu.Lock()
	defer d.mu.Unlock()

	c.SetParent(Drawable(d))

	d.children = append(d.children, c)

	if d.graphics != nil {
		c.SetGraphics(d.graphics)
	}
	d.SetShouldDraw(true)

	return nil
}

func (d *drawable) RemoveChild(c Drawable) error {
	if d == nil {
		return errors.New("parent is nil")
	}

	if c == nil {
		return errors.New("child is nil")
	}

	d.init()
	d.mu.Lock()
	defer d.mu.Unlock()

	index := -1
	for i, child := range d.children {
		if c == child {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("Could not find %v to remove", c)
	}

	d.children[index].SetParent(nil)
	d.children[index].SetGraphics(nil)

	d.children = append(d.children[:index], d.children[index+1:]...)

	d.SetShouldDraw(true)
	return nil
}

func (d *drawable) Children() []Drawable {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.children
}

func (d *drawable) SetUpdateHandler(handler func()) {
	d.updateHandler = handler
}

func (d *drawable) HandleUpdate() {
	if d.updateHandler != nil {
		d.updateHandler()
	}

	for _, child := range d.children {
		child.HandleUpdate()
	}
}

func (d *drawable) SetGraphicsReadyHandler(handler func()) {
	d.graphicsReadyHandler = handler
}

func (d *drawable) SetResizeHandler(handler func()) {
	d.resizeHandler = handler
}

func (d *drawable) ParentOffsetRect() Rect {
	pRect := Rect{}
	if d.parent != nil {
		pRect = d.parent.ParentOffsetRect()
	}

	return Rect{
		X: d.rect.X + pRect.X,
		Y: d.rect.Y + pRect.Y,
		W: d.rect.W,
		H: d.rect.H,
	}
}

func (d *drawable) ShouldDraw() bool {
	return d.drawNeeded
}

func (d *drawable) SetShouldDraw(b bool) {
	d.drawNeeded = b
	if d.Parent() != nil && b {
		d.Parent().SetShouldDraw(b)
	}
}

func (d *drawable) RedrawChildren() {
	for _, child := range d.children {
		child.SetShouldDraw(true)
		child.RedrawChildren()
	}
}

func (d *drawable) TextureInvalid() bool {
	return d.textureInvalid
}

func (d *drawable) SetTextureInvalid(invalid bool) {
	d.textureInvalid = invalid
}

func (d *drawable) InvalidateChildTextures() {
	for _, child := range d.children {
		child.SetTextureInvalid(true)
		child.InvalidateChildTextures()
	}
}

func (d *drawable) IsVisible() bool {
	if d.Parent() == nil {
		return false
	}

	rect := d.Rect()
	if rect.X+rect.W < 0 || rect.Y+rect.H < 0 {
		return false
	}
	for p := d.Parent(); p != nil; p = p.Parent() {
		rect.X += p.Rect().X
		rect.Y += p.Rect().Y
		if rect.X+rect.W < 0 || rect.Y+rect.H < 0 {
			return false
		}
	}

	return true
}
