// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContainerTrackingsColumns holds the columns for the "container_trackings" table.
	ContainerTrackingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "container_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ContainerTrackingsTable holds the schema information for the "container_trackings" table.
	ContainerTrackingsTable = &schema.Table{
		Name:        "container_trackings",
		Columns:     ContainerTrackingsColumns,
		PrimaryKey:  []*schema.Column{ContainerTrackingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "containertracking_container_id",
				Unique:  false,
				Columns: []*schema.Column{ContainerTrackingsColumns[1]},
			},
		},
	}
	// ContainerTrackingSuggestionsColumns holds the columns for the "container_tracking_suggestions" table.
	ContainerTrackingSuggestionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "container_id", Type: field.TypeString},
		{Name: "bic", Type: field.TypeString, Nullable: true},
		{Name: "serial", Type: field.TypeString, Nullable: true},
		{Name: "checksum", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "score", Type: field.TypeFloat32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "tracking_id", Type: field.TypeInt, Nullable: true},
	}
	// ContainerTrackingSuggestionsTable holds the schema information for the "container_tracking_suggestions" table.
	ContainerTrackingSuggestionsTable = &schema.Table{
		Name:       "container_tracking_suggestions",
		Columns:    ContainerTrackingSuggestionsColumns,
		PrimaryKey: []*schema.Column{ContainerTrackingSuggestionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "container_tracking_suggestions_container_trackings_suggestions",
				Columns:    []*schema.Column{ContainerTrackingSuggestionsColumns[8]},
				RefColumns: []*schema.Column{ContainerTrackingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "containertrackingsuggestion_container_id",
				Unique:  false,
				Columns: []*schema.Column{ContainerTrackingSuggestionsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContainerTrackingsTable,
		ContainerTrackingSuggestionsTable,
		UsersTable,
	}
)

func init() {
	ContainerTrackingSuggestionsTable.ForeignKeys[0].RefTable = ContainerTrackingsTable
}
