package sdl

import (
	"errors"
	"fmt"

	"github.com/forestgiant/eff"
)

type Graphics struct {
	renderer *renderer
	scale    float64
}

func NewGraphics() *Graphics {
	g := Graphics{}

	g.scale = 1

	return &g
}

// DrawPoint draw a point on the screen specifying what color
func (graphics *Graphics) DrawPoint(point eff.Point, color eff.Color) {
	point = point.Scale(graphics.scale)
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
	scale := graphics.scale
	for _, p := range points {
		sdlPoints = append(sdlPoints, point{X: int32(float64(p.X) * scale), Y: int32(float64(p.Y) * scale)})
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
func (graphics *Graphics) DrawColorPoints(points []eff.Point, colors []eff.Color) {
	if len(points) != len(colors) {
		fmt.Println("length of points and length of colors mismatch")
		return
	}

	scale := graphics.scale

	mainThread <- func() {
		for i := range points {
			p := points[i]
			c := colors[i]
			graphics.renderer.setDrawColor(
				uint8(c.R),
				uint8(c.G),
				uint8(c.B),
				uint8(c.A),
			)

			graphics.renderer.drawPoint(int(float64(p.X)*scale), int(float64(p.Y)*scale))
		}
	}
}

// DrawLine draw a line of to the screen with a color
func (graphics *Graphics) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	p1 = p1.Scale(graphics.scale)
	p2 = p2.Scale(graphics.scale)

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
	scale := graphics.scale
	for _, p := range points {
		p := point{X: int32(float64(p.X) * scale), Y: int32(float64(p.Y) * scale)}
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

// StrokeRect draw an outlined rectangle to the screen with a color
func (graphics *Graphics) StrokeRect(r eff.Rect, color eff.Color) {
	scale := graphics.scale
	sdlRect := rect{
		X: int32(float64(r.X) * scale),
		Y: int32(float64(r.Y) * scale),
		W: int32(float64(r.W) * scale),
		H: int32(float64(r.H) * scale),
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

// StrokeRects draw a slice of rectangles to the screen all the same color
func (graphics *Graphics) StrokeRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []rect
	scale := graphics.scale
	for _, r := range rects {
		r := rect{
			X: int32(float64(r.X) * scale),
			Y: int32(float64(r.Y) * scale),
			W: int32(float64(r.W) * scale),
			H: int32(float64(r.H) * scale),
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

// StrokeColorRects draw a slice of color rectangles to the screen
func (graphics *Graphics) StrokeColorRects(rects []eff.Rect, colors []eff.Color) {
	if len(rects) != len(colors) {
		fmt.Println("length of rects and length of colors mismatch")
		return
	}
	scale := graphics.scale
	mainThread <- func() {
		for i := range rects {
			r := rects[i]
			c := colors[i]
			graphics.renderer.setDrawColor(
				uint8(c.R),
				uint8(c.G),
				uint8(c.B),
				uint8(c.A),
			)

			sdlRect := rect{
				X: int32(float64(r.X) * scale),
				Y: int32(float64(r.Y) * scale),
				W: int32(float64(r.W) * scale),
				H: int32(float64(r.H) * scale),
			}

			graphics.renderer.drawRect(&sdlRect)
		}
	}
}

// FillRect draw a filled in rectangle to the screen
func (graphics *Graphics) FillRect(r eff.Rect, color eff.Color) {
	scale := graphics.scale
	sdlRect := rect{
		X: int32(float64(r.X) * scale),
		Y: int32(float64(r.Y) * scale),
		W: int32(float64(r.W) * scale),
		H: int32(float64(r.H) * scale),
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
	scale := graphics.scale
	for _, r := range rects {
		sdlRects = append(sdlRects,
			rect{
				X: int32(float64(r.X) * scale),
				Y: int32(float64(r.Y) * scale),
				W: int32(float64(r.W) * scale),
				H: int32(float64(r.H) * scale),
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

// FillColorRects draw a slice of color rectangles to the screen
func (graphics *Graphics) FillColorRects(rects []eff.Rect, colors []eff.Color) {
	if len(rects) != len(colors) {
		fmt.Println("length of rects and length of colors mismatch")
		return
	}
	scale := graphics.scale
	mainThread <- func() {
		for i := range rects {
			r := rects[i]
			c := colors[i]
			graphics.renderer.setDrawColor(
				uint8(c.R),
				uint8(c.G),
				uint8(c.B),
				uint8(c.A),
			)

			sdlRect := rect{
				X: int32(float64(r.X) * scale),
				Y: int32(float64(r.Y) * scale),
				W: int32(float64(r.W) * scale),
				H: int32(float64(r.H) * scale),
			}

			graphics.renderer.fillRect(&sdlRect)
		}
	}
}

// DrawText draws a string using a font to the screen, the point is the upper left hand corner
func (graphics *Graphics) DrawText(font eff.Font, text string, col eff.Color, point eff.Point) error {
	point.X = int(float64(point.X) * graphics.scale)
	point.Y = int(float64(point.Y) * graphics.scale)
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
