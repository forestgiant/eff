package shapes

import (
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
)

type SquareSpiral struct {
	eff.Shape

	color        eff.Color
	linePoints   []eff.Point
	renderPoints []eff.Point
	tweener      tween.Tweener
}

func (s *SquareSpiral) Init(width int, height int) {
	turnCount := 10
	spiralSize := 5
	separation := 5

	bottomLeft := eff.Point{
		X: (width - spiralSize) / 2,
		Y: (height-spiralSize)/2 + spiralSize,
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

	s.SetUpdateHandler(func() {
		s.tweener.Tween()
		s.Clear()
		s.DrawLines(s.renderPoints, s.color)
	})
}
