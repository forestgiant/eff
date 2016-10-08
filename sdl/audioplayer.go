package sdl

import "fmt"

type AudioPlayer struct{}

func (ap *AudioPlayer) PlayMusic(musicPath string) {
	music, err := LoadMusic(musicPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = PlayMusic(music, 1)
	if err != nil {
		fmt.Println(err)
	}
}
