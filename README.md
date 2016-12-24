# eff
Effulgent Media API
---
This API provides a way to easily create graphics programs in Go. Providing a framework to create [games](#examples), 
[ui](#examples), [animation](#examples), or any type of graphical application.

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
    1. Download mysys264 <https://msys2.github.io/>
    2. Install the git the x86_64 toolchain and SDL2 using pacman ```pacman -S git mingw-w64-x86_64-toolchain mingw64/mingw-w64-x86_64-SDL2 mingw64/mingw-w64-x86_64-SDL2_mixer mingw64/mingw-w64-x86_64-SDL2_image mingw64/mingw-w64-x86_64-SDL2_ttf mingw64/mingw-w64-x86_64-SDL2_net mingw64/mingw-w64-x86_64-cmake make``` You might need to restart the mysys64 shell once installed
    3. Download golang from the website as a zip not an installer
    4. Install in inside the mysys64 root, typically `C:\msys64\go`
	5. Inside the mysys64 enviroment update your `.bashrc` to have a `$GOROOT` (in the above example is `C:\mysys64\go` and `$GOPATH`
	6. Ensure that the `$PATH` is appended with `$GOROOT\bin` and `$GOPATH\bin`
	7. Either `go get github.com/forestgiant/eff` or clone the repo from git in your `$GOPATH/src`
	8. From the root of the `eff` source code type `make` this will build all the examples

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