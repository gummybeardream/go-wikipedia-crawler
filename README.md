# go-wikipedia-crawler

## Management Problem

A technology firm has decided to use a Go web crawler program to scrape Wikipedia pages on intelligent systems and robotics to create its own online library. The firm would like to make use of Go concurrency because it will be faster than their Python Scrapy tool which runs slowly due to sequential searching through ordered lists of web pages.

## Project Objective

This project is a Go concurrent web crawler and scraper that works with the Colly framework to extract text from Wikipedia articles and export structured data in JSON Lines format.

The scraper extracts paragraph text from each article and removes inline citation markers because the reference lists are not included in the output. The source Wikipedia URL is retained for attribution and verification.