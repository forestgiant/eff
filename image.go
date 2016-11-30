package eff

type Image interface {
	Drawable
	Container

	Path() string
}
