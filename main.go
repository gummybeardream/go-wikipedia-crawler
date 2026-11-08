package main

import (
	"fmt"
	"os"
)

func main() {
	// Wikipedia URLs for topics of interest.
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
	}

	article, err := scrapeArticle(urls[0])
	if err != nil {
		fmt.Println("Could not scrape article:", err)
		return
	}

	if err := os.MkdirAll("output", 0755); err != nil {
		fmt.Println("Could not create output directory:", err)
		return
	}

	outputFile, err := os.Create("output/wikipedia_articles.jl")
	if err != nil {
		fmt.Println("Could not create output file:", err)
		return
	}
	defer outputFile.Close()

	if err := writeArticle(outputFile, article); err != nil {
		fmt.Println("Could not write article:", err)
		return
	}

	fmt.Println("Wrote article to output/wikipedia_articles.jl")
	fmt.Println("Scraped article:", article.Title)
	fmt.Println("Text characters:", len(article.Text))

}
