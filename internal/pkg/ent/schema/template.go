package schema

import (
	"encoding/json"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("snapshot"),
		field.UUID("preview_file_id", uuid.UUID{}).
			Default(uuid.New).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return nil
}

type ProjectTemplateSnapshot struct {
	Pages []*PageTemplateSnapshot `json:"pages"`
}

type PageTemplateSnapshot struct {
	Name         string       `json:"name"`
	Route        string       `json:"route"`
	ComponentMap ComponentMap `json:"componentMap"`
}

// ParseStringToProjectTemplateSnapshot parses the snapshot string to a struct.
func ParseStringToProjectTemplateSnapshot(snapshot string) (*ProjectTemplateSnapshot, error) {
	templateSnapshot := ProjectTemplateSnapshot{}

	err := json.Unmarshal([]byte(snapshot), &templateSnapshot)
	if err != nil {
		return nil, err
	}

	return &templateSnapshot, nil
}
