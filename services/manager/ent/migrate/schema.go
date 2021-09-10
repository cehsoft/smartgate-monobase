// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CamSettingsColumns holds the columns for the "cam_settings" table.
	CamSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "rtsp_url", Type: field.TypeString},
		{Name: "webrtc_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "lane_id", Type: field.TypeInt, Nullable: true},
	}
	// CamSettingsTable holds the schema information for the "cam_settings" table.
	CamSettingsTable = &schema.Table{
		Name:       "cam_settings",
		Columns:    CamSettingsColumns,
		PrimaryKey: []*schema.Column{CamSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cam_settings_lanes_cams",
				Columns:    []*schema.Column{CamSettingsColumns[5]},
				RefColumns: []*schema.Column{LanesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ContainerTrackingsColumns holds the columns for the "container_trackings" table.
	ContainerTrackingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "container_id", Type: field.TypeString},
		{Name: "session_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ContainerTrackingsTable holds the schema information for the "container_trackings" table.
	ContainerTrackingsTable = &schema.Table{
		Name:       "container_trackings",
		Columns:    ContainerTrackingsColumns,
		PrimaryKey: []*schema.Column{ContainerTrackingsColumns[0]},
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
		{Name: "result", Type: field.TypeString, Nullable: true},
		{Name: "tracking_type", Type: field.TypeString, Nullable: true},
		{Name: "bic", Type: field.TypeString, Nullable: true},
		{Name: "serial", Type: field.TypeString, Nullable: true},
		{Name: "checksum", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "score", Type: field.TypeFloat32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "cam_id", Type: field.TypeInt, Nullable: true},
		{Name: "tracking_id", Type: field.TypeInt, Nullable: true},
	}
	// ContainerTrackingSuggestionsTable holds the schema information for the "container_tracking_suggestions" table.
	ContainerTrackingSuggestionsTable = &schema.Table{
		Name:       "container_tracking_suggestions",
		Columns:    ContainerTrackingSuggestionsColumns,
		PrimaryKey: []*schema.Column{ContainerTrackingSuggestionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "container_tracking_suggestions_cam_settings_suggestions",
				Columns:    []*schema.Column{ContainerTrackingSuggestionsColumns[10]},
				RefColumns: []*schema.Column{CamSettingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "container_tracking_suggestions_container_trackings_suggestions",
				Columns:    []*schema.Column{ContainerTrackingSuggestionsColumns[11]},
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
	// GatesColumns holds the columns for the "gates" table.
	GatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// GatesTable holds the schema information for the "gates" table.
	GatesTable = &schema.Table{
		Name:       "gates",
		Columns:    GatesColumns,
		PrimaryKey: []*schema.Column{GatesColumns[0]},
	}
	// LanesColumns holds the columns for the "lanes" table.
	LanesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "gate_id", Type: field.TypeInt, Nullable: true},
	}
	// LanesTable holds the schema information for the "lanes" table.
	LanesTable = &schema.Table{
		Name:       "lanes",
		Columns:    LanesColumns,
		PrimaryKey: []*schema.Column{LanesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lanes_gates_lanes",
				Columns:    []*schema.Column{LanesColumns[3]},
				RefColumns: []*schema.Column{GatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CamSettingsTable,
		ContainerTrackingsTable,
		ContainerTrackingSuggestionsTable,
		GatesTable,
		LanesTable,
		UsersTable,
	}
)

func init() {
	CamSettingsTable.ForeignKeys[0].RefTable = LanesTable
	ContainerTrackingSuggestionsTable.ForeignKeys[0].RefTable = CamSettingsTable
	ContainerTrackingSuggestionsTable.ForeignKeys[1].RefTable = ContainerTrackingsTable
	LanesTable.ForeignKeys[0].RefTable = GatesTable
}
