package sdl

import "github.com/forestgiant/eff"

type Shape struct {
	Container

	rect      eff.Rect
	parent    eff.Container
	scale     float64
	bgColor   eff.Color
	drawCalls []func()
	graphics  *Graphics
}

func (shape *Shape) Draw(canvas eff.Canvas) {
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

func (shape *Shape) SetParent(c eff.Container) {
	shape.parent = c
}

func (shape *Shape) Parent() eff.Container {
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

func (shape *Shape) DrawCalls() []func() {
	return shape.drawCalls
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
