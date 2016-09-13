package eff

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle  = "Effulgent"
	windowWidth  = 1280
	windowHeight = 720
	frameRate    = 90
	frameTime    = 1000 / frameRate
)

type Point struct {
	X int
	Y int
}

type Color struct {
	R int
	G int
	B int
	A int
}

type Drawable interface {
	Init()
	Draw()
	Update()
}

type Canvas struct {
	window    *sdl.Window
	renderer  *sdl.Renderer
	drawables []Drawable
}

func (canvas *Canvas) AddDrawable(drawable Drawable) {
	canvas.drawables = append(canvas.drawables, drawable)
}

func (canvas *Canvas) Run() int {
	var err error
	sdl.CallQueue <- func() {
		canvas.window, err = sdl.CreateWindow(
			windowTitle,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			windowWidth,
			windowHeight,
			sdl.WINDOW_OPENGL,
		)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		sdl.CallQueue <- func() {
			canvas.window.Destroy()
		}
	}()

	sdl.CallQueue <- func() {
		canvas.renderer, err = sdl.CreateRenderer(
			canvas.window,
			-1,
			sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC,
		)
	}
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.CallQueue <- func() {
			canvas.renderer.Destroy()
		}
	}()

	sdl.CallQueue <- func() {
		canvas.renderer.Clear()
	}

	// Init Code Goes Here
	for _, drawable := range canvas.drawables {
		drawable.Init()
	}

	running := true
	fullscreen := false
	var lastFrameTime uint32 = sdl.GetTicks()
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
						fullscreen = !fullscreen
						if fullscreen {
							canvas.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
						} else {
							canvas.window.SetFullscreen(0)
						}
					}
				}
			}

			canvas.renderer.SetDrawColor(0, 0, 0, 0xFF)
			canvas.renderer.Clear()
		}

		for _, drawable := range canvas.drawables {
			drawable.Draw()
			drawable.Update()
		}

		sdl.CallQueue <- func() {
			currentFrameTime := sdl.GetTicks()
			canvas.renderer.Present()
			if currentFrameTime-lastFrameTime < frameTime {
				sdl.Delay(frameTime - (currentFrameTime - lastFrameTime))
			}
			lastFrameTime = currentFrameTime
		}
	}
	return 0
}

func (canvas *Canvas) DrawPoints(points *[]Point, color Color) {
	sdl.CallQueue <- func() {
		canvas.renderer.SetDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		sdlPoints := make([]sdl.Point, len(*points))

		for i, point := range *points {
			sdlPoints[i] = sdl.Point{int32(point.X), int32(point.Y)}
		}

		canvas.renderer.DrawPoints(sdlPoints)
	}
}
