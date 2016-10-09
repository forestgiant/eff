package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/forestgiant/eff/eff"
	"github.com/forestgiant/eff/sdl"
)

type player struct {
	initialized bool
	audioPlayer sdl.AudioPlayer
	musicPath   string
}

func newPlayer(musicPath string) player {
	p := player{}
	p.musicPath = musicPath
	p.audioPlayer = sdl.NewAudioPlayer(musicPath, -1)
	return p
}

func (p *player) Init(c eff.Canvas) {
	font := eff.Font{
		Path: "../assets/fonts/Jellee-Roman.ttf",
	}
	err := c.SetFont(font, 24)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.audioPlayer.Play()
	p.initialized = true
}

func (p *player) Initialized() bool {
	return p.initialized
}

func (p *player) Draw(c eff.Canvas) {
	c.DrawText("Now Playing: "+path.Base(p.musicPath), eff.RandomColor(), eff.Point{X: 0, Y: 0})
	c.DrawText("Press p to pause", eff.RandomColor(), eff.Point{X: 0, Y: 40})
	c.DrawText("Press r to resume", eff.RandomColor(), eff.Point{X: 0, Y: 80})
}

func (p *player) Update(c eff.Canvas) {

}

func main() {
	usage := "Usage sound-player <PATH_TO_WAV>"
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	ext := path.Ext(os.Args[1])
	ext = strings.ToLower(ext)

	if ext != ".wav" {
		fmt.Println(usage)
		return
	}
	player := newPlayer(os.Args[1])

	canvas := sdl.NewCanvas("Sound Player", 800, 540, 60, true)

	canvas.AddDrawable(&player)

	canvas.AddKeyUpHandler(func(key string) {
		switch key {
		case "P":
			player.audioPlayer.Pause()
		case "R":
			player.audioPlayer.Resume()
		}
	})

	canvas.Run()
}
