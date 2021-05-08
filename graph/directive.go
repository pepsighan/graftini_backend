package graph

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/graph/generated"
	iauth "github.com/pepsighan/nocodepress_backend/internal/auth"
)

func NewDirective(entClient *ent.Client, firebaseAuth *auth.Client) generated.DirectiveRoot {
	return generated.DirectiveRoot{
		IsAuthenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			user, err := iauth.GetUserFromBearerAuthInContext(ctx, entClient, firebaseAuth)
			if err != nil {
				return nil, err
			}

			if user == nil {
				return nil, iauth.ErrUnauthorizedAccess
			}

			return next(iauth.WithAuthUser(ctx, user))
		},
	}
}
