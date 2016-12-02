package sdl

import "github.com/forestgiant/eff"

type Shape struct {
	drawable

	bgColor   eff.Color
	drawCalls []func()
}

func (shape *Shape) Draw(canvas eff.Canvas) {
	if shape.graphics.renderer == nil {
		return
	}

	shape.graphics.FillRect(shape.rect.Scale(shape.scale), shape.bgColor)

	for _, fn := range shape.drawCalls {
		fn()
	}

	for _, child := range shape.children {
		child.Draw(canvas)
	}
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
		p = p.Scale(shape.Scale())
		shape.graphics.DrawPoint(p, c)
	})
}

func (shape *Shape) DrawPoints(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.ScalePoints(p, shape.Scale())
		shape.graphics.DrawPoints(p, c)
	})
}

func (shape *Shape) DrawColorPoints(p []eff.Point, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.ScalePoints(p, shape.Scale())
		shape.graphics.DrawColorPoints(p, c)
	})
}

func (shape *Shape) DrawLine(p1 eff.Point, p2 eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p1 = p1.Scale(shape.Scale())
		p2 = p2.Scale(shape.Scale())
		shape.graphics.DrawLine(p1, p2, c)
	})
}

func (shape *Shape) DrawLines(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.ScalePoints(p, shape.Scale())
		shape.graphics.DrawLines(p, c)
	})
}

func (shape *Shape) StrokeRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = r.Scale(shape.Scale())
		shape.graphics.StrokeRect(r, c)
	})
}

func (shape *Shape) StrokeRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = eff.ScaleRects(r, shape.Scale())
		shape.graphics.StrokeRects(r, c)
	})
}

func (shape *Shape) StrokeColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = eff.ScaleRects(r, shape.Scale())
		shape.graphics.StrokeColorRects(r, c)
	})
}

func (shape *Shape) FillRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = r.Scale(shape.Scale())
		shape.graphics.FillRect(r, c)
	})
}

func (shape *Shape) FillRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = eff.ScaleRects(r, shape.Scale())
		shape.graphics.FillRects(r, c)
	})
}

func (shape *Shape) FillColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = eff.ScaleRects(r, shape.Scale())
		shape.graphics.FillColorRects(r, c)
	})
}
