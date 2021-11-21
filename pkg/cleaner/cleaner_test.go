package cleaner_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/teletype-crawler/pkg/cleaner"
)

func TestCleaner_Clean(t *testing.T) {
	html := `<p>foo</p><tags><tag>bar</tag><tag>fiz</tag><tag>baz</tag></tags>`
	expected := `<p>foo</p>`

	actual, err := cleaner.NewCleaner().Clean(html, "tags")

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}
