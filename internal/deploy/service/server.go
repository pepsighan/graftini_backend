package service

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// Server is used to implement the GRPC deploy service.
type Server struct {
	UnimplementedDeployServer
	Ent *ent.Client
}

func (s *Server) DeployProject(ctx context.Context, in *DeployRequest) (*DeployReply, error) {
	projectID, err := uuid.FromBytes(in.GetProjectID())
	if err != nil {
		return nil, fmt.Errorf("could not get the project id: %w", err)
	}

	project, err := s.Ent.Project.Get(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not find the project: %w", err)
	}

	projectPath, err := appgenerate.GenerateCodeBaseForProject(ctx, project)
	defer projectPath.Cleanup() // Cleanup the folder regardless of the error.

	if err != nil {
		return nil, fmt.Errorf("could not generate code base for project: %w", err)
	}

	deployment, err := recordDeployment(ctx, project, s.Ent)
	if err != nil {
		return nil, fmt.Errorf("could not record the deployment: %w", err)
	}

	deployID, err := deployment.ID.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("could not get the deployment id: %w", err)
	}
	return &DeployReply{DeploymentID: deployID}, nil
}

// recordDeployment records the deployment to be tracked later.
func recordDeployment(ctx context.Context, project *ent.Project, client *ent.Client) (*ent.Deployment, error) {
	return client.Deployment.Create().
		SetStatus(schema.DeploymentInitializing).
		SetDeploymentsOf(project).
		Save(ctx)
}
