package eff

// Drawable interface describing required methods for drawable objects
type Drawable interface {
	Init(canvas Canvas)
	Draw(canvas Canvas)
	Update(canvas Canvas)
	Initialized() bool
}
