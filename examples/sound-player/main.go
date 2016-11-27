package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type player struct {
	initialized bool
	audioPlayer sdl.AudioPlayer
	musicPath   string
	font        eff.Font
}

func newPlayer(musicPath string) player {
	p := player{}
	p.musicPath = musicPath
	p.audioPlayer = sdl.NewAudioPlayer(musicPath, -1)
	return p
}

func (p *player) Init(c eff.Canvas) {

	font, err := c.OpenFont("../assets/fonts/vcr_osd_mono.ttf", 24)
	if err != nil {
		log.Fatal(err)
	}
	p.font = font
	p.audioPlayer.Play()
	p.initialized = true
}

func (p *player) Initialized() bool {
	return p.initialized
}

func (p *player) Draw(c eff.Canvas) {
	margin := eff.Point{X: 10, Y: 10}
	yPos := 0
	c.DrawText(p.font, "Now Playing: "+path.Base(p.musicPath), eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y})
	yPos += 24 + margin.Y
	c.DrawText(p.font, "Press p to pause", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText(p.font, "Press z to fade in", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText(p.font, "Press x to fade out", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText(p.font, "Press r to resume", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	yPos += 24 + margin.Y
	c.DrawText(p.font, "Press q to quit", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
}

func (p *player) Update(c eff.Canvas) {

}

func main() {

	canvas := sdl.NewCanvas("Sound Player", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)

	canvas.Run(func() {
		usage := "Usage sound-player <PATH_TO_WAV>"
		if len(os.Args) < 2 {
			log.Fatal(usage)
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
			case "Z":
				player.audioPlayer.FadeVolume(100, 1)
			case "X":
				player.audioPlayer.FadeVolume(500, 0)
			}
		})
	})
}
