package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type CamSetting struct {
	ent.Schema
}

// Fields of the User.
func (CamSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("lane_id").Optional(),
		field.String("rtsp_url"),
		field.String("webrtc_url"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (CamSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lane", Lane.Type).Ref("cams").Field("lane_id").Unique(),
		edge.To("suggestions", ContainerTrackingSuggestion.Type), // One Tracking has many Suggestions
	}
}
