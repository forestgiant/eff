package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// type Image struct {
// 	drawable

// 	texture  *texture
// 	graphics *Graphics
// }

// func (image *Image) Draw(canvas eff.Canvas) {
// 	if image.graphics.renderer == nil {
// 		return
// 	}

// 	if image.texture == nil {
// 		fmt.Println("image texture is nil")
// 		return
// 	}

// 	r1 := rect{
// 		X: 0,
// 		Y: 0,
// 		W: int32(image.rect.W),
// 		H: int32(image.rect.H),
// 	}

// 	r := rect{
// 		X: int32(float64(image.rect.X) * image.graphics.scale),
// 		Y: int32(float64(image.rect.Y) * image.graphics.scale),
// 		W: int32(float64(image.rect.W) * image.graphics.scale),
// 		H: int32(float64(image.rect.H) * image.graphics.scale),
// 	}
// 	image.graphics.renderer.renderCopy(image.texture, r1, r)
// }

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
