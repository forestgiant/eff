package eff

// Clickable interface describing the required methods for clickable objects
type Clickable interface {
	Hitbox() Rect
	MouseDown(leftState bool, middleState bool, rightState bool)
	MouseUp(leftState bool, middleState bool, rightState bool)
	MouseOver()
	MouseOut()
	IsMouseOver() bool
}
