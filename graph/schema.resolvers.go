package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pepsighan/nocodepress_backend/auth"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/graph/generated"
	"github.com/pepsighan/nocodepress_backend/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*ent.Project, error) {
	user, err := auth.UserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	return r.Ent.Project.Create().
		SetName(input.Name).
		SetOwner(user).
		Save(ctx)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	return auth.UserFromContext(ctx, r.Ent, r.FirebaseAuth)
}

func (r *queryResolver) MyProjects(ctx context.Context) ([]*ent.Project, error) {
	user, err := auth.UserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	return user.QueryProjects().All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
