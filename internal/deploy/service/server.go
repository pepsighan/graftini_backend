package service

import (
	context "context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/deploy/vercel"
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

	vercelProj, err := createVercelProjectIfNotExists(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not create a vercel project: %w", err)
	}

	projectPath, err := appgenerate.GenerateCodeBaseForProject(ctx, project)
	defer projectPath.Cleanup() // Cleanup the folder regardless of the error.

	if err != nil {
		return nil, fmt.Errorf("could not generate code base for project: %w", err)
	}

	projectFiles, err := uploadProjectFiles(ctx, string(projectPath))
	if err != nil {
		return nil, fmt.Errorf("could not upload files to vercel: %w", err)
	}

	vercelDeployment, err := vercel.CreateNewDeployment(ctx, vercelProj.Name, projectFiles)
	if err != nil {
		return nil, fmt.Errorf("could not create a deployment on vercel: %w", err)
	}

	deployment, err := recordDeployment(ctx, vercelDeployment.ID, project, s.Ent)
	if err != nil {
		return nil, fmt.Errorf("could not record the deployment: %w", err)
	}

	deployID, err := deployment.ID.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("could not get the deployment id: %w", err)
	}
	return &DeployReply{DeploymentID: deployID}, nil
}

// createVercelProjectIfNotExists creates a vercel project if it does not exist.
func createVercelProjectIfNotExists(ctx context.Context, projectID uuid.UUID) (*vercel.Project, error) {
	projectName := generateProjectName(projectID)
	project, err := vercel.GetProject(ctx, projectName)
	if err != nil {
		return nil, err
	}

	if project != nil {
		return project, nil
	}

	return vercel.CreateProject(ctx, projectName)
}

// recordDeployment records the deployment to be tracked later.
func recordDeployment(ctx context.Context, vercelDeploymentID string, project *ent.Project, client *ent.Client) (*ent.Deployment, error) {
	return client.Deployment.Create().
		SetVercelDeploymentID(vercelDeploymentID).
		SetStatus(schema.DeploymentInitializing).
		SetDeploymentsOf(project).
		Save(ctx)
}

// generateProjectName generates a project name with the prefix `app` and its UUID.
func generateProjectName(projectID uuid.UUID) string {
	return fmt.Sprintf("app%v", projectID)
}

// uploadProjectFiles uploads all the files in the project path to vercel.
func uploadProjectFiles(ctx context.Context, projectPath string) ([]*vercel.ProjectFile, error) {
	files := []*vercel.ProjectFile{}

	err := filepath.WalkDir(projectPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// If an error has occurred, just short-circuit it. We have nothing else to
			// do now.
			return err
		}

		if d.IsDir() {
			// We do not upload a directory. It has no meaning in the context of vercel.
			return nil
		}

		hash, size, err := vercel.UploadDeploymentFile(ctx, path)
		if err != nil {
			return err
		}

		files = append(files, &vercel.ProjectFile{
			File: strings.Replace(path, projectPath, "", 1),
			SHA:  hash,
			Size: size,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
