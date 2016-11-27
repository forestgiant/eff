package util

import (
	"errors"

	"github.com/forestgiant/eff"
)

// EllipseText returns a substring of argument text that does not exceed the argument width, if string is shortened '...' is added
func EllipseText(font eff.Font, text string, width int, c eff.Canvas) (string, error) {
	if width <= 0 {
		return "", errors.New("Invalid text width")
	}

	textW, _, err := c.GetTextSize(font, text)
	if err != nil {
		return "", err
	}

	if textW <= width {
		return text, nil
	}

	for textW > width {
		text = text[:len(text)-1]
		textW, _, err = c.GetTextSize(font, text+"...")
		if err != nil {
			return "", err
		}
	}

	return text + "...", nil

}

// CenterTextInRect finds the point that would center the text in the canvas, assumes the font is already set
func CenterTextInRect(font eff.Font, text string, rect eff.Rect, c eff.Canvas) (eff.Point, error) {
	textW, textH, err := c.GetTextSize(font, text)
	if err != nil {
		return eff.Point{}, err
	}

	return eff.Point{
		X: rect.X + (rect.W-textW)/2,
		Y: rect.Y + (rect.H-textH)/2,
	}, nil
}
