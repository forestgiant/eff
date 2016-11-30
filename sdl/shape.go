package sdl

import "github.com/forestgiant/eff"

type Shape struct {
	Container

	rect   eff.Rect
	parent eff.Container
	scale  float64
}

func (shape *Shape) Draw(canvas eff.Canvas) {

}

func (shape *Shape) Rect() eff.Rect {
	return shape.rect
}

func (shape *Shape) SetParent(c eff.Container) {
	shape.parent = c
}

func (shape *Shape) Parent() eff.Container {
	return shape.parent
}
