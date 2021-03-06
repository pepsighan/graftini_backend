package graph

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql"
	iauth "github.com/pepsighan/graftini_backend/internal/backend/auth"
	"github.com/pepsighan/graftini_backend/internal/backend/errs"
	"github.com/pepsighan/graftini_backend/internal/backend/graph/generated"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

func NewDirective(entClient *ent.Client, firebaseAuth *auth.Client) generated.DirectiveRoot {
	return generated.DirectiveRoot{
		IsAuthenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			user, err := iauth.GetUserFromBearerAuthInContext(ctx, entClient, firebaseAuth)
			if err != nil {
				return nil, err
			}

			if user == nil {
				return nil, logger.Error(errs.ErrUnauthorizedAccess)
			}

			return next(iauth.WithAuthUser(ctx, user))
		},
	}
}
