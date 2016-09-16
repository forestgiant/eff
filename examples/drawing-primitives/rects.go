package main

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type rects struct {
	rects       []eff.Rect
	initialized bool
}

func (r *rects) randomRects(count int, maxX int, maxY int) *[]eff.Rect {
	rects := make([]eff.Rect, count)
	for i := 0; i < count; i++ {
		rects[i] = eff.Rect{X: rand.Intn(maxX), Y: rand.Intn(maxY), W: rand.Intn(maxX / 2), H: rand.Intn(maxY / 2)}
	}

	return &rects
}

func (r *rects) Init(canvas eff.Canvas) {
	r.rects = *r.randomRects(100, canvas.Width(), canvas.Height())
	r.initialized = true
}

func (r *rects) Draw(canvas eff.Canvas) {
	canvas.DrawRects(&r.rects, eff.Color{}.RandomColor())
}

func (r *rects) Update(canvas eff.Canvas) {
	updateRandomRects := func() {
		for i := range r.rects {
			r.rects[i].X = rand.Intn(canvas.Width())
			r.rects[i].Y = rand.Intn(canvas.Height())
			r.rects[i].W = rand.Intn(canvas.Width() / 2)
			r.rects[i].H = rand.Intn(canvas.Height() / 2)
		}
	}

	updateRandomRects()
}

func (r *rects) Initialized() bool {
	return r.initialized
}
