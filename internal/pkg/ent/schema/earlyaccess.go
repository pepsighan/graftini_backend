package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EarlyAccess holds the schema definition for the EarlyAccess entity.
type EarlyAccess struct {
	ent.Schema
}

// Fields of the EarlyAccess.
func (EarlyAccess) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the EarlyAccess.
func (EarlyAccess) Edges() []ent.Edge {
	return nil
}
