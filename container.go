package eff

type Container interface {
	AddChild(Drawable)
	RemoveChild(Drawable)
	Children() []Drawable
}
