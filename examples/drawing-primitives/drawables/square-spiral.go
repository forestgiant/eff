package drawables

import (
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
)

type SquareSpiral struct {
	color        eff.Color
	linePoints   []eff.Point
	renderPoints []eff.Point
	// t            float64
	tweener     tween.Tweener
	initialized bool
}

func (s *SquareSpiral) Init(canvas eff.Canvas) {
	turnCount := 10
	spiralSize := 5
	separation := 5

	bottomLeft := eff.Point{
		X: (canvas.Width() - spiralSize) / 2,
		Y: (canvas.Height()-spiralSize)/2 + spiralSize,
	}

	s.linePoints = append(s.linePoints, bottomLeft)

	for i := 0; i < (turnCount * 4); i++ {
		upperLeft := eff.Point{
			X: bottomLeft.X,
			Y: bottomLeft.Y - (spiralSize),
		}

		upperRight := eff.Point{
			X: upperLeft.X + (spiralSize),
			Y: upperLeft.Y,
		}

		spiralSize += separation

		bottomRight := eff.Point{
			X: upperRight.X,
			Y: upperRight.Y + (spiralSize),
		}

		bottomLeft = eff.Point{
			X: bottomRight.X - (spiralSize),
			Y: bottomRight.Y,
		}

		spiralSize += separation

		s.linePoints = append(s.linePoints, upperLeft)
		s.linePoints = append(s.linePoints, upperRight)
		s.linePoints = append(s.linePoints, bottomRight)
		s.linePoints = append(s.linePoints, bottomLeft)
	}
	s.color = eff.RandomColor()

	s.tweener = tween.NewTweener(time.Second*10, func(progress float64) {
		percentage := float64(len(s.linePoints)) * progress
		index := int(percentage)
		diff := percentage - float64(index)

		s.renderPoints = make([]eff.Point, index)
		copy(s.renderPoints, s.linePoints[:index])

		if diff > 0 {
			if index < len(s.linePoints)-1 && index > 0 {
				lastPoint := s.linePoints[index-1]
				nextPoint := s.linePoints[index]

				newPoint := eff.Point{
					X: lastPoint.X + int(float64(nextPoint.X-lastPoint.X)*diff),
					Y: lastPoint.Y + int(float64(nextPoint.Y-lastPoint.Y)*diff),
				}

				s.renderPoints = append(s.renderPoints, newPoint)
			}
		}
	}, true, true, func() {
		s.color = eff.RandomColor()
	}, nil)

	s.initialized = true
}

func (s *SquareSpiral) Draw(canvas eff.Canvas) {
	canvas.DrawLines(s.renderPoints, s.color)
}

func (s *SquareSpiral) Update(canvas eff.Canvas) {
	s.tweener.Tween()
}

func (s *SquareSpiral) Initialized() bool {
	return s.initialized
}
