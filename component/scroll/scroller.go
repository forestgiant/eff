package scroll

import (
	"fmt"

	"github.com/forestgiant/eff"
)

type Scroller struct {
	eff.Shape

	content *eff.Shape
}

func (s *Scroller) init(content *eff.Shape, r eff.Rect, c eff.Canvas) {
	s.content = content
	s.content.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: s.content.Rect().W,
		H: s.content.Rect().H,
	})
	s.SetRect(r)

	barWidth := 30
	scrollBar := NewScrollBar(eff.Rect{
		X: s.Rect().W - barWidth,
		Y: 0,
		W: barWidth,
		H: s.Rect().H,
	}, c)
	scrollBar.SetBackgroundColor(eff.Black())

	s.AddChild(content)
	s.AddChild(scrollBar)

	scrollBar.OnScrollHandler = func(p float64) {
		heightDiff := content.Rect().H - s.Rect().H
		s.content.SetRect(eff.Rect{
			X: 0,
			Y: int(-1 * float64(heightDiff) * p),
			W: s.content.Rect().W,
			H: s.content.Rect().H,
		})

	}
}

func NewScroller(content *eff.Shape, r eff.Rect, c eff.Canvas) *Scroller {
	s := &Scroller{}
	s.init(content, r, c)
	// s.SetRect(r)
	fmt.Println(s.Rect())
	return s
}
