package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
	"github.com/forestgiant/eff/util"
)

const (
	windowW = 1024
	windowH = 768
)

type grid struct {
	eff.Shape
	rows    int
	cols    int
	padding int
}

func (g *grid) rectForIndex(index int) eff.Rect {
	roundDivide := func(v1 int, v2 int) int {
		return util.RoundToInt(float64(v1) / float64(v2))
	}

	r := eff.Rect{}
	if g.rows == 0 || g.cols == 0 {
		return r
	}

	row := index / g.cols
	col := index % g.cols
	cellWidth := roundDivide(g.Rect().W-(g.padding*(g.cols+1)), g.cols)
	cellHeight := roundDivide(g.Rect().H-(g.padding*(g.rows+1)), g.rows)

	r.X = col*cellWidth + (g.padding * (col + 1))
	r.Y = row*cellHeight + (g.padding * (row + 1))
	r.W = cellWidth
	r.H = cellHeight
	return r
}

func (g *grid) AddChild(c eff.Drawable) error {
	c.SetRect(g.rectForIndex(len(g.Children())))
	return g.Shape.AddChild(c)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	canvas := sdl.NewCanvas("Grid", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		g := &grid{}
		g.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: 800,
			H: 600,
		})
		g.SetBackgroundColor(eff.Black())
		g.rows = 8
		g.cols = 6
		g.padding = 10
		canvas.AddChild(g)
		for i := 0; i < g.rows*g.cols; i++ {
			s := &eff.Shape{}
			s.SetBackgroundColor(eff.RandomColor())
			// s.SetBackgroundColor(eff.Black())
			s.SetRect(eff.Rect{X: 0, Y: 0, W: 100, H: 100})
			g.AddChild(s)
		}

	})
}
