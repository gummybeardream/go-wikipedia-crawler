package main

import (
	"encoding/json"
	"io"
)

func writeArticle(writer io.Writer, article Article) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(article)
}
