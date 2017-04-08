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
	drawCalls []func(Graphics)
}

// Draw this draws the shape and all of its children to the canvas, typically called by the canvas its added to
func (shape *Shape) Draw(canvas Canvas) {
	shape.mu.RLock()
	defer shape.mu.RUnlock()
	if shape.Graphics() == nil {
		return
	}

	shape.Graphics().Begin(shape)

	if shape.ShouldDraw() {

		shape.Graphics().FillRect(Rect{X: 0, Y: 0, W: shape.Rect().W, H: shape.Rect().H}, shape.bgColor)
		for _, fn := range shape.drawCalls {
			fn(shape.Graphics())
		}
		shape.SetShouldDraw(false)

		for _, child := range shape.children {
			rect := shape.Rect()
			if rect.LocalInside(child.Rect()) {
				child.Draw(canvas)
			}
		}
	}

	shape.Graphics().End(shape)
}

// SetBackgroundColor sets the background color of the shape
func (shape *Shape) SetBackgroundColor(c Color) {
	shape.bgColor = c
	shape.SetShouldDraw(true)
}

// BackgroundColor returns the current background color of the shape
func (shape *Shape) BackgroundColor() Color {
	return shape.bgColor
}

// Clear removes all the of current draw calls on the shape
func (shape *Shape) Clear() {
	shape.drawCalls = nil
	shape.SetShouldDraw(true)
}

// DrawPoint adds a single point to the shape
func (shape *Shape) DrawPoint(p Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawPoint(p, c)
	})
	shape.SetShouldDraw(true)
}

// DrawPoints draws a slice of points to the shape all the same color
func (shape *Shape) DrawPoints(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawPoints(p, c)
	})
	shape.SetShouldDraw(true)
}

// DrawColorPoints draws a slice of points to the shape using different colors, expects color slice to equal the point slice
func (shape *Shape) DrawColorPoints(p []Point, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawColorPoints(p, c)
	})
	shape.SetShouldDraw(true)
}

// DrawLine draws a line to the shape
func (shape *Shape) DrawLine(p1 Point, p2 Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawLine(p1, p2, c)
	})
	shape.SetShouldDraw(true)
}

// DrawLines draws a slice of lines to canvas using a single color.  The point slice length should be even since lines are defined by two points
func (shape *Shape) DrawLines(p []Point, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawLines(p, c)
	})
	shape.SetShouldDraw(true)
}

// StrokeRect strokes a rect to the canvas
func (shape *Shape) StrokeRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.StrokeRect(r, c)
	})
	shape.SetShouldDraw(true)
}

// StrokeRects strokes a slice of rects using a single color
func (shape *Shape) StrokeRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.StrokeRects(r, c)
	})
	shape.SetShouldDraw(true)
}

// StrokeColorRects strokes a slice of rects using different colors, expects the length of color slice to equal the length of the rect slice
func (shape *Shape) StrokeColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.StrokeColorRects(r, c)
	})
	shape.SetShouldDraw(true)
}

// FillRect fills a single rect to the canvas
func (shape *Shape) FillRect(r Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.FillRect(r, c)
	})
	shape.SetShouldDraw(true)
}

// FillRects fills a slice of rects to the canvas using a single color
func (shape *Shape) FillRects(r []Rect, c Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.FillRects(r, c)
	})
	shape.SetShouldDraw(true)
}

// FillColorRects fills a slice of rects to the canvas using a different color for each, expects the length of the color slice to equal the length of the rect slice
func (shape *Shape) FillColorRects(r []Rect, c []Color) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.FillColorRects(r, c)
	})
	shape.SetShouldDraw(true)
}

// DrawText draws a text string to the canvas
func (shape *Shape) DrawText(f Font, text string, c Color, p Point) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawText(f, text, c, p)
	})
	shape.SetShouldDraw(true)
}

// DrawImage draws an image to the canvas
func (shape *Shape) DrawImage(i Image, r Rect) {
	shape.drawCalls = append(shape.drawCalls, func(g Graphics) {
		g.DrawImage(i, r)
	})
	shape.SetShouldDraw(true)
}
