package wordcounter

import (
	"strings"
	"unicode/utf8"
)

type Counter struct{}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Count(text string) int {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", "")

	parts := strings.Split(text, " ")

	var words int
	for _, part := range parts {
		length := utf8.RuneCountInString(part)

		if length > proposalMaxLength {
			words++
		}
	}

	return words
}
