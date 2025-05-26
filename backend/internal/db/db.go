package db

import "trend-summary-engine/graph/model"

func Connect() (*DB, error) {
	// Implement your DB connection logic here
	return &DB{}, nil
}

type DB struct{}

func (db *DB) GetNextPendingArticle() (*model.Article, error) {
	// Query the database for one article where summary is null
	return &model.Article{
		ID:  1,
		URL: "https://example.com/article",
	}, nil
}

func (db *DB) SaveSummary(articleID int, summary string) error {
	// Update the article record with the summary text
	return nil
}
