package eff

import (
	"math"
	"math/rand"
)

const (
	windowTitle   = "Effulgent"
	defaultWidth  = 480
	defaultHeight = 320
	frameRate     = 90
	frameTime     = 1000 / frameRate
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

func (c Color) RandomColor() Color {
	return Color{
		R: rand.Intn(255),
		G: rand.Intn(255),
		B: rand.Intn(255),
		A: rand.Intn(255),
	}
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

func (r *Rect) Intersects(otherRect Rect) bool {
	return (int(math.Abs(float64(r.X-otherRect.X)))*2 < (r.W + otherRect.W)) &&
		(int(math.Abs(float64(r.Y-otherRect.Y)))*2 < (r.H + otherRect.H))
}

type KeyHandler func(key string, canvas Canvas)

// Canvas interface describing methods required for canvas renderers
type Canvas interface {
	AddDrawable(drawable Drawable)
	RemoveDrawable(drawable Drawable)
	Run() int

	DrawPoint(point Point, color Color)
	DrawPoints(points *[]Point, color Color)

	DrawLine(p1 Point, p2 Point, color Color)
	DrawLines(points *[]Point, color Color)

	DrawRect(rect Rect, color Color)
	DrawRects(rect *[]Rect, color Color)
	FillRect(rect Rect, color Color)
	FillRects(rect *[]Rect, color Color)

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
