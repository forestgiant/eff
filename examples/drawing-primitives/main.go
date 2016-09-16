package main

import (
	"math/rand"
	"os"

	"github.com/forestgiant/eff"
)

type dots struct {
	points      []eff.Point
	initialized bool
}

func (d *dots) randomPoints(count int, maxX int, maxY int) *[]eff.Point {
	points := make([]eff.Point, count)
	for i := 0; i < count; i++ {
		points[i] = eff.Point{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		}
	}
	return &points
}

func (d *dots) Init(canvas eff.Canvas) {
	d.points = *d.randomPoints(10000, canvas.Width(), canvas.Height())

	d.initialized = true
}

func (d *dots) Draw(canvas eff.Canvas) {
	//Draw Points in a random color
	canvas.DrawPoints(&d.points, eff.Color{}.RandomColor())
}

func (d *dots) Update(canvas eff.Canvas) {
	updateRandomPoints := func() {
		for i := range d.points {
			d.points[i].X = rand.Intn(canvas.Width())
			d.points[i].Y = rand.Intn(canvas.Height())
		}
	}

	updateRandomPoints()
}

func (d *dots) Initialized() bool {
	return d.initialized
}

type rects struct {
	rects       []eff.Rect
	initialized bool
}

func (r *rects) randomRects(count int, maxX int, maxY int) *[]eff.Rect {
	rects := make([]eff.Rect, count)
	for i := 0; i < count; i++ {
		rects[i] = eff.Rect{X: rand.Intn(maxX), Y: rand.Intn(maxY), W: rand.Intn(maxX / 2), H: rand.Intn(maxY / 2)}
	}

	return &rects
}

func (r *rects) Init(canvas eff.Canvas) {
	r.rects = *r.randomRects(100, canvas.Width(), canvas.Height())
	r.initialized = true
}

func (r *rects) Draw(canvas eff.Canvas) {
	canvas.DrawRects(&r.rects, eff.Color{}.RandomColor())
}

func (r *rects) Update(canvas eff.Canvas) {
	updateRandomRects := func() {
		for i := range r.rects {
			r.rects[i].X = rand.Intn(canvas.Width())
			r.rects[i].Y = rand.Intn(canvas.Height())
			r.rects[i].W = rand.Intn(canvas.Width() / 2)
			r.rects[i].H = rand.Intn(canvas.Height() / 2)
		}
	}

	updateRandomRects()
}

func (r *rects) Initialized() bool {
	return r.initialized
}

type block struct {
	rect  eff.Rect
	dir   eff.Point
	color eff.Color
}

func (b *block) applyDir() {
	b.rect.X += b.dir.X
	b.rect.Y += b.dir.Y
}

func (b *block) wallBounce(width int, height int) {
	if b.rect.X < 0 || b.rect.X+b.rect.W > width {
		b.dir.X *= -1
	}

	if b.rect.Y < 0 || b.rect.Y+b.rect.H > height {
		b.dir.Y *= -1
	}
}

type collidingBlocks struct {
	blocks      []block
	initialized bool
}

func (c *collidingBlocks) Init(canvas eff.Canvas) {
	blockCount := 200
	blockSize := 10
	c.blocks = make([]block, blockCount)
	for i := 0; i < blockCount; i++ {
		b := block{
			rect: eff.Rect{
				X: rand.Intn(canvas.Width() - blockSize),
				Y: rand.Intn(canvas.Height() - blockSize),
				W: blockSize,
				H: blockSize,
			},
			dir: eff.Point{
				X: rand.Intn(4) + 1,
				Y: rand.Intn(4) + 1,
			},
			color: eff.Color{}.RandomColor(),
		}
		c.blocks[i] = b
	}

	c.initialized = true
}

func (c *collidingBlocks) Draw(canvas eff.Canvas) {
	for _, block := range c.blocks {
		canvas.FillRect(block.rect, block.color)
	}
}

func (c *collidingBlocks) Initialized() bool {
	return c.initialized
}

func (c *collidingBlocks) Update(canvas eff.Canvas) {
	for i, block := range c.blocks {
		for j, otherBlock := range c.blocks {
			if otherBlock.rect.Equals(block.rect) {
				// fmt.Println("Skipping block because its the same")
				continue
			}

			if block.rect.Intersects(otherBlock.rect) {

				block.dir.X *= -1
				block.dir.Y *= -1

				otherBlock.dir.X *= -1
				otherBlock.dir.Y *= -1
				for block.rect.Intersects(otherBlock.rect) {
					block.applyDir()
					otherBlock.applyDir()
				}

				break
			}
			c.blocks[j] = otherBlock
		}

		block.applyDir()
		block.wallBounce(canvas.Width(), canvas.Height())
		c.blocks[i] = block
	}
}

func main() {
	//Create drawables
	drawables := make([]eff.Drawable, 3)
	drawables[0] = &dots{}
	drawables[1] = &rects{}
	drawables[2] = &collidingBlocks{}

	drawableIndex := 0

	//Create Eff Canvas
	canvas := eff.SDLCanvas{}
	canvas.SetWidth(1280)
	canvas.SetHeight(720)

	setDrawable := func(index int) {
		if index < 0 || index >= len(drawables) {
			return
		}

		if index == drawableIndex {
			return
		}

		if len(drawables) > 0 && drawableIndex >= 0 {
			canvas.RemoveDrawable(drawables[drawableIndex])
		}

		canvas.AddDrawable(drawables[index])

		drawableIndex = index
	}

	//Add drawables to canvas
	canvas.AddDrawable(drawables[0])

	canvas.AddKeyUpHandler(func(key string, canvas eff.Canvas) {
		// fmt.Println("Up", key)
		if key == "1" {
			setDrawable(0)
		} else if key == "2" {
			setDrawable(1)
		} else if key == "3" {
			setDrawable(2)
		}
	})

	//Start the run loop
	os.Exit(canvas.Run())
}
