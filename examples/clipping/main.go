package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW    = 800
	windowH    = 600
	parentSize = 300
)

type colorMarquee struct {
	eff.Shape
	colors []*eff.Shape
}

func (c *colorMarquee) init(f eff.Font) {
	c.SetRect(eff.Rect{
		X: (windowW - parentSize) / 2,
		Y: (windowH - parentSize) / 2,
		W: parentSize,
		H: parentSize,
	})
	c.SetBackgroundColor(eff.Black())
	colorCount := 10
	newColor := func() *eff.Shape {
		color := &eff.Shape{}
		color.SetRect(eff.Rect{
			X: 0,
			Y: 0,
			W: c.Rect().W,
			H: int(float64(c.Rect().H) / float64(colorCount)),
		})
		color.SetBackgroundColor(eff.RandomColor())
		return color
	}

	for i := 0; i < colorCount+1; i++ {
		color := newColor()
		color.SetRect(eff.Rect{
			X: color.Rect().X,
			Y: i * color.Rect().H,
			W: color.Rect().W,
			H: color.Rect().H,
		})

		color.SetUpdateHandler(func() {
			y := color.Rect().Y - 1
			if y < color.Rect().H*-1 {
				y = (colorCount) * color.Rect().H
			}

			color.SetRect(eff.Rect{
				X: color.Rect().X,
				Y: y,
				W: color.Rect().W,
				H: color.Rect().H,
			})
		})

		textW, textH, err := c.Graphics().GetTextSize(f, fmt.Sprintf("%d", i))
		if err != nil {
			log.Fatal(err)
		}

		textPoint := eff.Point{
			X: (color.Rect().W - textW) / 2,
			Y: (color.Rect().H - textH) / 2,
		}
		color.DrawText(f, fmt.Sprintf("%d", i), eff.White(), textPoint)
		c.AddChild(color)
	}
}

func main() {
	canvas := sdl.NewCanvas("Clipping", windowW, windowH, eff.White(), 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())

		c := &colorMarquee{}
		f, err := canvas.OpenFont("../assets/fonts/roboto/Roboto-Medium.ttf", 15)
		if err != nil {
			log.Fatal(err)
		}
		canvas.AddChild(c)
		c.init(f)
	})
}
