package sdl

// #include "wrapper.h"
import "C"
import (
	"errors"
	"unsafe"
)

const maxVolume int = C.MIX_MAX_VOLUME

// Music (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC86)
type music C.Mix_Music

// InitMix initalizes SDL_mixer
func initMix() error {
	soundFreq := 44100
	chunkSize := 4096
	numChannels := 2
	format := C.MIX_DEFAULT_FORMAT

	if C.Mix_OpenAudio(C.int(soundFreq), C.Uint16(format), C.int(numChannels), C.int(chunkSize)) == -1 {
		return getMixError()
	}

	return nil
}

// GetMixError (https://wiki.libsdl.org/SDL_GetError)
func getMixError() error {
	if err := C.Mix_GetError(); err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

// LoadMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC55)
func loadMusic(path string) (*music, error) {
	_musicPath := C.CString(path)
	var _music = C.Mix_LoadMUS(_musicPath)

	if _music == nil {
		return nil, getMixError()
	}
	return (*music)(unsafe.Pointer(_music)), nil
}

// PlayMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC57)
func playMusic(music *music, loopCount int) error {
	if C.Mix_PlayMusic(music, C.int(loopCount)) < 0 {
		return getMixError()
	}

	return nil
}

// PauseMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC62)
func pauseMusic() {
	C.Mix_PauseMusic()
}

// ResumeMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC63)
func resumeMusic() {
	C.Mix_ResumeMusic()
}

// HaltMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer.html#SEC67)
func haltMusic() {
	C.Mix_HaltMusic()
}

func fadeInMusic(music *music, loopCount int, fadeTimeMS int) error {
	if C.Mix_FadeInMusic(music, C.int(loopCount), C.int(fadeTimeMS)) < 0 {
		return getMixError()
	}

	return nil
}

func fadeOutMusic(fadeTimeMS int) error {
	if C.Mix_FadeOutMusic(C.int(fadeTimeMS)) < 1 {
		return getMixError()
	}

	return nil
}

func volumeMusic(volume int) int {
	return int(C.Mix_VolumeMusic(C.int(volume)))
}

func musicPlaying() bool {
	musicPlaying := C.Mix_PlayingMusic()

	if musicPlaying > 0 {
		return true
	}

	return false
}
