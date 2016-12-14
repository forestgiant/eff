# eff
Effulgent Media API
---
This API provides a way to easily create graphics programs in Go. Providing a framework to create [games](#examples), 
[ui](#examples), [animation](#examples), or any type of graphical application.

The sdl package is a partial wrapper of sdl for Go.  For a complete wrapper checkout go-sdl2 <https://github.com/veandco/go-sdl2>
> **NOTE:** This repository is under heavy ongoing development and
is likely to break over time. We currently do not have any releases
yet. If you are planning to use the repository, please consider vendoring
the packages in your project and update them when a stable tag is out.

### SDL setup
Eff uses the openGL renderer in SDL, any system it runs on will require opengl support
* OSX: `brew install sdl2{,_mixer,_image,_ttf}`
* Arch Linux: `sudo pacman -S sdl2{,_mixer,_image,_ttf}`
* Ubuntu/Debian: `sudo apt-get install libsdl2{,-mixer,-image,-ttf}-dev `
* Windows:
    1. Install mingw <http://www.mingw.org/>, ensure the `bin` folder is in the windows path
    2. Download the windows development sdl2 libraries:
        1. SDL2 <https://www.libsdl.org/download-2.0.php>
        2. SDL2_ttf <https://www.libsdl.org/projects/SDL_ttf/>
        3. SDL2_mixer <https://www.libsdl.org/projects/SDL_mixer/>
        4. SDL2_image <https://www.libsdl.org/projects/SDL_image/>
    3. Extract each tarball to the same directory (i.e. `c:\mingw_dev_lib`) and add the bin(i.e. `c:\mingw_dev_lib\bin`) folder to the PATH enviroment variable.  Currently only the 32bit(i686) libraries are supported.
    4. Update the cgo comment at the top of sdl.go to ensure that the include path and lib path match where you extracted the libraries
    5. When building set the `GOARCH=386` and `CGO_ENABLED=1`. Use the `SET` command if you are using the normal windows command line and not git-bash

### API Usage Example
```
const (
	windowW    = 800
	windowH    = 540
	squareSize = 100
)

type myShape struct {
	eff.Shape
}

func main() {
	canvas := sdl.NewCanvas("Boilerplate", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		m := &myShape{}
		m.SetRect(eff.Rect{
			X: (windowW - squareSize) / 2,
			Y: (windowH - squareSize) / 2,
			W: squareSize,
			H: squareSize,
		})
		minSpeed := 3
		maxSpeed := 10
		vec := eff.Point{X: rand.Intn(maxSpeed-minSpeed) + minSpeed, Y: rand.Intn(maxSpeed-minSpeed) + minSpeed}
		m.SetUpdateHandler(func() {
			x := m.Rect().X + vec.X
			y := m.Rect().Y + vec.Y
			if x <= 0 || x >= (canvas.Rect().W-m.Rect().W) {
				vec.X *= -1
			}

			if y <= 0 || y >= (canvas.Rect().H-m.Rect().H) {
				vec.Y *= -1
			}

			m.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: m.Rect().W,
				H: m.Rect().H,
			})
		})
		m.SetBackgroundColor(eff.RandomColor())
		canvas.AddChild(m)
	})
}
```

### Keyboard control
* Press `f` to toggle fullscreen
* Press `q` to quit the program

### Examples
* [Tetris](https://github.com/thales17/eff-tetris)
* [Animating Text](https://github.com/forestgiant/eff/tree/master/examples/animating-text)
* [Image Tiling](https://github.com/forestgiant/eff/tree/master/examples/image-tile)
* [Text Layout](https://github.com/forestgiant/eff/tree/master/examples/text-view)