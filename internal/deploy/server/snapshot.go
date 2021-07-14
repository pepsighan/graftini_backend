package server

import (
	context "context"

	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// createSnapshotOfProject creates a snapshot of a project for the deployment so that
// the deployment can be pin-pointed to the exact copy of the project.
// This is useful in two ways:
// 1. The user may be actively making changes to the project. What the user intended to
// 		deploy may not be deployed.
// 2. To create a history of deployments to support rollbacks.
func createSnapshotOfProject(ctx context.Context, project *ent.Project, client *ent.Client) (*ent.Deployment, error) {
	return client.Deployment.Create().
		SetVercelDeploymentID(""). // We do not have a deployment ID.
		SetStatus(schema.DeploymentInitializing).
		SetDeploymentsOf(project).
		Save(ctx)
}
