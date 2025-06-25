package graph

import (
	"errors"
	"time"

	"trend-summary-engine/internal/auth"
	"trend-summary-engine/internal/db"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var RegisterResolver = func(p graphql.ResolveParams) (interface{}, error) {
	input := p.Args["input"].(map[string]interface{})
	email := input["email"].(string)
	password := input["password"].(string)

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := db.User{Email: email, PasswordHash: hashedPassword}
	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"user":  user,
	}, nil
}

var LoginResolver = func(p graphql.ResolveParams) (interface{}, error) {
	input := p.Args["input"].(map[string]interface{})
	email := input["email"].(string)
	password := input["password"].(string)

	var user db.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"user":  user,
	}, nil
}

var SaveArticleResolver = func(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context
	userId, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	input := p.Args["input"].(map[string]interface{})
	url := input["url"].(string)

	article := db.Article{
		URL:       url,
		UserID:    userId,
		CreatedAt: time.Now(),
	}
	if err := db.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

var queryType = graphql.ObjectConfig{
	Name:   "Query",
	Fields: graphql.Fields{
		// Add getSavedArticles in Phase 2
	},
}

var mutationType = graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"register": &graphql.Field{
			Type:    authPayloadType,
			Args:    graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{Type: registerInputType}},
			Resolve: RegisterResolver,
		},
		"login": &graphql.Field{
			Type:    authPayloadType,
			Args:    graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{Type: loginInputType}},
			Resolve: LoginResolver,
		},
		"saveArticle": &graphql.Field{
			Type:    articleType,
			Args:    graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{Type: articleInputType}},
			Resolve: SaveArticleResolver,
		},
	},
}
