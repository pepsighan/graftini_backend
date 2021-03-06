package server

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/deploy/vercel"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// Server is used to implement the GRPC deploy service.
type Server struct {
	service.UnimplementedDeployServer
	Ent *ent.Client
}

func (s *Server) DeployProject(ctx context.Context, in *service.DeployRequest) (*service.DeployReply, error) {
	projectID, err := uuid.FromBytes(in.GetProjectID())
	if err != nil {
		return nil, logger.Errorf("could not get the project id: %w", err)
	}

	project, err := s.Ent.Project.Get(ctx, projectID)
	if err != nil {
		return nil, logger.Errorf("could not find the project: %w", err)
	}

	deployment, snapshot, err := initializeDeployment(ctx, project, s.Ent)
	if err != nil {
		return nil, logger.Errorf("could not create the deployment: %w", err)
	}

	reply, err := deployProject(ctx, project.ID.String(), deployment, snapshot, &appgenerate.GenerateContext{
		Ent: s.Ent,
	})
	if err != nil {
		if _, err := updateDeployment(ctx, deployment, "", schema.DeploymentError); err != nil {
			log.Errorf("failed to mark deployment as failed manually: %v", err)
		}

		return nil, err
	}

	return reply, nil
}

// CheckStatus checks the status of the deployment.
func (s *Server) CheckStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusReply, error) {
	deploymentID, err := uuid.FromBytes(in.GetDeploymentID())
	if err != nil {
		return nil, logger.Errorf("could not get the deployment id: %w", err)
	}

	deployment, err := s.Ent.Deployment.Get(ctx, deploymentID)
	if err != nil {
		return nil, logger.Errorf("could not find the deployment: %w", err)
	}

	// The deployment has not taken place yet or never will because it failed.
	if deployment.VercelDeploymentID == "" {
		// No new status to be found.
		return &service.StatusReply{DeploymentID: in.DeploymentID}, nil
	}

	vercelDeployment, err := vercel.GetDeployment(ctx, deployment.VercelDeploymentID)
	if err != nil {
		return nil, logger.Errorf("could not get vercel deployment: %w", err)
	}

	// Update the status if it has changed.
	if deployment.Status != schema.DeploymentStatus(vercelDeployment.ReadyState) {
		_, err = deployment.Update().
			SetStatus(schema.DeploymentStatus(vercelDeployment.ReadyState)).
			Save(ctx)
		if err != nil {
			return nil, logger.Errorf("could not update the deployment status")
		}
	}

	return &service.StatusReply{DeploymentID: in.DeploymentID}, nil
}

// DeleteProjectDeployment deletes the project deployment and its resources.
func (s *Server) DeleteProjectDeployment(ctx context.Context, in *service.DeleteProjectDeploymentRequest) (*service.DeleteProjectDeploymentReply, error) {
	projectID, err := uuid.FromBytes(in.GetProjectID())
	if err != nil {
		return nil, logger.Errorf("could not get the project id: %w", err)
	}

	projectName := generateVercelProjectName(projectID.String())

	vercelProject, err := vercel.GetProject(ctx, projectName)
	if err != nil {
		return nil, fmt.Errorf("could not find the deployed project: %w", err)
	}

	if vercelProject != nil {
		err := vercel.DeleteProject(ctx, projectName)
		if err != nil {
			return nil, fmt.Errorf("could not delete project: %w", err)
		}
	}

	return &service.DeleteProjectDeploymentReply{ProjectID: in.ProjectID}, nil
}
