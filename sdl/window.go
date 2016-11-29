package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

// Window SDL window
type Window C.SDL_Window

const (
	windowFullscreen        = C.SDL_WINDOW_FULLSCREEN
	windowOpenGl            = C.SDL_WINDOW_OPENGL
	windowShown             = C.SDL_WINDOW_SHOWN
	windowHidden            = C.SDL_WINDOW_HIDDEN
	windowBorderless        = C.SDL_WINDOW_BORDERLESS
	windowResizable         = C.SDL_WINDOW_RESIZABLE
	windowMinimized         = C.SDL_WINDOW_MINIMIZED
	windowMaximized         = C.SDL_WINDOW_MAXIMIZED
	windowInputGrabbed      = C.SDL_WINDOW_INPUT_GRABBED
	windowInputFocus        = C.SDL_WINDOW_INPUT_FOCUS
	windowMouseFocus        = C.SDL_WINDOW_MOUSE_FOCUS
	windowFullscreenDesktop = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	windowForeign           = C.SDL_WINDOW_FOREIGN
	windowAllowHighDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI
)

const (
	windowPosUndefinedMask = C.SDL_WINDOWPOS_UNDEFINED_MASK
	windowPosUndefined     = C.SDL_WINDOWPOS_UNDEFINED
	windowPosCenteredMask  = C.SDL_WINDOWPOS_CENTERED_MASK
	windowPosCentered      = C.SDL_WINDOWPOS_CENTERED
)

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

// SetFullscreen (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (w *Window) setFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(w.cptr(), C.Uint32(flags)) != 0 {
		return getError()
	}
	return nil
}

// CreateWindow (https://wiki.libsdl.org/SDL_CreateWindow)
func createWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error) {
	_title := C.CString(title)
	var _window = C.SDL_CreateWindow(_title, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, getError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

func (w *Window) getDrawableSize() (int, int) {
	var _w C.int
	var _h C.int
	C.SDL_GL_GetDrawableSize(w.cptr(), &_w, &_h)

	return int(_w), int(_h)
}

// Destroy (https://wiki.libsdl.org/SDL_DestroyWindow)
func (w *Window) destroy() {
	C.SDL_DestroyWindow(w.cptr())
}
