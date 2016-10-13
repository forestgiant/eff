package main

import "github.com/forestgiant/eff"

type squareSpiral struct {
	color        eff.Color
	linePoints   []eff.Point
	renderPoints []eff.Point
	t            float64
	initialized  bool
}

func (s *squareSpiral) Init(canvas eff.Canvas) {
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

	s.initialized = true
}

func (s *squareSpiral) Draw(canvas eff.Canvas) {
	canvas.DrawLines(s.renderPoints, s.color)
}

func (s *squareSpiral) Update(canvas eff.Canvas) {
	s.t += 0.0006
	if s.t > 1 {
		s.t = 0
		s.color = eff.RandomColor()
	}

	percentage := float64(len(s.linePoints)) * s.t
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

}

func (s *squareSpiral) Initialized() bool {
	return s.initialized
}
