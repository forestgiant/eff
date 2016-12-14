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

	DrawImage(Image, Rect)
}

// Shape struct that can be added as a child to a canvas or another Shape
type Shape struct {
	drawable

	bgColor   Color
	drawCalls []func()
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

// Draw this draws the shape and all of its children to the canvas, typically called by the canvas its added to
func (shape *Shape) Draw(canvas Canvas) {

	shape.Graphics().Begin(shape.Rect())
	shape.graphics.FillRect(Rect{X: 0, Y: 0, W: shape.Rect().W, H: shape.Rect().H}, shape.bgColor)

	for _, fn := range shape.drawCalls {
		fn()
	}

	for _, child := range shape.children {
		child.Draw(canvas)
	}

	shape.Graphics().End(shape.Rect())
}

// SetBackgroundColor sets the background color of the shape
func (shape *Shape) SetBackgroundColor(c Color) {
	shape.bgColor = c
}

// BackgroundColor returns the current background color of the shape
func (shape *Shape) BackgroundColor() Color {
	return shape.bgColor
}

// Clear removes all the of current draw calls on the shape
func (shape *Shape) Clear() {
	shape.drawCalls = []func(){}
}

// DrawPoint adds a single point to the shape
func (shape *Shape) DrawPoint(p Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawPoint(shape.offsetPoint(p), c)
	})
}

// DrawPoints draws a slice of points to the shape all the same color
func (shape *Shape) DrawPoints(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawPoints(shape.offsetPoints(p), c)
	})
}

// DrawColorPoints draws a slice of points to the shape using different colors, expects color slice to equal the point slice
func (shape *Shape) DrawColorPoints(p []Point, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawColorPoints(shape.offsetPoints(p), c)
	})
}

// DrawLine draws a line to the shape
func (shape *Shape) DrawLine(p1 Point, p2 Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawLine(shape.offsetPoint(p1), shape.offsetPoint(p2), c)
	})
}

// DrawLines draws a slice of lines to canvas using a single color.  The point slice length should be even since lines are defined by two points
func (shape *Shape) DrawLines(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().DrawLines(shape.offsetPoints(p), c)
	})
}

// StrokeRect strokes a rect to the canvas
func (shape *Shape) StrokeRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeRect(shape.offsetRect(r), c)
	})
}

// StrokeRects strokes a slice of rects using a single color
func (shape *Shape) StrokeRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeRects(shape.offsetRects(r), c)
	})
}

// StrokeColorRects strokes a slice of rects using different colors, expects the length of color slice to equal the length of the rect slice
func (shape *Shape) StrokeColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.Graphics().StrokeColorRects(shape.offsetRects(r), c)
	})
}

// FillRect fills a single rect to the canvas
func (shape *Shape) FillRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRect(shape.offsetRect(r), c)
	})
}

// FillRects fills a slice of rects to the canvas using a single color
func (shape *Shape) FillRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillRects(shape.offsetRects(r), c)
	})
}

// FillColorRects fills a slice of rects to the canvas using a different color for each, expects the length of the color slice to equal the length of the rect slice
func (shape *Shape) FillColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.FillColorRects(shape.offsetRects(r), c)
	})
}

// DrawText draws a text string to the canvas
func (shape *Shape) DrawText(f Font, text string, c Color, p Point) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawText(f, text, c, shape.offsetPoint(p))
	})
}

// DrawImage draws an image to the canvas
func (shape *Shape) DrawImage(i Image, r Rect) {
	shape.drawCalls = append(shape.drawCalls, func() {
		shape.graphics.DrawImage(i, r)
	})
}