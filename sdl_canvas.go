package eff

import (
	"fmt"
	"os"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// SDLCanvas creates window and renderer and calls all drawable methods
type SDLCanvas struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	drawablesMutex  sync.Mutex
	drawables       []Drawable
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
	sdlCanvas.drawablesMutex.Lock()
	sdlCanvas.drawables = append(sdlCanvas.drawables, drawable)
	sdlCanvas.drawablesMutex.Unlock()
}

func (sdlCanvas *SDLCanvas) RemoveDrawable(drawable Drawable) {
	index := -1
	for i, d := range sdlCanvas.drawables {
		if d == drawable {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	sdlCanvas.drawablesMutex.Lock()
	sdlCanvas.drawables = append(sdlCanvas.drawables[:index], sdlCanvas.drawables[index+1:]...)
	sdlCanvas.drawablesMutex.Unlock()
}

func (sdlCanvas *SDLCanvas) AddKeyUpHandler(handler KeyHandler) {
	sdlCanvas.keyUpHandlers = append(sdlCanvas.keyUpHandlers, handler)
}

func (sdlCanvas *SDLCanvas) AddKeyDownHandler(handler KeyHandler) {
	sdlCanvas.keyUpHandlers = append(sdlCanvas.keyUpHandlers, handler)
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
						handler(sdl.GetKeyName(t.Keysym.Sym), sdlCanvas)
					}
				case *sdl.KeyDownEvent:
					for _, handler := range sdlCanvas.keyDownHandlers {
						handler(sdl.GetKeyName(t.Keysym.Sym), sdlCanvas)
					}
				}
			}

			sdlCanvas.renderer.SetDrawColor(0, 0, 0, 0xFF)
			sdlCanvas.renderer.Clear()
		}

		sdlCanvas.drawablesMutex.Lock()
		for _, drawable := range sdlCanvas.drawables {
			if drawable == nil {
				continue
			}

			if !drawable.Initialized() {
				drawable.Init(sdlCanvas)
			}

			drawable.Draw(sdlCanvas)
			drawable.Update(sdlCanvas)
		}
		sdlCanvas.drawablesMutex.Unlock()

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
func (sdlCanvas *SDLCanvas) DrawPoints(points *[]Point, color Color) {
	sdlPoints := make([]sdl.Point, len(*points))

	for i, point := range *points {
		sdlPoints[i] = sdl.Point{X: int32(point.X), Y: int32(point.Y)}
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

func (sdlCanvas *SDLCanvas) FillRects(rects *[]Rect, color Color) {
	sdlRects := make([]sdl.Rect, len(*rects))

	for i, rect := range *rects {
		sdlRects[i] = sdl.Rect{
			X: int32(rect.X),
			Y: int32(rect.Y),
			W: int32(rect.W),
			H: int32(rect.H),
		}
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

func (sdlCanvas *SDLCanvas) DrawColorRects(colorRects *[]ColorRect) {
	sdl.CallQueue <- func() {
		for _, colorRect := range *colorRects {
			sdlCanvas.renderer.SetDrawColor(
				uint8(colorRect.Color.R),
				uint8(colorRect.Color.G),
				uint8(colorRect.Color.B),
				uint8(colorRect.Color.A),
			)

			sdlRect := sdl.Rect{
				X: int32(colorRect.Rect.X),
				Y: int32(colorRect.Rect.Y),
				W: int32(colorRect.Rect.W),
				H: int32(colorRect.Rect.H),
			}

			sdlCanvas.renderer.FillRect(&sdlRect)
		}
	}
}

func (sdlCanvas *SDLCanvas) DrawRects(rects *[]Rect, color Color) {
	sdlRects := make([]sdl.Rect, len(*rects))

	for i, rect := range *rects {
		sdlRects[i] = sdl.Rect{
			X: int32(rect.X),
			Y: int32(rect.Y),
			W: int32(rect.W),
			H: int32(rect.H),
		}
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

func (sdlCanvas *SDLCanvas) DrawLines(points *[]Point, color Color) {
	sdlPoints := make([]sdl.Point, len(*points))

	for i, point := range *points {
		sdlPoints[i] = sdl.Point{X: int32(point.X), Y: int32(point.Y)}
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

func (sdlCanvas *SDLCanvas) Fullscreen() bool {
	return sdlCanvas.fullscreen
}

func (sdlCanvas *SDLCanvas) SetFullscreen(fullscreen bool) {
	sdlCanvas.fullscreen = fullscreen
}
