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

	ShouldClip() bool
	SetShouldClip(bool)
}

type Shape struct {
	drawable

	bgColor    Color
	drawCalls  []func()
	shouldClip bool
}

func (shape *Shape) offsetPoint(p Point) Point {

	return Point{
		X: p.X,
		Y: p.Y,
	}
}

func (shape *Shape) offsetPoints(points []Point) []Point {
	var offsetPoints []Point
	for _, p := range points {
		offsetPoints = append(offsetPoints, Point{
			X: p.X,
			Y: p.Y,
		})
	}

	return offsetPoints
}

func (shape *Shape) offsetRect(r Rect) Rect {
	return Rect{
		X: r.X,
		Y: r.Y,
		W: r.W,
		H: r.H,
	}
}

func (shape *Shape) offsetRects(rects []Rect) []Rect {
	var offsetRects []Rect
	for _, r := range rects {
		offsetRects = append(offsetRects, Rect{
			X: r.X,
			Y: r.Y,
			W: r.W,
			H: r.H,
		})
	}
	return offsetRects
}

func (shape *Shape) Draw(canvas Canvas) {
	shape.Graphics().Begin(shape.Rect())
	shape.graphics.FillRect(Rect{X: 0, Y: 0, W: shape.Rect().W, H: shape.Rect().H}, shape.bgColor)

	for _, fn := range shape.drawCalls {
		fn()
	}
	pRect := shape.Rect()
	if shape.Parent() != nil {
		pRect = shape.Parent().Rect()
	}
	shape.Graphics().End(shape.shouldClip, shape.Rect(), pRect)

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
		shape.Graphics().DrawPoint(shape.offsetPoint(p), c)
	})
}

func (shape *Shape) DrawPoints(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawPoints(shape.offsetPoints(p), c)
	})
}

func (shape *Shape) DrawColorPoints(p []Point, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawColorPoints(shape.offsetPoints(p), c)
	})
}

func (shape *Shape) DrawLine(p1 Point, p2 Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawLine(shape.offsetPoint(p1), shape.offsetPoint(p2), c)
	})
}

func (shape *Shape) DrawLines(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawLines(shape.offsetPoints(p), c)
	})
}

func (shape *Shape) StrokeRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeRect(shape.offsetRect(r), c)
	})
}

func (shape *Shape) StrokeRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeRects(shape.offsetRects(r), c)
	})
}

func (shape *Shape) StrokeColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeColorRects(shape.offsetRects(r), c)
	})
}

func (shape *Shape) FillRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRect(shape.offsetRect(r), c)
	})
}

func (shape *Shape) FillRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRects(shape.offsetRects(r), c)
	})
}

func (shape *Shape) FillColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillColorRects(shape.offsetRects(r), c)
	})
}

func (shape *Shape) DrawText(f Font, text string, c Color, p Point) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawText(f, text, c, shape.offsetPoint(p))
	})
}

func (shape *Shape) ShouldClip() bool {
	return shape.shouldClip
}

func (shape *Shape) SetShouldClip(shouldClip bool) {
	shape.shouldClip = shouldClip
}
