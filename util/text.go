package util

import (
	"errors"
	"fmt"
	"log"
	"math"
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

	fmt.Println("Starting the text size query ")
	textW, _, err := g.GetTextSize(font, text)
	fmt.Println("Done getting the text size", textW)
	if err != nil {
		return "", err
	}

	if textW <= width {
		return text, nil
	}

	for textW > width {
		fmt.Println("Looping through the text")
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
func GetMultilineText(font eff.Font, text string, maxWidth int, g eff.Graphics) ([]string, int) {
	w, h, err := g.GetTextSize(font, text)
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
				w, _, err = g.GetTextSize(font, line+words[wordIndex]+" ")
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
