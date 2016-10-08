package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Music (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC86)
type Music C.Mix_Music

var mixInitialized = false

// InitMix initalizes SDL_mixer
func InitMix() error {
	soundFreq := 44100
	chunkSize := 4096
	numChannels := 2
	format := C.MIX_DEFAULT_FORMAT

	if C.Mix_OpenAudio(C.int(soundFreq), C.Uint16(format), C.int(numChannels), C.int(chunkSize)) == -1 {
		return GetMixError()
	}

	mixInitialized = true
	return nil
}

// GetMixError (https://wiki.libsdl.org/SDL_GetError)
func GetMixError() error {
	if err := C.Mix_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// LoadMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC55)
func LoadMusic(path string) (*Music, error) {
	_musicPath := C.CString(path)
	var _music = C.Mix_LoadMUS(_musicPath)

	if _music == nil {
		return nil, GetMixError()
	}
	return (*Music)(unsafe.Pointer(_music)), nil
}

// PlayMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC57)
func PlayMusic(music *Music, loopCount int) error {
	if C.Mix_PlayMusic(music, C.int(loopCount)) < 0 {
		return GetMixError()
	}

	return nil
}
