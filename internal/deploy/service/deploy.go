package service

import (
	context "context"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/vercel"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

func deployProject(ctx context.Context, project *ent.Project, deployment *ent.Deployment) (*DeployReply, error) {
	vercelProj, err := createVercelProjectIfNotExists(ctx, project.ID)
	if err != nil {
		return nil, fmt.Errorf("could not create a vercel project: %w", err)
	}

	err = attachGraftiniSubdomain(ctx, project)
	if err != nil {
		return nil, fmt.Errorf("could not attach a graftini subdomain: %w", err)
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

	_, err = updateDeployment(ctx, deployment, vercelDeployment.ID, schema.DeploymentStatus(vercelDeployment.ReadyState))
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

// recordDeployment records the deployment to be tracked later.
func recordNewDeployment(ctx context.Context, project *ent.Project, client *ent.Client) (*ent.Deployment, error) {
	return client.Deployment.Create().
		SetVercelDeploymentID(""). // We do not have a deployment ID.
		SetStatus(schema.DeploymentInitializing).
		SetDeploymentsOf(project).
		Save(ctx)
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
func generateVercelProjectName(projectID uuid.UUID) string {
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
func attachGraftiniSubdomain(ctx context.Context, project *ent.Project) error {
	domainName, err := getDomainNameForProject(ctx, project)
	if err != nil {
		return err
	}

	vercelProjectName := generateVercelProjectName(project.ID)
	isAttached, err := vercel.DoesDomainExistInProject(ctx, vercelProjectName, domainName)
	if err != nil {
		return err
	}

	if isAttached {
		return nil
	}

	return vercel.AddDomainToProject(ctx, vercelProjectName, domainName)
}

// getDomainNameForProject generates a new subdomain if it is not already.
func getDomainNameForProject(ctx context.Context, project *ent.Project) (string, error) {
	if project.RefID != nil {
		return GenerateDomainNameFromRefID(*project.RefID), nil
	}

	// Try to generate until there is a valid refID.
	for {
		newRefID, err := subdomainFromString(project.Name)
		if err != nil {
			return "", err
		}

		saved, err := project.Update().
			SetRefID(newRefID).
			Save(ctx)

		if ent.IsConstraintError(err) {
			// If the refID was not unique, regenerate it.
			continue
		}

		if err != nil {
			return "", err
		}

		return GenerateDomainNameFromRefID(*saved.RefID), nil
	}
}

// GenerateDomainNameFromRefID generates a full domain name from the Ref ID of the project.
func GenerateDomainNameFromRefID(refID string) string {
	return fmt.Sprintf("%v.%v", refID, suffixDomainName())
}

const graftiniAppDomain string = "graftini.app"

// suffixDomainName gives the domain suffix to use. We use [graftiniaAppDomain]
// for production and for others we append with development & local.
func suffixDomainName() string {
	if config.Env.IsProduction() {
		// This is hard-coded for the app. There is no other domain we use.
		return graftiniAppDomain
	}

	return fmt.Sprintf("%v.%v", config.Env, graftiniAppDomain)
}

// invalidSubdomainChars is any characters not alphanumeric and -.
var invalidSubdomainChars = regexp.MustCompile("[^a-zA-Z0-9-]+")

// invalidStartingDash is any - character in the start.
var invalidStartingDash = regexp.MustCompile("^-+")

// invalidEndingDash is any - character in the end.
var invalidEndingDash = regexp.MustCompile("-+$")

const nanoidCharacterSpace = "abcdefghijklmnopqrstuvwxyz0123456789"
const suffixLength = 8

// subdomainFromString gets a valid subdomain using the given string.
// It removes any invalid characters from the given name.
func subdomainFromString(name string) (string, error) {
	subdomain := invalidSubdomainChars.ReplaceAllString(name, "")
	subdomain = invalidStartingDash.ReplaceAllString(subdomain, "")
	subdomain = invalidEndingDash.ReplaceAllString(subdomain, "")

	subdomain = strings.ToLower(subdomain)

	randomSuffix, err := gonanoid.Generate(nanoidCharacterSpace, suffixLength)
	if err != nil {
		return "", fmt.Errorf("could not generate a random suffix: %w", err)
	}

	return fmt.Sprintf("%v-%v", subdomain, randomSuffix), nil
}
