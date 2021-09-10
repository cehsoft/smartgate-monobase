package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Gate struct {
	ent.Schema
}

// Fields of the User.
func (Gate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (Gate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lanes", Lane.Type), // One Tracking has many Suggestions
	}
}
