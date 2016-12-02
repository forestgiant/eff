package eff

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Draw(canvas Canvas)

	SetRect(Rect)
	Rect() Rect

	SetParent(Drawable)
	Parent() Drawable

	SetGraphics(Graphics)
	Graphics() Graphics

	SetUpdateHandler(func())
	HandleUpdate()

	AddChild(Drawable)
	RemoveChild(Drawable)
	Children() []Drawable
}
