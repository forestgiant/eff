package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
	initialized bool
	val         int
	rect        eff.Rect
	vec         eff.Point
	color       eff.Color
	textColor   eff.Color
	sizeDir     int
	font        eff.Font
}

func (m *movingText) Init(c eff.Canvas) {
	rand.Seed(time.Now().UnixNano())

	font, err := c.OpenFont("../assets/fonts/Jellee-Roman.ttf", 24)

	if err != nil {
		log.Fatal(err)
	}

	m.font = font

	m.rect = eff.Rect{
		X: rand.Intn(c.Width() - size),
		Y: rand.Intn(c.Height() - size),
		W: size,
		H: size,
	}
	m.vec = eff.Point{
		X: 1,
		Y: 1,
	}

	m.color = eff.RandomColor()
	m.textColor = eff.RandomColor()

	m.sizeDir = 1
	m.initialized = true
}

func (m *movingText) Initialized() bool {
	return m.initialized
}

func (m *movingText) Draw(c eff.Canvas) {
	c.FillRect(m.rect, m.color)
	valText := strconv.Itoa(m.val)
	valText, err := util.EllipseText(m.font, valText, m.rect.W, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	centeredPoint, err := util.CenterTextInRect(m.font, valText, m.rect, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.DrawText(m.font, valText, m.textColor, centeredPoint)
}

func (m *movingText) Update(c eff.Canvas) {
	m.val += (100 + rand.Intn(1000))
	if m.val > 100000 {
		m.val = 0
	}

	m.rect.X += m.vec.X
	if m.rect.X < sizeChange || m.rect.X+m.rect.W > (c.Width()-(sizeChange*2)) {
		m.vec.X *= -1
	}
	m.rect.Y += m.vec.Y
	if m.rect.Y < sizeChange || m.rect.Y+m.rect.H > (c.Height()-(sizeChange*2)) {
		m.vec.Y *= -1
	}

	m.rect.W += m.sizeDir
	m.rect.H += m.sizeDir
	if m.rect.W > size+sizeChange || m.rect.W < size-sizeChange {
		m.sizeDir *= -1
	}
}

func main() {
	canvas := sdl.NewCanvas("Moving Text", 800, 540, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.AddDrawable(&movingText{})
	})
}
