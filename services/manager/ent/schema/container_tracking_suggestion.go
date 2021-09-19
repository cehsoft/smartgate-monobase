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
		field.String("container_id").Default("").StructTag(`db:"container_id"`), // deprecated in favor of result
		field.String("result").Default(""),
		field.Int("cam_id").Optional().StructTag(`db:"cam_id"`),
		field.Int("tracking_id").Optional().StructTag(`db:"tracking_id"`),
		field.Bool("is_valid").Default(false).StructTag(`db:"is_valid"`),
		field.String("tracking_type").Optional().StructTag(`db:"tracking_type"`),
		field.String("bic").Default(""),
		field.String("serial").Default(""),
		field.String("checksum").Default(""),
		field.String("image_url").Optional().StructTag(`db:"image_url"`),
		field.Float32("score"),
		field.Time("created_at").Default(time.Now).StructTag(`db:"created_at"`),
	}
}

// Edges of the ContainerTracking.
func (ContainerTrackingSuggestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cam", CamSetting.Type).Ref("suggestions").Field("cam_id").Unique(),
		edge.From("tracking", ContainerTracking.Type).Ref("suggestions").Field("tracking_id").Unique(),
	}
}

func (ContainerTrackingSuggestion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("container_id"),
	}
}
