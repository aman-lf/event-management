package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/aman-lf/event-management/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user []*model.User
}
