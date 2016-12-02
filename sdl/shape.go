package sdl

import "github.com/forestgiant/eff"

type Shape struct {
	drawable

	bgColor   eff.Color
	drawCalls []func()
}

func (shape *Shape) offsetPoint(p eff.Point) eff.Point {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	return eff.Point{
		X: p.X + shape.Rect().X + px,
		Y: p.Y + shape.Rect().Y + py,
	}
}

func (shape *Shape) offsetPoints(points []eff.Point) []eff.Point {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	var offsetPoints []eff.Point
	for _, p := range points {
		offsetPoints = append(offsetPoints, eff.Point{
			X: p.X + shape.Rect().X + px,
			Y: p.Y + shape.Rect().Y + py,
		})
	}

	return offsetPoints
}

func (shape *Shape) offsetRect(r eff.Rect) eff.Rect {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	return eff.Rect{
		X: r.X + shape.Rect().X + px,
		Y: r.Y + shape.Rect().Y + py,
		W: r.W,
		H: r.H,
	}
}

func (shape *Shape) offsetRects(rects []eff.Rect) []eff.Rect {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	var offsetRects []eff.Rect
	for _, r := range rects {
		offsetRects = append(offsetRects, eff.Rect{
			X: r.X + shape.Rect().X + px,
			Y: r.Y + shape.Rect().Y + py,
			W: r.W,
			H: r.H,
		})
	}
	return offsetRects
}

func (shape *Shape) Draw(canvas eff.Canvas) {
	if shape.graphics.renderer == nil {
		return
	}
	shape.graphics.FillRect(shape.offsetRect(shape.Rect()), shape.bgColor)

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
		p = shape.offsetPoint(p)
		shape.graphics.DrawPoint(p, c)
	})
}

func (shape *Shape) DrawPoints(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.graphics.DrawPoints(p, c)
	})
}

func (shape *Shape) DrawColorPoints(p []eff.Point, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.graphics.DrawColorPoints(p, c)
	})
}

func (shape *Shape) DrawLine(p1 eff.Point, p2 eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p1 = p1.Offset(shape.Rect().X, shape.Rect().Y)
		p2 = p2.Offset(shape.Rect().X, shape.Rect().Y)
		shape.graphics.DrawLine(p1, p2, c)
	})
}

func (shape *Shape) DrawLines(p []eff.Point, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = eff.OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.graphics.DrawLines(p, c)
	})
}

func (shape *Shape) StrokeRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRect(r)
		shape.graphics.StrokeRect(r, c)
	})
}

func (shape *Shape) StrokeRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.StrokeRects(r, c)
	})
}

func (shape *Shape) StrokeColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.StrokeColorRects(r, c)
	})
}

func (shape *Shape) FillRect(r eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRect(r)
		shape.graphics.FillRect(r, c)
	})
}

func (shape *Shape) FillRects(r []eff.Rect, c eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.FillRects(r, c)
	})
}

func (shape *Shape) FillColorRects(r []eff.Rect, c []eff.Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.FillColorRects(r, c)
	})
}
