# go-wikipedia-crawler

## Management Problem

A technology firm has decided to use a Go web crawler program to scrape Wikipedia pages on intelligent systems and robotics to create its own online library. The firm would like to make use of Go concurrency because it will be faster than their Python Scrapy tool which runs slowly due to sequential searching through ordered lists of web pages.

## Project Objective

This project is a Go concurrent web crawler and scraper that works with the Colly framework to extract text from Wikipedia articles and export structured data in JSON Lines format.

The scraper extracts paragraph text from each article and removes inline citation markers because the reference lists are not included in the output. The source Wikipedia URL is retained for attribution and verification.

## Development Status

The current version visits one Wikipedia article, extracts its title and main paragraph text, removes inline citation markers, and writes the result to a JSON Lines file. Support for all ten target pages, automated tests, logging, and concurrent processing will be added in later stages.

## Output Format

The program writes extracted article data to `output/wikipedia_articles.jl`. Each line is an independent JSON object with the following fields:

- `title`: The Wikipedia article title
- `url`: The source Wikipedia URL
- `text`: Cleaned text extracted from the article paragraphs

Example:

```json
{"title":"Robotics","url":"https://en.wikipedia.org/wiki/Robotics","text":"Robotics is the interdisciplinary study..."}
```

## Running the Program

Requirements:

- Go 1.27.1
- Internet access

Clone the repository and enter its directory:

```bash
git clone https://github.com/gummybeardream/go-wikipedia-crawler.git
cd go-wikipedia-crawler
```

## GenAI Tools 
I used OpenAI Codex as a programming tutor and review assistant during this project. It helped to:

* break the assignment into milestones
* explain Go and Colly concepts
* interpret compiler and HTTP errors 
* review code and JSON Lines output

I reviewed, entered, tested, and revised the project code and documentation. AI-generated suggestions were validated with Go formatting, tests, vetting, benchmarks, and manual application runs.

## License

The Go source code in this repository is licensed under the MIT License. Extracted Wikipedia content remains subject to Wikipedia's applicable content license and attribution requirements.