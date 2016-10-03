package sdl

import "C"
import (
	"unsafe"
)

// Font SDL TTF Font
type Font C.TTF_Font

// OpenFont (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC14)
func OpenFont(fontPath string, pointSize int) (*Font, error) {
	_fontPath := C.CString(fontPath)
	var _font = C.TTF_OpenFont(_fontPath, C.int(pointSize))
	if _font == nil {
		return nil, sdl.GetError()
	}
	return (*Font)(unsafe.Pointer(_font)), nil
}

// RenderTextSolid (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC43)
func RenderTextSolid(renderer Renderer, font *Font, text string, color Color) (*Surface, error) {
	_text := C.CString(text)
	_color := color.cptr()
	_surface := C.TTF_RenderText_Solid(font, _text, _color)

	return _surface, nil
}
