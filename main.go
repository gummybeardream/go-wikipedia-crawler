package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Wikipedia URLs for topics of interest.
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
	}

	collector := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.UserAgent(
			"GoWikipediaCrawlerBot/1.0 (https://github.com/gummybeardream/go-wikipedia-crawler; educational project)",
		),
	)
	article := Article{}
	var paragraphs []string

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting:", request.URL)
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("Received status:", response.StatusCode)
	})

	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("Request failed:", response.Request.URL, err)
	})

	collector.OnHTML("h1#firstHeading", func(element *colly.HTMLElement) {
		article.Title = element.Text
		article.URL = element.Request.URL.String()
		fmt.Println("Stored title:", article.Title)
		fmt.Println("Stored URL:", article.URL)
	})

	collector.OnHTML("#mw-content-text .mw-parser-output p", func(element *colly.HTMLElement) {
		paragraph := cleanParagraph(element.DOM)

		if paragraph != "" {
			paragraphs = append(paragraphs, paragraph)
		}
	})
	if err := collector.Visit(urls[0]); err != nil {
		fmt.Println("Could not visit page:", err)
		return
	}

	article.Text = strings.Join(paragraphs, "\n\n")
	if article.Title == "" || article.URL == "" || article.Text == "" {
		fmt.Println("Article is missing required data")
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

	fmt.Println("Paragraphs extracted:", len(paragraphs))
	fmt.Println("Text characters:", len(article.Text))

	if len(paragraphs) > 0 {
		fmt.Println("First paragraph:", paragraphs[0])
	}

}
