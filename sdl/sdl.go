package sdl

// #cgo windows CFLAGS: -Ic:/mingw_dev_lib/include/SDL2
// #cgo windows LDFLAGS: -Lc:/mingw_dev_lib/lib -lSDL2 -lSDL2_ttf
// #cgo linux freebsd darwin pkg-config: sdl2
// #cgo linux freebsd darwin LDFLAGS: -lSDL2_ttf
// #include "wrapper.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func init() {
	InitTTF()
}

// Point is a structure that defines a two demensional point.
// (https://wiki.libsdl.org/SDL_Point)
type Point struct {
	X int32
	Y int32
}

func (p *Point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

// Rect is a structure that defines a rectangle, with the origin at the upper
// left.
// (https://wiki.libsdl.org/SDL_Rect)
type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (a *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(a))
}

// Color defines a color using r, g, b, a values from 0-255
// (https://wiki.libsdl.org/SDL_Color)
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (a *Color) cptr() *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(a))
}

//MainThread manages the thread that SDL calls execute on
var MainThread = make(chan func())

// LockMain calls runtime.LockOSThread on the calling thread.  This is intended to be the main thread since SDL on some platforms requires the main thread.  Use the MainThread channel to execute SDL calls.
func LockMain() {
	runtime.LockOSThread()

	for {
		f := <-MainThread
		f()
	}
}

// GetTicks (https://wiki.libsdl.org/SDL_GetTicks)
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// Delay (https://wiki.libsdl.org/SDL_Delay)
func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

// GetError (https://wiki.libsdl.org/SDL_GetError)
func GetError() error {
	if err := C.SDL_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// ClearError (https://wiki.libsdl.org/SDL_ClearError)
func ClearError() {
	C.SDL_ClearError()
}
