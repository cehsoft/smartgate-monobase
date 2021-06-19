package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ContainerTracking holds the schema definition for the ContainerTracking entity.
type ContainerTracking struct {
	ent.Schema
}

// Fields of the ContainerTracking.
func (ContainerTracking) Fields() []ent.Field {
	return []ent.Field{
		field.String("container_id"),
		field.String("image_url").Optional(),
		field.Bool("manual").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the ContainerTracking.
func (ContainerTracking) Edges() []ent.Edge {
	return nil
}

func (ContainerTracking) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("container_id"),
	}
}
