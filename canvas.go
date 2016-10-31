package eff

// KeyHandler function that is called when a key board event occurs
type KeyHandler func(key string)

type Image struct {
	Path string
	Rect Rect
}

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	AddDrawable(drawable Drawable)
	RemoveDrawable(drawable Drawable)
	Run(setup func())

	DrawPoint(point Point, color Color)
	DrawPoints(points []Point, color Color)
	DrawColorPoints(colorPoints []ColorPoint)

	DrawLine(p1 Point, p2 Point, color Color)
	DrawLines(points []Point, color Color)

	DrawRect(rect Rect, color Color)
	DrawRects(rect []Rect, color Color)
	DrawColorRects(colorRect []ColorRect)
	FillRect(rect Rect, color Color)
	FillRects(rect []Rect, color Color)

	SetFont(font Font, size int) error
	DrawText(text string, color Color, point Point) error
	GetTextSize(text string) (int, int, error)

	SetWidth(width int)
	SetHeight(height int)
	Width() int
	Height() int

	Fullscreen() bool
	SetFullscreen(fullscreen bool)

	AddKeyUpHandler(handler KeyHandler)
	AddKeyDownHandler(handler KeyHandler)

	AddImage(i *Image)
	RemoveImage(i *Image)
}
