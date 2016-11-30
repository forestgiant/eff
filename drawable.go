package eff

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Draw(canvas Canvas)

	SetRect(Rect)
	Rect() Rect

	SetParent(Container)
	Parent() Container

	SetScale(float64)
	Scale() float64

	SetGraphics(Graphics)
	Graphics() Graphics
}
