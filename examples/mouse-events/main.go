package main

import (
	"math"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	startSize = 20
	minSize   = 5
	maxSize   = 150
)

type mouseBox struct {
	initialized bool
	x           int
	y           int
	size        int
	color       eff.Color
}

func (m *mouseBox) Init(c eff.Canvas) {
	m.size = startSize
	m.color = eff.RandomColor()
	m.initialized = true
}

func (m *mouseBox) Initialized() bool {
	return m.initialized
}

func (m *mouseBox) Draw(c eff.Canvas) {
	r := eff.Rect{
		X: m.x,
		Y: m.y,
		W: m.size,
		H: m.size,
	}

	c.FillRect(r, m.color)
}

func (m *mouseBox) Update(c eff.Canvas) {

}

func main() {
	canvas := sdl.NewCanvas("Mouse Events", 800, 540, eff.Black(), 60, true)
	canvas.Run(func() {
		mb := mouseBox{}
		canvas.AddDrawable(&mb)

		canvas.AddMouseDownHandler(func(leftState bool, middleState bool, rightState bool) {
			canvas.SetClearColor(eff.RandomColor())
			mb.color = eff.Black()

			if leftState {
				mb.size = int(math.Min(float64(mb.size+1), float64(maxSize)))
			}

			if middleState {
				mb.size = startSize
			}

			if rightState {
				mb.size = int(math.Max(float64(mb.size-1), float64(minSize)))
			}

		})

		canvas.AddMouseUpHandler(func(leftState bool, middleState bool, rightState bool) {
			canvas.SetClearColor(eff.Black())
			mb.color = eff.RandomColor()
		})

		canvas.AddMouseMoveHandler(func(x int, y int) {
			mb.x = x
			mb.y = y
		})

		canvas.AddMouseWheelHandler(func(x int, y int) {
			mb.color = eff.RandomColor()
		})
	})

}
