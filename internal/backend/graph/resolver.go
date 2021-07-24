package graph

import (
	"cloud.google.com/go/storage"
	"firebase.google.com/go/v4/auth"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Ent          *ent.Client
	FirebaseAuth *auth.Client
	Storage      *storage.Client
}

func NewResolver(client *ent.Client, auth *auth.Client, storage *storage.Client) *Resolver {

	return &Resolver{
		Ent:          client,
		FirebaseAuth: auth,
		Storage:      storage,
	}
}
