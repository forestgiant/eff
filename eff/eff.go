package eff

import (
	"math"
	"math/rand"
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

// RandomColor genereate a random color struct.  The opacity is also random
func RandomColor() Color {
	return Color{
		R: rand.Intn(0xFF),
		G: rand.Intn(0xFF),
		B: rand.Intn(0xFF),
		A: rand.Intn(0xFF),
	}
}

// Rect container for rectangle
type Rect struct {
	X int
	Y int
	W int
	H int
}

// ColorRect container for rectange and color
type ColorRect struct {
	Rect
	Color
}

// ColorPoint container for point and color
type ColorPoint struct {
	Point
	Color
}

// Font describes a ttf font
type Font struct {
	Path string
}

// Equals test to see if two rectangles occupy the same location exactly
func (r *Rect) Equals(otherRect Rect) bool {
	return (r.X == otherRect.X &&
		r.Y == otherRect.Y &&
		r.W == otherRect.W &&
		r.H == otherRect.H)
}

// Intersects check to see if a rectangle is inside of this rectangle
func (r *Rect) Intersects(otherRect Rect) bool {
	return (int(math.Abs(float64(r.X-otherRect.X)))*2 < (r.W + otherRect.W)) &&
		(int(math.Abs(float64(r.Y-otherRect.Y)))*2 < (r.H + otherRect.H))
}

// KeyHandler function that is called when a key board event occurs
type KeyHandler func(key string)

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	AddDrawable(drawable Drawable)
	RemoveDrawable(drawable Drawable)
	Run()

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

	SetFont(font Font, size int)
	DrawText(text string, color Color, point Point)

	SetWidth(width int)
	SetHeight(height int)
	Width() int
	Height() int

	Fullscreen() bool
	SetFullscreen(fullscreen bool)

	AddKeyUpHandler(handler KeyHandler)
	AddKeyDownHandler(handler KeyHandler)
}

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Init(canvas Canvas)
	Draw(canvas Canvas)
	Update(canvas Canvas)
	Initialized() bool
}
