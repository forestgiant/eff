package eff

import "testing"

func TestAddRemoveChild(t *testing.T) {
	var tests = []struct {
		child      Drawable
		shouldFail bool
	}{
		{child: nil, shouldFail: true},
		{child: &drawable{}, shouldFail: false},
	}

	d := &drawable{}

	// Test race condition
	go func() {
		d.Children()
	}()

	// Remove a child before adding
	d.RemoveChild(&drawable{})

	var expected int
	for _, test := range tests {
		if err := d.AddChild(test.child); test.shouldFail != (err != nil) {
			t.Fatal(err)
		}

		// If the test is suppose to pass then increment expected
		if !test.shouldFail {
			expected++
		}
	}

	// Verify the children were added
	children := d.Children()
	if len(children) != expected {
		t.Fatal("The number of children should match the total children added")
	}

	// Now remove the children
	for _, test := range tests {
		if err := d.RemoveChild(test.child); test.shouldFail != (err != nil) {
			t.Fatal(err)
		}
	}

	// Test if a drawable gets set to nil
	d = nil
	if err := d.AddChild(&drawable{}); err == nil {
		t.Fatal(err)
	}
}
