package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func cleanParagraph(paragraph *goquery.Selection) string {
	cleanedParagraph := paragraph.Clone()

	// Remove Wikipedia's inline citation elements before extracting text.
	cleanedParagraph.Find("sup.reference").Remove()
	// strings.Fields splits text at whitespace and strings.Join rebuilds it with single spaces.
	return strings.Join(strings.Fields(cleanedParagraph.Text()), " ")
}
