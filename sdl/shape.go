package sdl

import (
	"errors"
	"fmt"

	"github.com/forestgiant/eff"
)

type Shape struct {
	Container

	rect     eff.Rect
	parent   eff.Container
	scale    float64
	renderer *renderer
}

func (shape *Shape) Draw(canvas eff.Canvas) {

}

func (shape *Shape) Rect() eff.Rect {
	return shape.rect
}

func (shape *Shape) SetParent(c eff.Container) {
	shape.parent = c
}

func (shape *Shape) Parent() eff.Container {
	return shape.parent
}

// DrawPoint draw a point on the screen specifying what color
func (shape *Shape) DrawPoint(point eff.Point, color eff.Color) {
	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		shape.renderer.drawPoint(point.X, point.Y)
	}
}

func (shape *Shape) DrawPoints(points []eff.Point, color eff.Color) {
	var sdlPoints []point

	for _, p := range points {
		sdlPoints = append(sdlPoints, point{X: int32(float64(p.X) * shape.scale), Y: int32(float64(p.Y) * shape.scale)})
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.drawPoints(sdlPoints)
	}
}

// DrawColorPoints draw a slide of colorPoints on the screen
func (shape *Shape) DrawColorPoints(colorPoints []eff.ColorPoint) {
	mainThread <- func() {
		for _, colorPoint := range colorPoints {
			colorPoint.X = int(float64(colorPoint.X) * shape.scale)
			colorPoint.Y = int(float64(colorPoint.Y) * shape.scale)

			shape.renderer.setDrawColor(
				uint8(colorPoint.R),
				uint8(colorPoint.G),
				uint8(colorPoint.B),
				uint8(colorPoint.A),
			)

			shape.renderer.drawPoint(colorPoint.X, colorPoint.Y)
		}
	}
}

// FillRect draw a filled in rectangle to the screen
func (shape *Shape) FillRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(float64(r.X) * shape.scale),
		Y: int32(float64(r.Y) * shape.scale),
		W: int32(float64(r.W) * shape.scale),
		H: int32(float64(r.H) * shape.scale),
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.fillRect(&sdlRect)
	}
}

// FillRects draw a slice of filled rectangles to the screen all the same color
func (shape *Shape) FillRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []rect

	for _, r := range rects {
		sdlRects = append(sdlRects,
			rect{
				X: int32(r.X),
				Y: int32(r.Y),
				W: int32(r.W),
				H: int32(r.H),
			},
		)
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.fillRects(sdlRects)
	}
}

// DrawRect draw an outlined rectangle to the screen with a color
func (shape *Shape) DrawRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(float64(r.X) * shape.scale),
		Y: int32(float64(r.Y) * shape.scale),
		W: int32(float64(r.W) * shape.scale),
		H: int32(float64(r.H) * shape.scale),
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.drawRect(&sdlRect)
	}
}

// DrawColorRects draw a slice of color rectangles to the screen
func (shape *Shape) DrawColorRects(colorRects []eff.ColorRect) {
	mainThread <- func() {
		for _, colorRect := range colorRects {
			shape.renderer.setDrawColor(
				uint8(colorRect.R),
				uint8(colorRect.G),
				uint8(colorRect.B),
				uint8(colorRect.A),
			)

			sdlRect := rect{
				X: int32(float64(colorRect.X) * shape.scale),
				Y: int32(float64(colorRect.Y) * shape.scale),
				W: int32(float64(colorRect.W) * shape.scale),
				H: int32(float64(colorRect.H) * shape.scale),
			}

			shape.renderer.fillRect(&sdlRect)
		}
	}
}

// DrawRects draw a slice of rectangles to the screen all the same color
func (shape *Shape) DrawRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []rect

	for _, r := range rects {
		r := rect{
			X: int32(float64(r.X) * shape.scale),
			Y: int32(float64(r.Y) * shape.scale),
			W: int32(float64(r.W) * shape.scale),
			H: int32(float64(r.H) * shape.scale),
		}

		sdlRects = append(sdlRects, r)
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.drawRects(sdlRects)
	}
}

// DrawLine draw a line of to the screen with a color
func (shape *Shape) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		shape.renderer.drawLine(
			int(float64(p1.X)*shape.scale),
			int(float64(p1.Y)*shape.scale),
			int(float64(p2.X)*shape.scale),
			int(float64(p2.Y)*shape.scale),
		)
	}
}

// DrawLines a slice of lines to the screen all the same color
func (shape *Shape) DrawLines(points []eff.Point, color eff.Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []point

	for _, p := range points {
		p := point{X: int32(float64(p.X) * shape.scale), Y: int32(float64(p.Y) * shape.scale)}
		sdlPoints = append(sdlPoints, p)
	}

	mainThread <- func() {
		shape.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		shape.renderer.drawLines(sdlPoints)
	}
}

// DrawText draws a string using a font to the screen, the point is the upper left hand corner
func (shape *Shape) DrawText(font eff.Font, text string, col eff.Color, point eff.Point) error {
	point.X = int(float64(point.X) * shape.scale)
	point.Y = int(float64(point.Y) * shape.scale)
	f := font.(*Font)
	if f.sdlFont == nil {
		return errors.New("Can't draw text no font set")
	}

	rgba := color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}

	mainThread <- func() {
		s, err := renderTextBlended(f, text, rgba)
		if err != nil {
			fmt.Println(err)
		}
		t, err := shape.renderer.createTextureFromSurface(s)

		if err != nil {
			fmt.Println(err)
		}

		r1 := rect{
			X: 0,
			Y: 0,
			W: int32(s.w),
			H: int32(s.h),
		}

		r := rect{
			X: int32(point.X),
			Y: int32(point.Y),
			W: int32(s.w),
			H: int32(s.h),
		}

		freeSurface(s)

		err = shape.renderer.renderCopy(t, r1, r)
		if err != nil {
			fmt.Println(err)
		}

		t.destroy()
	}

	return nil
}

// GetTextSize this uses the currently set font to determine the size of string rendered with that font, does not actually add the text to the canvas
func (shape *Shape) GetTextSize(font eff.Font, text string) (int, int, error) {
	f := font.(*Font)
	if f.sdlFont == nil {
		return -1, -1, errors.New("Can't get text size font not loaded")
	}

	errChan := make(chan error)
	sizeChan := make(chan point)

	mainThread <- func() {
		w, h, err := sizeText(f, text)
		if err != nil {
			errChan <- err
		}

		p := point{}
		p.X = int32(float64(w) / shape.scale)
		p.Y = int32(float64(h) / shape.scale)

		sizeChan <- p
	}

	for {
		select {
		case e := <-errChan:
			return -1, -1, e
		case p := <-sizeChan:
			return int(p.X), int(p.Y), nil
		}
	}
}
