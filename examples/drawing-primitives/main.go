package main

import (
	"strconv"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/examples/drawing-primitives/shapes"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW = 800
	windowH = 480
)

func main() {
	//Create Eff Canvas
	canvas := sdl.NewCanvas("Drawing Primitives", windowW, windowH, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)

	//Start the run loop
	canvas.Run(func() {

		var d []eff.Drawable
		dots := &shapes.Dots{}
		dots.Init(windowW, windowH)
		dots.SetRect(canvas.Rect())
		d = append(d, dots)
		rects := &shapes.Rects{}
		rects.Init(windowW, windowH)
		rects.SetRect(canvas.Rect())
		d = append(d, rects)
		collidingBlocks := &shapes.CollidingBlocks{}
		collidingBlocks.Init(windowW, windowH)
		collidingBlocks.SetRect(canvas.Rect())
		d = append(d, collidingBlocks)
		circleDots := &shapes.CircleDots{}
		circleDots.Init(windowW, windowH)
		circleDots.SetRect(canvas.Rect())
		d = append(d, circleDots)
		sqaureSpiral := &shapes.SquareSpiral{}
		sqaureSpiral.Init(windowW, windowH)
		sqaureSpiral.SetRect(canvas.Rect())
		d = append(d, sqaureSpiral)

		shapeIndex := 0

		setDrawable := func(index int) {
			if index < 0 || index >= len(d) {
				return
			}

			if index == shapeIndex {
				return
			}

			if len(d) > 0 && shapeIndex >= 0 {
				canvas.RemoveChild(d[shapeIndex])
			}

			canvas.AddChild(d[index])

			shapeIndex = index
		}
		canvas.AddChild(d[0])

		canvas.AddKeyUpHandler(func(key string) {

			// fmt.Println("Up", key)
			n, err := strconv.Atoi(key)
			if err == nil {
				setDrawable(n - 1)
			}
		})
	})
}
