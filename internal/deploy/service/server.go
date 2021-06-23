package service

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
)

// Server is used to implement the GRPC deploy service.
type Server struct {
	UnimplementedDeployServer
	Ent *ent.Client
}

func (s *Server) DeployProject(ctx context.Context, in *DeployRequest) (*DeployReply, error) {
	projectID, err := uuid.FromBytes(in.GetProjectID())
	if err != nil {
		return nil, err
	}

	project, err := s.Ent.Project.Get(ctx, projectID)
	projectPath, err := appgenerate.GenerateCodeBaseForProject(ctx, project)
	defer projectPath.Cleanup() // Cleanup the folder regardless of the error.

	if err != nil {
		return nil, err
	}

	fmt.Println("Deploying the project from " + projectPath)
	return nil, nil
}
