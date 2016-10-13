package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

const (
	// EventQuit (https://wiki.libsdl.org/SDL_EventType#SDL_QUIT)
	eventQuit = C.SDL_QUIT
	// EventKeyDown (https://wiki.libsdl.org/SDL_KeyboardEvent)
	eventKeyDown = C.SDL_KEYDOWN
	// EventKeyUp (https://wiki.libsdl.org/SDL_KeyboardEvent)
	eventKeyUp = C.SDL_KEYUP
)

// Event (https://wiki.libsdl.org/SDL_Event)
type event interface{}

type cEvent struct {
	Type uint32
	_    [52]byte // padding
}

// PollEvent (https://wiki.libsdl.org/SDL_PollEvent)
func pollEvent() event {
	var cevent C.SDL_Event
	ret := C.SDL_PollEvent(&cevent)
	if ret == 0 {
		return nil
	}
	return goEvent((*cEvent)(unsafe.Pointer(&cevent)))
}

func goEvent(cevent *cEvent) event {
	switch cevent.Type {
	case eventKeyDown:
		return (*keyDownEvent)(unsafe.Pointer(cevent))
	case eventKeyUp:
		return (*keyUpEvent)(unsafe.Pointer(cevent))
	case eventQuit:
		return (*quitEvent)(unsafe.Pointer(cevent))
	default:
		return (*commonEvent)(unsafe.Pointer(cevent))
	}
}

// QuitEvent (https://wiki.libsdl.org/SDL_QuitEvent)
type quitEvent struct {
	Type      uint32
	Timestamp uint32
}

// KeyUpEvent (https://wiki.libsdl.org/SDL_KeyboardEvent)
type keyUpEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	_         uint8 // padding
	_         uint8 // padding
	Keysym    Keysym
}

// KeyDownEvent (https://wiki.libsdl.org/SDL_KeyboardEvent)
type keyDownEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	_         uint8 // padding
	_         uint8 // padding
	Keysym    Keysym
}

// CommonEvent ()
type commonEvent struct {
	Type      uint32
	Timestamp uint32
}
type cKeyboardEvent C.SDL_KeyboardEvent
