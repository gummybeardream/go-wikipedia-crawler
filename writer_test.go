package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestWriteArticle(t *testing.T) {
	expected := Article{
		Title: "Robotics",
		URL:   "https://en.wikipedia.org/wiki/Robotics",
		Text:  `Robots can safely encode "quotation marks" and punctuation.`,
	}

	var output bytes.Buffer

	if err := writeArticle(&output, expected); err != nil {
		t.Fatalf("writeArticle returned an error: %v", err)
	}

	if !bytes.HasSuffix(output.Bytes(), []byte("\n")) {
		t.Error("JSON Lines record does not end with a newline")
	}

	var actual Article
	if err := json.Unmarshal(output.Bytes(), &actual); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	if actual != expected {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
