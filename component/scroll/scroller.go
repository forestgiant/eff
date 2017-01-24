package scroll

import "github.com/forestgiant/eff"

type Scroller struct {
	eff.Shape

	content          eff.Drawable
	scrollBarVisibie bool
}

func (s *Scroller) init(content eff.Drawable, r eff.Rect, c eff.Canvas) {
	s.SetBackgroundColor(eff.White())
	s.content = content
	s.content.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: s.content.Rect().W,
		H: s.content.Rect().H,
	})
	s.SetRect(r)

	barWidth := 10
	scrollBar := NewScrollBar(eff.Rect{
		X: s.Rect().W - barWidth,
		Y: 0,
		W: barWidth,
		H: s.Rect().H,
	}, c)
	scrollBar.SetBackgroundColor(eff.Color{R: 0x00, G: 0x00, B: 0x00, A: 0x66})
	s.content.SetResizeHandler(func() {
		if s.content.Rect().H > s.Rect().H && !s.scrollBarVisibie {
			s.AddChild(scrollBar)
			s.scrollBarVisibie = true
		} else if s.content.Rect().H < s.Rect().H && s.scrollBarVisibie {
			s.RemoveChild(scrollBar)
			s.scrollBarVisibie = false
			s.content.SetRect(eff.Rect{
				X: 0,
				Y: 0,
				W: s.content.Rect().W,
				H: s.content.Rect().H,
			})
		}
	})

	s.AddChild(content)
	if s.content.Rect().H > s.Rect().H && !s.scrollBarVisibie {
		s.AddChild(scrollBar)
		s.scrollBarVisibie = true
	}

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

func NewScroller(content eff.Drawable, r eff.Rect, c eff.Canvas) *Scroller {
	s := &Scroller{}
	s.init(content, r, c)
	// s.SetRect(r)
	// fmt.Println(s.Rect())
	return s
}
