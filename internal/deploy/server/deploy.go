package server

import (
	context "context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/deploy/vercel"
	"github.com/pepsighan/graftini_backend/internal/pkg/domain"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

func deployProject(ctx context.Context, projectID string, deployment *ent.Deployment, snapshot *schema.DeploymentSnapshot, generateCtx *appgenerate.GenerateContext) (*service.DeployReply, error) {
	vercelProj, err := createVercelProjectIfNotExists(ctx, projectID)
	if err != nil {
		return nil, logger.Errorf("could not create a vercel project: %w", err)
	}

	err = attachGraftiniSubdomain(ctx, projectID, snapshot.Project)
	if err != nil {
		return nil, logger.Errorf("could not attach a graftini subdomain: %w", err)
	}

	projectPath, err := appgenerate.GenerateCodeBaseForProject(ctx, snapshot.Pages, generateCtx)
	defer projectPath.Cleanup() // Cleanup the folder regardless of the error.

	if err != nil {
		return nil, logger.Errorf("could not generate code base for project: %w", err)
	}

	projectFiles, err := uploadProjectFiles(ctx, string(projectPath))
	if err != nil {
		return nil, logger.Errorf("could not upload files to vercel: %w", err)
	}

	vercelDeployment, err := vercel.CreateNewDeployment(ctx, vercelProj.Name, projectFiles)
	if err != nil {
		return nil, logger.Errorf("could not create a deployment on vercel: %w", err)
	}

	_, err = updateDeployment(ctx, deployment, vercelDeployment.ID, schema.DeploymentStatus(vercelDeployment.ReadyState))
	if err != nil {
		return nil, logger.Errorf("could not record the deployment: %w", err)
	}

	deployID, err := deployment.ID.MarshalBinary()
	if err != nil {
		return nil, logger.Errorf("could not get the deployment id: %w", err)
	}
	return &service.DeployReply{DeploymentID: deployID}, nil
}

// createVercelProjectIfNotExists creates a vercel project if it does not exist.
func createVercelProjectIfNotExists(ctx context.Context, projectID string) (*vercel.Project, error) {
	projectName := generateVercelProjectName(projectID)
	project, err := vercel.GetProject(ctx, projectName)
	if err != nil {
		return nil, err
	}

	if project != nil {
		return project, nil
	}

	return vercel.CreateProject(ctx, projectName)
}

// updateDeployment updates the deployment with the final status.
func updateDeployment(ctx context.Context, deployment *ent.Deployment, vercelDeploymentID string, status schema.DeploymentStatus) (*ent.Deployment, error) {
	return deployment.Update().
		SetVercelDeploymentID(vercelDeploymentID).
		SetStatus(status).
		Save(ctx)
}

// generateVercelProjectName generates a vercel project name with the prefix `{env}-graftini-app` and
// its UUID.
func generateVercelProjectName(projectID string) string {
	return fmt.Sprintf("%v-graftini-app-%v", config.Env, projectID)
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

// attachGraftiniSubdomain attaches a graftini.app subdomain to the project if it is
// not already attached.
func attachGraftiniSubdomain(ctx context.Context, projectID string, snapshot *schema.ProjectSnapshot) error {
	domainName := domain.GenerateDomainNameFromRefID(snapshot.RefID, config.Env)

	vercelProjectName := generateVercelProjectName(projectID)
	isAttached, err := vercel.DoesDomainExistInProject(ctx, vercelProjectName, domainName)
	if err != nil {
		return err
	}

	if isAttached {
		return nil
	}

	return vercel.AddDomainToProject(ctx, vercelProjectName, domainName)
}
