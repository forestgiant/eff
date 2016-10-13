package sdl

import (
	"errors"
	"fmt"
)

// NewAudioPlayer Creates a new audio player instance, the music path is the path to a wav file
func NewAudioPlayer(musicPath string, loopCount int) AudioPlayer {
	ap := AudioPlayer{
		musicPath: musicPath,
		loopCount: loopCount,
	}

	ap.load()
	return ap
}

// AudioPlayer container for a music file
type AudioPlayer struct {
	musicPath string
	loopCount int
	music     *music
}

func (ap *AudioPlayer) load() {
	music, err := loadMusic(ap.musicPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	ap.music = music
}

// Play begins playing the loaded music
func (ap *AudioPlayer) Play() error {
	if ap.music == nil {
		fmt.Println("cannot play, no music loaded")
		return errors.New("no music loaded")

	}
	err := playMusic(ap.music, ap.loopCount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Pause pauses the currently playing music
func (ap *AudioPlayer) Pause() {
	if ap.music == nil {
		fmt.Println("cannot pause, no music loaded")
		return
	}
	pauseMusic()
}

// Stop stops playing the currently playing music and resets the play head position to the beginning
func (ap *AudioPlayer) Stop() {
	if ap.music == nil {
		fmt.Println("cannot stop, no music loaded")
		return
	}
	haltMusic()
}

// Resume resumes the currently paused or stopped music
func (ap *AudioPlayer) Resume() {
	if ap.music == nil {
		fmt.Println("cannot resume, no music loaded")
		return
	}
	resumeMusic()
}

func (ap *AudioPlayer) FadeIn(fadeTimeMS int) error {
	if ap.music == nil {
		fmt.Println("cannot fade in, no music loaded")
		return errors.New("no music loaded")
	}

	err := fadeInMusic(ap.music, ap.loopCount, fadeTimeMS)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ap *AudioPlayer) FadeOut(fadeTimeMS int) {
	if ap.music == nil {
		fmt.Println("cannot fade out, no music loaded")
		return
	}

	fadeOutMusic(fadeTimeMS)
}
