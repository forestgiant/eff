package eff

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Wraps the Drawable to track whether or not it has been initialized
type sdlDrawable struct {
	initialized bool
	drawable    Drawable
}

func (s *sdlDrawable) Init(canvas Canvas) {
	s.drawable.Init(canvas)
	s.initialized = true
}

func (s *sdlDrawable) Initialized() bool {
	return s.initialized
}

func (s *sdlDrawable) Draw(canvas Canvas) {
	s.drawable.Draw(canvas)
}

func (s *sdlDrawable) Update(canvas Canvas) {
	s.drawable.Update(canvas)
}

// SDLCanvas creates window and renderer and calls all drawable methods
type SDLCanvas struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	// drawablesMutex  sync.Mutex
	drawables       []*sdlDrawable
	width           int
	height          int
	fullscreen      bool
	keyUpHandlers   []KeyHandler
	keyDownHandlers []KeyHandler
}

// SetWidth set the width of the canvas, must be called prior to run
func (sdlCanvas *SDLCanvas) SetWidth(width int) {
	sdlCanvas.width = width
}

// Width get the width of the canvas window
func (sdlCanvas *SDLCanvas) Width() int {
	return sdlCanvas.width
}

// SetHeight set the height of the canvas, must be called prior to run
func (sdlCanvas *SDLCanvas) SetHeight(height int) {
	sdlCanvas.height = height
}

// Height get the height of the canvas window
func (sdlCanvas *SDLCanvas) Height() int {
	return sdlCanvas.height
}

// AddDrawable adds a struct that implements the eff.Drawable interface
func (sdlCanvas *SDLCanvas) AddDrawable(drawable Drawable) {
	sdlCanvas.drawables = append(sdlCanvas.drawables, &sdlDrawable{drawable: drawable})
}

//RemoveDrawable removes struct from canvas that implements eff.Drawable
func (sdlCanvas *SDLCanvas) RemoveDrawable(drawable Drawable) {
	index := -1
	for i, d := range sdlCanvas.drawables {
		if d.drawable == drawable {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	sdlCanvas.drawables = append(sdlCanvas.drawables[:index], sdlCanvas.drawables[index+1:]...)
}

//AddKeyUpHandler adds key up event handler to the canvas
func (sdlCanvas *SDLCanvas) AddKeyUpHandler(handler KeyHandler) {
	sdlCanvas.keyUpHandlers = append(sdlCanvas.keyUpHandlers, handler)
}

//AddKeyDownHandler adds key down event handler to the canvas
func (sdlCanvas *SDLCanvas) AddKeyDownHandler(handler KeyHandler) {
	sdlCanvas.keyDownHandlers = append(sdlCanvas.keyDownHandlers, handler)
}

// Run creates an infinite loop that renders all drawables, init is only call once and draw and update are called once per frame
func (sdlCanvas *SDLCanvas) Run() int {
	if sdlCanvas.width == 0 {
		sdlCanvas.width = defaultWidth
	}

	if sdlCanvas.height == 0 {
		sdlCanvas.height = defaultHeight
	}

	var err error
	sdl.CallQueue <- func() {
		sdlCanvas.window, err = sdl.CreateWindow(
			windowTitle,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			sdlCanvas.Width(),
			sdlCanvas.Height(),
			sdl.WINDOW_OPENGL,
		)

		if sdlCanvas.fullscreen {
			sdlCanvas.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
		} else {
			sdlCanvas.window.SetFullscreen(0)
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		sdl.CallQueue <- func() {
			sdlCanvas.window.Destroy()
		}
	}()

	sdl.CallQueue <- func() {
		sdlCanvas.renderer, err = sdl.CreateRenderer(
			sdlCanvas.window,
			-1,
			sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC,
		)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create renderer: ", err)
		return 2
	}
	defer func() {
		sdl.CallQueue <- func() {
			sdlCanvas.renderer.Destroy()
		}
	}()

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.Clear()
	}

	running := true

	var lastFrameTime = sdl.GetTicks()
	for running {
		sdl.CallQueue <- func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.KeyUpEvent:
					switch t.Keysym.Sym {
					case sdl.K_q:
						running = false
					case sdl.K_f:
						sdlCanvas.fullscreen = !sdlCanvas.fullscreen
						if sdlCanvas.fullscreen {
							sdlCanvas.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
						} else {
							sdlCanvas.window.SetFullscreen(0)
						}
					}

					for _, handler := range sdlCanvas.keyUpHandlers {
						handler(sdl.GetKeyName(t.Keysym.Sym))
					}
				case *sdl.KeyDownEvent:
					for _, handler := range sdlCanvas.keyDownHandlers {
						handler(sdl.GetKeyName(t.Keysym.Sym))
					}
				}
			}

			sdlCanvas.renderer.SetDrawColor(0x0, 0x0, 0x0, 0xFF)
			sdlCanvas.renderer.Clear()
		}

		for _, drawable := range sdlCanvas.drawables {
			if drawable.drawable == nil {
				continue
			}

			if !drawable.Initialized() {
				drawable.Init(sdlCanvas)
			}

			drawable.Draw(sdlCanvas)
			drawable.Update(sdlCanvas)
		}

		sdl.CallQueue <- func() {
			currentFrameTime := sdl.GetTicks()
			sdlCanvas.renderer.Present()
			if currentFrameTime-lastFrameTime < frameTime {
				sdl.Delay(frameTime - (currentFrameTime - lastFrameTime))
			}
			lastFrameTime = currentFrameTime
		}
	}
	return 0
}

//DrawPoints draw a slice of points to the screen all the same color
func (sdlCanvas *SDLCanvas) DrawPoints(points []Point, color Color) {
	var sdlPoints []sdl.Point

	for _, point := range points {
		sdlPoints = append(sdlPoints, sdl.Point{X: int32(point.X), Y: int32(point.Y)})
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.DrawPoints(sdlPoints)
	}
}

//DrawPoint draw a point on the screen specifying what color
func (sdlCanvas *SDLCanvas) DrawPoint(point Point, color Color) {
	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		sdlCanvas.renderer.DrawPoint(point.X, point.Y)
	}
}

//DrawColorPoints draw a slide of colorPoints on the screen
func (sdlCanvas *SDLCanvas) DrawColorPoints(colorPoints []ColorPoint) {
	sdl.CallQueue <- func() {
		for _, colorPoint := range colorPoints {
			sdlCanvas.renderer.SetDrawColor(
				uint8(colorPoint.R),
				uint8(colorPoint.G),
				uint8(colorPoint.B),
				uint8(colorPoint.A),
			)

			sdlCanvas.renderer.DrawPoint(colorPoint.X, colorPoint.Y)
		}
	}
}

//FillRect draw a filled in rectangle to the screen
func (sdlCanvas *SDLCanvas) FillRect(rect Rect, color Color) {
	sdlRect := sdl.Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.FillRect(&sdlRect)
	}
}

//FillRects draw a slice of filled rectangles to the screen all the same color
func (sdlCanvas *SDLCanvas) FillRects(rects []Rect, color Color) {
	var sdlRects []sdl.Rect

	for _, rect := range rects {
		sdlRects = append(sdlRects,
			sdl.Rect{
				X: int32(rect.X),
				Y: int32(rect.Y),
				W: int32(rect.W),
				H: int32(rect.H),
			},
		)
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.FillRects(sdlRects)
	}
}

//DrawRect draw an outlined rectangle to the screen with a color
func (sdlCanvas *SDLCanvas) DrawRect(rect Rect, color Color) {
	sdlRect := sdl.Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.DrawRect(&sdlRect)
	}
}

//DrawColorRects draw a slice of color rectangles to the screen
func (sdlCanvas *SDLCanvas) DrawColorRects(colorRects []ColorRect) {
	sdl.CallQueue <- func() {
		for _, colorRect := range colorRects {
			sdlCanvas.renderer.SetDrawColor(
				uint8(colorRect.R),
				uint8(colorRect.G),
				uint8(colorRect.B),
				uint8(colorRect.A),
			)

			sdlRect := sdl.Rect{
				X: int32(colorRect.X),
				Y: int32(colorRect.Y),
				W: int32(colorRect.W),
				H: int32(colorRect.H),
			}

			sdlCanvas.renderer.FillRect(&sdlRect)
		}
	}
}

//DrawRects draw a slice of rectangles to the screen all the same color
func (sdlCanvas *SDLCanvas) DrawRects(rects []Rect, color Color) {
	var sdlRects []sdl.Rect

	for _, rect := range rects {
		r := sdl.Rect{
			X: int32(rect.X),
			Y: int32(rect.Y),
			W: int32(rect.W),
			H: int32(rect.H),
		}

		sdlRects = append(sdlRects, r)
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.DrawRects(sdlRects)
	}
}

//DrawLine draw a line of to the screen with a color
func (sdlCanvas *SDLCanvas) DrawLine(p1 Point, p2 Point, color Color) {
	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		sdlCanvas.renderer.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	}
}

//DrawLines a slice of lines to the screen all the same color
func (sdlCanvas *SDLCanvas) DrawLines(points []Point, color Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []sdl.Point

	for _, point := range points {
		p := sdl.Point{X: int32(point.X), Y: int32(point.Y)}
		sdlPoints = append(sdlPoints, p)
	}

	sdl.CallQueue <- func() {
		sdlCanvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlCanvas.renderer.DrawLines(sdlPoints)
	}
}

//Fullscreen get the full screen state of the window
func (sdlCanvas *SDLCanvas) Fullscreen() bool {
	return sdlCanvas.fullscreen
}

//SetFullscreen set the fullscreen state of the window
func (sdlCanvas *SDLCanvas) SetFullscreen(fullscreen bool) {
	sdlCanvas.fullscreen = fullscreen
}
