package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
	"github.com/forestgiant/eff/sdl"
)

type imageTiler struct {
	eff.Shape
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
		log.Fatal(usage)
	}

	img, err := c.OpenImage(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Since the W and H are set to -1 when adding the image they will be replaced with the image size
	w := img.Width()
	h := img.Height()

	rows := 3
	baseH := int(float64(c.Rect().H) / float64(rows))
	baseW := int(float64(baseH) * (float64(w) / float64(h)))
	cols := int(float64(c.Rect().W)/float64(baseW)) + 1

	totalImgs := rows * cols

	count := 0
	tweener := tween.NewTweener(time.Second*2, func(progress float64) {
		count = int(progress * float64(totalImgs))
	}, true, false, nil, nil)

	i.SetUpdateHandler(func() {
		tweener.Tween()
		i.Clear()
		i.SetBackgroundColor(eff.RandomColor())
		for j := 0; j < count; j++ {
			x := j % cols
			x *= baseW
			y := int(float64(j) / float64(cols))
			y *= baseH
			r := eff.Rect{
				X: x,
				Y: y,
				W: baseW,
				H: baseH,
			}
			i.DrawImage(img, r)
		}
	})
}

func main() {
	canvas := sdl.NewCanvas("image tile", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		i := &imageTiler{}
		i.SetRect(canvas.Rect())
		canvas.AddChild(i)
		i.Init(canvas)
	})
}
