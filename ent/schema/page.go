package schema

import (
	"encoding/json"
	"time"

	"entgo.io/ent"
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
		// This is where the design of the page would be stored in a serialized format.
		field.String("markup").
			Default(defaultMarkup()).
			Validate(func(s string) error {
				var markup Markup
				// It is valid only if it can be read into the Markup schema.
				return json.Unmarshal([]byte(s), &markup)
			}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pageOf", Project.Type).
			Ref("pages").
			Unique(),
	}
}
