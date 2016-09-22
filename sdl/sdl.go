package sdl

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "sdl_wrapper.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

// Point is a structure that defines a two demensional point.
// (https://wiki.libsdl.org/SDL_Point)
type Point struct {
	X int32
	Y int32
}

// Rect is a structure that defines a rectangle, with the origin at the upper
// left.
// (https://wiki.libsdl.org/SDL_Rect)
type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (p *Point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

func (a *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(a))
}

//Renderer SDL renderer
type Renderer C.SDL_Renderer

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

//Window SDL window
type Window C.SDL_Window

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

//CallQueue manages the thread that SDL calls execute on
var CallQueue = make(chan func(), 1)

const (
	WindowFullscreen        = C.SDL_WINDOW_FULLSCREEN
	WindowOpenGl            = C.SDL_WINDOW_OPENGL
	WindowShown             = C.SDL_WINDOW_SHOWN
	WindowHidden            = C.SDL_WINDOW_HIDDEN
	WindowBorderless        = C.SDL_WINDOW_BORDERLESS
	WindowResizable         = C.SDL_WINDOW_RESIZABLE
	WindowMinimized         = C.SDL_WINDOW_MINIMIZED
	WindowMaximized         = C.SDL_WINDOW_MAXIMIZED
	WindowInputGrabbed      = C.SDL_WINDOW_INPUT_GRABBED
	WindowInputFocus        = C.SDL_WINDOW_INPUT_FOCUS
	WindowMouseFocus        = C.SDL_WINDOW_MOUSE_FOCUS
	WindowFullscreenDesktop = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WindowForeign           = C.SDL_WINDOW_FOREIGN
	WindowAllowHighDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI
)

const (
	WindowPosUndefinedMask = C.SDL_WINDOWPOS_UNDEFINED_MASK
	WindowPosUndefined     = C.SDL_WINDOWPOS_UNDEFINED
	WindowPosCenteredMask  = C.SDL_WINDOWPOS_CENTERED_MASK
	WindowPosCentered      = C.SDL_WINDOWPOS_CENTERED
)

const (
	RendererSoftware      = C.SDL_RENDERER_SOFTWARE
	RendererAccelerated   = C.SDL_RENDERER_ACCELERATED
	RendererPresentVsync  = C.SDL_RENDERER_PRESENTVSYNC
	RendererTargetTexture = C.SDL_RENDERER_TARGETTEXTURE

	TextureAccessStatic    = C.SDL_TEXTUREACCESS_STATIC
	TextureAccessStreaming = C.SDL_TEXTUREACCESS_STREAMING
	TextureAccessTarget    = C.SDL_TEXTUREACCESS_TARGET

	TextureModulateNone  = C.SDL_TEXTUREMODULATE_NONE
	TextureModulateColor = C.SDL_TEXTUREMODULATE_COLOR
	TextureModulateAlpha = C.SDL_TEXTUREMODULATE_ALPHA
)

// ProcessCalls run through functions in CallQueue. Intended to be called as a goroutine.
func ProcessCalls() {
	runtime.LockOSThread()

	for {
		f := <-CallQueue
		f()
	}
}

// GetTicks (https://wiki.libsdl.org/SDL_GetTicks)
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// Delay (https://wiki.libsdl.org/SDL_Delay)
func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

// GetError (https://wiki.libsdl.org/SDL_GetError)
func GetError() error {
	if err := C.SDL_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// ClearError (https://wiki.libsdl.org/SDL_ClearError)
func ClearError() {
	C.SDL_ClearError()
}

// SetFullscreen (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (w *Window) SetFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(w.cptr(), C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// CreateWindow (https://wiki.libsdl.org/SDL_CreateWindow)
func CreateWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error) {
	_title := C.CString(title)
	var _window = C.SDL_CreateWindow(_title, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// Destroy (https://wiki.libsdl.org/SDL_DestroyWindow)
func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.cptr())
}

// CreateRenderer (https://wiki.libsdl.org/SDL_CreateRenderer)
func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error) {
	_renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if _renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(_renderer)), nil
}

// Clear (https://wiki.libsdl.org/SDL_RenderClear)
func (r *Renderer) Clear() error {
	_ret := C.SDL_RenderClear(r.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Present (https://wiki.libsdl.org/SDL_RenderPresent)
func (r *Renderer) Present() {
	C.SDL_RenderPresent(r.cptr())
}

// SetDrawColor (https://wiki.libsdl.org/SDL_SetRenderDrawColor)
func (r *Renderer) SetDrawColor(re, g, b, a uint8) error {
	_r := C.Uint8(re)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	_ret := C.SDL_SetRenderDrawColor(r.cptr(), _r, _g, _b, _a)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawPoint (https://wiki.libsdl.org/SDL_RenderDrawPoint)
func (r *Renderer) DrawPoint(x, y int) error {
	_ret := C.SDL_RenderDrawPoint(r.cptr(), C.int(x), C.int(y))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawPoints (https://wiki.libsdl.org/SDL_RenderDrawPoints)
func (r *Renderer) DrawPoints(points []Point) error {
	_ret := C.SDL_RenderDrawPoints(r.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawLine (https://wiki.libsdl.org/SDL_RenderDrawLine)
func (r *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	_x1 := C.int(x1)
	_y1 := C.int(y1)
	_x2 := C.int(x2)
	_y2 := C.int(y2)
	_ret := C.SDL_RenderDrawLine(r.cptr(), _x1, _y1, _x2, _y2)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawLines (https://wiki.libsdl.org/SDL_RenderDrawLines)
func (r *Renderer) DrawLines(points []Point) error {
	_ret := C.SDL_RenderDrawLines(r.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawRect (https://wiki.libsdl.org/SDL_RenderDrawRect)
func (r *Renderer) DrawRect(rect *Rect) error {
	_ret := C.SDL_RenderDrawRect(r.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// DrawRects (https://wiki.libsdl.org/SDL_RenderDrawRects)
func (r *Renderer) DrawRects(rects []Rect) error {
	_ret := C.SDL_RenderDrawRects(r.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// FillRect (https://wiki.libsdl.org/SDL_RenderFillRect)
func (r *Renderer) FillRect(rect *Rect) error {
	_ret := C.SDL_RenderFillRect(r.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// FillRects (https://wiki.libsdl.org/SDL_RenderFillRects)
func (r *Renderer) FillRects(rects []Rect) error {
	_ret := C.SDL_RenderFillRects(r.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Destroy (https://wiki.libsdl.org/SDL_DestroyRenderer)
func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer(r.cptr())
}
