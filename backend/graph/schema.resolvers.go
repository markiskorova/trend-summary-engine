package graph

import (
	"context"

	"github.com/markiskorova/trend-summary-engine/backend/graph/generated"
)

func (r *queryResolver) Health(ctx context.Context) (string, error) {
	return "OK", nil
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
