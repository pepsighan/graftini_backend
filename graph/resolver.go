package graph

import "github.com/pepsighan/nocodepress_backend/ent"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Ent *ent.Client
}

func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		Ent: client,
	}
}
