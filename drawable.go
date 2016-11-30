package eff

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Draw(canvas Canvas)
	Rect() Rect
	SetParent(Container)
	Parent() Container
}
