package sdl

import "github.com/forestgiant/eff"

type Shape struct {
	rect          eff.Rect
	parent        eff.Drawable
	scale         float64
	bgColor       eff.Color
	drawCalls     []func()
	graphics      *Graphics
	children      []eff.Drawable
	updateHandler func()
}

func (shape *Shape) Draw(canvas eff.Canvas) {
	if shape.graphics.renderer == nil {
		return
	}

	shape.graphics.FillRect(shape.rect, shape.bgColor)
	for _, fn := range shape.drawCalls {
		fn()
	}
}

func (shape *Shape) SetRect(r eff.Rect) {
	shape.rect = r
}

func (shape *Shape) Rect() eff.Rect {
	return shape.rect
}

func (shape *Shape) SetParent(d eff.Drawable) {
	shape.parent = d
}

func (shape *Shape) Parent() eff.Drawable {
	return shape.parent
}

func (shape *Shape) SetScale(s float64) {
	shape.scale = s
}

func (shape *Shape) Scale() float64 {
	return shape.scale
}

func (shape *Shape) SetGraphics(g eff.Graphics) {
	sdlGraphics, ok := g.(*Graphics)
	if ok {
		shape.graphics = sdlGraphics
	}
}

func (shape *Shape) Graphics() eff.Graphics {
	return shape.graphics
}

func (shape *Shape) SetBackgroundColor(c eff.Color) {
	shape.bgColor = c
}

func (shape *Shape) BackgroundColor() eff.Color {
	return shape.bgColor
}

func (shape *Shape) Clear() {
	shape.drawCalls = []func(){}
}

func (shape *Shape) DrawPoint(p eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawPoint(p, c)
	})
}

func (shape *Shape) DrawPoints(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawPoints(p, c)
	})
}

func (shape *Shape) DrawColorPoints(p []eff.Point, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawColorPoints(p, c)
	})
}

func (shape *Shape) DrawLine(p1 eff.Point, p2 eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawLine(p1, p2, c)
	})
}

func (shape *Shape) DrawLines(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawLines(p, c)
	})
}

func (shape *Shape) StrokeRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.StrokeRect(r, c)
	})
}

func (shape *Shape) StrokeRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.StrokeRects(r, c)
	})
}

func (shape *Shape) StrokeColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.StrokeColorRects(r, c)
	})
}

func (shape *Shape) FillRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRect(r, c)
	})
}

func (shape *Shape) FillRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRects(r, c)
	})
}

func (shape *Shape) FillColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillColorRects(r, c)
	})
}

func (shape *Shape) SetUpdateHandler(handler func()) {
	shape.updateHandler = handler
}

func (shape *Shape) HandleUpdate() {
	if shape.updateHandler != nil {
		shape.updateHandler()
	}
}

func (shape *Shape) AddChild(d eff.Drawable) {
	if d == nil {
		return
	}

	d.SetParent(eff.Drawable(shape))
	d.SetScale(shape.scale)
	d.SetGraphics(shape.graphics)

	shape.children = append(shape.children, d)
}

func (shape *Shape) RemoveChild(d eff.Drawable) {
	if d == nil {
		return
	}

	index := -1
	for i, child := range shape.children {
		if d == child {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	shape.children[index].SetParent(nil)
	shape.children[index].SetGraphics(nil)
	shape.children[index].SetScale(1)

	shape.children = append(shape.children[:index], shape.children[index+1:]...)
}

func (shape *Shape) Children() []eff.Drawable {
	return shape.children
}
