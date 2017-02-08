package eff

import (
	"errors"
	"math"
	"math/rand"

	"strconv"
)

const (
	// Version current semantic version of eff
	Version = "0.4.2"
)

// Point container for 2d points
type Point struct {
	X int
	Y int
}

// Scale returns a new scaled point
func (p *Point) Scale(s float64) Point {
	return Point{
		X: int(float64(p.X) * s),
		Y: int(float64(p.Y) * s),
	}
}

// Offset returns an offset point
func (p *Point) Offset(x int, y int) Point {
	return Point{
		X: p.X + x,
		Y: p.Y + y,
	}
}

// ScalePoints returns a new slice of scaled points
func ScalePoints(points []Point, s float64) []Point {
	var scaledPoints []Point
	for _, p := range points {
		scaledPoints = append(scaledPoints, p.Scale(s))
	}

	return scaledPoints
}

// OffsetPoints returns a new slice of offset points
func OffsetPoints(points []Point, x int, y int) []Point {
	var offsetPoints []Point
	for _, p := range points {
		offsetPoints = append(offsetPoints, p.Offset(x, y))
	}

	return offsetPoints
}

// Color container for argb colors
type Color struct {
	R int
	G int
	B int
	A int
}

// Add offset the rgb values of the color
func (c *Color) Add(v int) {
	c.R += v
	c.G += v
	c.B += v

	c.R = int(math.Min(0xFF, float64(c.R)))
	c.R = int(math.Max(0x00, float64(c.R)))
	c.G = int(math.Min(0xFF, float64(c.G)))
	c.G = int(math.Max(0x00, float64(c.G)))
	c.B = int(math.Min(0xFF, float64(c.B)))
	c.B = int(math.Max(0x00, float64(c.B)))
}

// RandomColor genereate a random color struct.  The opacity is also random
func RandomColor() Color {
	return Color{
		R: rand.Intn(0xFF),
		G: rand.Intn(0xFF),
		B: rand.Intn(0xFF),
		A: 0xFF,
	}
}

// Black returns a color struct that is black
func Black() Color {
	return Color{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	}
}

// White returns a color struct that is white
func White() Color {
	return Color{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}
}

// ColorWithHex creates an eff color w/ a hex string in the formant "#FF00FF"
func ColorWithHex(hex string) (Color, error) {
	if hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) < 6 {
		return Color{}, errors.New("Invalid hex color, too short")
	}

	r, err := strconv.ParseInt(hex[:2], 16, 32)
	if err != nil {
		return Color{}, err
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return Color{}, err
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return Color{}, err
	}

	return Color{
		R: int(r),
		G: int(g),
		B: int(b),
		A: 0xFF,
	}, nil
}

// Rect container for rectangle
type Rect struct {
	X int
	Y int
	W int
	H int
}

// Scale returns a new scaled Rect
func (r *Rect) Scale(s float64) Rect {
	return Rect{
		X: int(float64(r.X) * s),
		Y: int(float64(r.Y) * s),
		W: int(float64(r.W) * s),
		H: int(float64(r.H) * s),
	}
}

// LocalInside tests to see if rect is inside of this rect, assumes test rect has coordinates local to this rect
func (r *Rect) LocalInside(testRect Rect) bool {
	if testRect.X > r.W {
		return false
	}

	if testRect.Y > r.H {
		return false
	}

	if (testRect.X + testRect.W) < 0 {
		return false
	}

	if (testRect.Y + testRect.H) < 0 {
		return false
	}

	return true
}

// ScaleRects returns a new slice of scaled Rects
func ScaleRects(rects []Rect, s float64) []Rect {
	var scaledRects []Rect
	for _, r := range rects {
		scaledRects = append(scaledRects, r.Scale(s))
	}

	return scaledRects
}

// Font describes a ttf font
type Font interface {
	Path() string
	Size() int
}

// Image describes an image
type Image interface {
	Path() string
	Width() int
	Height() int
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

// Inside check to see if a point inside of this rectangle
func (r *Rect) Inside(p Point) bool {
	return (p.X > r.X) && (p.X < (r.X + r.W)) && (p.Y > r.Y) && (p.Y < (r.Y + r.H))
}
