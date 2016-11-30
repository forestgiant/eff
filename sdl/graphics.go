package sdl

import (
	"errors"
	"fmt"

	"github.com/forestgiant/eff"
)

type Graphics struct {
	renderer *renderer
}

// DrawPoint draw a point on the screen specifying what color
func (graphics *Graphics) DrawPoint(point eff.Point, color eff.Color) {
	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		graphics.renderer.drawPoint(point.X, point.Y)
	}
}

func (graphics *Graphics) DrawPoints(points []eff.Point, color eff.Color) {
	var sdlPoints []point

	for _, p := range points {
		sdlPoints = append(sdlPoints, point{X: int32(float64(p.X)), Y: int32(float64(p.Y))})
	}

	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.drawPoints(sdlPoints)
	}
}

// DrawColorPoints draw a slide of colorPoints on the screen
func (graphics *Graphics) DrawColorPoints(colorPoints []eff.ColorPoint) {
	mainThread <- func() {
		for _, colorPoint := range colorPoints {
			colorPoint.X = int(float64(colorPoint.X))
			colorPoint.Y = int(float64(colorPoint.Y))

			graphics.renderer.setDrawColor(
				uint8(colorPoint.R),
				uint8(colorPoint.G),
				uint8(colorPoint.B),
				uint8(colorPoint.A),
			)

			graphics.renderer.drawPoint(colorPoint.X, colorPoint.Y)
		}
	}
}

// FillRect draw a filled in rectangle to the screen
func (graphics *Graphics) FillRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(float64(r.X)),
		Y: int32(float64(r.Y)),
		W: int32(float64(r.W)),
		H: int32(float64(r.H)),
	}

	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.fillRect(&sdlRect)
	}
}

// FillRects draw a slice of filled rectangles to the screen all the same color
func (graphics *Graphics) FillRects(rects []eff.Rect, color eff.Color) {
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
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.fillRects(sdlRects)
	}
}

// DrawRect draw an outlined rectangle to the screen with a color
func (graphics *Graphics) DrawRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(float64(r.X)),
		Y: int32(float64(r.Y)),
		W: int32(float64(r.W)),
		H: int32(float64(r.H)),
	}

	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.drawRect(&sdlRect)
	}
}

// DrawColorRects draw a slice of color rectangles to the screen
func (graphics *Graphics) DrawColorRects(colorRects []eff.ColorRect) {
	mainThread <- func() {
		for _, colorRect := range colorRects {
			graphics.renderer.setDrawColor(
				uint8(colorRect.R),
				uint8(colorRect.G),
				uint8(colorRect.B),
				uint8(colorRect.A),
			)

			sdlRect := rect{
				X: int32(float64(colorRect.X)),
				Y: int32(float64(colorRect.Y)),
				W: int32(float64(colorRect.W)),
				H: int32(float64(colorRect.H)),
			}

			graphics.renderer.fillRect(&sdlRect)
		}
	}
}

// DrawRects draw a slice of rectangles to the screen all the same color
func (graphics *Graphics) DrawRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []rect

	for _, r := range rects {
		r := rect{
			X: int32(float64(r.X)),
			Y: int32(float64(r.Y)),
			W: int32(float64(r.W)),
			H: int32(float64(r.H)),
		}

		sdlRects = append(sdlRects, r)
	}

	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.drawRects(sdlRects)
	}
}

// DrawLine draw a line of to the screen with a color
func (graphics *Graphics) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		graphics.renderer.drawLine(
			int(float64(p1.X)),
			int(float64(p1.Y)),
			int(float64(p2.X)),
			int(float64(p2.Y)),
		)
	}
}

// DrawLines a slice of lines to the screen all the same color
func (graphics *Graphics) DrawLines(points []eff.Point, color eff.Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []point

	for _, p := range points {
		p := point{X: int32(float64(p.X)), Y: int32(float64(p.Y))}
		sdlPoints = append(sdlPoints, p)
	}

	mainThread <- func() {
		graphics.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		graphics.renderer.drawLines(sdlPoints)
	}
}

// DrawText draws a string using a font to the screen, the point is the upper left hand corner
func (graphics *Graphics) DrawText(font eff.Font, text string, col eff.Color, point eff.Point) error {
	point.X = int(float64(point.X))
	point.Y = int(float64(point.Y))
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
		t, err := graphics.renderer.createTextureFromSurface(s)

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

		err = graphics.renderer.renderCopy(t, r1, r)
		if err != nil {
			fmt.Println(err)
		}

		t.destroy()
	}

	return nil
}

// GetTextSize this uses the currently set font to determine the size of string rendered with that font, does not actually add the text to the canvas
func (graphics *Graphics) GetTextSize(font eff.Font, text string) (int, int, error) {
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
		p.X = int32(float64(w))
		p.Y = int32(float64(h))

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
