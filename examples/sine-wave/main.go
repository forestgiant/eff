package main

import (
	"math"
	"os"

	"github.com/forestgiant/eff"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cols         = 50
	rows         = 50
	windowWidth  = 1280
	windowHeight = 720
	numPoints    = (cols-1)*windowHeight + (rows-1)*windowWidth
)

type sineWaveDrawable struct {
	gridColor      eff.Color
	gridPoints     []eff.Point
	origGridPoints []eff.Point
	canvas         eff.Canvas

	tx       float32
	ty       float32
	xFreq    float32
	yFreq    float32
	xFreqDir float32
	yFreqDir float32
}

func (s *sineWaveDrawable) init() {
	s.tx = math.Pi / 9
	s.ty = math.Pi / 4
	s.xFreq = 1
	s.yFreq = 1
	s.xFreqDir = 1
	s.yFreqDir = 1

	var gridPoints = make([]sdl.Point, numPoints)
	var origGridPoints = make([]sdl.Point, numPoints)
	index := 0
	cellWidth := math.Ceil(float64(windowWidth) / float64(cols))
	cellHeight := math.Ceil(float64(windowHeight) / float64(rows))
	// Create Columns
	for i := 1; i < cols-1; i++ {
		x := i * int(cellWidth)
		for j := 0; j < windowHeight; j++ {
			gridPoints[index] = sdl.Point{X: int32(x), Y: int32(j)}
			origGridPoints[index] = sdl.Point{X: int32(x), Y: int32(j)}
			index++
		}
	}

	// Create Rows
	for i := 1; i < rows-1; i++ {
		y := i * int(cellHeight)
		for j := 0; j < windowWidth; j++ {
			gridPoints[index] = sdl.Point{X: int32(j), Y: int32(y)}
			origGridPoints[index] = sdl.Point{X: int32(j), Y: int32(y)}
			index++
		}
	}
}

func (s *sineWaveDrawable) draw() {
	s.canvas.DrawPoints(&s.gridPoints, s.gridColor)
}

func (s *sineWaveDrawable) update() {

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
		newX, newY := sineWaveDistortXY(point.X, point.Y, windowWidth, windowHeight)
		s.gridPoints[i] = eff.Point{X: newX, Y: newY}
	}

	updateDistortionState()
}

func main() {
	s := sineWaveDrawable{}
	os.Exit(s.canvas.Run())
}
