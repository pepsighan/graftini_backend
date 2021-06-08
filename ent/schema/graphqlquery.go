package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GraphQLQuery holds the schema definition for the GraphQLQuery entity.
type GraphQLQuery struct {
	ent.Schema
}

// Fields of the GraphQLQuery.
func (GraphQLQuery) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("variableName"),
		field.String("gqlAst"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GraphQLQuery.
func (GraphQLQuery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("queryOf", Project.Type).
			Ref("queries").
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
