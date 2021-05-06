package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pepsighan/nocodepress_backend/auth"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/ent/project"
	"github.com/pepsighan/nocodepress_backend/ent/user"
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

func (r *mutationResolver) CreatePage(ctx context.Context, input model.NewPage) (*ent.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePage(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Pages(ctx context.Context, obj *ent.Project) ([]*ent.Page, error) {
	return obj.QueryPages().All(ctx)
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

func (r *queryResolver) MyProject(ctx context.Context, id uuid.UUID) (*ent.Project, error) {
	owner, err := auth.UserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	return r.Ent.Project.Query().
		Where(project.And(
			project.IDEQ(id),
			project.HasOwnerWith(user.IDEQ(owner.ID)),
		)).
		First(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
