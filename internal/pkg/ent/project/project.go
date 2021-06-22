// Code generated by entc, DO NOT EDIT.

package project

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGraphqlEndpoint holds the string denoting the graphql_endpoint field in the database.
	FieldGraphqlEndpoint = "graphql_endpoint"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgePages holds the string denoting the pages edge name in mutations.
	EdgePages = "pages"
	// EdgeQueries holds the string denoting the queries edge name in mutations.
	EdgeQueries = "queries"
	// EdgeDeployments holds the string denoting the deployments edge name in mutations.
	EdgeDeployments = "deployments"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "projects"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_projects"
	// PagesTable is the table the holds the pages relation/edge.
	PagesTable = "pages"
	// PagesInverseTable is the table name for the Page entity.
	// It exists in this package in order to avoid circular dependency with the "page" package.
	PagesInverseTable = "pages"
	// PagesColumn is the table column denoting the pages relation/edge.
	PagesColumn = "project_pages"
	// QueriesTable is the table the holds the queries relation/edge.
	QueriesTable = "graph_ql_queries"
	// QueriesInverseTable is the table name for the GraphQLQuery entity.
	// It exists in this package in order to avoid circular dependency with the "graphqlquery" package.
	QueriesInverseTable = "graph_ql_queries"
	// QueriesColumn is the table column denoting the queries relation/edge.
	QueriesColumn = "project_queries"
	// DeploymentsTable is the table the holds the deployments relation/edge.
	DeploymentsTable = "deployments"
	// DeploymentsInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentsInverseTable = "deployments"
	// DeploymentsColumn is the table column denoting the deployments relation/edge.
	DeploymentsColumn = "project_deployments"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldGraphqlEndpoint,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "projects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_projects",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
