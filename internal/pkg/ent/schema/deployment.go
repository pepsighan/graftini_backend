package schema

import (
	"encoding/json"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		// This is required when creating. Default value is for existing ones.
		field.String("vercel_deployment_id").Default(""),
		field.String("status").GoType(DeploymentStatus("")),
		// This is actually a required field. Default value is for existing fields.
		field.String("project_snapshot").Default("").
			Validate(func(s string) error {
				var artifacts DeploymentSnapshot
				// It is valid only if it can be read into the schema.
				return json.Unmarshal([]byte(s), &artifacts)
			}),
		field.Time("created_at").Default(time.Now).Immutable().
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployments_of", Project.Type).
			Ref("deployments").
			Unique(),
	}
}

// DeploymentStatus is the current state the deployment is in.
type DeploymentStatus string

const (
	DeploymentInitializing DeploymentStatus = "INITIALIZING"
	DeploymentAnalyzing    DeploymentStatus = "ANALYZING"
	DeploymentBuilding     DeploymentStatus = "BUILDING"
	DeploymentDeploying    DeploymentStatus = "DEPLOYING"
	DeploymentReady        DeploymentStatus = "READY"
	DeploymentError        DeploymentStatus = "ERROR"
	DeploymentCancelled    DeploymentStatus = "CANCELED"
)

// DeploymentSnapshot is a snapshot of a project when deployed.
type DeploymentSnapshot struct {
	Project *ProjectSnapshot `json:"project"`
	Pages   []*PageSnapshot  `json:"pages"`
}

type ProjectSnapshot struct {
	RefID string `json:"refId"`
	Name  string `json:"name"`
}

type PageSnapshot struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Route        string       `json:"route"`
	ComponentMap ComponentMap `json:"componentMap"`
}
