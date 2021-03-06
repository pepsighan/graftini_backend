package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/backend/auth"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/backend/customer"
	"github.com/pepsighan/graftini_backend/internal/backend/deployclient"
	"github.com/pepsighan/graftini_backend/internal/backend/errs"
	"github.com/pepsighan/graftini_backend/internal/backend/graph/generated"
	model1 "github.com/pepsighan/graftini_backend/internal/backend/graph/model"
	"github.com/pepsighan/graftini_backend/internal/backend/sanitize"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/pkg/db"
	"github.com/pepsighan/graftini_backend/internal/pkg/domain"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/deployment"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/earlyaccess"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/graphqlquery"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/page"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/project"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/template"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/user"
	"github.com/pepsighan/graftini_backend/internal/pkg/imagekit"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
	"github.com/pepsighan/graftini_backend/internal/pkg/storage"
)

func (r *deploymentResolver) Status(ctx context.Context, obj *ent.Deployment) (string, error) {
	return string(obj.Status), nil
}

func (r *fileResolver) FileURL(ctx context.Context, obj *ent.File) (string, error) {
	return imagekit.GetImageKitURLForFile(config.ImageKitURLEndpoint, obj.ID, obj.Kind), nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input model1.UpdateProfile) (*ent.User, error) {
	authUser := auth.RequiredAuthenticatedUser(ctx)

	return authUser.Update().
		SetFirstName(input.FirstName).
		SetLastName(input.LastName).
		Save(ctx)
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model1.NewProject) (*ent.Project, error) {
	authUser := auth.RequiredAuthenticatedUser(ctx)

	projectCount, err := r.Ent.Project.Query().
		Where(project.HasOwnerWith(user.IDEQ(authUser.ID))).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// Do not allow production users (other than admin) to create any more than two projects.
	if config.Env.IsProduction() && projectCount >= 2 && !authUser.IsAdmin {
		return nil, logger.Error(errs.ErrProjectLimitExceeded)
	}

	var parsedTemplate *schema.ProjectTemplateSnapshot

	if input.TemplateID != nil {
		usedTemplate, err := r.Ent.Template.Get(ctx, *input.TemplateID)
		if err != nil {
			return nil, err
		}

		parsedTemplate, err = schema.ParseStringToProjectTemplateSnapshot(usedTemplate.Snapshot)
		if err != nil {
			return nil, err
		}
	}

	var project *ent.Project

	// Do not create a page if project fails.
	err = db.WithTx(ctx, r.Ent, func(tx *ent.Tx) error {
		savedPages := []*ent.Page{}

		if parsedTemplate == nil {
			// Use the blank project template when no template is provided.
			defaultPage, err := r.Ent.Page.
				Create().
				SetName("Home").
				SetRoute("/").
				SetComponentMap(*input.DefaultPageComponentMap).
				Save(ctx)
			if err != nil {
				return err
			}

			savedPages = append(savedPages, defaultPage)
		} else {
			// Otherwise use the template to create a new project.
			for _, templatePages := range parsedTemplate.Pages {
				componentMap, err := json.Marshal(templatePages.ComponentMap)
				if err != nil {
					return err
				}

				page, err := r.Ent.Page.
					Create().
					SetName(templatePages.Name).
					SetRoute(templatePages.Route).
					SetComponentMap(string(componentMap)).
					Save(ctx)
				if err != nil {
					return err
				}

				savedPages = append(savedPages, page)
			}
		}

		project, err = r.Ent.Project.Create().
			SetName(input.Name).
			SetOwner(authUser).
			AddPages(savedPages...).
			Save(ctx)

		return err
	})

	return project, err
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model1.UpdateProject) (*ent.Project, error) {
	owner := auth.RequiredAuthenticatedUser(ctx)

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

func (r *mutationResolver) DeleteProject(ctx context.Context, projectID uuid.UUID) (*ent.Project, error) {
	owner := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(projectID, owner.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	deploy, grpc, err := deployclient.GrpcClient()
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	err = db.WithTx(ctx, r.Ent, func(tx *ent.Tx) error {
		err = tx.Project.DeleteOne(prj).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = deploy.DeleteProjectDeployment(ctx, &service.DeleteProjectDeploymentRequest{
			ProjectID: prj.ID[:],
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return prj, nil
}

func (r *mutationResolver) DeployProject(ctx context.Context, projectID uuid.UUID) (*ent.Deployment, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	deploy, grpc, err := deployclient.GrpcClient()
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(projectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := deploy.DeployProject(ctx, &service.DeployRequest{
		ProjectID: prj.ID[:],
	})
	if err != nil {
		return nil, err
	}

	deploymentID, err := uuid.FromBytes(reply.DeploymentID)
	if err != nil {
		return nil, err
	}

	// Create a deployment here.
	return r.Ent.Deployment.Get(ctx, deploymentID)
}

func (r *mutationResolver) UpdateProjectDesign(ctx context.Context, input model1.UpdateProjectDesign) (*ent.Project, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ProjectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// Update all the pages in a transaction.
	err = db.WithTx(ctx, r.Ent, func(tx *ent.Tx) error {
		for _, pg := range input.Pages {
			_, err := r.Ent.Page.
				UpdateOneID(pg.PageID).
				SetComponentMap(pg.ComponentMap).
				Save(ctx)

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return prj, nil
}

func (r *mutationResolver) CreatePage(ctx context.Context, input model1.NewPage) (*ent.Page, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ProjectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	route := sanitize.CleanRoute(input.Route)
	pageExists, err := prj.QueryPages().
		Where(page.RouteEQ(route)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}

	if pageExists {
		return nil, logger.Errorf("cannot create a page with existing route")
	}

	return r.Ent.Page.Create().
		SetName(input.Name).
		SetRoute(route).
		SetComponentMap(input.ComponentMap).
		SetPageOf(prj).
		Save(ctx)
}

func (r *mutationResolver) DuplicatePage(ctx context.Context, input model1.DuplicatePage) (*ent.Page, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ProjectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := prj.QueryPages().
		Where(page.IDEQ(input.CopyPageID)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	route := sanitize.CleanRoute(input.Route)
	pageExists, err := prj.QueryPages().
		Where(page.RouteEQ(route)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}

	if pageExists {
		return nil, logger.Errorf("cannot create a page with existing route")
	}

	return r.Ent.Page.Create().
		SetName(input.Name).
		SetRoute(route).
		SetComponentMap(existing.ComponentMap).
		SetPageOf(prj).
		Save(ctx)
}

func (r *mutationResolver) UpdatePage(ctx context.Context, input model1.UpdatePage) (*ent.Page, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ProjectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := prj.QueryPages().
		Where(page.IDEQ(input.PageID)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// Check if the new updated route already exists elsewhere. We do not
	// allow duplicate routes.
	route := sanitize.CleanRoute(input.Route)
	pageExists, err := prj.QueryPages().
		Where(page.And(page.RouteEQ(route), page.IDNEQ(existing.ID))).
		Exist(ctx)
	if err != nil {
		return nil, err
	}

	if pageExists {
		return nil, logger.Errorf("cannot update a page with existing route")
	}

	return existing.Update().
		SetName(input.Name).
		SetRoute(input.Route).
		Save(ctx)
}

func (r *mutationResolver) DeletePage(ctx context.Context, projectID uuid.UUID, pageID uuid.UUID) (*ent.Page, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	prj, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(projectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	pgCount, err := prj.QueryPages().
		Count(ctx)
	if err != nil {
		return nil, err
	}

	if pgCount == 1 {
		return nil, logger.Errorf("cannot delete the only page")
	}

	pg, err := prj.QueryPages().
		Where(page.IDEQ(pageID)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	if pg.Route == "/" {
		return nil, logger.Errorf("cannot delete the default page")
	}

	err = r.Ent.Page.DeleteOne(pg).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (r *mutationResolver) CreateQuery(ctx context.Context, input model1.NewGraphQLQuery) (*ent.GraphQLQuery, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	pg, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(input.ProjectID, user.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return r.Ent.GraphQLQuery.Create().
		SetVariableName(input.VariableName).
		SetGqlAst(input.GqlAst).
		SetQueryOf(pg).
		Save(ctx)
}

func (r *mutationResolver) DeleteQuery(ctx context.Context, projectID uuid.UUID, queryID uuid.UUID) (*ent.GraphQLQuery, error) {
	user := auth.RequiredAuthenticatedUser(ctx)

	pg, err := r.Ent.Project.Query().ByIDAndOwnedBy(projectID, user.ID).First(ctx)
	if err != nil {
		return nil, err
	}

	query, err := pg.QueryQueries().Where(graphqlquery.IDEQ(queryID)).First(ctx)
	if err != nil {
		return nil, err
	}

	err = r.Ent.GraphQLQuery.DeleteOne(query).Exec(ctx)
	return query, err
}

func (r *mutationResolver) UploadFile(ctx context.Context, file graphql.Upload) (*ent.File, error) {
	return storage.UploadFile(ctx, file.File, file.ContentType, config.GoogleCloudStorageBucket, r.Storage, r.Ent)
}

func (r *mutationResolver) IsEarlyAccessAllowed(ctx context.Context, email string) (bool, error) {
	allowed, err := r.Ent.EarlyAccess.Query().
		Where(earlyaccess.EmailEQ(email)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

func (r *mutationResolver) ContactUs(ctx context.Context, input model1.ContactUsMessage) (*time.Time, error) {
	err := customer.SendContactUsEmail(ctx, input.Name, input.Email, input.Content, r.Ent)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &now, nil
}

func (r *projectResolver) Pages(ctx context.Context, obj *ent.Project) ([]*ent.Page, error) {
	return obj.QueryPages().
		Order(ent.Asc(page.FieldCreatedAt)).
		All(ctx)
}

func (r *projectResolver) Queries(ctx context.Context, obj *ent.Project) ([]*ent.GraphQLQuery, error) {
	return obj.QueryQueries().All(ctx)
}

func (r *projectResolver) AppURL(ctx context.Context, obj *ent.Project) (*string, error) {
	if obj.RefID == nil {
		return nil, nil
	}

	url := domain.GenerateDomainNameFromRefID(*obj.RefID, config.Env)
	return &url, nil
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	// This will return nil if there is no logged in user.
	return auth.GetUserFromBearerAuthInContext(ctx, r.Ent, r.FirebaseAuth)
}

func (r *queryResolver) MyProjects(ctx context.Context) ([]*ent.Project, error) {
	user := auth.RequiredAuthenticatedUser(ctx)
	return user.QueryProjects().
		Order(ent.Asc(project.FieldName)).
		All(ctx)
}

func (r *queryResolver) MyProject(ctx context.Context, id uuid.UUID) (*ent.Project, error) {
	owner := auth.RequiredAuthenticatedUser(ctx)

	return r.Ent.Project.Query().
		ByIDAndOwnedBy(id, owner.ID).
		First(ctx)
}

func (r *queryResolver) MyLastDeployment(ctx context.Context, projectID uuid.UUID) (*ent.Deployment, error) {
	owner := auth.RequiredAuthenticatedUser(ctx)

	deploy, grpc, err := deployclient.GrpcClient()
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	project, err := r.Ent.Project.Query().
		ByIDAndOwnedBy(projectID, owner.ID).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// Get the last deployment.
	deployment, err := project.QueryDeployments().
		Order(ent.Desc(deployment.FieldCreatedAt)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// The project has never been deployed.
			return nil, nil
		}

		return nil, err
	}

	// If the status has settled return the deployment.
	if deployment.Status == schema.DeploymentCancelled || deployment.Status == schema.DeploymentError || deployment.Status == schema.DeploymentReady {
		return deployment, nil
	}

	// Otherwise, fetch the current status.
	_, err = deploy.CheckStatus(ctx, &service.StatusRequest{DeploymentID: deployment.ID[:]})
	if err != nil {
		return nil, err
	}

	return r.Ent.Deployment.Get(ctx, deployment.ID)
}

func (r *queryResolver) File(ctx context.Context, fileID uuid.UUID) (*ent.File, error) {
	return r.Ent.File.Get(ctx, fileID)
}

func (r *queryResolver) Templates(ctx context.Context) ([]*ent.Template, error) {
	return r.Ent.Template.
		Query().
		Order(ent.Asc(template.FieldCreatedAt)).
		All(ctx)
}

func (r *templateResolver) FileURL(ctx context.Context, obj *ent.Template) (*string, error) {
	zeroUUID := uuid.UUID{}
	if obj.PreviewFileID == zeroUUID {
		return nil, nil
	}

	fileURL := imagekit.GetImageKitURLForFile(config.ImageKitURLEndpoint, obj.PreviewFileID, file.KindImage)
	return &fileURL, nil
}

// Deployment returns generated.DeploymentResolver implementation.
func (r *Resolver) Deployment() generated.DeploymentResolver { return &deploymentResolver{r} }

// File returns generated.FileResolver implementation.
func (r *Resolver) File() generated.FileResolver { return &fileResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Template returns generated.TemplateResolver implementation.
func (r *Resolver) Template() generated.TemplateResolver { return &templateResolver{r} }

type deploymentResolver struct{ *Resolver }
type fileResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type templateResolver struct{ *Resolver }
