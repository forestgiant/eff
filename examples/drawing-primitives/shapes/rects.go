package shapes

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type Rects struct {
	eff.Shape
	rects []eff.Rect
}

func (r *Rects) randomRects(count int, maxX int, maxY int) []eff.Rect {
	var rects []eff.Rect
	for i := 0; i < count; i++ {
		r := eff.Rect{X: rand.Intn(maxX), Y: rand.Intn(maxY), W: rand.Intn(maxX / 2), H: rand.Intn(maxY / 2)}
		rects = append(rects, r)
	}

	return rects
}

func (r *Rects) Init(width int, height int) {
	numRects := (width * height) / 1000
	r.rects = r.randomRects(numRects, width, height)
	r.SetUpdateHandler(func() {
		for i := range r.rects {
			r.rects[i].X = rand.Intn(width)
			r.rects[i].Y = rand.Intn(height)
			r.rects[i].W = rand.Intn(width / 2)
			r.rects[i].H = rand.Intn(height / 2)
		}
		r.Clear()
		r.StrokeRects(r.rects, eff.RandomColor())
	})
}
