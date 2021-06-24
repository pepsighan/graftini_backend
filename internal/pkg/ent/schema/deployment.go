package schema

import (
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
		field.String("status").GoType(DeploymentStatus("")),
		// These are newly added fields, so will require a default value for older
		// rows, hence `CURRENT_TIMESTAMP`.
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
