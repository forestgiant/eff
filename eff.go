package eff

import (
	"errors"
	"math"
	"math/rand"

	"strconv"
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
type Font interface {
	Path() string
	Size() int
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
