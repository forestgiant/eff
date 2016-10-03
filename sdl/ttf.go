package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// Font SDL TTF Font
type Font C.TTF_Font

var ttfInitialized bool

// InitTTF initialize the SDL_ttf
func InitTTF() error {
	fmt.Println("Initializing TTF")
	if C.TTF_Init() == -1 {
		return GetTTFError()
	}

	ttfInitialized = true
	return nil
}

// GetTTFError (https://wiki.libsdl.org/SDL_GetError)
func GetTTFError() error {
	if err := C.TTF_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// OpenFont (https://www.libsdl.org/projects/SDL_ttf)
func OpenFont(fontPath string, pointSize int) (*Font, error) {
	_fontPath := C.CString(fontPath)
	fmt.Println("Opening font")
	var _font = C.TTF_OpenFont(_fontPath, C.int(pointSize))
	fmt.Println("Done with font")

	if _font == nil {
		return nil, GetTTFError()
	}
	return (*Font)(unsafe.Pointer(_font)), nil
}

// RenderTextSolid (https://www.libsdl.org/projects/SDL_ttf)
func RenderTextSolid(font *Font, text string, color Color) (*Surface, error) {
	_text := C.CString(text)
	_color := color.cptr()
	_surface := C.TTF_RenderText_Solid(font, _text, *_color)

	if _surface == nil {
		return nil, GetTTFError()
	}

	return (*Surface)(unsafe.Pointer(_surface)), nil
}
