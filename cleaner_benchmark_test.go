package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var benchmarkCleanedParagraph string

func BenchmarkCleanParagraph(b *testing.B) {
	const html = `
		<p>
			Robotics is an interdisciplinary field.
			<sup class="reference">[1]</sup>
			Robots can assist people in manufacturing,
			medicine, transportation, and research.
			<sup class="reference">[2]</sup>
		</p>`

	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		b.Fatalf("could not parse benchmark HTML: %v", err)
	}

	paragraph := document.Find("p")

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		benchmarkCleanedParagraph = cleanParagraph(paragraph)
	}
}
