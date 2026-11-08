package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestCleanParagraph(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "normal paragraph",
			html:     "<p>Robotics is useful.</p>",
			expected: "Robotics is useful.",
		},
		{
			name:     "extra whitespace",
			html:     "<p>Robotics   is\n useful.</p>",
			expected: "Robotics is useful.",
		},
		{
			name:     "one citation",
			html:     `<p>Robotics is useful.<sup class="reference">[1]</sup> It helps people.</p>`,
			expected: "Robotics is useful. It helps people.",
		},
		{
			name:     "multiple citations",
			html:     `<p>Robots move.<sup class="reference">[1]</sup> They also sense.<sup class="reference">[2]</sup></p>`,
			expected: "Robots move. They also sense.",
		},
		{
			name:     "legitimate brackets",
			html:     "<p>The coordinates are [x, y, z].</p>",
			expected: "The coordinates are [x, y, z].",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			document, err := goquery.NewDocumentFromReader(strings.NewReader(test.html))
			if err != nil {
				t.Fatalf("could not parse test HTML: %v", err)
			}

			actual := cleanParagraph(document.Find("p"))

			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
