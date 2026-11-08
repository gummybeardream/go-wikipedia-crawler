package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func scrapeArticle(url string, logger *log.Logger) (Article, error) {
	article := Article{}
	var paragraphs []string

	collector := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.UserAgent(
			"GoWikipediaCrawlerBot/1.0 (https://github.com/gummybeardream/go-wikipedia-crawler; educational project)",
		),
	)
	collector.OnRequest(func(request *colly.Request) {
		logger.Printf("visiting %s", request.URL)
	})

	collector.OnResponse(func(response *colly.Response) {
		logger.Printf(
			"received status %d for %s",
			response.StatusCode,
			response.Request.URL,
		)
	})

	collector.OnError(func(response *colly.Response, err error) {
		logger.Printf("request failed for %s: %v", response.Request.URL, err)
	})

	collector.OnHTML("h1#firstHeading", func(element *colly.HTMLElement) {
		article.Title = element.Text
		article.URL = element.Request.URL.String()
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

	logger.Printf("scraped article %q", article.Title)

	return article, nil
}
