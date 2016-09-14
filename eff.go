package eff

const (
	windowTitle = "Effulgent"
	frameRate   = 90
	frameTime   = 1000 / frameRate
)

// Point container for 2d points
type Point struct {
	X int
	Y int
}

// Color container for argb colors
type Color struct {
	R int
	G int
	B int
	A int
}

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	AddDrawable(drawable Drawable)
	Run() int
	DrawPoints(points *[]Point, color Color)
	SetWidth(width int)
	SetHeight(height int)
	Width() int
	Height() int
}

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Init(canvas Canvas)
	Draw(canvas Canvas)
	Update(canvas Canvas)
}
