package shapes

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
	eff.Shape

	dots    []twoPositionDot
	colors  []eff.Color
	points  []eff.Point
	tweener tween.Tweener
}

func (dot *CircleDots) Init(width int, height int) {
	dotCount := (width * height) / 20

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
				X: rand.Intn(width),
				Y: rand.Intn(height),
			},
			p2: pointOnCirlce(100, i, dotCount, width, height),
		}

		dot.dots = append(dot.dots, d)
		dot.points = append(dot.points, d.linearInterpolate(0))
		dot.colors = append(dot.colors, eff.RandomColor())
	}

	dot.tweener = tween.NewTweener(time.Second*2, func(progress float64) {
		for i := range dot.points {
			dot.points[i] = dot.dots[i].linearInterpolate(progress)
		}
	}, true, true, nil, nil)

	dot.SetUpdateHandler(func() {
		dot.tweener.Tween()
		dot.Clear()
		dot.DrawColorPoints(dot.points, dot.colors)
	})
}
