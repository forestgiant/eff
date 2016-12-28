package main

import (
	"fmt"
	"log"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/scroll"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW   = 800
	windowH   = 540
	rowHeight = 30
	rowCount  = 50
)

func main() {
	canvas := sdl.NewCanvas("Scroller", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rowsCreated := 0
		font, err := canvas.OpenFont("../assets/fonts/roboto/Roboto-Medium.ttf", 15)
		if err != nil {
			log.Fatal(err)
		}
		createRow := func(text string) *eff.Shape {
			s := &eff.Shape{}
			s.SetBackgroundColor(eff.RandomColor())
			s.SetRect(eff.Rect{
				X: 0,
				Y: rowHeight * rowsCreated,
				W: windowW,
				H: rowHeight,
			})

			textW, textH, err := canvas.Graphics().GetTextSize(font, text)
			if err != nil {
				log.Fatal(err)
			}

			textPoint := eff.Point{
				X: (windowW - textW) / 2,
				Y: (rowHeight - textH) / 2,
			}

			s.DrawText(font, text, eff.White(), textPoint)

			rowsCreated++
			return s
		}
		content := &eff.Shape{}
		content.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: windowW,
			H: rowHeight * rowCount,
		})
		content.SetBackgroundColor(eff.RandomColor())
		for i := 0; i < rowCount; i++ {
			content.AddChild(createRow(fmt.Sprintf("Row: %d", i)))
		}

		scroller := scroll.NewScroller(content, canvas.Rect(), canvas)
		scroller.SetBackgroundColor(eff.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF})
		canvas.AddChild(scroller)
	})
}
