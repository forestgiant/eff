package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/forestgiant/eff"
)

type Image struct {
	rect          eff.Rect
	parent        eff.Drawable
	scale         float64
	graphics      *Graphics
	texture       *texture
	children      []eff.Drawable
	updateHandler func()
}

func (image *Image) Draw(canvas eff.Canvas) {
	if image.graphics.renderer == nil {
		return
	}

	if image.texture == nil {
		fmt.Println("image texture is nil")
		return
	}

	r1 := rect{
		X: 0,
		Y: 0,
		W: int32(image.rect.W),
		H: int32(image.rect.H),
	}

	r := rect{
		X: int32(float64(image.rect.X) * image.scale),
		Y: int32(float64(image.rect.Y) * image.scale),
		W: int32(float64(image.rect.W) * image.scale),
		H: int32(float64(image.rect.H) * image.scale),
	}
	image.graphics.renderer.renderCopy(image.texture, r1, r)
}

func (image *Image) SetUpdateHandler(handler func()) {
	image.updateHandler = handler
}

func (image *Image) HandleUpdate() {
	if image.updateHandler != nil {
		image.updateHandler()
	}
}

func (image *Image) AddChild(d eff.Drawable) {
	if d == nil {
		return
	}

	d.SetParent(eff.Drawable(image))
	d.SetScale(image.scale)
	d.SetGraphics(image.graphics)

	image.children = append(image.children, d)
}

func (image *Image) RemoveChild(d eff.Drawable) {
	if d == nil {
		return
	}

	index := -1
	for i, child := range image.children {
		if d == child {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	image.children[index].SetParent(nil)
	image.children[index].SetGraphics(nil)
	image.children[index].SetScale(1)

	image.children = append(image.children[:index], image.children[index+1:]...)
}

func (image *Image) Children() []eff.Drawable {
	return image.children
}

func (image *Image) Rect() eff.Rect {
	return image.rect
}

func (image *Image) SetParent(d eff.Drawable) {
	image.parent = d
}

func (image *Image) Parent() eff.Drawable {
	return image.parent
}

func (image *Image) SetGraphics(g eff.Graphics) {
	sdlGraphics, ok := g.(*Graphics)
	if ok {
		image.graphics = sdlGraphics
	}
}

func (image *Image) Graphics() eff.Graphics {
	return image.graphics
}

func (image *Image) SetScale(s float64) {
	image.scale = s
}

func (image *Image) Scale() float64 {
	return image.scale
}

func (image *Image) SetRect(r eff.Rect) {
	image.rect = r
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
