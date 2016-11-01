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
