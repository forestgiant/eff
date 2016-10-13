package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

type imageTex struct {
	texture *texture
	w       int32
	h       int32
}

// InitImg (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC8)
func initImg() error {
	flags := C.IMG_INIT_PNG | C.IMG_INIT_JPG
	if C.IMG_Init(C.int(flags))&C.int(flags) == 0 {
		return getImgError()
	}

	return nil
}

// GetImgError (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC45)
func getImgError() error {
	if err := C.IMG_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// LoadImg (https://www.libsdl.org/projects/SDL_image/docs/SDL_image.html#SEC11)
func loadImg(path string) (*surface, error) {
	_path := C.CString(path)
	_surface := C.IMG_Load(_path)
	if _surface == nil {
		return nil, getImgError()
	}

	return (*surface)(unsafe.Pointer(_surface)), nil
}
