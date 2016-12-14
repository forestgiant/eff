package shapes

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type block struct {
	dir   eff.Point
	color eff.Color
	rect  eff.Rect
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

type CollidingBlocks struct {
	eff.Shape

	blocks []block
}

func (c *CollidingBlocks) Init(width int, height int) {
	blockCount := (width * height) / 200
	blockSize := 5
	for i := 0; i < blockCount; i++ {
		b := block{
			rect: eff.Rect{
				X: rand.Intn(width - blockSize),
				Y: rand.Intn(height - blockSize),
				W: blockSize,
				H: blockSize,
			},
			color: eff.RandomColor(),
			dir: eff.Point{
				X: rand.Intn(4) + 1,
				Y: rand.Intn(4) + 1,
			},
		}
		c.blocks = append(c.blocks, b)
	}

	c.SetUpdateHandler(func() {
		var rects []eff.Rect
		var colors []eff.Color
		for i, block := range c.blocks {
			block.applyDir()
			block.wallBounce(width, height)
			c.blocks[i] = block

			rects = append(rects, block.rect)
			colors = append(colors, block.color)
		}

		c.Clear()
		c.FillColorRects(rects, colors)
	})
}
