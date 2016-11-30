package eff

// KeyHandler function that is called when a key board event occurs
type KeyHandler func(key string)

// MouseButtonHandler function that is called when a mouse button event occurs
type MouseButtonHandler func(leftState bool, middleState bool, rightState bool)

// MouseMoveHandler function that is called when a mouse move event occurs
type MouseMoveHandler func(x int, y int)

// MouseWheelHandler function that is called when a mouse wheel event occurs
type MouseWheelHandler func(x int, y int)

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	Shape

	AddClickable(clickable Clickable)
	RemoveClickable(clickable Clickable)

	Run(setup func())

	OpenFont(path string, size int) (Font, error)

	SetClearColor(color Color)

	Fullscreen() bool
	SetFullscreen(fullscreen bool)

	AddKeyUpHandler(handler KeyHandler)
	AddKeyDownHandler(handler KeyHandler)

	AddMouseDownHandler(handler MouseButtonHandler)
	AddMouseUpHandler(handler MouseButtonHandler)
	AddMouseMoveHandler(handler MouseMoveHandler)
	AddMouseWheelHandler(handler MouseWheelHandler)
}
