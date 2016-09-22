package sdl

import (
	"fmt"
	"os"

	"github.com/forestgiant/eff/eff"
)

const (
	windowTitle   = "Effulgent"
	defaultWidth  = 480
	defaultHeight = 320
	frameRate     = 90
	frameTime     = 1000 / frameRate
)

// Wraps the Drawable to track whether or not it has been initialized
type sdlDrawable struct {
	initialized bool
	drawable    eff.Drawable
}

func (s *sdlDrawable) Init(canvas eff.Canvas) {
	s.drawable.Init(canvas)
	s.initialized = true
}

func (s *sdlDrawable) Initialized() bool {
	return s.initialized
}

func (s *sdlDrawable) Draw(canvas eff.Canvas) {
	s.drawable.Draw(canvas)
}

func (s *sdlDrawable) Update(canvas eff.Canvas) {
	s.drawable.Update(canvas)
}

// Canvas creates window and renderer and calls all drawable methods
type Canvas struct {
	window   *Window
	renderer *Renderer
	// drawablesMutex  sync.Mutex
	drawables       []*sdlDrawable
	width           int
	height          int
	fullscreen      bool
	keyUpHandlers   []eff.KeyHandler
	keyDownHandlers []eff.KeyHandler
}

// SetWidth set the width of the canvas, must be called prior to run
func (sdlCanvas *Canvas) SetWidth(width int) {
	sdlCanvas.width = width
}

// Width get the width of the canvas window
func (sdlCanvas *Canvas) Width() int {
	return sdlCanvas.width
}

// SetHeight set the height of the canvas, must be called prior to run
func (sdlCanvas *Canvas) SetHeight(height int) {
	sdlCanvas.height = height
}

// Height get the height of the canvas window
func (sdlCanvas *Canvas) Height() int {
	return sdlCanvas.height
}

// AddDrawable adds a struct that implements the eff.Drawable interface
func (sdlCanvas *Canvas) AddDrawable(drawable eff.Drawable) {
	sdlCanvas.drawables = append(sdlCanvas.drawables, &sdlDrawable{drawable: drawable})
}

//RemoveDrawable removes struct from canvas that implements eff.Drawable
func (sdlCanvas *Canvas) RemoveDrawable(drawable eff.Drawable) {
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
func (sdlCanvas *Canvas) AddKeyUpHandler(handler eff.KeyHandler) {
	sdlCanvas.keyUpHandlers = append(sdlCanvas.keyUpHandlers, handler)
}

//AddKeyDownHandler adds key down event handler to the canvas
func (sdlCanvas *Canvas) AddKeyDownHandler(handler eff.KeyHandler) {
	sdlCanvas.keyDownHandlers = append(sdlCanvas.keyDownHandlers, handler)
}

// Run creates an infinite loop that renders all drawables, init is only call once and draw and update are called once per frame
func (sdlCanvas *Canvas) Run() int {

	init := func() int {
		if sdlCanvas.width == 0 {
			sdlCanvas.width = defaultWidth
		}

		if sdlCanvas.height == 0 {
			sdlCanvas.height = defaultHeight
		}

		var err error
		CallQueue <- func() {
			sdlCanvas.window, err = CreateWindow(
				windowTitle,
				WindowPosUndefined,
				WindowPosUndefined,
				sdlCanvas.Width(),
				sdlCanvas.Height(),
				WindowOpenGl,
			)

			if sdlCanvas.fullscreen {
				sdlCanvas.window.SetFullscreen(WindowFullscreen)
			} else {
				sdlCanvas.window.SetFullscreen(0)
			}
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
			return 1
		}

		CallQueue <- func() {
			sdlCanvas.renderer, err = CreateRenderer(
				sdlCanvas.window,
				-1,
				RendererAccelerated|RendererPresentVsync,
			)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create renderer: ", err)
			return 2
		}

		CallQueue <- func() {
			sdlCanvas.renderer.Clear()
		}

		return 0
	}

	run := func() {
		running := true
		var lastFrameTime uint32

		for running {
			CallQueue <- func() {
				for event := PollEvent(); event != nil; event = PollEvent() {
					switch t := event.(type) {
					case *QuitEvent:
						running = false
					case *KeyUpEvent:
						switch t.Keysym.Sym {
						case KeyQ:
							running = false
						case KeyF:
							sdlCanvas.fullscreen = !sdlCanvas.fullscreen
							if sdlCanvas.fullscreen {
								sdlCanvas.window.SetFullscreen(WindowFullscreen)
							} else {
								sdlCanvas.window.SetFullscreen(0)
							}
						}

						for _, handler := range sdlCanvas.keyUpHandlers {
							handler(GetKeyName(t.Keysym.Sym))
						}
					case *KeyDownEvent:
						for _, handler := range sdlCanvas.keyDownHandlers {
							handler(GetKeyName(t.Keysym.Sym))
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

			CallQueue <- func() {
				currentFrameTime := GetTicks()
				if lastFrameTime == 0 {
					lastFrameTime = currentFrameTime
				}
				sdlCanvas.renderer.Present()
				if currentFrameTime-lastFrameTime < frameTime {
					Delay(frameTime - (currentFrameTime - lastFrameTime))
				}
				lastFrameTime = currentFrameTime
			}
		}
	}

	go func() {
		initOK := init()
		if initOK != 0 {
			os.Exit(initOK)
		}
		run()
		CallQueue <- func() {
			sdlCanvas.renderer.Destroy()
			sdlCanvas.window.Destroy()
		}
		os.Exit(0)
	}()

	ProcessCalls()

	return 0
}

//DrawPoints draw a slice of points to the screen all the same color
func (sdlCanvas *Canvas) DrawPoints(points []eff.Point, color eff.Color) {
	var sdlPoints []Point

	for _, point := range points {
		sdlPoints = append(sdlPoints, Point{X: int32(point.X), Y: int32(point.Y)})
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawPoint(point eff.Point, color eff.Color) {
	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawColorPoints(colorPoints []eff.ColorPoint) {
	CallQueue <- func() {
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
func (sdlCanvas *Canvas) FillRect(rect eff.Rect, color eff.Color) {
	sdlRect := Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) FillRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []Rect

	for _, rect := range rects {
		sdlRects = append(sdlRects,
			Rect{
				X: int32(rect.X),
				Y: int32(rect.Y),
				W: int32(rect.W),
				H: int32(rect.H),
			},
		)
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawRect(rect eff.Rect, color eff.Color) {
	sdlRect := Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawColorRects(colorRects []eff.ColorRect) {
	CallQueue <- func() {
		for _, colorRect := range colorRects {
			sdlCanvas.renderer.SetDrawColor(
				uint8(colorRect.R),
				uint8(colorRect.G),
				uint8(colorRect.B),
				uint8(colorRect.A),
			)

			sdlRect := Rect{
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
func (sdlCanvas *Canvas) DrawRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []Rect

	for _, rect := range rects {
		r := Rect{
			X: int32(rect.X),
			Y: int32(rect.Y),
			W: int32(rect.W),
			H: int32(rect.H),
		}

		sdlRects = append(sdlRects, r)
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	CallQueue <- func() {
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
func (sdlCanvas *Canvas) DrawLines(points []eff.Point, color eff.Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []Point

	for _, point := range points {
		p := Point{X: int32(point.X), Y: int32(point.Y)}
		sdlPoints = append(sdlPoints, p)
	}

	CallQueue <- func() {
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
func (sdlCanvas *Canvas) Fullscreen() bool {
	return sdlCanvas.fullscreen
}

//SetFullscreen set the fullscreen state of the window
func (sdlCanvas *Canvas) SetFullscreen(fullscreen bool) {
	sdlCanvas.fullscreen = fullscreen
}
