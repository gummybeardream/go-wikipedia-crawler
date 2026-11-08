package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func scrapeArticle(url string) (Article, error) {
	article := Article{}
	var paragraphs []string

	collector := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.UserAgent(
			"GoWikipediaCrawlerBot/1.0 (https://github.com/gummybeardream/go-wikipedia-crawler; educational project)",
		),
	)
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
	if err := collector.Visit(url); err != nil {
		return Article{}, fmt.Errorf("visit %s: %w", url, err)
	}

	article.Text = strings.Join(paragraphs, "\n\n")

	if article.Title == "" || article.URL == "" || article.Text == "" {
		return Article{}, fmt.Errorf("article from %s is missing required data", url)
	}

	return article, nil
}
