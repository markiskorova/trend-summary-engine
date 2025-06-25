package graph

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.ID},
		"email": &graphql.Field{Type: graphql.String},
	},
})

var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"url":       &graphql.Field{Type: graphql.String},
		"userId":    &graphql.Field{Type: graphql.ID},
		"createdAt": &graphql.Field{Type: graphql.String},
	},
})

var authPayloadType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthPayload",
	Fields: graphql.Fields{
		"token": &graphql.Field{Type: graphql.String},
		"user":  &graphql.Field{Type: userType},
	},
})

var registerInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RegisterInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var loginInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "LoginInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var articleInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ArticleInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"url": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})
