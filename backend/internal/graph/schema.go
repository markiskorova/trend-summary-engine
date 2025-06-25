package graph

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    graphql.NewObject(queryType),
	Mutation: graphql.NewObject(mutationType),
})
