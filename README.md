# go-wikipedia-crawler

## Management Problem

A technology firm has decided to use a Go web crawler program to scrape Wikipedia pages on intelligent systems and robotics to create its own online library. The firm would like to make use of Go concurrency because it will be faster than their Python Scrapy tool which runs slowly due to sequential searching through ordered lists of web pages.

## Project Objective

This project is a Go concurrent web crawler and scraper that works with the Colly framework to extract text from Wikipedia articles and export structured data in JSON Lines format.

The current implementation processes pages sequentially; controlled concurrent processing will be added after measuring the sequential baseline.

## Development Status

The current version sequentially processes all ten required Wikipedia articles, extracts and cleans their paragraph text, and writes one JSON object per article to a JSON Lines file. Unit tests cover paragraph cleaning and JSON encoding. The application records crawl activity, failures, and sequential processing time. Controlled concurrent processing will be added in the next development stage.

## Logging

The application writes crawl activity to `logs/crawl.log` and displays the same messages in the terminal. The log records requested URLs, response statuses, successfully scraped articles, failures, output record counts, and sequential crawl time.

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

Run the program:

```bash
go run .
```

The generated JSON Lines file contains 10 records and is saved at `output/wikipedia_articles.jl`.

## Testing

Run the automated unit tests:

```bash
go test ./...
```

The unit tests use synthetic data and do not access Wikipedia. They verify paragraph cleaning, whitespace normalization, and JSON Lines encoding.

## Performance

The sequential crawler successfully processed all ten target pages.

| Processing mode | Attempted | Successful | Failed | Crawl time |
|---|---:|---:|---:|---:|
| Sequential | 10 | 10 | 0 | 1.192 seconds |

Processing time can vary because of network conditions, Wikipedia response times, and caching. The sequential result will be compared with the concurrent implementation.

## GenAI Tools
I used OpenAI Codex as a programming tutor and review assistant during this project. It helped to:

* break the assignment into milestones
* explain Go and Colly concepts
* interpret compiler and HTTP errors
* review code and JSON Lines output

I reviewed, entered, tested, and revised the project code and documentation. AI-generated suggestions were validated with Go formatting, tests, vetting, benchmarks, and manual application runs.

## License

The Go source code in this repository is licensed under the MIT License. Extracted Wikipedia content remains subject to Wikipedia's applicable content license and attribution requirements.