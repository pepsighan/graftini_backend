// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/project"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/user"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GraphqlEndpoint holds the value of the "graphql_endpoint" field.
	GraphqlEndpoint *string `json:"graphql_endpoint,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges         ProjectEdges `json:"edges"`
	user_projects *uuid.UUID
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Pages holds the value of the pages edge.
	Pages []*Page `json:"pages,omitempty"`
	// Queries holds the value of the queries edge.
	Queries []*GraphQLQuery `json:"queries,omitempty"`
	// Deployments holds the value of the deployments edge.
	Deployments []*Deployment `json:"deployments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// PagesOrErr returns the Pages value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) PagesOrErr() ([]*Page, error) {
	if e.loadedTypes[1] {
		return e.Pages, nil
	}
	return nil, &NotLoadedError{edge: "pages"}
}

// QueriesOrErr returns the Queries value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) QueriesOrErr() ([]*GraphQLQuery, error) {
	if e.loadedTypes[2] {
		return e.Queries, nil
	}
	return nil, &NotLoadedError{edge: "queries"}
}

// DeploymentsOrErr returns the Deployments value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) DeploymentsOrErr() ([]*Deployment, error) {
	if e.loadedTypes[3] {
		return e.Deployments, nil
	}
	return nil, &NotLoadedError{edge: "deployments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldName, project.FieldGraphqlEndpoint:
			values[i] = new(sql.NullString)
		case project.FieldCreatedAt, project.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case project.FieldID:
			values[i] = new(uuid.UUID)
		case project.ForeignKeys[0]: // user_projects
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Project", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldGraphqlEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field graphql_endpoint", values[i])
			} else if value.Valid {
				pr.GraphqlEndpoint = new(string)
				*pr.GraphqlEndpoint = value.String
			}
		case project.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case project.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case project.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_projects", values[i])
			} else if value != nil {
				pr.user_projects = value
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Project entity.
func (pr *Project) QueryOwner() *UserQuery {
	return (&ProjectClient{config: pr.config}).QueryOwner(pr)
}

// QueryPages queries the "pages" edge of the Project entity.
func (pr *Project) QueryPages() *PageQuery {
	return (&ProjectClient{config: pr.config}).QueryPages(pr)
}

// QueryQueries queries the "queries" edge of the Project entity.
func (pr *Project) QueryQueries() *GraphQLQueryQuery {
	return (&ProjectClient{config: pr.config}).QueryQueries(pr)
}

// QueryDeployments queries the "deployments" edge of the Project entity.
func (pr *Project) QueryDeployments() *DeploymentQuery {
	return (&ProjectClient{config: pr.config}).QueryDeployments(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return (&ProjectClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	if v := pr.GraphqlEndpoint; v != nil {
		builder.WriteString(", graphql_endpoint=")
		builder.WriteString(*v)
	}
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project

func (pr Projects) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}