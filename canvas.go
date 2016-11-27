package eff

// KeyHandler function that is called when a key board event occurs
type KeyHandler func(key string)

// MouseButtonHandler function that is called when a mouse button event occurs
type MouseButtonHandler func(leftState bool, middleState bool, rightState bool)

// MouseMoveHandler function that is called when a mouse move event occurs
type MouseMoveHandler func(x int, y int)

// MouseWheelHandler function that is called when a mouse wheel event occurs
type MouseWheelHandler func(x int, y int)

// Image struct that defines the size location and path to a image
type Image struct {
	Path string
	Rect Rect
}

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	AddDrawable(drawable Drawable)
	RemoveDrawable(drawable Drawable)

	AddClickable(clickable Clickable)
	RemoveClickable(clickable Clickable)

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

	OpenFont(path string, size int) (Font, error)
	DrawText(font Font, text string, color Color, point Point) error
	GetTextSize(font Font, text string) (int, int, error)

	SetWidth(width int)
	SetHeight(height int)
	Width() int
	Height() int

	SetClearColor(color Color)

	Fullscreen() bool
	SetFullscreen(fullscreen bool)

	AddKeyUpHandler(handler KeyHandler)
	AddKeyDownHandler(handler KeyHandler)

	AddMouseDownHandler(handler MouseButtonHandler)
	AddMouseUpHandler(handler MouseButtonHandler)
	AddMouseMoveHandler(handler MouseMoveHandler)
	AddMouseWheelHandler(handler MouseWheelHandler)

	AddImage(i *Image)
	RemoveImage(i *Image)
}
