package sdl

// #cgo windows CFLAGS: -Ic:/mingw_dev_lib/include/SDL2 -w -Wl,-subsystem,windows
// #cgo windows LDFLAGS: -Lc:/mingw_dev_lib/lib -lSDL2 -lSDL2_ttf -lSDL2_mixer -lSDL2_image
// #cgo linux freebsd darwin pkg-config: sdl2
// #cgo linux freebsd darwin LDFLAGS: -lSDL2_ttf -lSDL2_mixer -lSDL2_image
// #include "wrapper.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func init() {
	initTTF()
	initMix()
	initImg()
}

// Point is a structure that defines a two demensional point.
// (https://wiki.libsdl.org/SDL_Point)
type point struct {
	X int32
	Y int32
}

func (p *point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

// Rect is a structure that defines a rectangle, with the origin at the upper
// left.
// (https://wiki.libsdl.org/SDL_Rect)
type rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (a *rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(a))
}

// Color defines a color using r, g, b, a values from 0-255
// (https://wiki.libsdl.org/SDL_Color)
type color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (a *color) cptr() *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(a))
}

// Surface SDL Surface (https://wiki.libsdl.org/SDL_Surface)
type surface C.SDL_Surface

func (a *surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(a))
}

// mainThread manages the thread that SDL calls execute on
var mainThread = make(chan func())

type callback func()

// lockMain calls runtime.LockOSThread on the calling thread.  This is intended to be the main thread since SDL on some platforms requires the main thread.  Use the MainThread channel to execute SDL calls.
func lockMain(cb callback) {
	runtime.LockOSThread()
	go cb()
	for {
		f := <-mainThread
		f()
	}
}

// GetTicks (https://wiki.libsdl.org/SDL_GetTicks)
func getTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// Delay (https://wiki.libsdl.org/SDL_Delay)
func delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

// GetError (https://wiki.libsdl.org/SDL_GetError)
func getError() error {
	if err := C.SDL_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// ClearError (https://wiki.libsdl.org/SDL_ClearError)
func clearError() {
	C.SDL_ClearError()
}

// FreeSurface (https://wiki.libsdl.org/SDL_FreeSurface)
func freeSurface(s *surface) {
	C.SDL_FreeSurface(s.cptr())
}

func quit() {
	C.Mix_Quit()
	C.TTF_Quit()
	C.IMG_Quit()
	C.SDL_Quit()
}
