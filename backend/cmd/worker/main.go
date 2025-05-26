package main

import (
	"fmt"
	"log"
	"time"

	"trend-summary-engine/internal/db"
	"trend-summary-engine/internal/scraper"
	"trend-summary-engine/internal/summarizer"
)

func main() {
	fmt.Println("Starting background worker...")

	// Connect to DB
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	for {
		// Get next unsummarized article
		article, err := database.GetNextPendingArticle()
		if err != nil {
			log.Printf("No articles or error: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		log.Printf("Processing article ID %d: %s", article.ID, article.URL)

		// Step 1: Extract content
		content, err := scraper.Extract(article.URL)
		if err != nil {
			log.Printf("Failed to extract content: %v", err)
			continue
		}

		// Step 2: Summarize content
		summary, err := summarizer.Summarize(content)
		if err != nil {
			log.Printf("Failed to summarize: %v", err)
			continue
		}

		// Step 3: Save summary back to DB
		err = database.SaveSummary(article.ID, summary)
		if err != nil {
			log.Printf("Failed to save summary: %v", err)
		} else {
			log.Printf("Successfully summarized article ID %d", article.ID)
		}

		time.Sleep(3 * time.Second)
	}
}
