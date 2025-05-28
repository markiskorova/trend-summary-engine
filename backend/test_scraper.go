package main

import (
	"fmt"
	"log"

	"github.com/markiskorova/trend-summary-engine/backend/internal/scraper"
)

func main() {
	url := "https://www.bbc.com/news/articles/c0k3700zljjo" // Replace with any valid article URL
	text, err := scraper.Extract(url)
	if err != nil {
		log.Fatalf("Error scraping article: %v", err)
	}
	fmt.Println(text)
}
