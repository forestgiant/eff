package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

//Renderer SDL renderer

const (
	rendererSoftware      = C.SDL_RENDERER_SOFTWARE
	rendererAccelerated   = C.SDL_RENDERER_ACCELERATED
	rendererPresentVsync  = C.SDL_RENDERER_PRESENTVSYNC
	rendererTargetTexture = C.SDL_RENDERER_TARGETTEXTURE

	textureAccessStatic    = C.SDL_TEXTUREACCESS_STATIC
	textureAccessStreaming = C.SDL_TEXTUREACCESS_STREAMING
	textureAccessTarget    = C.SDL_TEXTUREACCESS_TARGET

	textureModulateNone  = C.SDL_TEXTUREMODULATE_NONE
	textureModulateColor = C.SDL_TEXTUREMODULATE_COLOR
	textureModulateAlpha = C.SDL_TEXTUREMODULATE_ALPHA

	blendModeNone  = C.SDL_BLENDMODE_NONE
	blendModeBlend = C.SDL_BLENDMODE_BLEND
	blendModeAdd   = C.SDL_BLENDMODE_ADD
	blendModeMod   = C.SDL_BLENDMODE_MOD
)

// Texture SDL Surface (https://wiki.libsdl.org/SDL_Texture)
type texture C.SDL_Texture

func (a *texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(a))
}

func (a *texture) destroy() {
	C.SDL_DestroyTexture(a.cptr())
}

// Renderer (https://wiki.libsdl.org/SDL_CreateRenderer)
type renderer C.SDL_Renderer

func (r *renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

// CreateRenderer (https://wiki.libsdl.org/SDL_CreateRenderer)
func createRenderer(window *Window, index int, flags uint32) (*renderer, error) {
	_renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if _renderer == nil {
		return nil, getError()
	}
	return (*renderer)(unsafe.Pointer(_renderer)), nil
}

// Clear (https://wiki.libsdl.org/SDL_RenderClear)
func (r *renderer) clear() error {
	_ret := C.SDL_RenderClear(r.cptr())
	if _ret < 0 {
		return getError()
	}
	return nil
}

// setDrawBlendMode (https://wiki.libsdl.org/SDL_SetRenderDrawBlendMode)
func (r *renderer) setDrawBlendMode(blendMode C.SDL_BlendMode) error {
	_ret := C.SDL_SetRenderDrawBlendMode(r.cptr(), blendMode)
	if _ret < 0 {
		return getError()
	}
	return nil
}

// Present (https://wiki.libsdl.org/SDL_RenderPresent)
func (r *renderer) present() {
	C.SDL_RenderPresent(r.cptr())
}

// SetDrawColor (https://wiki.libsdl.org/SDL_SetRenderDrawColor)
func (r *renderer) setDrawColor(re, g, b, a uint8) error {
	_r := C.Uint8(re)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	_ret := C.SDL_SetRenderDrawColor(r.cptr(), _r, _g, _b, _a)
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawPoint (https://wiki.libsdl.org/SDL_RenderDrawPoint)
func (r *renderer) drawPoint(x, y int) error {
	_ret := C.SDL_RenderDrawPoint(r.cptr(), C.int(x), C.int(y))
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawPoints (https://wiki.libsdl.org/SDL_RenderDrawPoints)
func (r *renderer) drawPoints(points []point) error {
	_ret := C.SDL_RenderDrawPoints(r.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawLine (https://wiki.libsdl.org/SDL_RenderDrawLine)
func (r *renderer) drawLine(x1, y1, x2, y2 int) error {
	_x1 := C.int(x1)
	_y1 := C.int(y1)
	_x2 := C.int(x2)
	_y2 := C.int(y2)
	_ret := C.SDL_RenderDrawLine(r.cptr(), _x1, _y1, _x2, _y2)
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawLines (https://wiki.libsdl.org/SDL_RenderDrawLines)
func (r *renderer) drawLines(points []point) error {
	_ret := C.SDL_RenderDrawLines(r.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawRect (https://wiki.libsdl.org/SDL_RenderDrawRect)
func (r *renderer) drawRect(rect *rect) error {
	_ret := C.SDL_RenderDrawRect(r.cptr(), rect.cptr())
	if _ret < 0 {
		return getError()
	}
	return nil
}

// DrawRects (https://wiki.libsdl.org/SDL_RenderDrawRects)
func (r *renderer) drawRects(rects []rect) error {
	_ret := C.SDL_RenderDrawRects(r.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return getError()
	}
	return nil
}

// FillRect (https://wiki.libsdl.org/SDL_RenderFillRect)
func (r *renderer) fillRect(rect *rect) error {
	_ret := C.SDL_RenderFillRect(r.cptr(), rect.cptr())
	if _ret < 0 {
		return getError()
	}
	return nil
}

// FillRects (https://wiki.libsdl.org/SDL_RenderFillRects)
func (r *renderer) fillRects(rects []rect) error {
	_ret := C.SDL_RenderFillRects(r.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return getError()
	}
	return nil
}

// Destroy (https://wiki.libsdl.org/SDL_DestroyRenderer)
func (r *renderer) destroy() {
	C.SDL_DestroyRenderer(r.cptr())
}

// CreateTextureFromSurface (https://wiki.libsdl.org/SDL_CreateTextureFromSurface)
func (r *renderer) createTextureFromSurface(surface *surface) (*texture, error) {
	_texture := C.SDL_CreateTextureFromSurface(r.cptr(), surface.cptr())
	if _texture == nil {
		return nil, getError()
	}

	return (*texture)(unsafe.Pointer(_texture)), nil
}

// RenderCopy (https://wiki.libsdl.org/SDL_CreateTextureFromSurface)
func (r *renderer) renderCopy(texture *texture, srcRect rect, destRect rect) error {
	err := C.SDL_RenderCopy(r.cptr(), texture.cptr(), srcRect.cptr(), destRect.cptr())
	if err < 0 {
		return getError()
	}

	return nil
}
