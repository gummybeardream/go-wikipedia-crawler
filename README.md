# go-wikipedia-crawler

A concurrent Go and Colly application that extracts cleaned text from ten Wikipedia articles and writes one structured record per article to a JSON Lines file.

## Management Problem

A technology firm has decided to use a Go web crawler program to scrape Wikipedia pages on intelligent systems and robotics to create its own online library. The firm would like to make use of Go concurrency because it will be faster than their Python Scrapy tool which runs slowly due to sequential searching through ordered lists of web pages.

## Project Objective

This project is a concurrent Go web crawler and scraper that works with the Colly framework to extract text from Wikipedia articles and export structured data in JSON Lines format. The crawler uses controlled concurrency to process multiple pages while limiting the number of simultaneous requests.

## Implementation Status

The current version concurrently processes all ten required Wikipedia articles, extracts and cleans their paragraph text, and writes one JSON object per article to a JSON Lines file. Unit tests cover paragraph cleaning and JSON encoding. Logging records crawl activity, failures, output counts, and processing time. Concurrent access is limited to three requests, and race-detector checks found no data races.

## Methodology

The scraper first requests each target page with a Wikimedia-compatible user agent. Colly callbacks extract the article title and paragraph elements. Goquery removes inline citation elements before the paragraph text is normalized and combined.

The application starts one goroutine per URL while a semaphore limits active requests to three. Completed results are sent through a channel and restored to their original URL order. The main goroutine writes the records sequentially to prevent concurrent file writes.

## Concurrency

The program starts one goroutine for each target URL and uses a semaphore to limit processing to three simultaneous requests. A `sync.WaitGroup` ensures that the program waits for all scraping tasks to finish. Workers send results through a buffered channel rather than modifying the shared result slice directly.

Because concurrent requests may finish in any order, each result retains its original URL index. The main goroutine restores the input order before writing the JSON Lines file. JSON writing remains sequential to avoid concurrent file writes.

## Repository Contents

| Path | Role |
|---|---|
| `main.go` | Coordinates concurrent crawling, result ordering, logging, and output |
| `scraper.go` | Downloads and extracts one Wikipedia article |
| `article.go` | Defines the JSON article record |
| `cleaner.go` | Removes citation elements and normalizes paragraph text |
| `writer.go` | Encodes an article as one JSON Lines record |
| `*_test.go` | Contains unit tests and benchmark code |
| `output/wikipedia_articles.jl` | Required 10-record crawl result |
| `logs/crawl.log` | Log from a completed concurrent crawl |
| `bin/go-wikipedia-crawler` | Compiled macOS ARM64 executable |

## Running the Program

Requirements:

- Go 1.27.1
- Internet access

Clone the repository and enter its directory:

```bash
git clone https://github.com/gummybeardream/go-wikipedia-crawler.git
cd go-wikipedia-crawler
```

Download module dependencies:

```bash
go mod download
```

Run the program:

```bash
go run .
```

The generated JSON Lines file contains 10 records and is saved at `output/wikipedia_articles.jl`.

## Building the Executable

Build the macOS command-line executable from the repository root:

```bash
mkdir -p bin
go build -o bin/go-wikipedia-crawler .
```

Run the executable from the repository root:

```bash
./bin/go-wikipedia-crawler
```

The program requires internet access and writes article data to `output/wikipedia_articles.jl` and crawl activity to `logs/crawl.log`.

The included executable was built for macOS on Apple Silicon (`arm64`). Users on other operating systems or processor architectures should build the executable from source.

## Output Format

The program writes extracted article data to `output/wikipedia_articles.jl`. Each line is an independent JSON object with the following fields:

- `title`: The Wikipedia article title
- `url`: The source Wikipedia URL
- `text`: Cleaned text extracted from the article paragraphs

Example:

```json
{"title":"Robotics","url":"https://en.wikipedia.org/wiki/Robotics","text":"Robotics is the interdisciplinary study..."}
```

## Testing

Run the unit tests:

```bash
go test ./...
```

The unit tests use synthetic data and do not access Wikipedia. They verify paragraph cleaning, whitespace normalization, and JSON Lines encoding.

Run the race detector:

```bash
go test -race ./...
go run -race .
```

## Logging

The application writes crawl activity to `logs/crawl.log` and displays the same messages in the terminal. The log records requested URLs, response statuses, successfully scraped articles, failures, output record counts, and concurrent crawl time.

## Performance

The sequential crawler successfully processed all ten target pages.

| Processing mode | Attempted | Successful | Failed | Crawl time |
|---|---:|---:|---:|---:|
| Sequential | 10 | 10 | 0 | 1.192 seconds |
| Concurrent | 10 | 10 | 0 | 0.655 seconds |

In these runs, controlled concurrency reduced crawl time from 1.192 seconds to 0.655 seconds, an approximate speedup of 1.82 times and a time reduction of 45 percent. Results can vary because of network conditions, Wikipedia response times, caching, and normal system activity.

### Software Metrics

The JSON Lines output contains 10 records and is approximately 248 KiB. Unit-test coverage is 6.1 percent. Tests focus on the deterministic paragraph-cleaning and JSON-encoding components. The live crawler is also checked through manual application runs and Go's race detector.

### Paragraph-Cleaning Benchmark

The benchmark uses synthetic HTML and measures citation removal and whitespace normalization without making network requests. It was run five times on an Apple M1 using:

```bash
go test -bench=BenchmarkCleanParagraph -benchmem -count=5
```

The median result from five runs was:

| Metric | Result |
|---|---:|
| Time per operation | 1,119 ns/op |
| Memory per operation | 1,896 B/op |
| Allocations per operation | 30 allocs/op |

The benchmark completed one paragraph-cleaning operation in approximately 1.12 microseconds. Results can vary based on hardware, Go version, and system activity.

## Results and Interpretation

The final crawl retrieved all ten required Wikipedia articles without request failures. The output contains one valid JSON object per page, with no empty required fields or detected HTML markup. Controlled concurrency reduced crawl time in the recorded comparison while preserving record completeness and input order.

## Limitations

- Results depend on network availability and Wikipedia response times.
- Wikipedia HTML changes could require selector updates.
- The performance comparison uses a small set of ten pages.
- The included executable runs only on macOS ARM64.
- The scraper extracts paragraph text but does not preserve headings, lists, tables, images, or reference lists.
- The application removes citation markers, so individual claims are not linked to their references.

## GenAI Tools

I used OpenAI Codex as a programming tutor and review assistant during this project. It helped to:

* break the assignment into milestones
* explain Go and Colly concepts
* interpret compiler and HTTP errors
* review code and JSON Lines output

I reviewed, entered, tested, and revised the project code and documentation. AI-generated suggestions were validated with Go formatting, unit tests, race detection, vetting, benchmarks, and manual application runs.

## License

The original Go source code in this repository is licensed under the MIT License. This project uses the Colly web-scraping framework, which is licensed under the Apache License 2.0. Third-party dependencies remain subject to their respective licenses.

Extracted Wikipedia content remains subject to Wikipedia's applicable content license and attribution requirements.
