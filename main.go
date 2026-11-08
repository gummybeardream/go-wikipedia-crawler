package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const maxConcurrentRequests = 3

type scrapeResult struct {
	index   int
	article Article
	err     error
}

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

	results := make(chan scrapeResult, len(urls))
	semaphore := make(chan struct{}, maxConcurrentRequests)

	var waitGroup sync.WaitGroup

	logger.Printf("starting concurrent crawl of %d URLs with a limit of %d requests",
		len(urls),
		maxConcurrentRequests,
	)
	crawlStart := time.Now()

	failedCount := 0
	for index, url := range urls {
		waitGroup.Add(1)

		go func(index int, url string) {
			defer waitGroup.Done()

			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			article, err := scrapeArticle(url, logger)

			results <- scrapeResult{
				index:   index,
				article: article,
				err:     err,
			}
		}(index, url)
	}

	waitGroup.Wait()
	close(results)

	crawlDuration := time.Since(crawlStart)

	orderedArticles := make([]Article, len(urls))
	successful := make([]bool, len(urls))

	for result := range results {
		if result.err != nil {
			logger.Printf("could not scrape article: %v", result.err)
			failedCount++
			continue
		}

		orderedArticles[result.index] = result.article
		successful[result.index] = true
	}

	articles := make([]Article, 0, len(urls))

	for index, article := range orderedArticles {
		if successful[index] {
			articles = append(articles, article)
		}
	}

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
		"concurrent crawl complete: %d succeeded, %d failed, crawl time %s",
		len(articles),
		failedCount,
		crawlDuration,
	)
}
