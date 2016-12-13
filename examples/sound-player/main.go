package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type player struct {
	eff.Shape
	audioPlayer sdl.AudioPlayer
}

func (p *player) Init(c eff.Canvas, musicPath string) {

	font, err := c.OpenFont("../assets/fonts/vcr_osd_mono.ttf", 24)
	if err != nil {
		log.Fatal(err)
	}

	p.audioPlayer = sdl.NewAudioPlayer(musicPath, -1)
	p.audioPlayer.Play()

	p.SetUpdateHandler(func() {
		p.Clear()
		margin := eff.Point{X: 10, Y: 10}
		yPos := 0
		p.DrawText(font, "Now Playing: "+path.Base(musicPath), eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y})
		yPos += 24 + margin.Y
		p.DrawText(font, "Press p to pause", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
		yPos += 24 + margin.Y
		p.DrawText(font, "Press z to fade in", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
		yPos += 24 + margin.Y
		p.DrawText(font, "Press x to fade out", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
		yPos += 24 + margin.Y
		p.DrawText(font, "Press r to resume", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
		yPos += 24 + margin.Y
		p.DrawText(font, "Press q to quit", eff.RandomColor(), eff.Point{X: margin.X, Y: margin.Y + yPos})
	})
}

func main() {

	canvas := sdl.NewCanvas("Sound Player", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)

	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
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
		player := &player{}
		player.SetRect(canvas.Rect())
		canvas.AddChild(player)
		player.Init(canvas, os.Args[1])

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
