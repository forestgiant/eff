package sdl

import (
	"fmt"
	"os"

	"github.com/forestgiant/eff"
)

const (
	defaultWindowTitle = "Effulgent"
	defaultWidth       = 480
	defaultHeight      = 320
	defaultFrameRate   = 60
)

var startTime uint32
var delta uint32
var currentFPS uint32

// KeyEnumHandler function that is called when a key board event occurs
type KeyEnumHandler func(key Keycode)

// Canvas creates window and renderer and calls all drawable methods
type Canvas struct {
	Shape

	window              *Window
	windowTitle         string
	clickables          []eff.Clickable
	fullscreen          bool
	keyUpHandlers       []eff.KeyHandler
	keyDownHandlers     []eff.KeyHandler
	keyDownEnumHandlers []KeyEnumHandler
	keyUpEnumHandlers   []KeyEnumHandler
	mouseDownHandlers   []eff.MouseButtonHandler
	mouseUpHandlers     []eff.MouseButtonHandler
	mouseWheelHandlers  []eff.MouseWheelHandler
	mouseMoveHandlers   []eff.MouseMoveHandler
	frameRate           int
	useVsync            bool
}

// NewCanvas creates a new SDL canvas instance
func NewCanvas(title string, width int, height int, clearColor eff.Color, frameRate int, useVsync bool) *Canvas {
	c := Canvas{}
	c.windowTitle = title
	c.rect = eff.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	}
	c.bgColor = clearColor
	c.frameRate = frameRate
	c.useVsync = useVsync
	c.graphics = NewGraphics()
	return &c
}

// AddClickable adds a struct that implements the eff.Clickable interface
func (c *Canvas) AddClickable(clickable eff.Clickable) {
	c.clickables = append(c.clickables, clickable)
}

// RemoveClickable removes a struct that implements the eff.Clickable interface
func (c *Canvas) RemoveClickable(clickable eff.Clickable) {
	index := -1
	for i, d := range c.clickables {
		if d == clickable {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	c.clickables = append(c.clickables[:index], c.clickables[index+1:]...)
}

// AddKeyUpHandler adds key up event handler to the canvas
func (c *Canvas) AddKeyUpHandler(handler eff.KeyHandler) {
	c.keyUpHandlers = append(c.keyUpHandlers, handler)
}

// AddKeyDownHandler adds key down event handler to the canvas
func (c *Canvas) AddKeyDownHandler(handler eff.KeyHandler) {
	c.keyDownHandlers = append(c.keyDownHandlers, handler)
}

// AddKeyUpEnumHandler adds key up event handler to the canvas
func (c *Canvas) AddKeyUpEnumHandler(handler KeyEnumHandler) {
	c.keyUpEnumHandlers = append(c.keyUpEnumHandlers, handler)
}

// AddKeyDownEnumHandler adds key down event handler to the canvas
func (c *Canvas) AddKeyDownEnumHandler(handler KeyEnumHandler) {
	c.keyDownEnumHandlers = append(c.keyDownEnumHandlers, handler)
}

// AddMouseDownHandler adds mouse down event handler to canvas
func (c *Canvas) AddMouseDownHandler(handler eff.MouseButtonHandler) {
	c.mouseDownHandlers = append(c.mouseDownHandlers, handler)
}

// AddMouseUpHandler adds mouse up event handler to canvas
func (c *Canvas) AddMouseUpHandler(handler eff.MouseButtonHandler) {
	c.mouseUpHandlers = append(c.mouseUpHandlers, handler)
}

// AddMouseMoveHandler adds mouse move event handler to canvas
func (c *Canvas) AddMouseMoveHandler(handler eff.MouseMoveHandler) {
	c.mouseMoveHandlers = append(c.mouseMoveHandlers, handler)
}

// AddMouseWheelHandler adds mouse wheel event handler to canvas
func (c *Canvas) AddMouseWheelHandler(handler eff.MouseWheelHandler) {
	c.mouseWheelHandlers = append(c.mouseWheelHandlers, handler)
}

// Run creates an infinite loop that renders all drawables, init is only call once and draw and update are called once per frame
func (c *Canvas) Run(setup func()) {
	lastFPSPrintTime := getTicks()
	init := func() int {
		if c.rect.W == 0 {
			c.rect.W = defaultWidth
		}

		if c.rect.H == 0 {
			c.rect.H = defaultHeight
		}

		if len(c.windowTitle) == 0 {
			c.windowTitle = defaultWindowTitle
		}

		if c.frameRate == 0 {
			c.frameRate = defaultFrameRate
		}

		var err error
		mainThread <- func() {
			c.window, err = createWindow(
				c.windowTitle,
				windowPosUndefined,
				windowPosUndefined,
				c.rect.W,
				c.rect.H,
				windowOpenGl|windowAllowHighDPI,
			)
			drawableW, _ := c.window.getDrawableSize()
			c.graphics.scale = float64(drawableW) / float64(c.rect.W)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
				// return 1
				return
			}

			if c.fullscreen {
				c.window.setFullscreen(windowFullscreen)
			} else {
				c.window.setFullscreen(0)
			}
		}

		mainThread <- func() {
			windowFlags := rendererAccelerated | rendererPresentVsync
			if !c.useVsync {
				windowFlags = rendererAccelerated
			}

			c.graphics.renderer, err = createRenderer(
				c.window,
				-1,
				uint32(windowFlags),
			)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to create renderer: ", err)
				// return 2
				return
			}
		}

		mainThread <- func() {
			c.graphics.renderer.setDrawBlendMode(blendModeBlend)
			c.graphics.renderer.clear()
		}

		startTime = getTicks()
		return 0
	}

	run := func() {
		running := true

		for running {
			mainThread <- func() {
				for event := pollEvent(); event != nil; event = pollEvent() {
					switch t := event.(type) {
					case *quitEvent:
						running = false
					case *keyUpEvent:
						switch t.Keysym.Sym {
						case KeyQ:
							running = false
						case KeyF:
							c.fullscreen = !c.fullscreen
							if c.fullscreen {
								c.window.setFullscreen(windowFullscreen)
							} else {
								c.window.setFullscreen(0)
							}
						}

						for _, handler := range c.keyUpHandlers {
							handler(getKeyName(t.Keysym.Sym))
						}
						for _, handler := range c.keyUpEnumHandlers {
							handler(t.Keysym.Sym)
						}
					case *keyDownEvent:
						for _, handler := range c.keyDownHandlers {
							handler(getKeyName(t.Keysym.Sym))
						}
						for _, handler := range c.keyDownEnumHandlers {
							handler(t.Keysym.Sym)
						}
					case *mouseDownEvent:
						leftState := t.Button == mouseLeft
						middleState := t.Button == mouseMiddle
						rightState := t.Button == mouseRight

						mousePoint := eff.Point{
							X: int(t.X),
							Y: int(t.Y),
						}

						for _, handler := range c.mouseDownHandlers {
							handler(leftState, middleState, rightState)
						}

						for _, clickable := range c.clickables {
							if clickable == nil {
								continue
							}

							hb := clickable.Hitbox()
							if hb.Inside(mousePoint) {
								clickable.MouseDown(leftState, middleState, rightState)
							}
						}

					case *mouseUpEvent:
						leftState := t.Button == mouseLeft
						middleState := t.Button == mouseMiddle
						rightState := t.Button == mouseRight

						mousePoint := eff.Point{
							X: int(t.X),
							Y: int(t.Y),
						}

						for _, handler := range c.mouseUpHandlers {
							handler(leftState, middleState, rightState)
						}

						for _, clickable := range c.clickables {
							if clickable == nil {
								continue
							}

							hb := clickable.Hitbox()
							if hb.Inside(mousePoint) {
								clickable.MouseUp(leftState, middleState, rightState)
							}
						}
					case *mouseMotionEvent:
						mousePoint := eff.Point{
							X: int(t.X),
							Y: int(t.Y),
						}

						for _, handler := range c.mouseMoveHandlers {
							handler(mousePoint.X, mousePoint.Y)
						}

						for _, clickable := range c.clickables {
							if clickable == nil {
								continue
							}

							hb := clickable.Hitbox()
							if hb.Inside(mousePoint) {
								if !clickable.IsMouseOver() {
									clickable.MouseOver()
								}
							} else {
								if clickable.IsMouseOver() {
									clickable.MouseOut()
								}
							}
						}

					case *mouseWheelEvent:
						for _, handler := range c.mouseWheelHandlers {
							handler(int(t.X), int(t.Y))
						}
					}

				}

				c.graphics.renderer.clear()
			}

			c.Draw(c)
			c.HandleUpdate()

			mainThread <- func() {

				printFPS := func() {
					delta = getTicks() - startTime
					if delta != 0 {
						currentFPS = 1000 / delta
					}
					if getTicks()-lastFPSPrintTime >= 1000 {
						fmt.Println(currentFPS, "fps")
						lastFPSPrintTime = getTicks()
					}
				}

				enforceFPS := func() {
					timeBetweenFrames := getTicks() - startTime
					targetTimeBetweenFrames := 1000 / uint32(c.frameRate)

					if timeBetweenFrames < targetTimeBetweenFrames {
						delay(targetTimeBetweenFrames - timeBetweenFrames)
					}

				}

				c.graphics.renderer.present()
				enforceFPS()
				printFPS()

				startTime = getTicks()
			}
		}
	}

	lockMain(func() {
		setup()
		initOK := init()
		if initOK != 0 {
			os.Exit(initOK)
		}
		run()
		mainThread <- func() {
			// Clean up goes here
			c.graphics.renderer.destroy()
			c.window.destroy()

			//Quit SDL
			quit()
			close(mainDone) // stop mainThread
		}
	})
}

// DrawPoints draw a slice of points to the screen all the same color

// Fullscreen get the full screen state of the window
func (c *Canvas) Fullscreen() bool {
	return c.fullscreen
}

// SetFullscreen set the fullscreen state of the window
func (c *Canvas) SetFullscreen(fullscreen bool) {
	c.fullscreen = fullscreen
}

// OpenFont creates a eff.Font object, used for rendering text
func (c *Canvas) OpenFont(path string, size int) (eff.Font, error) {
	size = int(float64(size) * c.graphics.scale)
	f, err := openFont(path, size)
	if err != nil {
		return nil, err
	}
	return f, nil
}
