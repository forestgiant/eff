package main

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type block struct {
	dir       eff.Point
	colorRect eff.ColorRect
}

func (b *block) applyDir() {
	b.colorRect.Rect.X += b.dir.X
	b.colorRect.Rect.Y += b.dir.Y
}

func (b *block) wallBounce(width int, height int) {
	if b.colorRect.Rect.X < 0 || b.colorRect.Rect.X+b.colorRect.Rect.W > width {
		b.dir.X *= -1
	}

	if b.colorRect.Rect.Y < 0 || b.colorRect.Rect.Y+b.colorRect.Rect.H > height {
		b.dir.Y *= -1
	}
}

type collidingBlocks struct {
	blocks []block
}

func (c *collidingBlocks) Init(canvas eff.Canvas) {
	blockCount := 5000
	blockSize := 5
	for i := 0; i < blockCount; i++ {
		b := block{
			colorRect: eff.ColorRect{
				Rect: eff.Rect{
					X: rand.Intn(canvas.Width() - blockSize),
					Y: rand.Intn(canvas.Height() - blockSize),
					W: blockSize,
					H: blockSize,
				},
				Color: eff.RandomColor(),
			},
			dir: eff.Point{
				X: rand.Intn(4) + 1,
				Y: rand.Intn(4) + 1,
			},
		}
		c.blocks = append(c.blocks, b)
	}
}

func (c *collidingBlocks) Draw(canvas eff.Canvas) {
	var colorRects []eff.ColorRect
	for _, block := range c.blocks {
		colorRects = append(colorRects, block.colorRect)
	}

	canvas.DrawColorRects(colorRects)
}

func (c *collidingBlocks) Update(canvas eff.Canvas) {
	for i, block := range c.blocks {
		block.applyDir()
		block.wallBounce(canvas.Width(), canvas.Height())
		c.blocks[i] = block
	}
}
