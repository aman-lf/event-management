package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/aman-lf/event-management/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.User{
		ID:      randNumber.String(),
		Name:    input.Name,
		Email:   input.Email,
		Phoneno: "123",
	}
	r.user = append(r.user, todo)
	return todo, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
