package main

import (
	"strconv"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/examples/drawing-primitives/drawables"
	"github.com/forestgiant/eff/sdl"
)

func main() {
	//Create Eff Canvas
	canvas := sdl.NewCanvas("Drawing Primitives", 800, 480, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)

	//Start the run loop
	canvas.Run(func() {
		//Create drawables
		var d []eff.Drawable
		d = append(d, &drawables.Dots{})
		d = append(d, &drawables.Rects{})
		d = append(d, &drawables.CollidingBlocks{})
		d = append(d, &drawables.CircleDots{})
		d = append(d, &drawables.SquareSpiral{})

		drawableIndex := 0

		setDrawable := func(index int) {
			if index < 0 || index >= len(d) {
				return
			}

			if index == drawableIndex {
				return
			}

			if len(d) > 0 && drawableIndex >= 0 {
				canvas.RemoveDrawable(d[drawableIndex])
			}

			canvas.AddDrawable(d[index])

			drawableIndex = index
		}

		//Add drawables to canvas
		canvas.AddDrawable(d[0])

		canvas.AddKeyUpHandler(func(key string) {

			// fmt.Println("Up", key)
			n, err := strconv.Atoi(key)
			if err == nil {
				setDrawable(n - 1)
			}
		})
	})
}
