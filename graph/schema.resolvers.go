package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/ent/page"
	"github.com/pepsighan/nocodepress_backend/ent/project"
	"github.com/pepsighan/nocodepress_backend/ent/user"
	"github.com/pepsighan/nocodepress_backend/graph/generated"
	"github.com/pepsighan/nocodepress_backend/graph/model"
	"github.com/pepsighan/nocodepress_backend/internal/auth"
	"github.com/pepsighan/nocodepress_backend/internal/db"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*ent.Project, error) {
	user, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	var project *ent.Project

	// Do not create a page if project fails.
	err = db.WithTx(ctx, r.Ent, func(tx *ent.Tx) error {
		defaultPage, err := r.Ent.Page.Create().SetName("Default").SetRoute("/").Save(ctx)
		if err != nil {
			return err
		}

		project, err = r.Ent.Project.Create().
			SetName(input.Name).
			SetOwner(user).
			AddPages(defaultPage).
			Save(ctx)

		return err
	})

	return project, err
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProject) (*ent.Project, error) {
	owner, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ID, owner.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	err = prj.Update().
		SetName(input.Name).
		SetNillableGraphqlEndpoint(input.GraphqlEndpoint).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) CreatePage(ctx context.Context, input model.NewPage) (*ent.Page, error) {
	user, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	prj, err := user.QueryProjects().Where(project.IDEQ(input.ProjectID)).First(ctx)
	if err != nil {
		return nil, err
	}

	return r.Ent.Page.Create().
		SetName(input.Name).
		SetRoute(input.Route).
		SetPageOf(prj).
		Save(ctx)
}

func (r *mutationResolver) DeletePage(ctx context.Context, projectID uuid.UUID, pageID uuid.UUID) (*ent.Page, error) {
	user, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	prj, err := user.QueryProjects().Where(project.IDEQ(projectID)).First(ctx)
	if err != nil {
		return nil, err
	}

	pgCount, err := prj.QueryPages().Count(ctx)
	if err != nil {
		return nil, err
	}

	if pgCount == 1 {
		return nil, fmt.Errorf("cannot delete the only page")
	}

	pg, err := prj.QueryPages().Where(page.IDEQ(pageID)).First(ctx)
	if err != nil {
		return nil, err
	}

	err = r.Ent.Page.DeleteOne(pg).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (r *projectResolver) Pages(ctx context.Context, obj *ent.Project) ([]*ent.Page, error) {
	return obj.QueryPages().All(ctx)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	// This will return nil if there is no logged in user.
	return auth.UserFromContext(ctx, r.Ent, r.FirebaseAuth)
}

func (r *queryResolver) MyProjects(ctx context.Context) ([]*ent.Project, error) {
	user, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
	if err != nil {
		return nil, err
	}

	return user.QueryProjects().All(ctx)
}

func (r *queryResolver) MyProject(ctx context.Context, id uuid.UUID) (*ent.Project, error) {
	owner, err := auth.RequireUserFromContext(ctx, r.Ent, r.FirebaseAuth)
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
