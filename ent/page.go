// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/pepsighan/nocodepress_backend/ent/page"
	"github.com/pepsighan/nocodepress_backend/ent/project"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Route holds the value of the "route" field.
	Route string `json:"route,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PageQuery when eager-loading is set.
	Edges         PageEdges `json:"edges"`
	project_pages *int
}

// PageEdges holds the relations/edges for other nodes in the graph.
type PageEdges struct {
	// PageOf holds the value of the pageOf edge.
	PageOf *Project `json:"pageOf,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PageOfOrErr returns the PageOf value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PageEdges) PageOfOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.PageOf == nil {
			// The edge pageOf was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.PageOf, nil
	}
	return nil, &NotLoadedError{edge: "pageOf"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			values[i] = new(sql.NullInt64)
		case page.FieldName, page.FieldRoute:
			values[i] = new(sql.NullString)
		case page.ForeignKeys[0]: // project_pages
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Page", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case page.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case page.FieldRoute:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field route", values[i])
			} else if value.Valid {
				pa.Route = value.String
			}
		case page.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_pages", value)
			} else if value.Valid {
				pa.project_pages = new(int)
				*pa.project_pages = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPageOf queries the "pageOf" edge of the Page entity.
func (pa *Page) QueryPageOf() *ProjectQuery {
	return (&PageClient{config: pa.config}).QueryPageOf(pa)
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return (&PageClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Page is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", route=")
	builder.WriteString(pa.Route)
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page

func (pa Pages) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}