package main

import (
	"os"
	"strconv"

	"github.com/forestgiant/eff"
)

func main() {
	//Create drawables
	var drawables []eff.Drawable
	drawables = append(drawables, &dots{})
	drawables = append(drawables, &rects{})
	drawables = append(drawables, &collidingBlocks{})
	drawables = append(drawables, &circleDots{})
	drawables = append(drawables, &squareSpiral{})

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

	canvas.AddKeyUpHandler(func(key string) {
		// fmt.Println("Up", key)
		n, err := strconv.Atoi(key)
		if err == nil {
			setDrawable(n - 1)
		}
	})

	//Start the run loop
	os.Exit(canvas.Run())
}
