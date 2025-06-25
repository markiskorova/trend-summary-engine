package main

import (
	"log"
	"net/http"
	"os"

	"trend-summary-engine/internal/auth"
	"trend-summary-engine/internal/graph"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &graph.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Inject JWT-auth middleware
	http.Handle("/graphql", auth.Middleware(h))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s/graphql", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
