package sdl

import "github.com/forestgiant/eff"

type drawable struct {
	rect          eff.Rect
	parent        eff.Drawable
	scale         float64
	graphics      *Graphics
	children      []eff.Drawable
	updateHandler func()
}

func (d *drawable) SetRect(r eff.Rect) {
	d.rect = r
}

func (d *drawable) Rect() eff.Rect {
	return d.rect
}

func (d *drawable) SetParent(p eff.Drawable) {
	d.parent = p
}

func (d *drawable) Parent() eff.Drawable {
	return d.parent
}

func (d *drawable) SetScale(s float64) {
	d.scale = s
	for _, child := range d.children {
		child.SetScale(s)
	}
}

func (d *drawable) Scale() float64 {
	return d.scale
}

func (d *drawable) SetGraphics(g eff.Graphics) {
	sdlGraphics, ok := g.(*Graphics)
	if ok {
		d.graphics = sdlGraphics
	}
}

func (d *drawable) Graphics() eff.Graphics {
	return d.graphics
}

func (d *drawable) Draw(c eff.Canvas) {}

func (d *drawable) AddChild(c eff.Drawable) {
	if d == nil {
		return
	}

	c.SetParent(eff.Drawable(d))
	c.SetScale(d.scale)
	c.SetGraphics(d.graphics)

	d.children = append(d.children, c)
}

func (d *drawable) RemoveChild(c eff.Drawable) {
	if d == nil {
		return
	}

	index := -1
	for i, child := range d.children {
		if c == child {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	d.children[index].SetParent(nil)
	d.children[index].SetGraphics(nil)
	d.children[index].SetScale(1)

	d.children = append(d.children[:index], d.children[index+1:]...)
}

func (d *drawable) Children() []eff.Drawable {
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
