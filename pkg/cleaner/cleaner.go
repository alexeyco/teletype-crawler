package cleaner

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Cleaner struct{}

func NewCleaner() *Cleaner {
	return &Cleaner{}
}

func (c *Cleaner) Clean(html string, selectors ...string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", fmt.Errorf(`error making document: %w`, err)
	}

	for _, sel := range selectors {
		doc.Find(sel).Each(func(_ int, sel *goquery.Selection) {
			sel.Remove()
		})
	}

	return doc.Find("body").Html()
}
