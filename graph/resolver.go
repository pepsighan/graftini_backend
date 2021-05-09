package graph

import (
	"firebase.google.com/go/v4/auth"
	"github.com/pepsighan/graftini_backend/ent"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Ent          *ent.Client
	FirebaseAuth *auth.Client
}

func NewResolver(client *ent.Client, auth *auth.Client) *Resolver {
	return &Resolver{
		Ent:          client,
		FirebaseAuth: auth,
	}
}
