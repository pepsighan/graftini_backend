// Code generated by entc, DO NOT EDIT.

package page

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the page type in the database.
	Label = "page"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRoute holds the string denoting the route field in the database.
	FieldRoute = "route"
	// EdgePageOf holds the string denoting the pageof edge name in mutations.
	EdgePageOf = "pageOf"
	// Table holds the table name of the page in the database.
	Table = "pages"
	// PageOfTable is the table the holds the pageOf relation/edge.
	PageOfTable = "pages"
	// PageOfInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	PageOfInverseTable = "projects"
	// PageOfColumn is the table column denoting the pageOf relation/edge.
	PageOfColumn = "project_pages"
)

// Columns holds all SQL columns for page fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRoute,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_pages",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
