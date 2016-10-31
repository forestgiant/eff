package sdl

import (
	"errors"
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

// Canvas creates window and renderer and calls all drawable methods
type Canvas struct {
	window          *Window
	renderer        *renderer
	drawables       []eff.Drawable
	width           int
	height          int
	fullscreen      bool
	keyUpHandlers   []eff.KeyHandler
	keyDownHandlers []eff.KeyHandler
	windowTitle     string
	frameRate       int
	useVsync        bool
	font            *font
	images          map[*eff.Image]*imageTex
	clearColor      eff.Color
}

// NewCanvas creates a new SDL canvas instance
func NewCanvas(title string, width int, height int, clearColor eff.Color, frameRate int, useVsync bool) *Canvas {
	c := Canvas{}
	c.windowTitle = title
	c.width = width
	c.height = height
	c.frameRate = frameRate
	c.useVsync = useVsync
	c.images = make(map[*eff.Image]*imageTex)
	c.clearColor = clearColor
	return &c
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

// RemoveDrawable removes struct from canvas that implements eff.Drawable
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

// AddKeyUpHandler adds key up event handler to the canvas
func (c *Canvas) AddKeyUpHandler(handler eff.KeyHandler) {
	c.keyUpHandlers = append(c.keyUpHandlers, handler)
}

// AddKeyDownHandler adds key down event handler to the canvas
func (c *Canvas) AddKeyDownHandler(handler eff.KeyHandler) {
	c.keyDownHandlers = append(c.keyDownHandlers, handler)
}

// Run creates an infinite loop that renders all drawables, init is only call once and draw and update are called once per frame
func (c *Canvas) Run(setup func()) {
	lastFPSPrintTime := getTicks()
	init := func() int {
		if c.width == 0 {
			c.width = defaultWidth
		}

		if c.height == 0 {
			c.height = defaultHeight
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
				c.Width(),
				c.Height(),
				windowOpenGl,
			)
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

			c.renderer, err = createRenderer(
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
			c.renderer.setDrawColor(uint8(c.clearColor.R), uint8(c.clearColor.G), uint8(c.clearColor.B), uint8(c.clearColor.A))
			c.renderer.clear()
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
						case keyQ:
							running = false
						case keyF:
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
					case *keyDownEvent:
						for _, handler := range c.keyDownHandlers {
							handler(getKeyName(t.Keysym.Sym))
						}
					}
				}

				c.renderer.setDrawColor(uint8(c.clearColor.R), uint8(c.clearColor.G), uint8(c.clearColor.B), uint8(c.clearColor.A))
				c.renderer.clear()
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

			mainThread <- func() {
				for i, iT := range c.images {
					if iT.texture == nil {
						fmt.Println("texture is nil")
						continue
					}

					r1 := rect{
						X: 0,
						Y: 0,
						W: iT.w,
						H: iT.h,
					}

					r := rect{
						X: int32(i.Rect.X),
						Y: int32(i.Rect.Y),
						W: int32(i.Rect.W),
						H: int32(i.Rect.H),
					}
					c.renderer.renderCopy(iT.texture, r1, r)
				}
			}

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

				c.renderer.present()
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
			c.renderer.destroy()
			c.window.destroy()

			//Quit SDL
			quit()
			close(mainDone) // stop mainThread
		}
	})
}

// DrawPoints draw a slice of points to the screen all the same color
func (c *Canvas) DrawPoints(points []eff.Point, color eff.Color) {
	var sdlPoints []point

	for _, p := range points {
		sdlPoints = append(sdlPoints, point{X: int32(p.X), Y: int32(p.Y)})
	}

	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.drawPoints(sdlPoints)
	}
}

// DrawPoint draw a point on the screen specifying what color
func (c *Canvas) DrawPoint(point eff.Point, color eff.Color) {
	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		c.renderer.drawPoint(point.X, point.Y)
	}
}

// DrawColorPoints draw a slide of colorPoints on the screen
func (c *Canvas) DrawColorPoints(colorPoints []eff.ColorPoint) {
	mainThread <- func() {
		for _, colorPoint := range colorPoints {
			c.renderer.setDrawColor(
				uint8(colorPoint.R),
				uint8(colorPoint.G),
				uint8(colorPoint.B),
				uint8(colorPoint.A),
			)

			c.renderer.drawPoint(colorPoint.X, colorPoint.Y)
		}
	}
}

// FillRect draw a filled in rectangle to the screen
func (c *Canvas) FillRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.W),
		H: int32(r.H),
	}

	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.fillRect(&sdlRect)
	}
}

// FillRects draw a slice of filled rectangles to the screen all the same color
func (c *Canvas) FillRects(rects []eff.Rect, color eff.Color) {
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
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.fillRects(sdlRects)
	}
}

// DrawRect draw an outlined rectangle to the screen with a color
func (c *Canvas) DrawRect(r eff.Rect, color eff.Color) {
	sdlRect := rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.W),
		H: int32(r.H),
	}

	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.drawRect(&sdlRect)
	}
}

// DrawColorRects draw a slice of color rectangles to the screen
func (c *Canvas) DrawColorRects(colorRects []eff.ColorRect) {
	mainThread <- func() {
		for _, colorRect := range colorRects {
			c.renderer.setDrawColor(
				uint8(colorRect.R),
				uint8(colorRect.G),
				uint8(colorRect.B),
				uint8(colorRect.A),
			)

			sdlRect := rect{
				X: int32(colorRect.X),
				Y: int32(colorRect.Y),
				W: int32(colorRect.W),
				H: int32(colorRect.H),
			}

			c.renderer.fillRect(&sdlRect)
		}
	}
}

// DrawRects draw a slice of rectangles to the screen all the same color
func (c *Canvas) DrawRects(rects []eff.Rect, color eff.Color) {
	var sdlRects []rect

	for _, r := range rects {
		r := rect{
			X: int32(r.X),
			Y: int32(r.Y),
			W: int32(r.W),
			H: int32(r.H),
		}

		sdlRects = append(sdlRects, r)
	}

	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.drawRects(sdlRects)
	}
}

// DrawLine draw a line of to the screen with a color
func (c *Canvas) DrawLine(p1 eff.Point, p2 eff.Point, color eff.Color) {
	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)
		c.renderer.drawLine(p1.X, p1.Y, p2.X, p2.Y)
	}
}

// DrawLines a slice of lines to the screen all the same color
func (c *Canvas) DrawLines(points []eff.Point, color eff.Color) {
	if len(points) == 0 {
		return
	}
	var sdlPoints []point

	for _, p := range points {
		p := point{X: int32(p.X), Y: int32(p.Y)}
		sdlPoints = append(sdlPoints, p)
	}

	mainThread <- func() {
		c.renderer.setDrawColor(
			uint8(color.R),
			uint8(color.G),
			uint8(color.B),
			uint8(color.A),
		)

		c.renderer.drawLines(sdlPoints)
	}
}

// Fullscreen get the full screen state of the window
func (c *Canvas) Fullscreen() bool {
	return c.fullscreen
}

// SetFullscreen set the fullscreen state of the window
func (c *Canvas) SetFullscreen(fullscreen bool) {
	c.fullscreen = fullscreen
}

// SetFont sets the font on the Canvas used for DrawText
func (c *Canvas) SetFont(font eff.Font, size int) error {
	f, err := openFont(font.Path, size)
	c.font = f
	if err != nil {
		return err
	}
	return nil
}

// DrawText draws a string using a font to the screen, the point is the upper left hand corner
func (c *Canvas) DrawText(text string, col eff.Color, point eff.Point) error {
	if c.font == nil {
		return errors.New("Can't draw text no font set")
	}

	rgba := color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}

	mainThread <- func() {
		s, err := renderTextBlended(c.font, text, rgba)
		if err != nil {
			fmt.Println(err)
		}
		t, err := c.renderer.createTextureFromSurface(s)

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

		err = c.renderer.renderCopy(t, r1, r)
		if err != nil {
			fmt.Println(err)
		}

		t.destroy()
	}

	return nil
}

// GetTextSize this uses the currently set font to determine the size of string rendered with that font, does not actually add the text to the canvas
func (c *Canvas) GetTextSize(text string) (int, int, error) {
	if c.font == nil {
		return -1, -1, errors.New("Can't get text size no font set")
	}

	errChan := make(chan error)
	sizeChan := make(chan point)

	rgba := color{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}

	mainThread <- func() {
		s, err := renderTextBlended(c.font, text, rgba)
		if err != nil {
			errChan <- err
		}

		p := point{}
		p.X = int32(s.w)
		p.Y = int32(s.h)

		freeSurface(s)

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

// AddImage load and store an image in this canvas instance, set the image height and width to -1 and they will be replaced with the images native height and width
func (c *Canvas) AddImage(i *eff.Image) {
	if c.images[i] != nil {
		//Texture already exists for this image
		fmt.Println("Image already in the canvas")
		return
	}

	// Load the texture
	s, err := loadImg(i.Path)
	if err != nil {
		fmt.Println(err)
	}

	if i.Rect.W == -1 {
		i.Rect.W = int(s.w)
	}

	if i.Rect.H == -1 {
		i.Rect.H = int(s.h)
	}

	t, err := c.renderer.createTextureFromSurface(s)
	if err != nil {
		fmt.Println(err)
	}

	c.images[i] = &imageTex{
		texture: t,
		w:       int32(s.w),
		h:       int32(s.h),
	}
}

// RemoveImage remove the image from this canvas instance
func (c *Canvas) RemoveImage(i *eff.Image) {
	if c.images[i] != nil {
		iT := c.images[i]
		delete(c.images, i)

		iT.texture.destroy()
	}
}
