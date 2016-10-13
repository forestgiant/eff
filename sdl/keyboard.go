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
	keyUnknown = C.SDLK_UNKNOWN

	keyReturn             = C.SDLK_RETURN
	keyEscape             = C.SDLK_ESCAPE
	keyBackspace          = C.SDLK_BACKSPACE
	keyTab                = C.SDLK_TAB
	keySpace              = C.SDLK_SPACE
	keyExclaim            = C.SDLK_EXCLAIM
	keyQuoteDbl           = C.SDLK_QUOTEDBL
	keyHash               = C.SDLK_HASH
	keyPercent            = C.SDLK_PERCENT
	keyDollar             = C.SDLK_DOLLAR
	keyAmpersand          = C.SDLK_AMPERSAND
	keyQuote              = C.SDLK_QUOTE
	keyLeftParen          = C.SDLK_LEFTPAREN
	keyRightParen         = C.SDLK_RIGHTPAREN
	keyAsterisk           = C.SDLK_ASTERISK
	keyPlus               = C.SDLK_PLUS
	keyComma              = C.SDLK_COMMA
	keyMinus              = C.SDLK_MINUS
	keyPeriod             = C.SDLK_PERIOD
	keySlash              = C.SDLK_SLASH
	key0                  = C.SDLK_0
	key1                  = C.SDLK_1
	key2                  = C.SDLK_2
	key3                  = C.SDLK_3
	key4                  = C.SDLK_4
	key5                  = C.SDLK_5
	key6                  = C.SDLK_6
	key7                  = C.SDLK_7
	key8                  = C.SDLK_8
	key9                  = C.SDLK_9
	keyColon              = C.SDLK_COLON
	keySemicolon          = C.SDLK_SEMICOLON
	keyLess               = C.SDLK_LESS
	keyEquals             = C.SDLK_EQUALS
	keyGreater            = C.SDLK_GREATER
	keyQuestion           = C.SDLK_QUESTION
	keyAt                 = C.SDLK_AT
	keyLeftBracket        = C.SDLK_LEFTBRACKET
	keyBackslash          = C.SDLK_BACKSLASH
	keyRightBracket       = C.SDLK_RIGHTBRACKET
	keyCaret              = C.SDLK_CARET
	keyUnderscore         = C.SDLK_UNDERSCORE
	keyBackQuote          = C.SDLK_BACKQUOTE
	keyA                  = C.SDLK_a
	keyB                  = C.SDLK_b
	keyC                  = C.SDLK_c
	keyD                  = C.SDLK_d
	keyE                  = C.SDLK_e
	keyF                  = C.SDLK_f
	keyG                  = C.SDLK_g
	keyH                  = C.SDLK_h
	keyI                  = C.SDLK_i
	keyJ                  = C.SDLK_j
	keyK                  = C.SDLK_k
	keyL                  = C.SDLK_l
	keyM                  = C.SDLK_m
	keyN                  = C.SDLK_n
	keyO                  = C.SDLK_o
	keyP                  = C.SDLK_p
	keyQ                  = C.SDLK_q
	keyR                  = C.SDLK_r
	keyS                  = C.SDLK_s
	keyT                  = C.SDLK_t
	keyU                  = C.SDLK_u
	keyV                  = C.SDLK_v
	keyW                  = C.SDLK_w
	keyX                  = C.SDLK_x
	keyY                  = C.SDLK_y
	keyZ                  = C.SDLK_z
	keyCapslock           = C.SDLK_CAPSLOCK
	keyF1                 = C.SDLK_F1
	keyF2                 = C.SDLK_F2
	keyF3                 = C.SDLK_F3
	keyF4                 = C.SDLK_F4
	keyF5                 = C.SDLK_F5
	keyF6                 = C.SDLK_F6
	keyF7                 = C.SDLK_F7
	keyF8                 = C.SDLK_F8
	keyF9                 = C.SDLK_F9
	keyF10                = C.SDLK_F10
	keyF11                = C.SDLK_F11
	keyF12                = C.SDLK_F12
	keyPrintScreen        = C.SDLK_PRINTSCREEN
	keyScrollLock         = C.SDLK_SCROLLLOCK
	keyPause              = C.SDLK_PAUSE
	keyInsert             = C.SDLK_INSERT
	keyHome               = C.SDLK_HOME
	keyPageUp             = C.SDLK_PAGEUP
	keyDelete             = C.SDLK_DELETE
	keyEnd                = C.SDLK_END
	keyPageDown           = C.SDLK_PAGEDOWN
	keyRight              = C.SDLK_RIGHT
	keyLeft               = C.SDLK_LEFT
	keyDown               = C.SDLK_DOWN
	keyUp                 = C.SDLK_UP
	keyNumLockClear       = C.SDLK_NUMLOCKCLEAR
	keyKpDivide           = C.SDLK_KP_DIVIDE
	keyKpMultiply         = C.SDLK_KP_MULTIPLY
	keyKpMinus            = C.SDLK_KP_MINUS
	keyKpPlus             = C.SDLK_KP_PLUS
	keyKpEnter            = C.SDLK_KP_ENTER
	keyKp1                = C.SDLK_KP_1
	keyKp2                = C.SDLK_KP_2
	keyKp3                = C.SDLK_KP_3
	keyKp4                = C.SDLK_KP_4
	keyKp5                = C.SDLK_KP_5
	keyKp6                = C.SDLK_KP_6
	keyKp7                = C.SDLK_KP_7
	keyKp8                = C.SDLK_KP_8
	keyKp9                = C.SDLK_KP_9
	keyKp0                = C.SDLK_KP_0
	keyKpPeriod           = C.SDLK_KP_PERIOD
	keyKpEquals           = C.SDLK_KP_EQUALS
	keyApplication        = C.SDLK_APPLICATION
	keyPower              = C.SDLK_POWER
	keyF13                = C.SDLK_F13
	keyF14                = C.SDLK_F14
	keyF15                = C.SDLK_F15
	keyF16                = C.SDLK_F16
	keyF17                = C.SDLK_F17
	keyF18                = C.SDLK_F18
	keyF19                = C.SDLK_F19
	keyF20                = C.SDLK_F20
	keyF21                = C.SDLK_F21
	keyF22                = C.SDLK_F22
	keyF23                = C.SDLK_F23
	keyF24                = C.SDLK_F24
	keyExecute            = C.SDLK_EXECUTE
	keyHelp               = C.SDLK_HELP
	keyMenu               = C.SDLK_MENU
	keySelect             = C.SDLK_SELECT
	keyStop               = C.SDLK_STOP
	keyAgain              = C.SDLK_AGAIN
	keyUndo               = C.SDLK_UNDO
	keyCut                = C.SDLK_CUT
	keyCopy               = C.SDLK_COPY
	keyPaste              = C.SDLK_PASTE
	keyFind               = C.SDLK_FIND
	keyMute               = C.SDLK_MUTE
	keyVolumeUp           = C.SDLK_VOLUMEUP
	keyVolumeDown         = C.SDLK_VOLUMEDOWN
	keyKpComma            = C.SDLK_KP_COMMA
	keyKpEqualsAS400      = C.SDLK_KP_EQUALSAS400
	keyAltErase           = C.SDLK_ALTERASE
	keySysReq             = C.SDLK_SYSREQ
	keyCancel             = C.SDLK_CANCEL
	keyClear              = C.SDLK_CLEAR
	keyPrior              = C.SDLK_PRIOR
	keyReturn2            = C.SDLK_RETURN2
	keySepartor           = C.SDLK_SEPARATOR
	keyOut                = C.SDLK_OUT
	keyOper               = C.SDLK_OPER
	keyClearAgain         = C.SDLK_CLEARAGAIN
	keyCrSel              = C.SDLK_CRSEL
	keyExSel              = C.SDLK_EXSEL
	keyKp00               = C.SDLK_KP_00
	keyKp000              = C.SDLK_KP_000
	keyThousandsSeparator = C.SDLK_THOUSANDSSEPARATOR
	keyDecimalSeparator   = C.SDLK_DECIMALSEPARATOR
	keyCurrencyUnit       = C.SDLK_CURRENCYUNIT
	keyCurrencySubUnit    = C.SDLK_CURRENCYSUBUNIT
	keyKpLeftParen        = C.SDLK_KP_LEFTPAREN
	keyKpRightParen       = C.SDLK_KP_RIGHTPAREN
	keyKpLeftBrace        = C.SDLK_KP_LEFTBRACE
	keyKpRightBrace       = C.SDLK_KP_RIGHTBRACE
	keyKpTab              = C.SDLK_KP_TAB
	keyKpBackspace        = C.SDLK_KP_BACKSPACE
	keyKpA                = C.SDLK_KP_A
	keyKpB                = C.SDLK_KP_B
	keyKpC                = C.SDLK_KP_C
	keyKpD                = C.SDLK_KP_D
	keyKpE                = C.SDLK_KP_E
	keyKpF                = C.SDLK_KP_F
	keyKpXOR              = C.SDLK_KP_XOR
	keyKpPower            = C.SDLK_KP_POWER
	keyKpPercent          = C.SDLK_KP_PERCENT
	keyKpLess             = C.SDLK_KP_LESS
	keyKpGreater          = C.SDLK_KP_GREATER
	keyKpAmpersand        = C.SDLK_KP_AMPERSAND
	keyKpDblAmpersand     = C.SDLK_KP_DBLAMPERSAND
	keyKpVerticalBar      = C.SDLK_KP_VERTICALBAR
	keyKpDblVerticalBar   = C.SDLK_KP_DBLVERTICALBAR
	keyKpColon            = C.SDLK_KP_COLON
	keyKpHash             = C.SDLK_KP_HASH
	keyKpSpace            = C.SDLK_KP_SPACE
	keyKpAt               = C.SDLK_KP_AT
	keyKpExclam           = C.SDLK_KP_EXCLAM
	keyKpMemStore         = C.SDLK_KP_MEMSTORE
	keyKpMemRecall        = C.SDLK_KP_MEMRECALL
	keyKpMemClear         = C.SDLK_KP_MEMCLEAR
	keyKpMemAdd           = C.SDLK_KP_MEMADD
	keyKpMemSubtract      = C.SDLK_KP_MEMSUBTRACT
	keyKpMemMultiply      = C.SDLK_KP_MEMMULTIPLY
	keyKpMemDivide        = C.SDLK_KP_MEMDIVIDE
	keyKpPlusMinus        = C.SDLK_KP_PLUSMINUS
	keyKpClear            = C.SDLK_KP_CLEAR
	keyKpClearEntry       = C.SDLK_KP_CLEARENTRY
	keyKpBinary           = C.SDLK_KP_BINARY
	keyKpOctal            = C.SDLK_KP_OCTAL
	keyKpDecimal          = C.SDLK_KP_DECIMAL
	keyKpHexadecimal      = C.SDLK_KP_HEXADECIMAL
	keyLCtrl              = C.SDLK_LCTRL
	keyLShift             = C.SDLK_LSHIFT
	keyLAlt               = C.SDLK_LALT
	keyLGui               = C.SDLK_LGUI
	keyRCtrl              = C.SDLK_RCTRL
	keyRShift             = C.SDLK_RSHIFT
	keyRAlt               = C.SDLK_RALT
	keyRGui               = C.SDLK_RGUI
	keyMode               = C.SDLK_MODE
	keyAudioNext          = C.SDLK_AUDIONEXT
	keyAudioPrev          = C.SDLK_AUDIOPREV
	keyAudioStop          = C.SDLK_AUDIOSTOP
	keyAudioPlay          = C.SDLK_AUDIOPLAY
	keyAudioMute          = C.SDLK_AUDIOMUTE
	keyMediaSelect        = C.SDLK_MEDIASELECT
	keyWWW                = C.SDLK_WWW
	keyMail               = C.SDLK_MAIL
	keyCalculator         = C.SDLK_CALCULATOR
	keyComputer           = C.SDLK_COMPUTER
	keyAcSearch           = C.SDLK_AC_SEARCH
	keyAcHome             = C.SDLK_AC_HOME
	keyAcBack             = C.SDLK_AC_BACK
	keyAcForward          = C.SDLK_AC_FORWARD
	keyAcStop             = C.SDLK_AC_STOP
	keyAcRefresh          = C.SDLK_AC_REFRESH
	keyAcBookmarks        = C.SDLK_AC_BOOKMARKS
	keyBrightnessDown     = C.SDLK_BRIGHTNESSDOWN
	keyBrightnessUp       = C.SDLK_BRIGHTNESSUP
	keyDisplaySwitch      = C.SDLK_DISPLAYSWITCH
	keyKbdIllumToggle     = C.SDLK_KBDILLUMTOGGLE
	keyKbdIllumDown       = C.SDLK_KBDILLUMDOWN
	keyKbdIllumUp         = C.SDLK_KBDILLUMUP
	keyEject              = C.SDLK_EJECT
	keySleep              = C.SDLK_SLEEP
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
