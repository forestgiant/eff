package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Font SDL TTF Font
type Font struct {
	path    string
	size    int
	sdlFont *C.TTF_Font
}

// Path the file path to the font
func (f *Font) Path() string {
	return f.path
}

// Size the size of the font
func (f *Font) Size() int {
	return f.size
}

func (f *Font) refresh(scale float64) error {
	if f.path == "" {
		return errors.New("fontPath is empty")
	}

	_fontPath := C.CString(f.path)
	pointSize := int(float64(f.size) * scale)
	var _font = C.TTF_OpenFont(_fontPath, C.int(pointSize))

	if _font == nil {
		return getTTFError()
	}

	f.sdlFont = _font

	return nil
}

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
func openFont(fontPath string, points int, scale float64) (*Font, error) {
	if fontPath == "" {
		return nil, errors.New("fontPath is empty")
	}

	_fontPath := C.CString(fontPath)
	pointSize := int(float64(points) * scale)
	var _font = C.TTF_OpenFont(_fontPath, C.int(pointSize))

	if _font == nil {
		return nil, getTTFError()
	}

	f := Font{
		path:    fontPath,
		size:    points,
		sdlFont: (*C.TTF_Font)(unsafe.Pointer(_font)),
	}
	return &f, nil
}

// RenderTextSolid (https://www.libsdl.org/projects/SDL_ttf)
func renderTextSolid(font *Font, text string, c color) (*surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_color := c.cptr()
	_surface := C.TTF_RenderText_Solid(font.sdlFont, _text, *_color)

	if _surface == nil {
		return nil, getTTFError()
	}

	return (*surface)(unsafe.Pointer(_surface)), nil
}

// RenderTextBlended (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC51)
func renderTextBlended(font *Font, text string, c color) (*surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_color := c.cptr()
	// defer C.free(unsafe.Pointer(_color))
	_surface := C.TTF_RenderText_Blended(font.sdlFont, _text, *_color)

	if _surface == nil {
		return nil, getTTFError()
	}

	return (*surface)(unsafe.Pointer(_surface)), nil
}

// SizeText https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC39
func sizeText(font *Font, text string) (int, int, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var _w C.int
	var _h C.int
	err := C.TTF_SizeText(font.sdlFont, _text, &_w, &_h)
	if err != 0 {
		return 0, 0, getTTFError()
	}

	return int(_w), int(_h), nil
}
