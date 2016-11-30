package eff

type Shape interface {
	Drawable
	Container

	SetBackgroundColor(Color)
	BackgroundColor() Color

	Clear()

	DrawPoint(Point, Color)
	DrawPoints([]Point, Color)
	DrawColorPoints([]Point, []Color)

	DrawLine(Point, Point, Color)
	DrawLines([]Point, Color)
	DrawColorLines([]Point, []Color)

	StrokeRect(Rect, Color)
	StrokeRects([]Rect, Color)
	StrokeColorRects([]Rect, []Color)
	FillRect(Rect, Color)
	FillRects([]Rect, Color)
	FillColorRects([]Rect, []Color)

	DrawCalls() []func()
}
