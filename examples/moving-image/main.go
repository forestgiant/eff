package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type imageMover struct {
	eff.Shape
}

func (i *imageMover) Init(c eff.Canvas) {
	usage := "Usage moving image(PNG or JPG) <PATH_TO_IMG>"

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	ext := path.Ext(os.Args[1])
	ext = strings.ToLower(ext)

	if ext != ".png" && ext != ".jpg" {
		log.Fatal(usage)
	}

	img, err := c.OpenImage(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	baseW := int(float64(c.Rect().W) / float64(4))
	baseH := int(float64(baseW) * (float64(img.Height()) / float64(img.Width())))
	imgRect := eff.Rect{
		X: 0,
		Y: 0,
		W: baseW,
		H: baseH,
	}

	vec := eff.Point{X: 1, Y: 1}
	i.SetUpdateHandler(func() {
		x := imgRect.X + vec.X
		y := imgRect.Y + vec.Y

		if x <= 0 || x >= (c.Rect().W-baseW) {
			vec.X *= -1
		}

		if y <= 0 || y >= (c.Rect().H-baseH) {
			vec.Y *= -1
		}

		imgRect.X = x
		imgRect.Y = y

		i.Clear()
		i.SetBackgroundColor(eff.RandomColor())
		i.DrawImage(img, imgRect)
	})
}

func main() {
	canvas := sdl.NewCanvas("moving image", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		imageMover := &imageMover{}
		imageMover.SetRect(canvas.Rect())
		canvas.AddChild(imageMover)
		imageMover.Init(canvas)
	})
}
