package sdl

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "sdl_wrapper.h"
import "C"
import "unsafe"

//Renderer SDL renderer
type Renderer C.SDL_Renderer

//Window SDL window
type Window C.SDL_Window

//CallQueue manages the thread that SDL calls execute on
var CallQueue = make(chan func(), 1)

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

const (
	RendererSoftware      = C.SDL_RENDERER_SOFTWARE
	RendererAccelerated   = C.SDL_RENDERER_ACCELERATED
	RendererPresentVsync  = C.SDL_RENDERER_PRESENTVSYNC
	RendererTargetTexture = C.SDL_RENDERER_TARGETTEXTURE

	TextureAccessStatic    = C.SDL_TEXTUREACCESS_STATIC
	TextureAccessStreaming = C.SDL_TEXTUREACCESS_STREAMING
	TextureAccessTarget    = C.SDL_TEXTUREACCESS_TARGET

	TextureModulateNone  = C.SDL_TEXTUREMODULATE_NONE
	TextureModulateColor = C.SDL_TEXTUREMODULATE_COLOR
	TextureModulateAlpha = C.SDL_TEXTUREMODULATE_ALPHA
)

// Window (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (window *Window) SetFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(window.cptr(), C.Uint32(flags)) != 0 {
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

// CreateRenderer (https://wiki.libsdl.org/SDL_CreateRenderer)
func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error) {
	_renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if _renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(_renderer)), nil
}

// Clear (https://wiki.libsdl.org/SDL_RenderClear)
func (renderer *Renderer) Clear() error {
	_ret := C.SDL_RenderClear(renderer.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

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
	case WINDOWEVENT:
		return (*WindowEvent)(unsafe.Pointer(cevent))
	case SYSWMEVENT:
		return (*SysWMEvent)(unsafe.Pointer(cevent))
	case KEYDOWN:
		return (*KeyDownEvent)(unsafe.Pointer(cevent))
	case KEYUP:
		return (*KeyUpEvent)(unsafe.Pointer(cevent))
	case TEXTEDITING:
		return (*TextEditingEvent)(unsafe.Pointer(cevent))
	case TEXTINPUT:
		return (*TextInputEvent)(unsafe.Pointer(cevent))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(unsafe.Pointer(cevent))
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return (*MouseButtonEvent)(unsafe.Pointer(cevent))
	case MOUSEWHEEL:
		return (*MouseWheelEvent)(unsafe.Pointer(cevent))
	case JOYAXISMOTION:
		return (*JoyAxisEvent)(unsafe.Pointer(cevent))
	case JOYBALLMOTION:
		return (*JoyBallEvent)(unsafe.Pointer(cevent))
	case JOYHATMOTION:
		return (*JoyHatEvent)(unsafe.Pointer(cevent))
	case JOYBUTTONDOWN, JOYBUTTONUP:
		return (*JoyButtonEvent)(unsafe.Pointer(cevent))
	case JOYDEVICEADDED, JOYDEVICEREMOVED:
		return (*JoyDeviceEvent)(unsafe.Pointer(cevent))
	case CONTROLLERAXISMOTION:
		return (*ControllerAxisEvent)(unsafe.Pointer(cevent))
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		return (*ControllerButtonEvent)(unsafe.Pointer(cevent))
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		return (*ControllerDeviceEvent)(unsafe.Pointer(cevent))
	case FINGERDOWN, FINGERUP, FINGERMOTION:
		return (*TouchFingerEvent)(unsafe.Pointer(cevent))
	case DOLLARGESTURE, DOLLARRECORD:
		return (*DollarGestureEvent)(unsafe.Pointer(cevent))
	case MULTIGESTURE:
		return (*MultiGestureEvent)(unsafe.Pointer(cevent))
	case DROPFILE:
		return (*DropEvent)(unsafe.Pointer(cevent))
	case RENDER_TARGETS_RESET:
		return (*RenderEvent)(unsafe.Pointer(cevent))
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	case USEREVENT:
		return (*UserEvent)(unsafe.Pointer(cevent))
	case CLIPBOARDUPDATE:
		return (*ClipboardEvent)(unsafe.Pointer(cevent))
	default:
		return (*CommonEvent)(unsafe.Pointer(cevent))
	}
}

// QuitEvent (https://wiki.libsdl.org/SDL_QuitEvent)
type QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

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
