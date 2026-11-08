package main

import (
	"fmt"
	"os"
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

	var articles []Article

	for _, url := range urls {
		article, err := scrapeArticle(url)
		if err != nil {
			fmt.Println("Could not scrape article:", err)
			continue
		}

		articles = append(articles, article)
	}

	if len(articles) == 0 {
		fmt.Println("No articles were scraped successfully")
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

	for _, article := range articles {
		if err := writeArticle(outputFile, article); err != nil {
			fmt.Println("Could not write article:", err)
			return
		}
	}

	fmt.Println("Wrote article to output/wikipedia_articles.jl")

	fmt.Printf(
		"Wrote %d of %d articles to output/wikipedia_articles.jl\n",
		len(articles),
		len(urls),
	)

}
