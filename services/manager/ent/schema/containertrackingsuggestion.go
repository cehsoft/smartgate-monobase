package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ContainerTracking holds the schema definition for the ContainerTracking entity.
type ContainerTrackingSuggestion struct {
	ent.Schema
}

// Fields of the ContainerTracking.
func (ContainerTrackingSuggestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tracking_id").Optional(),
		field.String("container_id"),
		field.String("image_url").Optional(),
		field.Float32("score"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the ContainerTracking.
func (ContainerTrackingSuggestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tracking", ContainerTracking.Type).Ref("suggestions").Field("tracking_id").Unique(),
	}
}

func (ContainerTrackingSuggestion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("container_id"),
	}
}
