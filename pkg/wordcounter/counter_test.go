package wordcounter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/teletype-crawler/pkg/wordcounter"
)

func TestCounter_Count(t *testing.T) {
	text := "foo b    bar\n\nfix\r\nbaz"

	actual := wordcounter.NewCounter().Count(text)

	assert.Equal(t, 4, actual)
}
