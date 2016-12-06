package sdl

import (
	"fmt"
	"os"
	"testing"

	"github.com/forestgiant/eff"
)

var testCanvas *Canvas

func TestMain(m *testing.M) {
	// Create testCanvas for all test to use
	testCanvas = NewCanvas("test", 640, 480, eff.Black(), 60, true)
	testCanvas.Run(
	//
	)
	t := m.Run()
	os.Exit(t)

}

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
	}

	for _, test := range tests {
		g := &Graphics{}
		f, err := openFont(test.fontPath, test.pointSize)
		if test.shouldFail != (err != nil) {
			t.Fatal(err)
		}

		w, h, err := g.GetTextSize(f, test.text)
		if test.shouldFail != (err != nil) {
			t.Fatal(err)
		}
		fmt.Println(w, h)
	}
}
