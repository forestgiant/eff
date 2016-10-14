package sdl

import (
	"errors"
	"fmt"
	"time"
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
	volume    int
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

// FadeMute fades the music volume to 0 over the argument time
func (ap *AudioPlayer) FadeMute(fadeTimeMS int) {
	frameTime := float64(1000) / float64(60)
	stepCount := float64(fadeTimeMS) / frameTime
	currentVolume := float64(volumeMusic(-1))
	fadeAmount := float64(currentVolume) / stepCount

	go func() {
		for currentVolume > 0 {
			currentVolume -= fadeAmount
			if currentVolume < 0 {
				currentVolume = 0
			}

			volumeMusic(int(currentVolume))
			time.Sleep(time.Millisecond * time.Duration(frameTime))
		}
	}()

}

// FadeUnmute fades the music volume to 128 over the argument time
func (ap *AudioPlayer) FadeUnmute(fadeTimeMS int) {
	targetVolume := float64(128)
	frameTime := float64(1000) / float64(60)
	stepCount := float64(fadeTimeMS) / frameTime
	currentVolume := float64(volumeMusic(-1))
	fadeAmount := (targetVolume - currentVolume) / stepCount

	go func() {
		for currentVolume < targetVolume {
			currentVolume += fadeAmount

			volumeMusic(int(currentVolume))
			time.Sleep(time.Millisecond * time.Duration(frameTime))
		}
	}()
}

// FadeIn fades the music in from zero starts from the beginning
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

// FadeOut fades the music down to zero and then stops playing
func (ap *AudioPlayer) FadeOut(fadeTimeMS int) {
	if ap.music == nil {
		fmt.Println("cannot fade out, no music loaded")
		return
	}
	fadeOutMusic(fadeTimeMS)
}

// Playing returns true if music is playing false otherwise
func (ap *AudioPlayer) Playing() bool {
	return musicPlaying()
}
