package sdl

import "testing"

func TestGetTextSize(t *testing.T) {
	var tests = []struct {
		fontPath   string
		text       string
		pointSize  int
		shouldFail bool
	}{
		{
			fontPath:   "",
			text:       "",
			pointSize:  -1,
			shouldFail: true,
		},
		{
			fontPath:   "../examples/assets/fonts/roboto/Roboto-Medium.ttf",
			text:       "text",
			pointSize:  -1,
			shouldFail: false,
		},
		{
			fontPath:   "../examples/assets/fonts/roboto/Roboto-Medium.ttf",
			text:       "text",
			pointSize:  10,
			shouldFail: false,
		},
	}

	// Create listener for mainThread
	closeChan := make(chan struct{})
	go func() {
		for {
			select {
			case f := <-mainThread:
				f()
			case <-closeChan:
				return
			}
		}
	}()

	for _, test := range tests {
		g := &Graphics{}
		f, err := openFont(test.fontPath, test.pointSize)
		if test.shouldFail != (err != nil) {
			t.Fatal(err)
		}

		_, _, err = g.GetTextSize(f, test.text)
		if test.shouldFail != (err != nil) {
			t.Fatal(err)
		}
	}

	close(closeChan)
}
