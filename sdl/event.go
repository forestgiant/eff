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

	eventMouseMotion     = C.SDL_MOUSEMOTION
	eventMouseButtonDown = C.SDL_MOUSEBUTTONDOWN
	eventMouseButtonUp   = C.SDL_MOUSEBUTTONUP
	eventMouseWheel      = C.SDL_MOUSEWHEEL
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
	case eventMouseButtonDown:
		return (*mouseDownEvent)(unsafe.Pointer(cevent))
	case eventMouseButtonUp:
		return (*mouseUpEvent)(unsafe.Pointer(cevent))
	case eventMouseMotion:
		return (*mouseMotionEvent)(unsafe.Pointer(cevent))
	case eventMouseWheel:
		return (*mouseWheelEvent)(unsafe.Pointer(cevent))
	default:
		return (*commonEvent)(unsafe.Pointer(cevent))
	}
}

// CommonEvent ()
type commonEvent struct {
	Type      uint32
	Timestamp uint32
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

// MouseMotionEvent (https://wiki.libsdl.org/SDL_MouseMotionEvent)
type mouseMotionEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	State     uint32
	X         int32
	Y         int32
	XRel      int32
	YRel      int32
}

// MouseButtonEvent (https://wiki.libsdl.org/SDL_MouseButtonEvent)
type mouseUpEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	Clicks    uint8 // padding
	_         uint8 // padding
	X         int32
	Y         int32
}

// MouseButtonEvent (https://wiki.libsdl.org/SDL_MouseButtonEvent)
type mouseDownEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	Clicks    uint8 // padding
	_         uint8 // padding
	X         int32
	Y         int32
}

// MouseWheelEvent (https://wiki.libsdl.org/SDL_MouseWheelEvent)
type mouseWheelEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
	Direction uint32
}
