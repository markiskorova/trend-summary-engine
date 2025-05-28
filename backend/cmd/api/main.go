package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/markiskorova/trend-summary-engine/backend/graph"
	"github.com/markiskorova/trend-summary-engine/backend/graph/generated"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authmw.Middleware(srv)) // <--- wrap GraphQL server in auth middleware

	log.Println("🚀 Server ready at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
