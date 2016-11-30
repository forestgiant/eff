package sdl

import "github.com/forestgiant/eff"

type Container struct {
	children []eff.Drawable
}

func (c *Container) AddChild(d eff.Drawable) {
	if d == nil {
		return
	}

	d.SetParent(c)
	//Need to copy scale
	//Need to copy renderer
	c.children = append(c.children, d)
}

func (c *Container) RemoveChild(d eff.Drawable) {
	if d == nil {
		return
	}

	index := -1
	for i, child := range c.children {
		if d == child {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	c.children = append(c.children[:index], c.children[index+1:]...)
}

func (c *Container) Children() []eff.Drawable {
	return c.children
}
