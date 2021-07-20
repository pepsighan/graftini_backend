package server

import (
	context "context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// initializeDeployment initializes deployment and makes a project ready to deploy.
func initializeDeployment(ctx context.Context, project *ent.Project, client *ent.Client) (*ent.Deployment, *schema.DeploymentSnapshot, error) {
	snapshot, err := createDeploymentSnapshot(ctx, project)
	if err != nil {
		return nil, nil, err
	}

	snapshotJSON, err := convertSnapshotToJSON(snapshot)
	if err != nil {
		return nil, nil, err
	}

	deployment, err := client.Deployment.Create().
		SetVercelDeploymentID(""). // We do not have a deployment ID.
		SetStatus(schema.DeploymentInitializing).
		SetProjectSnapshot(snapshotJSON).
		SetDeploymentsOf(project).
		Save(ctx)

	if err != nil {
		return nil, nil, err
	}

	return deployment, snapshot, nil
}

// createDeploymentSnapshot creates a snapshot of a project for the deployment so that
// the deployment can be pin-pointed to the exact copy of the project.
// This is useful in two ways:
// 1. The user may be actively making changes to the project. What the user intended to
// 		deploy may not be deployed.
// 2. To create a history of deployments to support rollbacks.
func createDeploymentSnapshot(ctx context.Context, project *ent.Project) (*schema.DeploymentSnapshot, error) {
	pages, err := project.QueryPages().All(ctx)
	if err != nil {
		return nil, err
	}

	snapshot := &schema.DeploymentSnapshot{
		Project: &schema.ProjectSnapshot{
			Name: project.Name,
		},
	}

	// Assign a new ref ID if it does not exist in the project.
	if project.RefID == nil {
		subdomain, err := subdomainFromString(project.Name)
		if err != nil {
			return nil, err
		}

		snapshot.Project.RefID = subdomain

		// Save the subdomain for later deployments too.
		_, err = project.Update().
			SetRefID(subdomain).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		snapshot.Project.RefID = *project.RefID
	}

	for _, page := range pages {
		var compMap schema.ComponentMap
		err := json.Unmarshal([]byte(page.ComponentMap), &compMap)
		if err != nil {
			return nil, err
		}

		snapshot.Pages = append(snapshot.Pages, &schema.PageSnapshot{
			ID:           page.ID.String(),
			Name:         page.Name,
			Route:        page.Route,
			ComponentMap: compMap,
		})
	}

	return snapshot, nil
}

func convertSnapshotToJSON(snapshot *schema.DeploymentSnapshot) (string, error) {
	bytes, err := json.Marshal(snapshot)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
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
		return "", logger.Errorf("could not generate a random suffix: %w", err)
	}

	return fmt.Sprintf("%v-%v", subdomain, randomSuffix), nil
}
