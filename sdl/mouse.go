package sdl

// #include "wrapper.h"
import "C"

const (
	mouseLeft   = C.SDL_BUTTON_LEFT
	mouseMiddle = C.SDL_BUTTON_MIDDLE
	mouseRight  = C.SDL_BUTTON_RIGHT
	mouseX1     = C.SDL_BUTTON_X1
	mouseX2     = C.SDL_BUTTON_X2
)

func captureMouse(capture bool) error {
	major, minor, patch := getVersion()
	if major == 2 && minor == 0 && patch < 4 {
		return nil
	}
	var _capture C.SDL_bool
	if capture {
		_capture = 1
	}
	err := C.SDL_CaptureMouse(_capture)
	if err != 0 {
		return getError()
	}

	return nil
}
