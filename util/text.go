package util

import (
	"errors"
	"log"
	"math"
	"strings"

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

// GetMultilineText creates and array of strings that do not exceed the maxWidth.  Returns the slice of strings and the height of the text per line
func GetMultilineText(font eff.Font, text string, maxWidth int, c eff.Canvas) ([]string, int) {
	w, h, err := c.GetTextSize(font, text)
	if err != nil {
		log.Fatal(err)
	}

	lineCount := w / maxWidth
	maxRunesPerLine := len(text) / lineCount
	checkSize := int(float64(maxRunesPerLine) * float64(0.75))
	checkSize = int(math.Max(float64(checkSize), 1))
	var lines []string
	words := strings.Split(text, " ")
	wordIndex := 0
	for wordIndex < len(words) {
		w := 0
		line := ""
		for wordIndex < len(words) {
			if len(line) >= checkSize {
				w, _, err = c.GetTextSize(font, line+words[wordIndex]+" ")
				if w > maxWidth {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
			}

			line += words[wordIndex]
			line += " "
			wordIndex++
		}

		lines = append(lines, line)
	}

	return lines, h

}
