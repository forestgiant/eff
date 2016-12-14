package eff

// Graphics describes the Graphics struct which is used to add graphics to the canvas
type Graphics interface {
	DrawPoints(points []Point, color Color)
	DrawPoint(Point, Color)
	DrawColorPoints([]Point, []Color)

	DrawLine(Point, Point, Color)
	DrawLines([]Point, Color)

	StrokeRect(Rect, Color)
	StrokeRects([]Rect, Color)
	StrokeColorRects([]Rect, []Color)
	FillRect(Rect, Color)
	FillRects([]Rect, Color)
	FillColorRects([]Rect, []Color)

	DrawText(Font, string, Color, Point) error
	GetTextSize(Font, string) (int, int, error)

	DrawImage(Image, Rect) error

	Begin(Rect)
	End(Rect)
}
