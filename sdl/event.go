package sdl

// #include "wrapper.h"
import "C"
import "unsafe"

const (
	// EventQuit (https://wiki.libsdl.org/SDL_EventType#SDL_QUIT)
	EventQuit = C.SDL_QUIT
	// EventKeyDown (https://wiki.libsdl.org/SDL_KeyboardEvent)
	EventKeyDown = C.SDL_KEYDOWN
	// EventKeyUp (https://wiki.libsdl.org/SDL_KeyboardEvent)
	EventKeyUp = C.SDL_KEYUP
)

// Event (https://wiki.libsdl.org/SDL_Event)
type Event interface{}

type cEvent struct {
	Type uint32
	_    [52]byte // padding
}

// PollEvent (https://wiki.libsdl.org/SDL_PollEvent)
func PollEvent() Event {
	var cevent C.SDL_Event
	ret := C.SDL_PollEvent(&cevent)
	if ret == 0 {
		return nil
	}
	return goEvent((*cEvent)(unsafe.Pointer(&cevent)))
}

func goEvent(cevent *cEvent) Event {
	switch cevent.Type {
	case EventKeyDown:
		return (*KeyDownEvent)(unsafe.Pointer(cevent))
	case EventKeyUp:
		return (*KeyUpEvent)(unsafe.Pointer(cevent))
	case EventQuit:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	default:
		return (*CommonEvent)(unsafe.Pointer(cevent))
	}
}

// QuitEvent (https://wiki.libsdl.org/SDL_QuitEvent)
type QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

// KeyUpEvent (https://wiki.libsdl.org/SDL_KeyboardEvent)
type KeyUpEvent struct {
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
type KeyDownEvent struct {
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
type CommonEvent struct {
	Type      uint32
	Timestamp uint32
}
type cKeyboardEvent C.SDL_KeyboardEvent
