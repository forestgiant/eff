package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/forestgiant/eff/sdl"
)

func main() {
	usage := "Usage sound-check <PATH_TO_WAV>"
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

	ap := sdl.NewAudioPlayer(os.Args[1], -1)
	ap.Play()

	select {}
}
