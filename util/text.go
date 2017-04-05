package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/forestgiant/eff"
)

// EllipseText returns a substring of argument text that does not exceed the argument width, if string is shortened '...' is added
func EllipseText(font eff.Font, text string, width int, g eff.Graphics) (string, error) {
	if width <= 0 {
		return "", errors.New("EllipseText Error: Invalid text width")
	}

	if font == nil {
		return "", errors.New("EllipseText Error: Font is nil")
	}

	if g == nil {
		return "", errors.New("EllipseText Error: Graphics is nil")
	}
	if len(text) == 0 {
		return "", errors.New("EllipseText Error: Text string is empty")
	}

	textW, _, err := g.GetTextSize(font, text)
	if err != nil {
		return "", err
	}

	if textW <= width {
		return text, nil
	}

	for textW > width {
		text = text[:len(text)-1]
		textW, _, err = g.GetTextSize(font, text+"...")
		if err != nil {
			return "", err
		}
	}

	return text + "...", nil

}

// CenterTextInRect finds the point that would center the text in the canvas, assumes the font is already set
func CenterTextInRect(font eff.Font, text string, rect eff.Rect, g eff.Graphics) (eff.Point, error) {
	textW, textH, err := g.GetTextSize(font, text)
	if err != nil {
		return eff.Point{}, err
	}

	return eff.Point{
		X: rect.X + (rect.W-textW)/2,
		Y: rect.Y + (rect.H-textH)/2,
	}, nil
}

// GetMultilineText creates and array of strings that do not exceed the maxWidth.  Returns the slice of strings and the height of the text per line
func GetMultilineText(font eff.Font, text string, maxWidth int, g eff.Graphics) ([]string, int, error) {
	var lines []string

	words := strings.Split(text, " ")
	var maxH int
	var line string
	for _, word := range words {
		oldLine := line
		line = fmt.Sprintf("%s%s ", line, word)
		w, h, err := g.GetTextSize(font, line)
		if err != nil {
			return nil, 0, err
		}

		if w > maxWidth {
			lines = append(lines, oldLine)
			line = word
		}

		if h > maxH {
			maxH = h
		}
	}

	if len(line) > 0 {
		lines = append(lines, line)
	}

	return lines, maxH, nil
}
