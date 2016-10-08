package main

import (
	"flag"
	"fmt"

	"github.com/forestgiant/eff/sdl"
)

func main() {

	sound := flag.String("path", "", "path to wav file")
	flag.Parse()
	ap := sdl.AudioPlayer{}

	if len(*sound) > 0 {
		fmt.Println("Playing", *sound)
		ap.PlayMusic(*sound)

		for {
		}
	}

}
