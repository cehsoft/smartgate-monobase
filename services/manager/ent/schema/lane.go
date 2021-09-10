package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Lane struct {
	ent.Schema
}

// Fields of the User.
func (Lane) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("gate_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (Lane) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cams", CamSetting.Type), // One Tracking has many Suggestions
		edge.From("gate", Gate.Type).Ref("lanes").Field("gate_id").Unique(),
	}
}
