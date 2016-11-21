package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
	"github.com/forestgiant/eff/sdl"
)

type imageTiler struct {
	imgPath      string
	imgs         []*eff.Image
	onScreenImgs []*eff.Image
	initialized  bool
	baseW        int
	baseH        int
	rows         int
	cols         int
	index        int
	tweener      tween.Tweener
}

func (i *imageTiler) Init(c eff.Canvas) {
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
	i.imgPath = os.Args[1]

	img := &eff.Image{
		Path: i.imgPath,
		Rect: eff.Rect{
			X: 0,
			Y: 0,
			W: -1,
			H: -1,
		},
	}

	c.AddImage(img)
	c.RemoveImage(img)
	// Since the W and H are set to -1 when adding the image they will be replaced with the image size
	w := img.Rect.W
	h := img.Rect.H

	i.rows = 3
	i.baseH = int(float64(c.Height()) / float64(i.rows))
	i.baseW = int(float64(i.baseH) * (float64(w) / float64(h)))
	i.cols = int(float64(c.Width())/float64(i.baseW)) + 1

	totalImgs := i.rows * i.cols
	for j := 0; j < totalImgs; j++ {
		x := j % i.cols
		x *= i.baseW
		y := int(float64(j) / float64(i.cols))
		y *= i.baseH

		img := &eff.Image{
			Path: i.imgPath,
			Rect: eff.Rect{
				X: x,
				Y: y,
				W: 0,
				H: 0,
			},
		}

		i.imgs = append(i.imgs, img)
		c.AddImage(img)
	}

	i.tweener = tween.NewTweener(time.Second*2, func(progress float64) {
		i.index = int(progress * float64(len(i.imgs)))
	}, true, false, nil, nil)
	i.initialized = true
}

func (i *imageTiler) Initialized() bool {
	return i.initialized
}

func (i *imageTiler) Draw(c eff.Canvas) {
	c.FillRect(eff.Rect{X: 0, Y: 0, W: c.Width(), H: c.Height()}, eff.RandomColor())

	for _, img := range i.onScreenImgs {
		img.Rect.W = 0
		img.Rect.H = 0
	}

	i.onScreenImgs = i.imgs[:i.index]

	for _, img := range i.onScreenImgs {
		img.Rect.W = i.baseW
		img.Rect.H = i.baseH
	}
}

func (i *imageTiler) Update(c eff.Canvas) {
	i.tweener.Tween()
}

func main() {
	canvas := sdl.NewCanvas("image tile", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.AddDrawable(&imageTiler{})
	})
}
