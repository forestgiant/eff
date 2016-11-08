package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
	"github.com/forestgiant/eff/util"
)

const (
	boxW = 100
	boxH = 100
)

type movingText struct {
	initialized bool
	val         int
	rect        eff.Rect
	vec         eff.Point
	color       eff.Color
	textColor   eff.Color
}

func (m *movingText) Init(c eff.Canvas) {
	rand.Seed(time.Now().UnixNano())
	font := eff.Font{
		Path: "../assets/fonts/Jellee-Roman.ttf",
	}

	err := c.SetFont(font, 24)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m.rect = eff.Rect{
		X: rand.Intn(c.Width() - boxW),
		Y: rand.Intn(c.Height() - boxH),
		W: boxW,
		H: boxH,
	}
	m.vec = eff.Point{
		X: 1,
		Y: 1,
	}

	m.color = eff.RandomColor()
	m.textColor = eff.RandomColor()

	m.initialized = true
}

func (m *movingText) Initialized() bool {
	return m.initialized
}

func (m *movingText) Draw(c eff.Canvas) {
	c.FillRect(m.rect, m.color)
	valText := strconv.Itoa(m.val)
	valText, err := util.EllipseText(valText, m.rect.W, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	centeredPoint, err := util.CenterTextInRect(valText, m.rect, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.DrawText(valText, m.textColor, centeredPoint)
}

func (m *movingText) Update(c eff.Canvas) {
	m.val += 10000
	if m.val > 10000000 {
		m.val = 0
	}

	m.rect.X += m.vec.X
	if m.rect.X < 0 || m.rect.X+m.rect.W > c.Width() {
		m.vec.X *= -1
	}
	m.rect.Y += m.vec.Y
	if m.rect.Y < 0 || m.rect.Y+m.rect.H > c.Height() {
		m.vec.Y *= -1
	}

}

func main() {
	canvas := sdl.NewCanvas("Moving Text", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.AddDrawable(&movingText{})
	})
}
