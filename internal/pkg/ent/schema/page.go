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

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("route"),
		// This is where the design of the page would be stored in a serialized standard format.
		field.String("component_map").
			Validate(func(s string) error {
				var compMap ComponentMap
				// It is valid only if it can be read into the schema.
				return json.Unmarshal([]byte(s), &compMap)
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

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("page_of", Project.Type).
			Ref("pages").
			Unique(),
	}
}
