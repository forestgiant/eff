package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Font SDL TTF Font
type font C.TTF_Font

var ttfInitialized bool

// InitTTF initialize the SDL_ttf
func initTTF() error {
	if C.TTF_Init() == -1 {
		return getTTFError()
	}

	ttfInitialized = true
	return nil
}

// GetTTFError (https://wiki.libsdl.org/SDL_GetError)
func getTTFError() error {
	if err := C.TTF_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// OpenFont (https://www.libsdl.org/projects/SDL_ttf)
func openFont(fontPath string, pointSize int) (*font, error) {
	_fontPath := C.CString(fontPath)
	var _font = C.TTF_OpenFont(_fontPath, C.int(pointSize))

	if _font == nil {
		return nil, getTTFError()
	}
	return (*font)(unsafe.Pointer(_font)), nil
}

// RenderTextSolid (https://www.libsdl.org/projects/SDL_ttf)
func renderTextSolid(font *font, text string, c color) (*surface, error) {
	_text := C.CString(text)
	_color := c.cptr()
	_surface := C.TTF_RenderText_Solid(font, _text, *_color)

	if _surface == nil {
		return nil, getTTFError()
	}

	return (*surface)(unsafe.Pointer(_surface)), nil
}

// RenderTextBlended (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC51)
func renderTextBlended(font *font, text string, c color) (*surface, error) {
	_text := C.CString(text)
	_color := c.cptr()
	_surface := C.TTF_RenderText_Blended(font, _text, *_color)

	if _surface == nil {
		return nil, getTTFError()
	}

	return (*surface)(unsafe.Pointer(_surface)), nil
}
