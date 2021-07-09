// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/deployment"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/file"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/graphqlquery"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/page"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/project"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescVercelDeploymentID is the schema descriptor for vercel_deployment_id field.
	deploymentDescVercelDeploymentID := deploymentFields[1].Descriptor()
	// deployment.DefaultVercelDeploymentID holds the default value on creation for the vercel_deployment_id field.
	deployment.DefaultVercelDeploymentID = deploymentDescVercelDeploymentID.Default.(string)
	// deploymentDescCreatedAt is the schema descriptor for created_at field.
	deploymentDescCreatedAt := deploymentFields[3].Descriptor()
	// deployment.DefaultCreatedAt holds the default value on creation for the created_at field.
	deployment.DefaultCreatedAt = deploymentDescCreatedAt.Default.(func() time.Time)
	// deploymentDescUpdatedAt is the schema descriptor for updated_at field.
	deploymentDescUpdatedAt := deploymentFields[4].Descriptor()
	// deployment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deployment.DefaultUpdatedAt = deploymentDescUpdatedAt.Default.(func() time.Time)
	// deployment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deployment.UpdateDefaultUpdatedAt = deploymentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deploymentDescID is the schema descriptor for id field.
	deploymentDescID := deploymentFields[0].Descriptor()
	// deployment.DefaultID holds the default value on creation for the id field.
	deployment.DefaultID = deploymentDescID.Default.(func() uuid.UUID)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[3].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileFields[4].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileFields[0].Descriptor()
	// file.DefaultID holds the default value on creation for the id field.
	file.DefaultID = fileDescID.Default.(func() uuid.UUID)
	graphqlqueryFields := schema.GraphQLQuery{}.Fields()
	_ = graphqlqueryFields
	// graphqlqueryDescCreatedAt is the schema descriptor for created_at field.
	graphqlqueryDescCreatedAt := graphqlqueryFields[3].Descriptor()
	// graphqlquery.DefaultCreatedAt holds the default value on creation for the created_at field.
	graphqlquery.DefaultCreatedAt = graphqlqueryDescCreatedAt.Default.(func() time.Time)
	// graphqlqueryDescUpdatedAt is the schema descriptor for updated_at field.
	graphqlqueryDescUpdatedAt := graphqlqueryFields[4].Descriptor()
	// graphqlquery.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	graphqlquery.DefaultUpdatedAt = graphqlqueryDescUpdatedAt.Default.(func() time.Time)
	// graphqlquery.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	graphqlquery.UpdateDefaultUpdatedAt = graphqlqueryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// graphqlqueryDescID is the schema descriptor for id field.
	graphqlqueryDescID := graphqlqueryFields[0].Descriptor()
	// graphqlquery.DefaultID holds the default value on creation for the id field.
	graphqlquery.DefaultID = graphqlqueryDescID.Default.(func() uuid.UUID)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescComponentMap is the schema descriptor for component_map field.
	pageDescComponentMap := pageFields[3].Descriptor()
	// page.ComponentMapValidator is a validator for the "component_map" field. It is called by the builders before save.
	page.ComponentMapValidator = pageDescComponentMap.Validators[0].(func(string) error)
	// pageDescCreatedAt is the schema descriptor for created_at field.
	pageDescCreatedAt := pageFields[4].Descriptor()
	// page.DefaultCreatedAt holds the default value on creation for the created_at field.
	page.DefaultCreatedAt = pageDescCreatedAt.Default.(func() time.Time)
	// pageDescUpdatedAt is the schema descriptor for updated_at field.
	pageDescUpdatedAt := pageFields[5].Descriptor()
	// page.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	page.DefaultUpdatedAt = pageDescUpdatedAt.Default.(func() time.Time)
	// page.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	page.UpdateDefaultUpdatedAt = pageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// pageDescID is the schema descriptor for id field.
	pageDescID := pageFields[0].Descriptor()
	// page.DefaultID holds the default value on creation for the id field.
	page.DefaultID = pageDescID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectFields[4].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectFields[5].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
