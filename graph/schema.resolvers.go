package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	"github.com/gustavopcr/touchdown/graph/model"
	"github.com/gustavopcr/touchdown/internal"
)

// Verify is the resolver for the verify field.
func (r *mutationResolver) Verify(ctx context.Context, score string) (*model.Combination, error) {
	arr, err := internal.ParseScore(score)
	if err != nil {
		return nil, err
	}
	value := internal.GetPos(arr[0]) * internal.GetPos(arr[1])
	return &model.Combination{Combinations: value}, nil
}

// Scores is the resolver for the scores field.
func (r *queryResolver) Scores(ctx context.Context) ([]*model.Score, error) {
	panic(fmt.Errorf("not implemented: Scores - scores"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
