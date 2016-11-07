package drawable

import "github.com/forestgiant/eff"

// MyDrawable Boilerplate drawable struct
type MyDrawable struct {
	initialized bool
}

// Init Boilerplate Init function, used to setup the drawable
func (m *MyDrawable) Init(c eff.Canvas) {

	m.initialized = true
}

// Initialized returns true if the drawable has been initialized
func (m *MyDrawable) Initialized() bool { return m.initialized }

// Draw called once per frame, this function calls the canvas draw functions
func (m *MyDrawable) Draw(c eff.Canvas) {

}

// Update called once per frame, this function is responsible for updating the state of the drawable
func (m *MyDrawable) Update(c eff.Canvas) {

}
