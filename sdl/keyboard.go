package sdl

// #include "wrapper.h"
import "C"

// Keysym (https://wiki.libsdl.org/SDL_Keysym)
type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      uint16
	Unicode  uint32
}

// Scancode (https://wiki.libsdl.org/SDL_Scancode)
type Scancode uint32

// Keycode (https://wiki.libsdl.org/SDL_Keycode)
type Keycode C.SDL_Keycode

func (code Keycode) c() C.SDL_Keycode {
	return C.SDL_Keycode(code)
}

const keyScancodeMask = 1 << 30

const (
	KeyUnknown            = C.SDLK_UNKNOWN
	KeyReturn             = C.SDLK_RETURN
	KeyEscape             = C.SDLK_ESCAPE
	KeyBackspace          = C.SDLK_BACKSPACE
	KeyTab                = C.SDLK_TAB
	KeySpace              = C.SDLK_SPACE
	KeyExclaim            = C.SDLK_EXCLAIM
	KeyQuoteDbl           = C.SDLK_QUOTEDBL
	KeyHash               = C.SDLK_HASH
	KeyPercent            = C.SDLK_PERCENT
	KeyDollar             = C.SDLK_DOLLAR
	KeyAmpersand          = C.SDLK_AMPERSAND
	KeyQuote              = C.SDLK_QUOTE
	KeyLeftParen          = C.SDLK_LEFTPAREN
	KeyRightParen         = C.SDLK_RIGHTPAREN
	KeyAsterisk           = C.SDLK_ASTERISK
	KeyPlus               = C.SDLK_PLUS
	KeyComma              = C.SDLK_COMMA
	KeyMinus              = C.SDLK_MINUS
	KeyPeriod             = C.SDLK_PERIOD
	KeySlash              = C.SDLK_SLASH
	Key0                  = C.SDLK_0
	Key1                  = C.SDLK_1
	Key2                  = C.SDLK_2
	Key3                  = C.SDLK_3
	Key4                  = C.SDLK_4
	Key5                  = C.SDLK_5
	Key6                  = C.SDLK_6
	Key7                  = C.SDLK_7
	Key8                  = C.SDLK_8
	Key9                  = C.SDLK_9
	KeyColon              = C.SDLK_COLON
	KeySemicolon          = C.SDLK_SEMICOLON
	KeyLess               = C.SDLK_LESS
	KeyEquals             = C.SDLK_EQUALS
	KeyGreater            = C.SDLK_GREATER
	KeyQuestion           = C.SDLK_QUESTION
	KeyAt                 = C.SDLK_AT
	KeyLeftBracket        = C.SDLK_LEFTBRACKET
	KeyBackslash          = C.SDLK_BACKSLASH
	KeyRightBracket       = C.SDLK_RIGHTBRACKET
	KeyCaret              = C.SDLK_CARET
	KeyUnderscore         = C.SDLK_UNDERSCORE
	KeyBackQuote          = C.SDLK_BACKQUOTE
	KeyA                  = C.SDLK_a
	KeyB                  = C.SDLK_b
	KeyC                  = C.SDLK_c
	KeyD                  = C.SDLK_d
	KeyE                  = C.SDLK_e
	KeyF                  = C.SDLK_f
	KeyG                  = C.SDLK_g
	KeyH                  = C.SDLK_h
	KeyI                  = C.SDLK_i
	KeyJ                  = C.SDLK_j
	KeyK                  = C.SDLK_k
	KeyL                  = C.SDLK_l
	KeyM                  = C.SDLK_m
	KeyN                  = C.SDLK_n
	KeyO                  = C.SDLK_o
	KeyP                  = C.SDLK_p
	KeyQ                  = C.SDLK_q
	KeyR                  = C.SDLK_r
	KeyS                  = C.SDLK_s
	KeyT                  = C.SDLK_t
	KeyU                  = C.SDLK_u
	KeyV                  = C.SDLK_v
	KeyW                  = C.SDLK_w
	KeyX                  = C.SDLK_x
	KeyY                  = C.SDLK_y
	KeyZ                  = C.SDLK_z
	KeyCapslock           = C.SDLK_CAPSLOCK
	KeyF1                 = C.SDLK_F1
	KeyF2                 = C.SDLK_F2
	KeyF3                 = C.SDLK_F3
	KeyF4                 = C.SDLK_F4
	KeyF5                 = C.SDLK_F5
	KeyF6                 = C.SDLK_F6
	KeyF7                 = C.SDLK_F7
	KeyF8                 = C.SDLK_F8
	KeyF9                 = C.SDLK_F9
	KeyF10                = C.SDLK_F10
	KeyF11                = C.SDLK_F11
	KeyF12                = C.SDLK_F12
	KeyPrintScreen        = C.SDLK_PRINTSCREEN
	KeyScrollLock         = C.SDLK_SCROLLLOCK
	KeyPause              = C.SDLK_PAUSE
	KeyInsert             = C.SDLK_INSERT
	KeyHome               = C.SDLK_HOME
	KeyPageUp             = C.SDLK_PAGEUP
	KeyDelete             = C.SDLK_DELETE
	KeyEnd                = C.SDLK_END
	KeyPageDown           = C.SDLK_PAGEDOWN
	KeyRight              = C.SDLK_RIGHT
	KeyLeft               = C.SDLK_LEFT
	KeyDown               = C.SDLK_DOWN
	KeyUp                 = C.SDLK_UP
	KeyNumLockClear       = C.SDLK_NUMLOCKCLEAR
	KeyKpDivide           = C.SDLK_KP_DIVIDE
	KeyKpMultiply         = C.SDLK_KP_MULTIPLY
	KeyKpMinus            = C.SDLK_KP_MINUS
	KeyKpPlus             = C.SDLK_KP_PLUS
	KeyKpEnter            = C.SDLK_KP_ENTER
	KeyKp1                = C.SDLK_KP_1
	KeyKp2                = C.SDLK_KP_2
	KeyKp3                = C.SDLK_KP_3
	KeyKp4                = C.SDLK_KP_4
	KeyKp5                = C.SDLK_KP_5
	KeyKp6                = C.SDLK_KP_6
	KeyKp7                = C.SDLK_KP_7
	KeyKp8                = C.SDLK_KP_8
	KeyKp9                = C.SDLK_KP_9
	KeyKp0                = C.SDLK_KP_0
	KeyKpPeriod           = C.SDLK_KP_PERIOD
	KeyKpEquals           = C.SDLK_KP_EQUALS
	KeyApplication        = C.SDLK_APPLICATION
	KeyPower              = C.SDLK_POWER
	KeyF13                = C.SDLK_F13
	KeyF14                = C.SDLK_F14
	KeyF15                = C.SDLK_F15
	KeyF16                = C.SDLK_F16
	KeyF17                = C.SDLK_F17
	KeyF18                = C.SDLK_F18
	KeyF19                = C.SDLK_F19
	KeyF20                = C.SDLK_F20
	KeyF21                = C.SDLK_F21
	KeyF22                = C.SDLK_F22
	KeyF23                = C.SDLK_F23
	KeyF24                = C.SDLK_F24
	KeyExecute            = C.SDLK_EXECUTE
	KeyHelp               = C.SDLK_HELP
	KeyMenu               = C.SDLK_MENU
	KeySelect             = C.SDLK_SELECT
	KeyStop               = C.SDLK_STOP
	KeyAgain              = C.SDLK_AGAIN
	KeyUndo               = C.SDLK_UNDO
	KeyCut                = C.SDLK_CUT
	KeyCopy               = C.SDLK_COPY
	KeyPaste              = C.SDLK_PASTE
	KeyFind               = C.SDLK_FIND
	KeyMute               = C.SDLK_MUTE
	KeyVolumeUp           = C.SDLK_VOLUMEUP
	KeyVolumeDown         = C.SDLK_VOLUMEDOWN
	KeyKpComma            = C.SDLK_KP_COMMA
	KeyKpEqualsAS400      = C.SDLK_KP_EQUALSAS400
	KeyAltErase           = C.SDLK_ALTERASE
	KeySysReq             = C.SDLK_SYSREQ
	KeyCancel             = C.SDLK_CANCEL
	KeyClear              = C.SDLK_CLEAR
	KeyPrior              = C.SDLK_PRIOR
	KeyReturn2            = C.SDLK_RETURN2
	KeySepartor           = C.SDLK_SEPARATOR
	KeyOut                = C.SDLK_OUT
	KeyOper               = C.SDLK_OPER
	KeyClearAgain         = C.SDLK_CLEARAGAIN
	KeyCrSel              = C.SDLK_CRSEL
	KeyExSel              = C.SDLK_EXSEL
	KeyKp00               = C.SDLK_KP_00
	KeyKp000              = C.SDLK_KP_000
	KeyThousandsSeparator = C.SDLK_THOUSANDSSEPARATOR
	KeyDecimalSeparator   = C.SDLK_DECIMALSEPARATOR
	KeyCurrencyUnit       = C.SDLK_CURRENCYUNIT
	KeyCurrencySubUnit    = C.SDLK_CURRENCYSUBUNIT
	KeyKpLeftParen        = C.SDLK_KP_LEFTPAREN
	KeyKpRightParen       = C.SDLK_KP_RIGHTPAREN
	KeyKpLeftBrace        = C.SDLK_KP_LEFTBRACE
	KeyKpRightBrace       = C.SDLK_KP_RIGHTBRACE
	KeyKpTab              = C.SDLK_KP_TAB
	KeyKpBackspace        = C.SDLK_KP_BACKSPACE
	KeyKpA                = C.SDLK_KP_A
	KeyKpB                = C.SDLK_KP_B
	KeyKpC                = C.SDLK_KP_C
	KeyKpD                = C.SDLK_KP_D
	KeyKpE                = C.SDLK_KP_E
	KeyKpF                = C.SDLK_KP_F
	KeyKpXOR              = C.SDLK_KP_XOR
	KeyKpPower            = C.SDLK_KP_POWER
	KeyKpPercent          = C.SDLK_KP_PERCENT
	KeyKpLess             = C.SDLK_KP_LESS
	KeyKpGreater          = C.SDLK_KP_GREATER
	KeyKpAmpersand        = C.SDLK_KP_AMPERSAND
	KeyKpDblAmpersand     = C.SDLK_KP_DBLAMPERSAND
	KeyKpVerticalBar      = C.SDLK_KP_VERTICALBAR
	KeyKpDblVerticalBar   = C.SDLK_KP_DBLVERTICALBAR
	KeyKpColon            = C.SDLK_KP_COLON
	KeyKpHash             = C.SDLK_KP_HASH
	KeyKpSpace            = C.SDLK_KP_SPACE
	KeyKpAt               = C.SDLK_KP_AT
	KeyKpExclam           = C.SDLK_KP_EXCLAM
	KeyKpMemStore         = C.SDLK_KP_MEMSTORE
	KeyKpMemRecall        = C.SDLK_KP_MEMRECALL
	KeyKpMemClear         = C.SDLK_KP_MEMCLEAR
	KeyKpMemAdd           = C.SDLK_KP_MEMADD
	KeyKpMemSubtract      = C.SDLK_KP_MEMSUBTRACT
	KeyKpMemMultiply      = C.SDLK_KP_MEMMULTIPLY
	KeyKpMemDivide        = C.SDLK_KP_MEMDIVIDE
	KeyKpPlusMinus        = C.SDLK_KP_PLUSMINUS
	KeyKpClear            = C.SDLK_KP_CLEAR
	KeyKpClearEntry       = C.SDLK_KP_CLEARENTRY
	KeyKpBinary           = C.SDLK_KP_BINARY
	KeyKpOctal            = C.SDLK_KP_OCTAL
	KeyKpDecimal          = C.SDLK_KP_DECIMAL
	KeyKpHexadecimal      = C.SDLK_KP_HEXADECIMAL
	KeyLCtrl              = C.SDLK_LCTRL
	KeyLShift             = C.SDLK_LSHIFT
	KeyLAlt               = C.SDLK_LALT
	KeyLGui               = C.SDLK_LGUI
	KeyRCtrl              = C.SDLK_RCTRL
	KeyRShift             = C.SDLK_RSHIFT
	KeyRAlt               = C.SDLK_RALT
	KeyRGui               = C.SDLK_RGUI
	KeyMode               = C.SDLK_MODE
	KeyAudioNext          = C.SDLK_AUDIONEXT
	KeyAudioPrev          = C.SDLK_AUDIOPREV
	KeyAudioStop          = C.SDLK_AUDIOSTOP
	KeyAudioPlay          = C.SDLK_AUDIOPLAY
	KeyAudioMute          = C.SDLK_AUDIOMUTE
	KeyMediaSelect        = C.SDLK_MEDIASELECT
	KeyWWW                = C.SDLK_WWW
	KeyMail               = C.SDLK_MAIL
	KeyCalculator         = C.SDLK_CALCULATOR
	KeyComputer           = C.SDLK_COMPUTER
	KeyAcSearch           = C.SDLK_AC_SEARCH
	KeyAcHome             = C.SDLK_AC_HOME
	KeyAcBack             = C.SDLK_AC_BACK
	KeyAcForward          = C.SDLK_AC_FORWARD
	KeyAcStop             = C.SDLK_AC_STOP
	KeyAcRefresh          = C.SDLK_AC_REFRESH
	KeyAcBookmarks        = C.SDLK_AC_BOOKMARKS
	KeyBrightnessDown     = C.SDLK_BRIGHTNESSDOWN
	KeyBrightnessUp       = C.SDLK_BRIGHTNESSUP
	KeyDisplaySwitch      = C.SDLK_DISPLAYSWITCH
	KeyKbdIllumToggle     = C.SDLK_KBDILLUMTOGGLE
	KeyKbdIllumDown       = C.SDLK_KBDILLUMDOWN
	KeyKbdIllumUp         = C.SDLK_KBDILLUMUP
	KeyEject              = C.SDLK_EJECT
	KeySleep              = C.SDLK_SLEEP
)

const (
	modNone     = C.KMOD_NONE
	modLShift   = C.KMOD_LSHIFT
	modRShift   = C.KMOD_RSHIFT
	modLCtrl    = C.KMOD_LCTRL
	modRCtrl    = C.KMOD_RCTRL
	modLAlt     = C.KMOD_LALT
	modRAlt     = C.KMOD_RALT
	modLGui     = C.KMOD_LGUI
	modRGui     = C.KMOD_RGUI
	modNum      = C.KMOD_NUM
	modCaps     = C.KMOD_CAPS
	modMode     = C.KMOD_MODE
	modCtrl     = C.KMOD_CTRL
	modShift    = C.KMOD_SHIFT
	modAlt      = C.KMOD_ALT
	modGui      = C.KMOD_GUI
	modReserved = C.KMOD_RESERVED
)

// GetKeyName (https://wiki.libsdl.org/SDL_GetKeyName)
func getKeyName(code Keycode) string {
	return (C.GoString)(C.SDL_GetKeyName(code.c()))
}
