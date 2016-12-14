package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
	"github.com/forestgiant/eff/util"
)

const (
	size       = 100
	sizeChange = 50
)

type movingText struct {
	eff.Shape
}

func (m *movingText) Init(c eff.Canvas) {
	rand.Seed(time.Now().UnixNano())

	font, err := c.OpenFont("../assets/fonts/Jellee-Roman.ttf", 24)

	if err != nil {
		log.Fatal(err)
	}

	rect := eff.Rect{
		X: rand.Intn(c.Rect().W - size),
		Y: rand.Intn(c.Rect().H - size),
		W: size,
		H: size,
	}
	vec := eff.Point{
		X: 1,
		Y: 1,
	}

	color := eff.RandomColor()
	textColor := eff.RandomColor()

	sizeDir := 1
	var val int

	m.SetUpdateHandler(func() {
		val += (100 + rand.Intn(1000))
		if val > 100000 {
			val = 0
		}

		x := rect.X + vec.X
		y := rect.Y + vec.Y

		if x < sizeChange || x+rect.W > (c.Rect().W-(sizeChange*2)) {
			vec.X *= -1
		}
		if y < sizeChange || y+rect.H > (c.Rect().H-(sizeChange*2)) {
			vec.Y *= -1
		}

		w := rect.W + sizeDir
		h := rect.H + sizeDir
		if w > size+sizeChange || w < size-sizeChange {
			sizeDir *= -1
		}

		rect.X = x
		rect.Y = y
		rect.W = w
		rect.H = h

		valText := strconv.Itoa(val)
		valText, err := util.EllipseText(font, valText, rect.W, m.Graphics())
		if err != nil {
			log.Fatal(err)
		}

		centeredPoint, err := util.CenterTextInRect(font, valText, rect, m.Graphics())
		if err != nil {
			log.Fatal(err)
		}

		m.Clear()
		m.FillRect(rect, color)
		m.DrawText(font, valText, textColor, centeredPoint)
	})

}

func main() {
	canvas := sdl.NewCanvas("Moving Text", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		movingText := &movingText{}
		movingText.SetRect(canvas.Rect())
		canvas.AddChild(movingText)
		movingText.Init(canvas)
	})
}
