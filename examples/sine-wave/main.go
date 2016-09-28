package main

import (
	"fmt"
	"math"

	"github.com/forestgiant/eff/eff"
	"github.com/forestgiant/eff/sdl"
)

type sineWaveDrawable struct {
	gridPoints     []eff.Point
	origGridPoints []eff.Point
	tx             float32
	ty             float32
	xFreq          float32
	yFreq          float32
	xFreqDir       float32
	yFreqDir       float32
	initialized    bool
}

func (s *sineWaveDrawable) Init(canvas eff.Canvas) {
	cols := int(math.Ceil(float64(canvas.Width()) / 100))
	rows := int(math.Ceil(float64(canvas.Height()) / 100))
	fmt.Println(cols, rows)
	s.tx = math.Pi / 9
	s.ty = math.Pi / 4
	s.xFreq = 1
	s.yFreq = 1
	s.xFreqDir = 1
	s.yFreqDir = 1

	cellWidth := int(math.Ceil(float64(canvas.Width()) / float64(cols)))
	cellHeight := int(math.Ceil(float64(canvas.Height()) / float64(rows)))
	fmt.Println(cellWidth, cellHeight)
	// Create Columns
	for i := 0; i < cols-1; i++ {
		x := i*cellWidth + cellWidth
		for j := 0; j < canvas.Height(); j++ {
			s.gridPoints = append(s.gridPoints, eff.Point{X: (x), Y: (j)})
			s.origGridPoints = append(s.origGridPoints, eff.Point{X: (x), Y: (j)})
		}
	}

	// Create Rows
	for i := 0; i < rows-1; i++ {
		y := i*cellHeight + cellHeight
		for j := 0; j < canvas.Width(); j++ {
			s.gridPoints = append(s.gridPoints, eff.Point{X: (j), Y: (y)})
			s.origGridPoints = append(s.origGridPoints, eff.Point{X: (j), Y: (y)})
		}
	}
	s.initialized = true
}

func (s *sineWaveDrawable) Draw(canvas eff.Canvas) {
	color := eff.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	canvas.DrawPoints(s.gridPoints, color)
}

func (s *sineWaveDrawable) Update(canvas eff.Canvas) {
	updateDistortionState := func() {
		s.xFreq += (0.1) * s.xFreqDir
		if s.xFreq > 25 || s.xFreq < 1 {
			s.xFreqDir *= -1
		}
		s.yFreq += (0.1) * s.yFreqDir
		if s.yFreq > 30 || s.yFreq < 1 {
			s.yFreqDir *= -1
		}
	}

	sineWaveDistortXY := func(x int, y int, w int, h int) (int, int) {
		var normalizedX = float32(x) / float32(w)
		var normalizedY = float32(y) / float32(h)

		var xOffset = int(50 * (math.Sin(float64(s.xFreq*normalizedY+s.yFreq*normalizedX+2*math.Pi*s.tx)) * 0.5))
		var yOffset = int(50 * (math.Sin(float64(s.xFreq*normalizedY+s.yFreq*normalizedX+2*math.Pi*s.ty)) * 0.5))

		return x + xOffset, y + yOffset
	}

	for i, point := range s.origGridPoints {
		newX, newY := sineWaveDistortXY(point.X, point.Y, canvas.Width(), canvas.Height())
		s.gridPoints[i] = eff.Point{X: newX, Y: newY}
	}

	updateDistortionState()
}

func (s *sineWaveDrawable) Initialized() bool {
	return s.initialized
}

func main() {
	//Create drawables
	s := sineWaveDrawable{}

	//Create canvas
	canvas := sdl.Canvas{}
	canvas.SetWidth(1280)
	canvas.SetHeight(720)

	//Add drawables to canvas
	canvas.AddDrawable(&s)

	//Start the run loop
	canvas.Run()
}
