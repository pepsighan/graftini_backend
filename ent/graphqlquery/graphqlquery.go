// Code generated by entc, DO NOT EDIT.

package graphqlquery

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the graphqlquery type in the database.
	Label = "graph_ql_query"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVariableName holds the string denoting the variablename field in the database.
	FieldVariableName = "variable_name"
	// FieldGqlAst holds the string denoting the gqlast field in the database.
	FieldGqlAst = "gql_ast"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeQueryOf holds the string denoting the queryof edge name in mutations.
	EdgeQueryOf = "queryOf"
	// Table holds the table name of the graphqlquery in the database.
	Table = "graph_ql_queries"
	// QueryOfTable is the table the holds the queryOf relation/edge.
	QueryOfTable = "graph_ql_queries"
	// QueryOfInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	QueryOfInverseTable = "projects"
	// QueryOfColumn is the table column denoting the queryOf relation/edge.
	QueryOfColumn = "project_queries"
)

// Columns holds all SQL columns for graphqlquery fields.
var Columns = []string{
	FieldID,
	FieldVariableName,
	FieldGqlAst,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "graph_ql_queries"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_queries",
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
