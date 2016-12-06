package eff

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Draw(canvas Canvas)

	SetRect(Rect)
	Rect() Rect

	SetParent(Drawable)
	Parent() Drawable

	SetGraphics(Graphics)
	Graphics() Graphics

	SetUpdateHandler(func())
	HandleUpdate()

	SetGraphicsReadyHandler(func())

	AddChild(Drawable)
	RemoveChild(Drawable)
	Children() []Drawable
}

type drawable struct {
	rect                 Rect
	parent               Drawable
	graphics             Graphics
	children             []Drawable
	updateHandler        func()
	graphicsReadyHandler func()
}

func (d *drawable) SetRect(r Rect) {
	d.rect = r
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

func (d *drawable) AddChild(c Drawable) {
	if d == nil {
		return
	}

	c.SetParent(Drawable(d))

	d.children = append(d.children, c)

	if d.graphics != nil {
		c.SetGraphics(d.graphics)
	}
}

func (d *drawable) RemoveChild(c Drawable) {
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

	d.children = append(d.children[:index], d.children[index+1:]...)
}

func (d *drawable) Children() []Drawable {
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
