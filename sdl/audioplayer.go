package sdl

import (
	"errors"
	"fmt"
)

func NewAudioPlayer(musicPath string, loopCount int) AudioPlayer {
	ap := AudioPlayer{
		musicPath: musicPath,
		loopCount: loopCount,
	}

	ap.load()
	return ap
}

type AudioPlayer struct {
	musicPath string
	loopCount int
	music     *Music
}

func (ap *AudioPlayer) load() {
	music, err := LoadMusic(ap.musicPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	ap.music = music
}

func (ap *AudioPlayer) Play() error {
	if ap.music == nil {
		fmt.Println("cannot play, no music loaded")
		return errors.New("no music loaded")

	}
	err := PlayMusic(ap.music, ap.loopCount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ap *AudioPlayer) Pause() {
	if ap.music == nil {
		fmt.Println("cannot pause, no music loaded")
		return
	}
	PauseMusic()
}

func (ap *AudioPlayer) Stop() {
	if ap.music == nil {
		fmt.Println("cannot stop, no music loaded")
		return
	}
	HaltMusic()
}

func (ap *AudioPlayer) Resume() {
	if ap.music == nil {
		fmt.Println("cannot resume, no music loaded")
		return
	}
	ResumeMusic()
}
