package eff

type shape interface {
	Drawable

	SetBackgroundColor(Color)
	BackgroundColor() Color

	Clear()

	DrawPoint(Point, Color)
	DrawPoints([]Point, Color)
	DrawColorPoints([]Point, []Color)

	DrawLine(Point, Point, Color)
	DrawLines([]Point, Color)

	StrokeRect(Rect, Color)
	StrokeRects([]Rect, Color)
	StrokeColorRects([]Rect, []Color)
	FillRect(Rect, Color)
	FillRects([]Rect, Color)
	FillColorRects([]Rect, []Color)
}

type Shape struct {
	drawable

	bgColor   Color
	drawCalls []func()
}

func (shape *Shape) offsetPoint(p Point) Point {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	return Point{
		X: p.X + shape.Rect().X + px,
		Y: p.Y + shape.Rect().Y + py,
	}
}

func (shape *Shape) offsetPoints(points []Point) []Point {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	var offsetPoints []Point
	for _, p := range points {
		offsetPoints = append(offsetPoints, Point{
			X: p.X + shape.Rect().X + px,
			Y: p.Y + shape.Rect().Y + py,
		})
	}

	return offsetPoints
}

func (shape *Shape) offsetRect(r Rect) Rect {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	return Rect{
		X: r.X + shape.Rect().X + px,
		Y: r.Y + shape.Rect().Y + py,
		W: r.W,
		H: r.H,
	}
}

func (shape *Shape) offsetRects(rects []Rect) []Rect {
	px := 0
	py := 0
	if shape.parent != nil {
		px = shape.parent.Rect().X
		py = shape.parent.Rect().Y
	}
	var offsetRects []Rect
	for _, r := range rects {
		offsetRects = append(offsetRects, Rect{
			X: r.X + shape.Rect().X + px,
			Y: r.Y + shape.Rect().Y + py,
			W: r.W,
			H: r.H,
		})
	}
	return offsetRects
}

func (shape *Shape) Draw(canvas Canvas) {
	shape.graphics.FillRect(shape.offsetRect(shape.Rect()), shape.bgColor)

	for _, fn := range shape.drawCalls {
		fn()
	}

	for _, child := range shape.children {
		child.Draw(canvas)
	}
}

func (shape *Shape) SetBackgroundColor(c Color) {
	shape.bgColor = c
}

func (shape *Shape) BackgroundColor() Color {
	return shape.bgColor
}

func (shape *Shape) Clear() {
	shape.drawCalls = []func(){}
}

func (shape *Shape) DrawPoint(p Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = shape.offsetPoint(p)
		shape.Graphics().DrawPoint(p, c)
	})
}

func (shape *Shape) DrawPoints(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.Graphics().DrawPoints(p, c)
	})
}

func (shape *Shape) DrawColorPoints(p []Point, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.Graphics().DrawColorPoints(p, c)
	})
}

func (shape *Shape) DrawLine(p1 Point, p2 Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p1 = p1.Offset(shape.Rect().X, shape.Rect().Y)
		p2 = p2.Offset(shape.Rect().X, shape.Rect().Y)
		shape.Graphics().DrawLine(p1, p2, c)
	})
}

func (shape *Shape) DrawLines(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = OffsetPoints(p, shape.Rect().X, shape.Rect().Y)
		shape.Graphics().DrawLines(p, c)
	})
}

func (shape *Shape) StrokeRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRect(r)
		shape.Graphics().StrokeRect(r, c)
	})
}

func (shape *Shape) StrokeRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.Graphics().StrokeRects(r, c)
	})
}

func (shape *Shape) StrokeColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.Graphics().StrokeColorRects(r, c)
	})
}

func (shape *Shape) FillRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRect(r)
		shape.graphics.FillRect(r, c)
	})
}

func (shape *Shape) FillRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.FillRects(r, c)
	})
}

func (shape *Shape) FillColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		r = shape.offsetRects(r)
		shape.graphics.FillColorRects(r, c)
	})
}

func (shape *Shape) DrawText(f Font, text string, c Color, p Point) {
	shape.drawCalls = append(shape.drawCalls, func() {
		p = shape.offsetPoint(p)
		shape.graphics.DrawText(f, text, c, p)
	})
}
