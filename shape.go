package eff

type Shape interface {
	Drawable
	Container

	DrawPoints(points []Point, color Color)
	DrawPoint(Point, Color)
	DrawColorPoints([]ColorPoint)

	DrawRect(Rect, Color)
	DrawRects([]Rect, Color)
	FillRect(Rect, Color)
	FillRects([]Rect, Color)
	DrawColorRects([]ColorRect)

	DrawLine(Point, Point, Color)
	DrawLines([]Point, Color)

	DrawText(Font, string, Color, Point) error
	GetTextSize(Font, string) (int, int, error)
}
