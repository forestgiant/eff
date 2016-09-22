package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

//Renderer SDL renderer

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

// Renderer (https://wiki.libsdl.org/SDL_CreateRenderer)
type Renderer C.SDL_Renderer

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
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
