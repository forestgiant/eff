package eff

type Image interface {
	Drawable

	Path() string
}
