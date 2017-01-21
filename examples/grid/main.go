package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/grid"
	"github.com/forestgiant/eff/component/scroll"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW = 1024
	windowH = 768
)

func main() {
	rand.Seed(time.Now().UnixNano())
	canvas := sdl.NewCanvas("Grid", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		// Grid
		g := grid.NewGrid(2, 2, 10, 0)
		g.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: windowW,
			H: 400,
		})
		g.SetBackgroundColor(eff.Black())

		canvas.AddChild(g)
		for i := 0; i < g.Rows()*g.Cols(); i++ {
			subG := grid.NewGrid((i + 2), (i + 2), 20, 0)
			subG.SetBackgroundColor(eff.RandomColor())
			g.AddChild(subG)
			for j := 0; j < subG.Rows()*subG.Cols(); j++ {
				s := &eff.Shape{}
				s.SetBackgroundColor(eff.RandomColor())
				subG.AddChild(s)
			}
		}

		// Scroller Grid
		content := grid.NewGrid(2, 3, 20, 50)
		content.SetRect(eff.Rect{X: 0, Y: 0, W: canvas.Rect().W, H: 0})
		for i := 0; i < 50; i++ {
			s := &eff.Shape{}
			s.SetBackgroundColor(eff.RandomColor())
			content.AddChild(s)
		}
		scroller := scroll.NewScroller(content, eff.Rect{X: 0, Y: 400, W: canvas.Rect().W, H: canvas.Rect().H - 400}, canvas)
		canvas.AddChild(scroller)
	})
}
