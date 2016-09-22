package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	// FIRSTEVENT              = C.SDL_FIRSTEVENT
	EventQuit = C.SDL_QUIT
	// APP_TERMINATING         = C.SDL_APP_TERMINATING
	// APP_LOWMEMORY           = C.SDL_APP_LOWMEMORY
	// APP_WILLENTERBACKGROUND = C.SDL_APP_WILLENTERBACKGROUND
	// APP_DIDENTERBACKGROUND  = C.SDL_APP_DIDENTERBACKGROUND
	// APP_WILLENTERFOREGROUND = C.SDL_APP_WILLENTERFOREGROUND
	// APP_DIDENTERFOREGROUND  = C.SDL_APP_DIDENTERFOREGROUND

	/* Window events */
	// WINDOWEVENT = C.SDL_WINDOWEVENT
	// SYSWMEVENT  = C.SDL_SYSWMEVENT

	/* Keyboard events */
	EventKeyDown = C.SDL_KEYDOWN
	EventKeyUp   = C.SDL_KEYUP
	// TEXTEDITING = C.SDL_TEXTEDITING
	// TEXTINPUT   = C.SDL_TEXTINPUT

	/* Mouse events */
	// MOUSEMOTION     = C.SDL_MOUSEMOTION
	// MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	// MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	// MOUSEWHEEL      = C.SDL_MOUSEWHEEL
	//
	// /* Joystick events */
	// JOYAXISMOTION    = C.SDL_JOYAXISMOTION
	// JOYBALLMOTION    = C.SDL_JOYBALLMOTION
	// JOYHATMOTION     = C.SDL_JOYHATMOTION
	// JOYBUTTONDOWN    = C.SDL_JOYBUTTONDOWN
	// JOYBUTTONUP      = C.SDL_JOYBUTTONUP
	// JOYDEVICEADDED   = C.SDL_JOYDEVICEADDED
	// JOYDEVICEREMOVED = C.SDL_JOYDEVICEREMOVED
	//
	// /* Game controller events */
	// CONTROLLERAXISMOTION     = C.SDL_CONTROLLERAXISMOTION
	// CONTROLLERBUTTONDOWN     = C.SDL_CONTROLLERBUTTONDOWN
	// CONTROLLERBUTTONUP       = C.SDL_CONTROLLERBUTTONUP
	// CONTROLLERDEVICEADDED    = C.SDL_CONTROLLERDEVICEADDED
	// CONTROLLERDEVICEREMOVED  = C.SDL_CONTROLLERDEVICEREMOVED
	// CONTROLLERDEVICEREMAPPED = C.SDL_CONTROLLERDEVICEREMAPPED

// 	/* Touch events */
// 	FINGERDOWN   = C.SDL_FINGERDOWN
// 	FINGERUP     = C.SDL_FINGERUP
// 	FINGERMOTION = C.SDL_FINGERMOTION
//
// 	/* Gesture events */
// 	DOLLARGESTURE = C.SDL_DOLLARGESTURE
// 	DOLLARRECORD  = C.SDL_DOLLARRECORD
// 	MULTIGESTURE  = C.SDL_MULTIGESTURE
//
// 	/* Clipboard events */
// 	CLIPBOARDUPDATE = C.SDL_CLIPBOARDUPDATE
//
// 	/* Drag and drop events */
// 	DROPFILE = C.SDL_DROPFILE
//
// 	/* Render events */
// 	RENDER_TARGETS_RESET = C.SDL_RENDER_TARGETS_RESET
//
// 	USEREVENT = C.SDL_USEREVENT
// 	LASTEVENT = C.SDL_LASTEVENT
)

//
// const (
// 	ADDEVENT  = C.SDL_ADDEVENT
// 	PEEKEVENT = C.SDL_PEEKEVENT
// 	GETEVENT  = C.SDL_GETEVENT
// )
//
// const (
// 	QUERY   = C.SDL_QUERY
// 	IGNORE  = C.SDL_IGNORE
// 	DISABLE = C.SDL_DISABLE
// 	ENABLE  = C.SDL_ENABLE
// )

// Event (https://wiki.libsdl.org/SDL_Event)
type Event interface{}

type CEvent struct {
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
	return goEvent((*CEvent)(unsafe.Pointer(&cevent)))
}

func goEvent(cevent *CEvent) Event {
	switch cevent.Type {
	// case WINDOWEVENT:
	// 	return (*WindowEvent)(unsafe.Pointer(cevent))
	// case SYSWMEVENT:
	// 	return (*SysWMEvent)(unsafe.Pointer(cevent))
	case EventKeyDown:
		return (*KeyDownEvent)(unsafe.Pointer(cevent))
	case EventKeyUp:
		return (*KeyUpEvent)(unsafe.Pointer(cevent))
		// case TEXTEDITING:
		// 	return (*TextEditingEvent)(unsafe.Pointer(cevent))
		// case TEXTINPUT:
		// 	return (*TextInputEvent)(unsafe.Pointer(cevent))
		// case MOUSEMOTION:
		// 	return (*MouseMotionEvent)(unsafe.Pointer(cevent))
		// case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		// 	return (*MouseButtonEvent)(unsafe.Pointer(cevent))
		// case MOUSEWHEEL:
		// 	return (*MouseWheelEvent)(unsafe.Pointer(cevent))
		// case JOYAXISMOTION:
		// 	return (*JoyAxisEvent)(unsafe.Pointer(cevent))
		// case JOYBALLMOTION:
		// 	return (*JoyBallEvent)(unsafe.Pointer(cevent))
		// case JOYHATMOTION:
		// 	return (*JoyHatEvent)(unsafe.Pointer(cevent))
		// case JOYBUTTONDOWN, JOYBUTTONUP:
		// 	return (*JoyButtonEvent)(unsafe.Pointer(cevent))
		// case JOYDEVICEADDED, JOYDEVICEREMOVED:
		// 	return (*JoyDeviceEvent)(unsafe.Pointer(cevent))
		// case CONTROLLERAXISMOTION:
		// 	return (*ControllerAxisEvent)(unsafe.Pointer(cevent))
		// case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		// 	return (*ControllerButtonEvent)(unsafe.Pointer(cevent))
		// case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		// 	return (*ControllerDeviceEvent)(unsafe.Pointer(cevent))
		// case FINGERDOWN, FINGERUP, FINGERMOTION:
		// 	return (*TouchFingerEvent)(unsafe.Pointer(cevent))
		// case DOLLARGESTURE, DOLLARRECORD:
		// 	return (*DollarGestureEvent)(unsafe.Pointer(cevent))
		// case MULTIGESTURE:
		// 	return (*MultiGestureEvent)(unsafe.Pointer(cevent))
		// case DROPFILE:
		// 	return (*DropEvent)(unsafe.Pointer(cevent))
		// case RENDER_TARGETS_RESET:
		// 	return (*RenderEvent)(unsafe.Pointer(cevent))
	case EventQuit:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	// case USEREVENT:
	// 	return (*UserEvent)(unsafe.Pointer(cevent))
	// case CLIPBOARDUPDATE:
	// 	return (*ClipboardEvent)(unsafe.Pointer(cevent))
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

// CommonEvent
type CommonEvent struct {
	Type      uint32
	Timestamp uint32
}
type cKeyboardEvent C.SDL_KeyboardEvent
