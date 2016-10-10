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
		Path: "../assets/fonts/vcr_osd_mono.ttf",
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
	margin := eff.Point{X: 10, Y: 10}
	yPos := 0
	c.DrawText("Now Playing: "+path.Base(p.musicPath), eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y})
	yPos += 24 + margin.Y
	c.DrawText("Press p to pause", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText("Press r to resume", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText("Press q to quit", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
}

func (p *player) Update(c eff.Canvas) {

}

func main() {

	canvas := sdl.NewCanvas("Sound Player", 800, 540, 60, true)

	canvas.Run(func() {
		usage := "Usage sound-player <PATH_TO_WAV>"
		if len(os.Args) < 2 {
			fmt.Println(usage)
			os.Exit(1)
		}

		ext := path.Ext(os.Args[1])
		ext = strings.ToLower(ext)

		if ext != ".wav" {
			fmt.Println(usage)
			os.Exit(1)
		}
		player := newPlayer(os.Args[1])

		canvas.AddDrawable(&player)

		canvas.AddKeyUpHandler(func(key string) {
			switch key {
			case "P":
				player.audioPlayer.Pause()
			case "R":
				player.audioPlayer.Resume()
			}
		})
	})
}
