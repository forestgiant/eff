package main

import (
	"os"

	"github.com/forestgiant/eff"
)

func main() {
	//Create drawables
	drawables := make([]eff.Drawable, 3)
	drawables[0] = &dots{}
	drawables[1] = &rects{}
	drawables[2] = &collidingBlocks{}

	drawableIndex := 0

	//Create Eff Canvas
	canvas := eff.SDLCanvas{}
	canvas.SetWidth(1280)
	canvas.SetHeight(720)

	setDrawable := func(index int) {
		if index < 0 || index >= len(drawables) {
			return
		}

		if index == drawableIndex {
			return
		}

		if len(drawables) > 0 && drawableIndex >= 0 {
			canvas.RemoveDrawable(drawables[drawableIndex])
		}

		canvas.AddDrawable(drawables[index])

		drawableIndex = index
	}

	//Add drawables to canvas
	canvas.AddDrawable(drawables[0])

	canvas.AddKeyUpHandler(func(key string, canvas eff.Canvas) {
		// fmt.Println("Up", key)
		if key == "1" {
			setDrawable(0)
		} else if key == "2" {
			setDrawable(1)
		} else if key == "3" {
			setDrawable(2)
		}
	})

	//Start the run loop
	os.Exit(canvas.Run())
}
