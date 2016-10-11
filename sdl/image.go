package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// InitImg (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC8)
func InitImg() error {
	flags := C.IMG_INIT_PNG | C.IMG_INIT_JPG
	if C.IMG_Init(C.int(flag))&flags == 0 {
		return GetMixError()
	}

	return nil
}

// GetImgError (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC45)
func GetImgError() error {
	if err := C.IMG_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// LoadImg (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC11)
func LoadImg(path string) (*Surface, error) {
	_path := C.CString(path)
	_surface := IMG_Load("sample.png")
	if _surface == nil {
		return nil, GetImgError()
	}

	return (*Surface)(unsafe.Pointer(_surface)), nil
}
