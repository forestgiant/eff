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

// Canvas creates window and renderer and calls all drawable methods
type Canvas struct {
	window          *Window
	renderer        *Renderer
	drawables       []eff.Drawable
	width           int
	height          int
	fullscreen      bool
	keyUpHandlers   []eff.KeyHandler
	keyDownHandlers []eff.KeyHandler
}

// SetWidth set the width of the canvas, must be called prior to run
func (c *Canvas) SetWidth(width int) {
	c.width = width
}

// Width get the width of the canvas window
func (c *Canvas) Width() int {
	return c.width
}

// SetHeight set the height of the canvas, must be called prior to run
func (c *Canvas) SetHeight(height int) {
	c.height = height
}

// Height get the height of the canvas window
func (c *Canvas) Height() int {
	return c.height
}

// AddDrawable adds a struct that implements the eff.Drawable interface
func (c *Canvas) AddDrawable(drawable eff.Drawable) {
	c.drawables = append(c.drawables, drawable)
}

//RemoveDrawable removes struct from canvas that implements eff.Drawable
func (c *Canvas) RemoveDrawable(drawable eff.Drawable) {
	index := -1
	for i, d := range c.drawables {
		if d == drawable {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	c.drawables = append(c.drawables[:index], c.drawables[index+1:]...)
}

//AddKeyUpHandler adds key up event handler to the canvas
func (c *Canvas) AddKeyUpHandler(handler eff.KeyHandler) {
	c.keyUpHandlers = append(c.keyUpHandlers, handler)
}

//AddKeyDownHandler adds key down event handler to the canvas
func (c *Canvas) AddKeyDownHandler(handler eff.KeyHandler) {
	c.keyDownHandlers = append(c.keyDownHandlers, handler)
}

// Run creates an infinite loop that renders all drawables, init is only call once and draw and update are called once per frame
func (c *Canvas) Run() {

	init := func() int {
		if c.width == 0 {
			c.width = defaultWidth
		}

		if c.height == 0 {
			c.height = defaultHeight
		}

		var err error
		MainThread <- func() {
			c.window, err = CreateWindow(
				windowTitle,
				WindowPosUndefined,
				WindowPosUndefined,
				c.Width(),
				c.Height(),
				WindowOpenGl,
			)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
				// return 1
				return
			}

			if c.fullscreen {
				c.window.SetFullscreen(WindowFullscreen)
			} else {
				c.window.SetFullscreen(0)
			}
		}

		MainThread <- func() {
			c.renderer, err = CreateRenderer(
				c.window,
				-1,
				RendererAccelerated|RendererPresentVsync,
			)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to create renderer: ", err)
				// return 2
				return
			}
		}

		MainThread <- func() {
			c.renderer.Clear()
		}

		return 0
	}

	run := func() {
		running := true
		var lastFrameTime uint32

		for running {
			MainThread <- func() {
				for event := PollEvent(); event != nil; event = PollEvent() {
					switch t := event.(type) {
					case *QuitEvent:
						running = false
					case *KeyUpEvent:
						switch t.Keysym.Sym {
						case KeyQ:
							running = false
						case KeyF:
							c.fullscreen = !c.fullscreen
							if c.fullscreen {
								c.window.SetFullscreen(WindowFullscreen)
							} else {
								c.window.SetFullscreen(0)
							}
						}

						for _, handler := range c.keyUpHandlers {
							handler(GetKeyName(t.Keysym.Sym))
						}
					case *KeyDownEvent:
						for _, handler := range c.keyDownHandlers {
							handler(GetKeyName(t.Keysym.Sym))
						}
					}
				}

				c.renderer.SetDrawColor(0x0, 0x0, 0x0, 0xFF)
				c.renderer.Clear()
			}

			for _, drawable := range c.drawables {
				if drawable == nil {
					continue
				}

				if !drawable.Initialized() {
					drawable.Init(c)
				}

				drawable.Draw(c)
				drawable.Update(c)
			}

			MainThread <- func() {
				currentFrameTime := GetTicks()
				if lastFrameTime == 0 {
					lastFrameTime = currentFrameTime
				}
				c.renderer.Present()
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
		MainThread <- func() {
			c.renderer.Destroy()
			c.window.Destroy()
		}
		os.Exit(0)
	}()

	LockMain()
}

//DrawPoints draw a slice of points to the screen all the same color
func (c *Canvas) DrawPoints(points []eff.Point, color eff.Color) {
	var sdlPoints []Point

	for _, point := range points {
		sdlPoints = append(sdlPoints, Point{X: int32(point.X), Y: int32(point.Y)})
	}

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.DrawPoints(sdlPoints)
	}
}

//DrawPoint draw a point on the screen specifying what color
func (c *Canvas) DrawPoint(point eff.Point, color eff.Color) {
	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		c.renderer.DrawPoint(point.X, point.Y)
	}
}

//DrawColorPoints draw a slide of colorPoints on the screen
func (c *Canvas) DrawColorPoints(colorPoints []eff.ColorPoint) {
	MainThread <- func() {
		for _, colorPoint := range colorPoints {
			c.renderer.SetDrawColor(
				uint8(colorPoint.R),
				uint8(colorPoint.G),
				uint8(colorPoint.B),
				uint8(colorPoint.A),
			)

			c.renderer.DrawPoint(colorPoint.X, colorPoint.Y)
		}
	}
}

//FillRect draw a filled in rectangle to the screen
func (c *Canvas) FillRect(rect eff.Rect, color eff.Color) {
	sdlRect := Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.FillRect(&sdlRect)
	}
}

//FillRects draw a slice of filled rectangles to the screen all the same color
func (c *Canvas) FillRects(rects []eff.Rect, color eff.Color) {
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

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.FillRects(sdlRects)
	}
}

//DrawRect draw an outlined rectangle to the screen with a color
func (c *Canvas) DrawRect(rect eff.Rect, color eff.Color) {
	sdlRect := Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.DrawRect(&sdlRect)
	}
}

//DrawColorRects draw a slice of color rectangles to the screen
func (c *Canvas) DrawColorRects(colorRects []eff.ColorRect) {
	MainThread <- func() {
		for _, colorRect := range colorRects {
			c.renderer.SetDrawColor(
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

			c.renderer.FillRect(&sdlRect)
		}
	}
}

//DrawRects draw a slice of rectangles to the screen all the same color
func (c *Canvas) DrawRects(rects []eff.Rect, color eff.Color) {
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

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.DrawRects(sdlRects)
	}
}

//DrawLine draw a line of to the screen with a color
func (c *Canvas) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		c.renderer.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	}
}

//DrawLines a slice of lines to the screen all the same color
func (c *Canvas) DrawLines(points []eff.Point, color eff.Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []Point

	for _, point := range points {
		p := Point{X: int32(point.X), Y: int32(point.Y)}
		sdlPoints = append(sdlPoints, p)
	}

	MainThread <- func() {
		c.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.DrawLines(sdlPoints)
	}
}

//Fullscreen get the full screen state of the window
func (c *Canvas) Fullscreen() bool {
	return c.fullscreen
}

//SetFullscreen set the fullscreen state of the window
func (c *Canvas) SetFullscreen(fullscreen bool) {
	c.fullscreen = fullscreen
}