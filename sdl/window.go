package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

// Window SDL window
type Window C.SDL_Window

const (
	WindowFullscreen        = C.SDL_WINDOW_FULLSCREEN
	WindowOpenGl            = C.SDL_WINDOW_OPENGL
	WindowShown             = C.SDL_WINDOW_SHOWN
	WindowHidden            = C.SDL_WINDOW_HIDDEN
	WindowBorderless        = C.SDL_WINDOW_BORDERLESS
	WindowResizable         = C.SDL_WINDOW_RESIZABLE
	WindowMinimized         = C.SDL_WINDOW_MINIMIZED
	WindowMaximized         = C.SDL_WINDOW_MAXIMIZED
	WindowInputGrabbed      = C.SDL_WINDOW_INPUT_GRABBED
	WindowInputFocus        = C.SDL_WINDOW_INPUT_FOCUS
	WindowMouseFocus        = C.SDL_WINDOW_MOUSE_FOCUS
	WindowFullscreenDesktop = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WindowForeign           = C.SDL_WINDOW_FOREIGN
	WindowAllowHighDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI
)

const (
	WindowPosUndefinedMask = C.SDL_WINDOWPOS_UNDEFINED_MASK
	WindowPosUndefined     = C.SDL_WINDOWPOS_UNDEFINED
	WindowPosCenteredMask  = C.SDL_WINDOWPOS_CENTERED_MASK
	WindowPosCentered      = C.SDL_WINDOWPOS_CENTERED
)

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

// SetFullscreen (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (w *Window) SetFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(w.cptr(), C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// CreateWindow (https://wiki.libsdl.org/SDL_CreateWindow)
func CreateWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error) {
	_title := C.CString(title)
	var _window = C.SDL_CreateWindow(_title, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// Destroy (https://wiki.libsdl.org/SDL_DestroyWindow)
func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.cptr())
}
