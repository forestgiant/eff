package scroll

import (
	"math"

	"github.com/forestgiant/eff"
)

type nub struct {
	eff.Shape

	mouseDown bool
	mouseOver bool
	scrolling bool
	onNubMove func()
}

func (n *nub) Hitbox() eff.Rect {
	return n.ParentOffsetRect()
}

func (n *nub) MouseDown(leftState bool, middleState bool, rightState bool) {
	n.mouseDown = true
	n.scrolling = !n.scrolling
}

func (n *nub) MouseUp(leftState bool, middleState bool, rightState bool) {
	if n.mouseDown {
		n.mouseDown = false
	}
}

func (n *nub) MouseOver() {
	n.mouseOver = true
}

func (n *nub) MouseOut() {
	n.mouseOver = false
}

func (n *nub) IsMouseOver() bool { return n.mouseOver }

type ScrollBar struct {
	eff.Shape

	OnScrollHandler func(float64)
}

func (s *ScrollBar) init(c eff.Canvas) {
	nub := &nub{}
	nub.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: s.Rect().W,
		H: s.Rect().H / 10,
	})
	nub.SetBackgroundColor(eff.Color{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF})
	s.AddChild(nub)

	c.AddClickable(nub)
	c.AddMouseMoveHandler(func(x int, y int) {
		if !nub.scrolling {
			return
		}
		pRect := s.ParentOffsetRect()

		minY := pRect.Y
		maxY := minY + s.Rect().H

		y = int(math.Max(float64(minY), float64(y)))
		y = int(math.Min(float64(maxY), float64(y)))

		y -= minY
		percentage := float64(y) / float64(s.Rect().H)
		// fmt.Println(percentage)
		nubPercentage := float64(nub.Rect().H) / float64(s.Rect().H)
		if percentage <= 1-nubPercentage {
			nub.SetRect(eff.Rect{
				X: 0,
				Y: int(percentage * float64(s.Rect().H)),
				W: nub.Rect().W,
				H: nub.Rect().H,
			})
		}

		if s.OnScrollHandler != nil {
			s.OnScrollHandler(percentage)
		}

	})

	c.AddMouseUpHandler(func(leftState bool, middleState bool, rightState bool) {
		nub.scrolling = false
	})
}

func NewScrollBar(rect eff.Rect, c eff.Canvas) *ScrollBar {
	scrollBar := &ScrollBar{}
	scrollBar.SetRect(rect)
	scrollBar.init(c)
	return scrollBar
}
