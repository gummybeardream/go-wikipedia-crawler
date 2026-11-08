package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Wikipedia URLs for topics of interest.
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	if err := os.MkdirAll("logs", 0755); err != nil {
		fmt.Println("Could not create log directory:", err)
		return
	}
	logFile, err := os.Create("logs/crawl.log")
	if err != nil {
		fmt.Println("Could not create log file:", err)
		return
	}
	defer logFile.Close()

	logger := log.New(
		io.MultiWriter(os.Stdout, logFile),
		"",
		log.LstdFlags,
	)

	var articles []Article

	logger.Printf("starting sequential crawl of %d URLs", len(urls))
	crawlStart := time.Now()

	failedCount := 0
	for _, url := range urls {
		article, err := scrapeArticle(url, logger)
		if err != nil {
			logger.Printf("could not scrape article: %v", err)
			failedCount++
			continue
		}

		articles = append(articles, article)
	}

	crawlDuration := time.Since(crawlStart)

	if len(articles) == 0 {
		logger.Printf("no articles were scraped successfully")
		return
	}
	if err := os.MkdirAll("output", 0755); err != nil {
		logger.Printf("could not create output directory: %v", err)
		return
	}

	outputFile, err := os.Create("output/wikipedia_articles.jl")
	if err != nil {
		logger.Printf("could not create output file: %v", err)
		return
	}
	defer outputFile.Close()

	for _, article := range articles {
		if err := writeArticle(outputFile, article); err != nil {
			logger.Printf("could not write article: %v", err)
			return
		}
	}

	logger.Printf(
		"wrote %d of %d articles to output/wikipedia_articles.jl",
		len(articles),
		len(urls),
	)

	logger.Printf(
		"sequential crawl complete: %d succeeded, %d failed, crawl time %s",
		len(articles),
		failedCount,
		crawlDuration,
	)
}
