package drawables

import (
	"math"
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
)

type twoPositionDot struct {
	eff.Color
	p1 eff.Point
	p2 eff.Point
}

func (dot *twoPositionDot) linearInterpolate(normalizedPercentage float64) eff.Point {
	xDiff := dot.p2.X - dot.p1.X
	yDiff := dot.p2.Y - dot.p1.Y

	return eff.Point{
		X: dot.p1.X + int(math.Ceil((float64(xDiff) * normalizedPercentage))),
		Y: dot.p1.Y + int(math.Ceil((float64(yDiff) * normalizedPercentage))),
	}
}

type CircleDots struct {
	dots        []twoPositionDot
	colorDots   []eff.ColorPoint
	tweener     tween.Tweener
	initialized bool
}

func (dot *CircleDots) Init(canvas eff.Canvas) {
	dotCount := (canvas.Width() * canvas.Height()) / 100

	pointOnCirlce := func(radius int, index int, totalPoints int, w int, h int) eff.Point {
		cx := w / 2
		cy := h / 2
		angle := math.Pi * 2 * (float64(index) / float64(totalPoints))
		return eff.Point{
			X: cx + int(float64(radius)*math.Cos(angle)),
			Y: cy + int(float64(radius)*math.Sin(angle)),
		}
	}

	for i := 0; i < dotCount; i++ {
		d := twoPositionDot{
			Color: eff.RandomColor(),
			p1: eff.Point{
				X: rand.Intn(canvas.Width()),
				Y: rand.Intn(canvas.Height()),
			},
			p2: pointOnCirlce(100, i, dotCount, canvas.Width(), canvas.Height()),
		}

		dot.dots = append(dot.dots, d)
		colorDot := eff.ColorPoint{
			Point: d.linearInterpolate(0),
			Color: eff.RandomColor(),
		}
		dot.colorDots = append(dot.colorDots, colorDot)
	}

	dot.tweener = tween.NewTweener(time.Second*2, func(progress float64) {
		for i := range dot.colorDots {
			dot.colorDots[i].Point = dot.dots[i].linearInterpolate(progress)
		}
	}, true, true, nil, nil)

	dot.initialized = true
}

func (dot *CircleDots) Draw(canvas eff.Canvas) {
	canvas.DrawColorPoints(dot.colorDots)
}

func (dot *CircleDots) Update(canvas eff.Canvas) {
	dot.tweener.Tween()
}

func (dot *CircleDots) Initialized() bool {
	return dot.initialized
}
