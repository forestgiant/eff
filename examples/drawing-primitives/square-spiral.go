package main

import "github.com/forestgiant/eff"

type squareSpiral struct {
	initialized bool
	color       eff.Color
	linePoints  []eff.Point
}

func (s *squareSpiral) Init(canvas eff.Canvas) {
	turnCount := 20
	spiralSize := 5
	separation := 4

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
	s.color = eff.Color{}.RandomColor()
	s.initialized = true
}

func (s *squareSpiral) Draw(canvas eff.Canvas) {
	canvas.DrawLines(s.linePoints, s.color)
}

func (s *squareSpiral) Update(canvas eff.Canvas) {

}

func (s *squareSpiral) Initialized() bool {
	return s.initialized
}
