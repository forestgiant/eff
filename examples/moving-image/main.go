package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

type imageMover struct {
	img         *eff.Image
	xDir        int
	yDir        int
	initialized bool
	baseW       int
	baseH       int
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
		fmt.Println(usage)
		os.Exit(1)
	}

	i.img = &eff.Image{
		Path: os.Args[1],
		Rect: eff.Rect{
			X: 0,
			Y: 0,
			W: -1,
			H: -1,
		},
	}

	c.AddImage(i.img)
	// Since the W and H are set to -1 when adding the image they will be replaced with the image size
	w := i.img.Rect.W
	h := i.img.Rect.H

	i.baseW = int(float64(c.Width()) / float64(4))
	i.baseH = int(float64(i.baseW) * (float64(h) / float64(w)))

	i.img.Rect.W = i.baseW
	i.img.Rect.H = i.baseH

	i.xDir = 1
	i.yDir = 1
	i.initialized = true
}

func (i *imageMover) Initialized() bool {
	return i.initialized
}

func (i *imageMover) Draw(c eff.Canvas) {
	c.FillRect(eff.Rect{X: 0, Y: 0, W: c.Width(), H: c.Height()}, eff.RandomColor())
}

func (i *imageMover) Update(c eff.Canvas) {
	i.img.Rect.X += i.xDir
	i.img.Rect.Y += i.yDir

	if i.img.Rect.X < 0 || i.img.Rect.X+i.img.Rect.W > c.Width() {
		i.xDir *= -1
	}

	if i.img.Rect.Y < 0 || i.img.Rect.Y+i.img.Rect.H > c.Height() {
		i.yDir *= -1
	}
}

func main() {
	canvas := sdl.NewCanvas("moving image", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.AddDrawable(&imageMover{})
	})
}
