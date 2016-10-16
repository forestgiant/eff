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
	quitFade  chan struct{}
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

// FadeVolume fades the music volume to 0 over the argument time
func (ap *AudioPlayer) FadeVolume(fadeTimeMS int, percentage float64) {
	if ap.quitFade != nil {
		close(ap.quitFade)
	}
	ap.quitFade = make(chan struct{})

	if percentage < 0 || percentage > 1 {
		fmt.Println("invalid volume", percentage)
		return
	}

	frameTime := float64(1000) / float64(60)
	stepCount := float64(fadeTimeMS) / frameTime
	currentVolume := float64(volumeMusic(-1))
	newVolume := float64(maxVolume) * percentage
	fadeAmount := (newVolume - currentVolume) / stepCount
	steps := 0

	go func(quit chan struct{}) {
		for steps < int(stepCount) {
			select {
			case <-quit:
				return
			default:
				currentVolume += fadeAmount

				volumeMusic(int(currentVolume))
				time.Sleep(time.Millisecond * time.Duration(frameTime))
				steps++
			}
		}
	}(ap.quitFade)
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

// SetVolume adjusts the volume of the currently playing music, percentage is a normalized value between 0-1
func (ap *AudioPlayer) SetVolume(percentage float64) {
	if percentage < 0 || percentage > 1 {
		fmt.Println("invalid volume", percentage)
		return
	}

	volumeMusic(int(float64(maxVolume) * percentage))
}
